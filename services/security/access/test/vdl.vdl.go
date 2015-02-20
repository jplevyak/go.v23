// This file was auto-generated by the veyron vdl tool.
// Source: vdl.vdl

// Package test provides a VDL specification for a service used in the unittest of the acl package.
package test

import (
	// VDL system imports
	"v.io/core/veyron2"
	"v.io/core/veyron2/context"
	"v.io/core/veyron2/ipc"
	"v.io/core/veyron2/vdl"
)

// Any package can define tags (of arbitrary types) to be attached to methods.
// This type can be used to index into a TaggedACLMap.
type MyTag string

func (MyTag) __VDLReflect(struct {
	Name string "v.io/core/veyron2/services/security/access/test.MyTag"
}) {
}

func init() {
	vdl.Register((*MyTag)(nil))
}

// For this example/unittest, there are three possible values of MyTag,
// each represented by a single-character string.
const Read = MyTag("R")

const Write = MyTag("W")

const Execute = MyTag("X")

// MyObjectClientMethods is the client interface
// containing MyObject methods.
//
// MyObject demonstrates how tags are attached to methods.
type MyObjectClientMethods interface {
	Get(*context.T, ...ipc.CallOpt) error
	Put(*context.T, ...ipc.CallOpt) error
	Resolve(*context.T, ...ipc.CallOpt) error
	NoTags(*context.T, ...ipc.CallOpt) error // No tags attached to this.
	AllTags(*context.T, ...ipc.CallOpt) error
}

// MyObjectClientStub adds universal methods to MyObjectClientMethods.
type MyObjectClientStub interface {
	MyObjectClientMethods
	ipc.UniversalServiceMethods
}

// MyObjectClient returns a client stub for MyObject.
func MyObjectClient(name string, opts ...ipc.BindOpt) MyObjectClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implMyObjectClientStub{name, client}
}

type implMyObjectClientStub struct {
	name   string
	client ipc.Client
}

func (c implMyObjectClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return veyron2.GetClient(ctx)
}

func (c implMyObjectClientStub) Get(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Get", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMyObjectClientStub) Put(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Put", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMyObjectClientStub) Resolve(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Resolve", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMyObjectClientStub) NoTags(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "NoTags", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMyObjectClientStub) AllTags(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "AllTags", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// MyObjectServerMethods is the interface a server writer
// implements for MyObject.
//
// MyObject demonstrates how tags are attached to methods.
type MyObjectServerMethods interface {
	Get(ipc.ServerContext) error
	Put(ipc.ServerContext) error
	Resolve(ipc.ServerContext) error
	NoTags(ipc.ServerContext) error // No tags attached to this.
	AllTags(ipc.ServerContext) error
}

// MyObjectServerStubMethods is the server interface containing
// MyObject methods, as expected by ipc.Server.
// There is no difference between this interface and MyObjectServerMethods
// since there are no streaming methods.
type MyObjectServerStubMethods MyObjectServerMethods

// MyObjectServerStub adds universal methods to MyObjectServerStubMethods.
type MyObjectServerStub interface {
	MyObjectServerStubMethods
	// Describe the MyObject interfaces.
	Describe__() []ipc.InterfaceDesc
}

// MyObjectServer returns a server stub for MyObject.
// It converts an implementation of MyObjectServerMethods into
// an object that may be used by ipc.Server.
func MyObjectServer(impl MyObjectServerMethods) MyObjectServerStub {
	stub := implMyObjectServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implMyObjectServerStub struct {
	impl MyObjectServerMethods
	gs   *ipc.GlobState
}

func (s implMyObjectServerStub) Get(ctx ipc.ServerContext) error {
	return s.impl.Get(ctx)
}

func (s implMyObjectServerStub) Put(ctx ipc.ServerContext) error {
	return s.impl.Put(ctx)
}

func (s implMyObjectServerStub) Resolve(ctx ipc.ServerContext) error {
	return s.impl.Resolve(ctx)
}

func (s implMyObjectServerStub) NoTags(ctx ipc.ServerContext) error {
	return s.impl.NoTags(ctx)
}

func (s implMyObjectServerStub) AllTags(ctx ipc.ServerContext) error {
	return s.impl.AllTags(ctx)
}

func (s implMyObjectServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implMyObjectServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{MyObjectDesc}
}

// MyObjectDesc describes the MyObject interface.
var MyObjectDesc ipc.InterfaceDesc = descMyObject

// descMyObject hides the desc to keep godoc clean.
var descMyObject = ipc.InterfaceDesc{
	Name:    "MyObject",
	PkgPath: "v.io/core/veyron2/services/security/access/test",
	Doc:     "// MyObject demonstrates how tags are attached to methods.",
	Methods: []ipc.MethodDesc{
		{
			Name: "Get",
			Tags: []*vdl.Value{vdl.ValueOf(MyTag("R"))},
		},
		{
			Name: "Put",
			Tags: []*vdl.Value{vdl.ValueOf(MyTag("W"))},
		},
		{
			Name: "Resolve",
			Tags: []*vdl.Value{vdl.ValueOf(MyTag("X"))},
		},
		{
			Name: "NoTags",
		},
		{
			Name: "AllTags",
			Tags: []*vdl.Value{vdl.ValueOf(MyTag("R")), vdl.ValueOf(MyTag("W")), vdl.ValueOf(MyTag("X"))},
		},
	},
}
