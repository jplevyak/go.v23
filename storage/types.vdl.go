// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

package storage

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_vdlutil "veyron2/vdl/vdlutil"
)

// ID is the type of keys for the key/value store.  The key is a numeric
// identifier that is globally unique, meaning that if two stores contain an
// entry with the same ID, then the entries represent the same thing.  To
// ensure uniqueness, stores create unique IDs that have never before been used
// as identifier.  The ID is large enough that collisions are extremely
// unlikely.
type ID [16]byte

// Version identifies the value in the store for a key at some point in time.
// The version is a numeric identifier that is globally unique within the space
// of a single ID, meaning that if two stores contain an entry with the same ID
// and version, then the entries represent the same thing, at the same point in
// time (as agreed upon by the two stores).
type Version uint64

// DEntry is a directory entry.
type DEntry struct {
	Name string
	ID   ID
}

// Stat provides information about an entry in the store.
//
// TODO(jyh): Specify versioning more precisely.
type Stat struct {
	// ID is the unique identifier of the entry.
	ID ID
	// MTimeNS is the last modification time in Unix nanoseconds (see time.UnixNano).
	//
	// TODO(jyh): Use Veyron Time when it gets implemented.
	MTimeNS int64
	// Attrs are the attributes associated with the entry.
	Attrs []_gen_vdlutil.Any
}

// Entry represents a value at some point in time in the store.
type Entry struct {
	// Stat is the entry's metadata.
	Stat Stat
	// Value is the value of the entry.
	Value _gen_vdlutil.Any
}

const (
	// NoVersion means the entry is not present in the store.
	NoVersion = Version(0)
)
