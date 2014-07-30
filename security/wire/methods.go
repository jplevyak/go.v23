package wire

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"

	"veyron2/security"
	"veyron2/vom"
)

// ErrNoIntegrity is the error returned when bytes of an object seem to have been tampered with.
var ErrNoIntegrity = errors.New("signature does not match bytes, possible tampering")

// errInvalidBlessingName returns an error specifying that the provided blessing name is invalid.
func errInvalidBlessingName(blessingName string) error {
	return fmt.Errorf("invalid blessing name:%q", blessingName)
}

// WriteString writes the length and contents of the provided string to the provided Writer.
func WriteString(w io.Writer, tmp []byte, s string) {
	w.Write(tmp[:binary.PutVarint(tmp, int64(len(s)))])
	io.WriteString(w, s)
}

// WriteBytes writes the length and contents of the provided byte slice to the provided Writer.
func WriteBytes(w io.Writer, tmp, b []byte) {
	w.Write(tmp[:binary.PutVarint(tmp, int64(len(b)))])
	w.Write(b)
}

func ellipticCurve(t KeyCurve) (elliptic.Curve, error) {
	switch t {
	case KeyCurveP256:
		return elliptic.P256(), nil
	default:
		return nil, fmt.Errorf("unrecognized elliptic curve %v", t)
	}
}

// Decode unmarshals the contents of the PublicKey object and returns
// and crypto.ecdsa.PublicKey object.
func (p *PublicKey) Decode() (*ecdsa.PublicKey, error) {
	curve, err := ellipticCurve(p.Curve)
	if err != nil {
		return nil, err
	}
	x, y := elliptic.Unmarshal(curve, p.XY)
	return &ecdsa.PublicKey{Curve: curve, X: x, Y: y}, nil
}

// Encode takes a crypto.ecdsa.PublicKey object, marshals its contents
// and populates the PublicKey object with them.
func (p *PublicKey) Encode(key *ecdsa.PublicKey) error {
	if key.Curve != elliptic.P256() {
		return fmt.Errorf("unrecognized elliptic curve %T", p.Curve)
	}
	p.Curve = KeyCurveP256
	p.XY = elliptic.Marshal(key.Curve, key.X, key.Y)
	return nil
}

// encode serializes a security.Caveat object and sets the resulting bytes on the Caveat object.
func (c *Caveat) encode(caveat security.Caveat) error {
	var b bytes.Buffer
	if err := vom.NewEncoder(&b).Encode(caveat); err != nil {
		return err
	}
	c.Bytes = b.Bytes()
	return nil
}

// EncodeCaveats encodes the provided set of security.ServiceCaveat objects into Caveat objects.
func EncodeCaveats(serviceCaveats []security.ServiceCaveat) ([]Caveat, error) {
	caveats := make([]Caveat, len(serviceCaveats))
	for i, c := range serviceCaveats {
		caveats[i].Service = c.Service
		if err := caveats[i].encode(c.Caveat); err != nil {
			return nil, err
		}
	}
	return caveats, nil
}

// Decode deserializes the contents of the Caveat object to obtain a security.Caveat object.
func (c *Caveat) Decode() (security.Caveat, error) {
	var caveat security.Caveat
	if err := vom.NewDecoder(bytes.NewReader(c.Bytes)).Decode(&caveat); err != nil {
		return nil, err
	}
	return caveat, nil
}

// DecodeThirdPartyCaveats decodes the provided Caveat objects into security.ThirdPartyCaveat
// objects. The resulting objects are wrapped in security.ServiceCaveat objects according
// to the services they are bound to.
func DecodeThirdPartyCaveats(caveats []Caveat) (thirdPartyCaveats []security.ServiceCaveat) {
	for _, wireCav := range caveats {
		cav, err := wireCav.Decode()
		if err != nil {
			continue
		}
		tpCav, ok := cav.(security.ThirdPartyCaveat)
		if !ok {
			continue
		}
		thirdPartyCaveats = append(thirdPartyCaveats, security.ServiceCaveat{Service: wireCav.Service, Caveat: tpCav})
	}
	return
}

// Validate verifies the restriction embedded inside the security.Caveat if the label
// is an empty string (indicating a universal caveat) or if the label matches the Name
// of the LocalID present in the provided context.
func (c *Caveat) Validate(ctx security.Context) error {
	// TODO(ataly): Is checking that the localID matches the caveat's Service pattern
	// the right choice here?
	if c.Service != security.AllPrincipals && (ctx.LocalID() == nil || !ctx.LocalID().Match(c.Service)) {
		return nil
	}
	cav, err := c.Decode()
	if err != nil {
		return err
	}
	return cav.Validate(ctx)
}

// -- Helper methods on the wire format for the chain implementation of Identity --

// contentHash returns a SHA256 hash of the contents of the certificate along with the
// provided signature.
func (c *Certificate) contentHash(issuerSignature security.Signature) []byte {
	h := sha256.New()
	tmp := make([]byte, binary.MaxVarintLen64)
	if issuerSignature.Hash != security.NoHash {
		WriteBytes(h, tmp, []byte(issuerSignature.Hash))
	}
	WriteBytes(h, tmp, issuerSignature.R)
	WriteBytes(h, tmp, issuerSignature.S)
	WriteString(h, tmp, c.Name)
	h.Write([]byte{byte(c.PublicKey.Curve)})
	WriteBytes(h, tmp, c.PublicKey.XY)
	binary.Write(h, binary.BigEndian, uint32(len(c.Caveats)))
	for _, cav := range c.Caveats {
		WriteString(h, tmp, string(cav.Service))
		WriteBytes(h, tmp, cav.Bytes)
	}
	return h.Sum(nil)
}

// cryptoHash hashes the provided data using the specified hash function, returning nil
// if the hash couldn't be applied.
func cryptoHash(hash security.Hash, data []byte) []byte {
	switch hash {
	case security.NoHash:
		return data
	case security.SHA1Hash:
		h := sha1.Sum(data)
		return h[:]
	case security.SHA256Hash:
		h := sha256.Sum256(data)
		return h[:]
	case security.SHA384Hash:
		h := sha512.Sum384(data)
		return h[:]
	case security.SHA512Hash:
		h := sha512.Sum512(data)
		return h[:]
	}
	return nil
}

// Sign uses the given Signer to sign the signature of the last certificate in
// the provided PublicID, storing the new signature in the current certificate.
func (c *Certificate) Sign(signer security.Signer, pubID *ChainPublicID) error {
	numCerts := len(pubID.Certificates)
	if numCerts == 0 {
		return errors.New("cannot sign a ChainPublicID with no certificates")
	}
	var err error
	c.Signature, err = signer.Sign(c.contentHash(pubID.Certificates[numCerts-1].Signature))
	return err
}

func (c *Certificate) verify(issuerSignature security.Signature, key *ecdsa.PublicKey) bool {
	msg := c.contentHash(issuerSignature)
	if msg = cryptoHash(issuerSignature.Hash, msg); msg == nil {
		return false
	}
	var r, s big.Int
	return ecdsa.Verify(key, msg, r.SetBytes(c.Signature.R), s.SetBytes(c.Signature.S))
}

// ValidateCaveats verifies if all caveats present on the certificate validate with
// respect to the provided context.
func (c *Certificate) ValidateCaveats(ctx security.Context) error {
	for _, cav := range c.Caveats {
		if err := cav.Validate(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Name returns the chained name obtained by joining all names along the ChainPublicID's
// certificate chain.
func (id *ChainPublicID) Name() string {
	var buf bytes.Buffer
	for i, c := range id.Certificates {
		if i > 0 {
			buf.WriteString(ChainSeparator)
		}
		buf.WriteString(c.Name)
	}
	return buf.String()
}

// VerifyIntegrity verifies that the ChainPublicID has a valid certificate chain, i.e,
// (1) each certificate on the chain has a signature that can be verified using the
// public key specified in the previous certificate, (2) the first certificate's
// signature can be verified using its own public key, and (3) all certificate names
// are valid blessing names.
func (id *ChainPublicID) VerifyIntegrity() error {
	nCerts := len(id.Certificates)
	if nCerts == 0 {
		return ErrNoIntegrity
	}
	verificationKey, err := id.Certificates[0].PublicKey.Decode()
	if err != nil {
		return ErrNoIntegrity
	}
	issuerSignature := security.Signature{}
	for _, c := range id.Certificates {
		if err := ValidateBlessingName(c.Name); err != nil {
			return err
		}
		// TODO(ashankar, ataly): Do we worry about timing attacks by
		// early exiting on invalid certificate?
		if !c.verify(issuerSignature, verificationKey) {
			return ErrNoIntegrity
		}
		if verificationKey, err = c.PublicKey.Decode(); err != nil {
			return ErrNoIntegrity
		}
		issuerSignature = c.Signature
	}
	return nil
}

// ValidateBlessingName verifies if the provided name is fit to be the name of a blessing.
func ValidateBlessingName(name string) error {
	// TODO(ataly, ashankar): Define the list of reserved characters (such as  "*", "#",
	// "/", "\", etc.) and ensure that the check below ensures absence of all of them.
	if name == "" || strings.ContainsAny(name, ChainSeparator) {
		return errInvalidBlessingName(name)
	}
	return nil
}
