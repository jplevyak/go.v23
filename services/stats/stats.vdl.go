// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: stats.vdl

// Package stats defines an interface to access statistical information for
// troubleshooting and monitoring purposes.
package stats

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/security/access"
	"v.io/v23/services/watch"
)

var (
	ErrNoValue = verror.Register("v.io/v23/services/stats.NoValue", verror.NoRetry, "{1:}{2:} object has no value, suffix: {3}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoValue.ID), "{1:}{2:} object has no value, suffix: {3}")
}

// NewErrNoValue returns an error with the ErrNoValue ID.
func NewErrNoValue(ctx *context.T, suffix string) error {
	return verror.New(ErrNoValue, ctx, suffix)
}

// StatsClientMethods is the client interface
// containing Stats methods.
//
// The Stats interface is used to access stats for troubleshooting and
// monitoring purposes. The stats objects are discoverable via the Globbable
// interface and watchable via the GlobWatcher interface.
//
// The types of the object values are implementation specific, but should be
// primarily numeric in nature, e.g. counters, memory usage, latency metrics,
// etc.
type StatsClientMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherClientMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, ...rpc.CallOpt) (*vdl.Value, error)
}

// StatsClientStub adds universal methods to StatsClientMethods.
type StatsClientStub interface {
	StatsClientMethods
	rpc.UniversalServiceMethods
}

// StatsClient returns a client stub for Stats.
func StatsClient(name string) StatsClientStub {
	return implStatsClientStub{name, watch.GlobWatcherClient(name)}
}

type implStatsClientStub struct {
	name string

	watch.GlobWatcherClientStub
}

func (c implStatsClientStub) Value(ctx *context.T, opts ...rpc.CallOpt) (o0 *vdl.Value, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Value", nil, []interface{}{&o0}, opts...)
	return
}

// StatsServerMethods is the interface a server writer
// implements for Stats.
//
// The Stats interface is used to access stats for troubleshooting and
// monitoring purposes. The stats objects are discoverable via the Globbable
// interface and watchable via the GlobWatcher interface.
//
// The types of the object values are implementation specific, but should be
// primarily numeric in nature, e.g. counters, memory usage, latency metrics,
// etc.
type StatsServerMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherServerMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, rpc.ServerCall) (*vdl.Value, error)
}

// StatsServerStubMethods is the server interface containing
// Stats methods, as expected by rpc.Server.
// The only difference between this interface and StatsServerMethods
// is the streaming methods.
type StatsServerStubMethods interface {
	// GlobWatcher allows a client to receive updates for changes to objects
	// that match a pattern.  See the package comments for details.
	watch.GlobWatcherServerStubMethods
	// Value returns the current value of an object, or an error. The type
	// of the value is implementation specific.
	// Some objects may not have a value, in which case, Value() returns
	// a NoValue error.
	Value(*context.T, rpc.ServerCall) (*vdl.Value, error)
}

// StatsServerStub adds universal methods to StatsServerStubMethods.
type StatsServerStub interface {
	StatsServerStubMethods
	// Describe the Stats interfaces.
	Describe__() []rpc.InterfaceDesc
}

// StatsServer returns a server stub for Stats.
// It converts an implementation of StatsServerMethods into
// an object that may be used by rpc.Server.
func StatsServer(impl StatsServerMethods) StatsServerStub {
	stub := implStatsServerStub{
		impl: impl,
		GlobWatcherServerStub: watch.GlobWatcherServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implStatsServerStub struct {
	impl StatsServerMethods
	watch.GlobWatcherServerStub
	gs *rpc.GlobState
}

func (s implStatsServerStub) Value(ctx *context.T, call rpc.ServerCall) (*vdl.Value, error) {
	return s.impl.Value(ctx, call)
}

func (s implStatsServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implStatsServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{StatsDesc, watch.GlobWatcherDesc}
}

// StatsDesc describes the Stats interface.
var StatsDesc rpc.InterfaceDesc = descStats

// descStats hides the desc to keep godoc clean.
var descStats = rpc.InterfaceDesc{
	Name:    "Stats",
	PkgPath: "v.io/v23/services/stats",
	Doc:     "// The Stats interface is used to access stats for troubleshooting and\n// monitoring purposes. The stats objects are discoverable via the Globbable\n// interface and watchable via the GlobWatcher interface.\n//\n// The types of the object values are implementation specific, but should be\n// primarily numeric in nature, e.g. counters, memory usage, latency metrics,\n// etc.",
	Embeds: []rpc.EmbedDesc{
		{"GlobWatcher", "v.io/v23/services/watch", "// GlobWatcher allows a client to receive updates for changes to objects\n// that match a pattern.  See the package comments for details."},
	},
	Methods: []rpc.MethodDesc{
		{
			Name: "Value",
			Doc:  "// Value returns the current value of an object, or an error. The type\n// of the value is implementation specific.\n// Some objects may not have a value, in which case, Value() returns\n// a NoValue error.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // *vdl.Value
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
	},
}
