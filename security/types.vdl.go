// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package security

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/uniqueid"
)

// BlessingPattern is a pattern that is matched by specific blessings.
//
// A pattern can either be a blessing (slash-separated human-readable string)
// or a blessing ending in "/$". A pattern ending in "/$" is matched exactly
// by the blessing specified by the pattern string with the "/$" suffix stripped
// out. For example, the pattern "a/b/c/$" is matched by exactly by the blessing
// "a/b/c".
//
// A pattern not ending in "/$" is more permissive, and is also matched by blessings
// that are extensions of the pattern (including the pattern itself). For example, the
// pattern "a/b/c" is matched by the blessings "a/b/c", "a/b/c/x", "a/b/c/x/y", etc.
//
// TODO(ataly, ashankar): Define a formal BNF grammar for blessings and blessing patterns.
type BlessingPattern string

func (BlessingPattern) __VDLReflect(struct {
	Name string "v.io/v23/security.BlessingPattern"
}) {
}

// Hash identifies a cryptographic hash function approved for use in signature algorithms.
type Hash string

func (Hash) __VDLReflect(struct {
	Name string "v.io/v23/security.Hash"
}) {
}

// Signature represents a digital signature.
type Signature struct {
	// Purpose of the signature. Can be used to prevent type attacks.
	// (See Section 4.2 of http://www-users.cs.york.ac.uk/~jac/PublishedPapers/reviewV1_1997.pdf for example).
	// The actual signature (R, S values for ECDSA keys) is produced by signing: Hash(Hash(message), Hash(Purpose)).
	Purpose []byte
	// Cryptographic hash function applied to the message before computing the signature.
	Hash Hash
	// Pair of integers that make up an ECDSA signature.
	R []byte
	S []byte
}

func (Signature) __VDLReflect(struct {
	Name string "v.io/v23/security.Signature"
}) {
}

// ThirdPartyRequirements specifies the information required by the third-party
// that will issue discharges for third-party caveats.
//
// These requirements are typically used to construct a DischargeImpetus, which
// will be sent to the third-party.
type ThirdPartyRequirements struct {
	ReportServer    bool // The blessings presented by the server of an IPC call.
	ReportMethod    bool // The name of the method being invoked.
	ReportArguments bool // Arguments to the method being invoked.
}

func (ThirdPartyRequirements) __VDLReflect(struct {
	Name string "v.io/v23/security.ThirdPartyRequirements"
}) {
}

// DischargeImpetus encapsulates the motivation for a discharge being sought.
//
// These values are reported by a principal that is requesting a Discharge for
// a third-party caveat on one of its blessings. The third-party issues
// discharges cannot safely assume that all these values are provided, or that
// they are provided honestly.
//
// Implementations of services that issue discharges are encouraged to add
// caveats to the discharge that bind the discharge to the impetus, thereby
// rendering the discharge unsuable for any other purpose.
type DischargeImpetus struct {
	Server    []BlessingPattern // The client intends to use the discharge to communicate with a server that has a blessing matching one of the patterns in this set.
	Method    string            // Name of the method being invoked by the client.
	Arguments []*vdl.Value      // Arguments to the method invocation.
}

func (DischargeImpetus) __VDLReflect(struct {
	Name string "v.io/v23/security.DischargeImpetus"
}) {
}

// Certificate represents the cryptographic proof of the binding of
// extensions of a blessing held by one principal to another (represented by
// a public key) under specific caveats.
//
// For example, if a principal P1 has a blessing "alice", then it can
// extend it with a Certificate to generate the blessing "alice/friend" for
// another principal P2.
type Certificate struct {
	Extension string    // Human-readable string extension bound to PublicKey.
	PublicKey []byte    // DER-encoded PKIX public key.
	Caveats   []Caveat  // Caveats on the binding of Name to PublicKey.
	Signature Signature // Signature by the blessing principal that binds the extension to the public key.
}

func (Certificate) __VDLReflect(struct {
	Name string "v.io/v23/security.Certificate"
}) {
}

// CaveatDescriptor defines an association between a caveat validation function
// (addressed by globally unique identifier) and the data needed by the
// validation function.
//
// For a validator to be invoked, a validation function must be registered with
// the validator description in the language that the function is defined in.
type CaveatDescriptor struct {
	Id        uniqueid.Id // The identifier of the caveat validation function.
	ParamType *vdl.Type   // The type of the parameter expected by the validation function.
}

func (CaveatDescriptor) __VDLReflect(struct {
	Name string "v.io/v23/security.CaveatDescriptor"
}) {
}

// Caveat is a condition on the validity of a blessing/discharge.
//
// These conditions are provided when asking a principal to create
// a blessing/discharge and are verified when extracting blessings
// (Blessings.ForName in the Go API).
//
// Given a Hash, the message digest of a caveat is:
// Hash(Hash(Id), Hash(ParamVom))
type Caveat struct {
	Id       uniqueid.Id // The identifier of the caveat validation function.
	ParamVom []byte      // VOM-encoded bytes of the parameters to be provided to the validation function.
}

func (Caveat) __VDLReflect(struct {
	Name string "v.io/v23/security.Caveat"
}) {
}

// WireBlessings encapsulates wire format of a set of blessings and the
// corresponding cryptographic proof that binds them to a principal
// (identified by a public key).
//
// This structure is the "wire" format for sending and receiving blessings
// in RPCs or marshaling to persistent storage. Typically, languages will
// provide a factory function that converts this wire representation to
// a more usable object to inspect and manipulate these blessings. For
// example, the NewBlessings factory function in Go.
type WireBlessings struct {
	// CertificateChains is an array of chains of certificates that bind
	// a blessing to the public key in the last certificate of the chain.
	CertificateChains [][]Certificate
}

func (WireBlessings) __VDLReflect(struct {
	Name string "v.io/v23/security.WireBlessings"
}) {
}

type (
	// WireDischarge represents any single field of the WireDischarge union type.
	//
	// WireDischarge encapsulates the wire format of a third-party caveat
	// Discharge.
	WireDischarge interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the WireDischarge union type.
		__VDLReflect(__WireDischargeReflect)
	}
	// WireDischargePublicKey represents field PublicKey of the WireDischarge union type.
	WireDischargePublicKey struct{ Value publicKeyDischarge } // Discharge for PublicKeyThirdPartyCaveat
	// __WireDischargeReflect describes the WireDischarge union type.
	__WireDischargeReflect struct {
		Name  string "v.io/v23/security.WireDischarge"
		Type  WireDischarge
		Union struct {
			PublicKey WireDischargePublicKey
		}
	}
)

func (x WireDischargePublicKey) Index() int                          { return 0 }
func (x WireDischargePublicKey) Interface() interface{}              { return x.Value }
func (x WireDischargePublicKey) Name() string                        { return "PublicKey" }
func (x WireDischargePublicKey) __VDLReflect(__WireDischargeReflect) {}

// RejectedBlessing describes why a blessing failed validation.
type RejectedBlessing struct {
	Blessing string
	Err      error
}

func (RejectedBlessing) __VDLReflect(struct {
	Name string "v.io/v23/security.RejectedBlessing"
}) {
}

func init() {
	vdl.RegisterNative(wireBlessingsToNative, wireBlessingsFromNative)
	vdl.RegisterNative(wireDischargeToNative, wireDischargeFromNative)
	vdl.Register((*BlessingPattern)(nil))
	vdl.Register((*Hash)(nil))
	vdl.Register((*Signature)(nil))
	vdl.Register((*ThirdPartyRequirements)(nil))
	vdl.Register((*DischargeImpetus)(nil))
	vdl.Register((*Certificate)(nil))
	vdl.Register((*CaveatDescriptor)(nil))
	vdl.Register((*Caveat)(nil))
	vdl.Register((*WireBlessings)(nil))
	vdl.Register((*WireDischarge)(nil))
	vdl.Register((*RejectedBlessing)(nil))
}

// Type-check WireBlessings conversion functions.
var _ func(WireBlessings, *Blessings) error = wireBlessingsToNative
var _ func(*WireBlessings, Blessings) error = wireBlessingsFromNative

// Type-check WireDischarge conversion functions.
var _ func(WireDischarge, *Discharge) error = wireDischargeToNative
var _ func(*WireDischarge, Discharge) error = wireDischargeFromNative

// NoExtension is an optional terminator for a blessing pattern indicating that the pattern
// cannot match any extensions of the blessing from that point onwards.
const NoExtension = BlessingPattern("$")

// TODO(ataly, ashankar): The semantics of AllPrincipals breaks monotonicity in
// AccessLists with NotIn clauses. For instance, the AccessList "In: {AllPrincipals}, NotIn: {"foo"}
// matches the principal that presents no recognizable blessings ([]) however does not
// match the principal that presents "foo" as the only recognizable blessings (["foo"])
// We need to sort this out.
const AllPrincipals = BlessingPattern("...") // Glob pattern that matches all blessings.

const ChainSeparator = "/" // ChainSeparator joins blessing names to form a blessing chain name.

const SHA1Hash = Hash("SHA1") // SHA1 cryptographic hash function defined in RFC3174.

const SHA256Hash = Hash("SHA256") // SHA256 cryptographic hash function defined  in FIPS 180-4.

const SHA384Hash = Hash("SHA384") // SHA384 cryptographic hash function defined in FIPS 180-2.

const SHA512Hash = Hash("SHA512") // SHA512 cryptographic hash function defined in FIPS 180-2.

const SignatureForMessageSigning = "S" // Signature.Purpose used by a Principal to sign arbitrary messages.

const SignatureForBlessingCertificates = "B" // Signature.Purpose used by a Principal when signing Certificates for creating blessings.

const SignatureForDischarge = "D" // Signature.Purpose used by a Principal when signing discharges for public-key based third-party caveats.

var (
	ErrUnrecognizedRoot    = verror.Register("v.io/v23/security.UnrecognizedRoot", verror.NoRetry, "{1:}{2:} unrecognized public key {3} in root certificate{:4}")
	ErrAuthorizationFailed = verror.Register("v.io/v23/security.AuthorizationFailed", verror.NoRetry, "{1:}{2:} principal with blessings {3} (rejected {4}) is not authorized by principal with blessings {5}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnrecognizedRoot.ID), "{1:}{2:} unrecognized public key {3} in root certificate{:4}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrAuthorizationFailed.ID), "{1:}{2:} principal with blessings {3} (rejected {4}) is not authorized by principal with blessings {5}")
}

// NewErrUnrecognizedRoot returns an error with the ErrUnrecognizedRoot ID.
func NewErrUnrecognizedRoot(ctx *context.T, rootKey string, details error) error {
	return verror.New(ErrUnrecognizedRoot, ctx, rootKey, details)
}

// NewErrAuthorizationFailed returns an error with the ErrAuthorizationFailed ID.
func NewErrAuthorizationFailed(ctx *context.T, remote []string, remoteErr []RejectedBlessing, local []string) error {
	return verror.New(ErrAuthorizationFailed, ctx, remote, remoteErr, local)
}
