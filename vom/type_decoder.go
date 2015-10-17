// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vom

import (
	"io"
	"sync"

	"v.io/v23/vdl"
	"v.io/v23/verror"
)

var (
	errTypeInvalid        = verror.Register(pkgPath+".errTypeInvalid", verror.NoRetry, "{1:}{2:} vom: type {3} id {4} invalid, the min user type id is {5}{:_}")
	errAlreadyDefined     = verror.Register(pkgPath+".errAlreadyDefined", verror.NoRetry, "{1:}{2:} vom: type {3} id {4} already defined as {5}{:_}")
	errUnknownType        = verror.Register(pkgPath+".errUnknownType", verror.NoRetry, "{1:}{2:} vom: unknown type id {3}{:_}")
	errEmptyName          = verror.Register(pkgPath+".errEmptyName", verror.NoRetry, "{1:}{2:} vom: NamedType has empty name{:_}")
	errUnknownWireTypeDef = verror.Register(pkgPath+".errUnknownWireTypeDef", verror.NoRetry, "{1:}{2:} vom: unknown wire type definition {3}{:_}")
)

// TypeDecoder manages the receipt and unmarshalling of types from the other
// side of a connection.  Start must be called to start decoding types,
// and Stop must be called to reclaim resources.
type TypeDecoder struct {
	// The type encoder uses a 2-lock strategy for decoding. We use typeMu to lock
	// type definitions, and use buildMu to allow only one worker to build types at
	// a time. This is for simplifying the workflow and avoid unnecessary blocking
	// for type lookups.
	typeMu   sync.RWMutex
	idToType map[typeId]*vdl.Type // GUARDED_BY(typeMu)

	buildMu             sync.Mutex
	buildCond           *sync.Cond
	err                 error               // GUARDED_BY(buildMu)
	idToWire            map[typeId]wireType // GUARDED_BY(buildMu)
	dec                 *Decoder            // GUARDED_BY(buildMu)
	receivedVersionByte bool                // GUARDED_BY(buildMu)

	processingControlMu sync.Mutex
	goroutineRunning    bool // GUARDED_BY(processingControlMu)
	goroutineShouldStop bool // GUARDED_BY(processingControlMu)
}

// NewTypeDecoder returns a new TypeDecoder that reads from the given reader.
// The TypeDecoder understands all wire type formats generated by the TypeEncoder.
func NewTypeDecoder(r io.Reader) *TypeDecoder {
	mr := newMessageReader(newDecbuf(r))
	td := newTypeDecoderInternal(mr)
	mr.SetCallbacks(td.lookupType, nil)
	return td
}

func newTypeDecoderInternal(mr *messageReader) *TypeDecoder {
	td := &TypeDecoder{
		idToType: make(map[typeId]*vdl.Type),
		idToWire: make(map[typeId]wireType),
		dec: &Decoder{
			mr:      mr,
			typeDec: nil,
		},
	}
	td.buildCond = sync.NewCond(&td.buildMu)
	return td
}

func newDerivedTypeDecoderInternal(mr *messageReader, orig *TypeDecoder) *TypeDecoder {
	td := &TypeDecoder{
		idToType: orig.idToType,
		idToWire: orig.idToWire,
		dec: &Decoder{
			mr:      mr,
			typeDec: nil,
		},
	}
	td.buildCond = sync.NewCond(&td.buildMu)
	return td
}

func (d *TypeDecoder) processLoop() {
	var err error
	for {
		d.processingControlMu.Lock()
		if d.goroutineShouldStop || err != nil {
			d.goroutineShouldStop = false
			d.goroutineRunning = false
			d.processingControlMu.Unlock()
			return
		}
		d.processingControlMu.Unlock()
		// Note that we will block indefinitely if the underlying
		// read blocks on the io.Reader.
		err = d.readSingleType()
		d.buildMu.Lock()
		d.err = err
		d.buildCond.Broadcast()
		d.buildMu.Unlock()
		// TODO(toddw): Reconsider d.err and d.buildCond strategy.
	}
}

// Start must be called to start decoding types.
func (d *TypeDecoder) Start() {
	d.processingControlMu.Lock()
	d.goroutineShouldStop = false
	if !d.goroutineRunning {
		d.goroutineRunning = true
		go d.processLoop()
	}
	d.processingControlMu.Unlock()
}

// Stop must be called after Start, to stop decoding types
// and reclaim resources.  Once Stop is called,
// subsequent Decode calls on Decoders initialized with d
// will return errors.
func (d *TypeDecoder) Stop() {
	d.processingControlMu.Lock()
	d.goroutineShouldStop = true
	d.processingControlMu.Unlock()
}

// readSingleType reads a single wire type
func (d *TypeDecoder) readSingleType() error {
	var wt wireType
	curTypeID, err := d.dec.decodeWireType(&wt)
	if err != nil {
		return err
	}

	// Add the wire type and wake up waiters.
	return d.addWireType(curTypeID, wt)
}

// LookupType returns the type for tid. If the type is not yet available,
// this will wait until it arrives and is built.
func (d *TypeDecoder) lookupType(tid typeId) (*vdl.Type, error) {
	if tt := d.lookupKnownType(tid); tt != nil {
		return tt, nil
	}

	d.buildMu.Lock()
	defer d.buildMu.Unlock()
	for {
		if d.err != nil && d.err != io.EOF {
			// Return any existing error immediately. Skip EOF because it
			// may still be possible to lookup a type.
			return nil, d.err
		}

		if tt := d.lookupKnownType(tid); tt != nil {
			return tt, nil
		}

		// If the wire type is available, build it now.
		if _, exists := d.idToWire[tid]; exists {
			if err := d.buildType(tid); err != nil {
				// TODO(bprosnitz) This will fail if we receive types out of order. We should reconsider our type
				// format and ensure that we have all types that we need.
				return nil, err
			}

			if tt := d.lookupKnownType(tid); tt != nil {
				return tt, nil
			}
		}

		if d.err != nil {
			return nil, d.err
		}

		d.buildCond.Wait()
	}
}

// addWireType adds the wire type wt with the type id tid.
func (d *TypeDecoder) addWireType(tid typeId, wt wireType) error {
	d.buildMu.Lock()
	err := d.addWireTypeBuildLocked(tid, wt)
	d.buildMu.Unlock()
	return err
}

func (d *TypeDecoder) addWireTypeBuildLocked(tid typeId, wt wireType) error {
	if tid < WireIdFirstUserType {
		return verror.New(errTypeInvalid, nil, wt, tid, WireIdFirstUserType)
	}
	// TODO(toddw): Allow duplicates according to some heuristic (e.g. only
	// identical, or only if the later one is a "superset", etc).
	if dup := d.lookupKnownType(tid); dup != nil {
		return verror.New(errAlreadyDefined, nil, wt, tid, dup)
	}
	if dup := d.idToWire[tid]; dup != nil {
		return verror.New(errAlreadyDefined, nil, wt, tid, dup)
	}
	d.idToWire[tid] = wt
	return nil
}

func (d *TypeDecoder) lookupKnownType(tid typeId) *vdl.Type {
	if tt := bootstrapIdToType[tid]; tt != nil {
		return tt
	}
	d.typeMu.RLock()
	tt := d.idToType[tid]
	d.typeMu.RUnlock()
	return tt
}

// buildType builds the type from the given wire type.
func (d *TypeDecoder) buildType(tid typeId) error {
	builder := vdl.TypeBuilder{}
	pending := make(map[typeId]vdl.PendingType)
	_, err := d.makeType(tid, &builder, pending)
	if err != nil {
		return err
	}
	builder.Build()
	types := make(map[typeId]*vdl.Type)
	for tid, pt := range pending {
		tt, err := pt.Built()
		if err != nil {
			return err
		}
		types[tid] = tt
	}
	// Add the types to idToType map.
	d.typeMu.Lock()
	for tid, tt := range types {
		d.idToType[tid] = tt
	}
	d.typeMu.Unlock()
	return nil
}

// makeType makes the pending type from its wire type representation.
func (d *TypeDecoder) makeType(tid typeId, builder *vdl.TypeBuilder, pending map[typeId]vdl.PendingType) (vdl.PendingType, error) {
	wt := d.idToWire[tid]
	if wt == nil {
		return nil, verror.New(errUnknownType, nil, tid)
	}
	// Make the type from its wireType representation. First remove it from
	// dt.idToWire, and add it to pending, so that subsequent lookups will get the
	// pending type. Eventually the built type will be added to dt.idToType.
	delete(d.idToWire, tid)
	if name := wt.(wireTypeGeneric).TypeName(); name != "" {
		// Named types may be recursive, so we must create the named type first and
		// add it to pending, before we make the base type. The base type may refer
		// back to this named type, and will find it in pending.
		namedType := builder.Named(name)
		pending[tid] = namedType
		if wtNamed, ok := wt.(wireTypeNamedT); ok {
			// This is a NamedType pointing at a base type.
			baseType, err := d.lookupOrMakeType(wtNamed.Value.Base, builder, pending)
			if err != nil {
				return nil, err
			}
			namedType.AssignBase(baseType)
			return namedType, nil
		}
		// This isn't NamedType, but has a non-empty name.
		baseType, err := d.makeBaseType(wt, builder, pending)
		if err != nil {
			return nil, err
		}
		namedType.AssignBase(baseType)
		return namedType, nil
	}
	// Unnamed types are made directly from their base type.  It's fine to update
	// pending after making the base type, since there's no way to create a
	// recursive type based solely on unnamed vdl.
	baseType, err := d.makeBaseType(wt, builder, pending)
	if err != nil {
		return nil, err
	}
	pending[tid] = baseType
	return baseType, nil
}

func (d *TypeDecoder) makeBaseType(wt wireType, builder *vdl.TypeBuilder, pending map[typeId]vdl.PendingType) (vdl.PendingType, error) {
	switch wt := wt.(type) {
	case wireTypeNamedT:
		return nil, verror.New(errEmptyName, nil, wt)
	case wireTypeEnumT:
		enumType := builder.Enum()
		for _, label := range wt.Value.Labels {
			enumType.AppendLabel(label)
		}
		return enumType, nil
	case wireTypeArrayT:
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Array().AssignElem(elemType).AssignLen(int(wt.Value.Len)), nil
	case wireTypeListT:
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.List().AssignElem(elemType), nil
	case wireTypeSetT:
		keyType, err := d.lookupOrMakeType(wt.Value.Key, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Set().AssignKey(keyType), nil
	case wireTypeMapT:
		keyType, err := d.lookupOrMakeType(wt.Value.Key, builder, pending)
		if err != nil {
			return nil, err
		}
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Map().AssignKey(keyType).AssignElem(elemType), nil
	case wireTypeStructT:
		structType := builder.Struct()
		for _, field := range wt.Value.Fields {
			fieldType, err := d.lookupOrMakeType(field.Type, builder, pending)
			if err != nil {
				return nil, err
			}
			structType.AppendField(field.Name, fieldType)
		}
		return structType, nil
	case wireTypeUnionT:
		unionType := builder.Union()
		for _, field := range wt.Value.Fields {
			fieldType, err := d.lookupOrMakeType(field.Type, builder, pending)
			if err != nil {
				return nil, err
			}
			unionType.AppendField(field.Name, fieldType)
		}
		return unionType, nil
	case wireTypeOptionalT:
		elemType, err := d.lookupOrMakeType(wt.Value.Elem, builder, pending)
		if err != nil {
			return nil, err
		}
		return builder.Optional().AssignElem(elemType), nil
	default:
		return nil, verror.New(errUnknownWireTypeDef, nil, wt)
	}
}

func (d *TypeDecoder) lookupOrMakeType(tid typeId, builder *vdl.TypeBuilder, pending map[typeId]vdl.PendingType) (vdl.TypeOrPending, error) {
	if tt := d.lookupKnownType(tid); tt != nil {
		return tt, nil
	}
	if p, ok := pending[tid]; ok {
		return p, nil
	}
	return d.makeType(tid, builder, pending)
}
