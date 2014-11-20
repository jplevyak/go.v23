// This file was auto-generated by the veyron vdl tool.
// Source: arith.vdl

// Package arith is an example of an IDL definition in veyron.  The syntax for
// IDL files is similar to, but not identical to, Go.  Here are the main
// concepts:
//   * PACKAGES - Just like in Go you must define the package at the beginning
//     of an IDL file, and everything defined in the file is part of this
//     package.  By convention all files in the same dir should be in the same
//     package.
//   * IMPORTS - Just like in Go you can import other idl packages, and you may
//     assign a local package name, or if unspecified the basename of the import
//     path is used as the import package name.
//   * DATA TYPES - Just like in Go you can define data types.  You get most of
//     the primitives (int32, float64, string, etc), the "error" built-in, and a
//     special "any" built-in described below.  In addition you can create
//     composite types like arrays, structs, etc.
//   * CONSTS - Just like in Go you can define constants, and numerics are
//     "infinite precision" within expressions.  Unlike Go numerics must be
//     typed to be used as const definitions or tags.
//   * INTERFACES - Just like in Go you can define interface types, which are
//     just a set of methods.  Interfaces can embed other interfaces.  Unlike
//     Go, you cannot use an interface as a data type; interfaces are purely
//     method sets.
//   * ERRORS - Errors may be defined in IDL files, and unlike Go they work
//     across separate address spaces.
package arith

import (
	"veyron.io/veyron/veyron2/vdl/testdata/arith/exp"

	"veyron.io/veyron/veyron2/vdl/testdata/base"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// Yes shows that bools may be untyped.
const Yes = true // yes trailing doc

// No shows explicit boolean typing.
const No = false

const Hello = "hello"

// Int32Const shows explicit integer typing.
const Int32Const = int32(123)

// Int64Const shows explicit integer conversion from another type, and referencing
// a constant from another package.
const Int64Const = int64(128)

// FloatConst shows arithmetic expressions may be used.
const FloatConst = float64(2)

// Mask shows bitwise operations.
const Mask = uint64(256)

// ArithClientMethods is the client interface
// containing Arith methods.
//
// Arith is an example of an interface definition for an arithmetic service.
// Things to note:
//   * There must be at least 1 out-arg, and the last out-arg must be error.
type ArithClientMethods interface {
	// Add is a typical method with multiple input and output arguments.
	Add(ctx __context.T, a int32, b int32, opts ...__ipc.CallOpt) (int32, error)
	// DivMod shows that runs of args with the same type can use the short form,
	// just like Go.
	DivMod(ctx __context.T, a int32, b int32, opts ...__ipc.CallOpt) (quot int32, rem int32, err error)
	// Sub shows that you can use data types defined in other packages.
	Sub(ctx __context.T, args base.Args, opts ...__ipc.CallOpt) (int32, error)
	// Mul tries another data type defined in another package.
	Mul(ctx __context.T, nested base.NestedArgs, opts ...__ipc.CallOpt) (int32, error)
	// GenError shows that it's fine to have no in args, and no out args other
	// than "error".  In addition GenError shows the usage of tags.  Tags are a
	// sequence of constants.  There's no requirement on uniqueness of types or
	// values, and regular const expressions may also be used.
	GenError(__context.T, ...__ipc.CallOpt) error
	// Count shows using only an int32 out-stream type, with no in-stream type.
	Count(ctx __context.T, start int32, opts ...__ipc.CallOpt) (ArithCountCall, error)
	// StreamingAdd shows a bidirectional stream.
	StreamingAdd(__context.T, ...__ipc.CallOpt) (ArithStreamingAddCall, error)
	// QuoteAny shows the any built-in type, representing a value of any type.
	QuoteAny(ctx __context.T, a __vdlutil.Any, opts ...__ipc.CallOpt) (__vdlutil.Any, error)
}

// ArithClientStub adds universal methods to ArithClientMethods.
type ArithClientStub interface {
	ArithClientMethods
	__ipc.UniversalServiceMethods
}

// ArithClient returns a client stub for Arith.
func ArithClient(name string, opts ...__ipc.BindOpt) ArithClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implArithClientStub{name, client}
}

type implArithClientStub struct {
	name   string
	client __ipc.Client
}

func (c implArithClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implArithClientStub) Add(ctx __context.T, i0 int32, i1 int32, opts ...__ipc.CallOpt) (o0 int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Add", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implArithClientStub) DivMod(ctx __context.T, i0 int32, i1 int32, opts ...__ipc.CallOpt) (o0 int32, o1 int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "DivMod", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &o1, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implArithClientStub) Sub(ctx __context.T, i0 base.Args, opts ...__ipc.CallOpt) (o0 int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Sub", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implArithClientStub) Mul(ctx __context.T, i0 base.NestedArgs, opts ...__ipc.CallOpt) (o0 int32, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Mul", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implArithClientStub) GenError(ctx __context.T, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GenError", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implArithClientStub) Count(ctx __context.T, i0 int32, opts ...__ipc.CallOpt) (ocall ArithCountCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Count", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implArithCountCall{Call: call}
	return
}

func (c implArithClientStub) StreamingAdd(ctx __context.T, opts ...__ipc.CallOpt) (ocall ArithStreamingAddCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "StreamingAdd", nil, opts...); err != nil {
		return
	}
	ocall = &implArithStreamingAddCall{Call: call}
	return
}

func (c implArithClientStub) QuoteAny(ctx __context.T, i0 __vdlutil.Any, opts ...__ipc.CallOpt) (o0 __vdlutil.Any, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "QuoteAny", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implArithClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// ArithCountClientStream is the client stream for Arith.Count.
type ArithCountClientStream interface {
	// RecvStream returns the receiver side of the Arith.Count client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// ArithCountCall represents the call returned from Arith.Count.
type ArithCountCall interface {
	ArithCountClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implArithCountCall struct {
	__ipc.Call
	valRecv int32
	errRecv error
}

func (c *implArithCountCall) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implArithCountCallRecv{c}
}

type implArithCountCallRecv struct {
	c *implArithCountCall
}

func (c implArithCountCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implArithCountCallRecv) Value() int32 {
	return c.c.valRecv
}
func (c implArithCountCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implArithCountCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// ArithStreamingAddClientStream is the client stream for Arith.StreamingAdd.
type ArithStreamingAddClientStream interface {
	// RecvStream returns the receiver side of the Arith.StreamingAdd client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Arith.StreamingAdd client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending, or if Send is called after Close or Cancel.  Blocks if
		// there is no buffer space; will unblock when buffer space is available or
		// after Cancel.
		Send(item int32) error
		// Close indicates to the server that no more items will be sent; server
		// Recv calls will receive io.EOF after all sent items.  This is an optional
		// call - e.g. a client might call Close if it needs to continue receiving
		// items from the server after it's done sending.  Returns errors
		// encountered while closing, or if Close is called after Cancel.  Like
		// Send, blocks if there is no buffer space available.
		Close() error
	}
}

// ArithStreamingAddCall represents the call returned from Arith.StreamingAdd.
type ArithStreamingAddCall interface {
	ArithStreamingAddClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (total int32, err error)
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implArithStreamingAddCall struct {
	__ipc.Call
	valRecv int32
	errRecv error
}

func (c *implArithStreamingAddCall) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implArithStreamingAddCallRecv{c}
}

type implArithStreamingAddCallRecv struct {
	c *implArithStreamingAddCall
}

func (c implArithStreamingAddCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implArithStreamingAddCallRecv) Value() int32 {
	return c.c.valRecv
}
func (c implArithStreamingAddCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implArithStreamingAddCall) SendStream() interface {
	Send(item int32) error
	Close() error
} {
	return implArithStreamingAddCallSend{c}
}

type implArithStreamingAddCallSend struct {
	c *implArithStreamingAddCall
}

func (c implArithStreamingAddCallSend) Send(item int32) error {
	return c.c.Send(item)
}
func (c implArithStreamingAddCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implArithStreamingAddCall) Finish() (o0 int32, err error) {
	if ierr := c.Call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// ArithServerMethods is the interface a server writer
// implements for Arith.
//
// Arith is an example of an interface definition for an arithmetic service.
// Things to note:
//   * There must be at least 1 out-arg, and the last out-arg must be error.
type ArithServerMethods interface {
	// Add is a typical method with multiple input and output arguments.
	Add(ctx __ipc.ServerContext, a int32, b int32) (int32, error)
	// DivMod shows that runs of args with the same type can use the short form,
	// just like Go.
	DivMod(ctx __ipc.ServerContext, a int32, b int32) (quot int32, rem int32, err error)
	// Sub shows that you can use data types defined in other packages.
	Sub(ctx __ipc.ServerContext, args base.Args) (int32, error)
	// Mul tries another data type defined in another package.
	Mul(ctx __ipc.ServerContext, nested base.NestedArgs) (int32, error)
	// GenError shows that it's fine to have no in args, and no out args other
	// than "error".  In addition GenError shows the usage of tags.  Tags are a
	// sequence of constants.  There's no requirement on uniqueness of types or
	// values, and regular const expressions may also be used.
	GenError(__ipc.ServerContext) error
	// Count shows using only an int32 out-stream type, with no in-stream type.
	Count(ctx ArithCountContext, start int32) error
	// StreamingAdd shows a bidirectional stream.
	StreamingAdd(ArithStreamingAddContext) (total int32, err error)
	// QuoteAny shows the any built-in type, representing a value of any type.
	QuoteAny(ctx __ipc.ServerContext, a __vdlutil.Any) (__vdlutil.Any, error)
}

// ArithServerStubMethods is the server interface containing
// Arith methods, as expected by ipc.Server.
// The only difference between this interface and ArithServerMethods
// is the streaming methods.
type ArithServerStubMethods interface {
	// Add is a typical method with multiple input and output arguments.
	Add(ctx __ipc.ServerContext, a int32, b int32) (int32, error)
	// DivMod shows that runs of args with the same type can use the short form,
	// just like Go.
	DivMod(ctx __ipc.ServerContext, a int32, b int32) (quot int32, rem int32, err error)
	// Sub shows that you can use data types defined in other packages.
	Sub(ctx __ipc.ServerContext, args base.Args) (int32, error)
	// Mul tries another data type defined in another package.
	Mul(ctx __ipc.ServerContext, nested base.NestedArgs) (int32, error)
	// GenError shows that it's fine to have no in args, and no out args other
	// than "error".  In addition GenError shows the usage of tags.  Tags are a
	// sequence of constants.  There's no requirement on uniqueness of types or
	// values, and regular const expressions may also be used.
	GenError(__ipc.ServerContext) error
	// Count shows using only an int32 out-stream type, with no in-stream type.
	Count(ctx *ArithCountContextStub, start int32) error
	// StreamingAdd shows a bidirectional stream.
	StreamingAdd(*ArithStreamingAddContextStub) (total int32, err error)
	// QuoteAny shows the any built-in type, representing a value of any type.
	QuoteAny(ctx __ipc.ServerContext, a __vdlutil.Any) (__vdlutil.Any, error)
}

// ArithServerStub adds universal methods to ArithServerStubMethods.
type ArithServerStub interface {
	ArithServerStubMethods
	// Describe the Arith interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// ArithServer returns a server stub for Arith.
// It converts an implementation of ArithServerMethods into
// an object that may be used by ipc.Server.
func ArithServer(impl ArithServerMethods) ArithServerStub {
	stub := implArithServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implArithServerStub struct {
	impl ArithServerMethods
	gs   *__ipc.GlobState
}

func (s implArithServerStub) Add(ctx __ipc.ServerContext, i0 int32, i1 int32) (int32, error) {
	return s.impl.Add(ctx, i0, i1)
}

func (s implArithServerStub) DivMod(ctx __ipc.ServerContext, i0 int32, i1 int32) (int32, int32, error) {
	return s.impl.DivMod(ctx, i0, i1)
}

func (s implArithServerStub) Sub(ctx __ipc.ServerContext, i0 base.Args) (int32, error) {
	return s.impl.Sub(ctx, i0)
}

func (s implArithServerStub) Mul(ctx __ipc.ServerContext, i0 base.NestedArgs) (int32, error) {
	return s.impl.Mul(ctx, i0)
}

func (s implArithServerStub) GenError(ctx __ipc.ServerContext) error {
	return s.impl.GenError(ctx)
}

func (s implArithServerStub) Count(ctx *ArithCountContextStub, i0 int32) error {
	return s.impl.Count(ctx, i0)
}

func (s implArithServerStub) StreamingAdd(ctx *ArithStreamingAddContextStub) (int32, error) {
	return s.impl.StreamingAdd(ctx)
}

func (s implArithServerStub) QuoteAny(ctx __ipc.ServerContext, i0 __vdlutil.Any) (__vdlutil.Any, error) {
	return s.impl.QuoteAny(ctx, i0)
}

func (s implArithServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implArithServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{ArithDesc}
}

// ArithDesc describes the Arith interface.
var ArithDesc __ipc.InterfaceDesc = descArith

// descArith hides the desc to keep godoc clean.
var descArith = __ipc.InterfaceDesc{
	Name:    "Arith",
	PkgPath: "veyron.io/veyron/veyron2/vdl/testdata/arith",
	Doc:     "// Arith is an example of an interface definition for an arithmetic service.\n// Things to note:\n//   * There must be at least 1 out-arg, and the last out-arg must be error.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Add",
			Doc:  "// Add is a typical method with multiple input and output arguments.",
			InArgs: []__ipc.ArgDesc{
				{"a", ``}, // int32
				{"b", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // int32
				{"", ``}, // error
			},
		},
		{
			Name: "DivMod",
			Doc:  "// DivMod shows that runs of args with the same type can use the short form,\n// just like Go.",
			InArgs: []__ipc.ArgDesc{
				{"a", ``}, // int32
				{"b", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"quot", ``}, // int32
				{"rem", ``},  // int32
				{"err", ``},  // error
			},
		},
		{
			Name: "Sub",
			Doc:  "// Sub shows that you can use data types defined in other packages.",
			InArgs: []__ipc.ArgDesc{
				{"args", ``}, // base.Args
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // int32
				{"", ``}, // error
			},
		},
		{
			Name: "Mul",
			Doc:  "// Mul tries another data type defined in another package.",
			InArgs: []__ipc.ArgDesc{
				{"nested", ``}, // base.NestedArgs
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // int32
				{"", ``}, // error
			},
		},
		{
			Name: "GenError",
			Doc:  "// GenError shows that it's fine to have no in args, and no out args other\n// than \"error\".  In addition GenError shows the usage of tags.  Tags are a\n// sequence of constants.  There's no requirement on uniqueness of types or\n// values, and regular const expressions may also be used.",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{"foo", "barz", "hello", int32(129), uint64(36)},
		},
		{
			Name: "Count",
			Doc:  "// Count shows using only an int32 out-stream type, with no in-stream type.",
			InArgs: []__ipc.ArgDesc{
				{"start", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
		{
			Name: "StreamingAdd",
			Doc:  "// StreamingAdd shows a bidirectional stream.",
			OutArgs: []__ipc.ArgDesc{
				{"total", ``}, // int32
				{"err", ``},   // error
			},
		},
		{
			Name: "QuoteAny",
			Doc:  "// QuoteAny shows the any built-in type, representing a value of any type.",
			InArgs: []__ipc.ArgDesc{
				{"a", ``}, // __vdlutil.Any
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // __vdlutil.Any
				{"", ``}, // error
			},
		},
	},
}

func (s implArithServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Add"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "a", Type: 36},
			{Name: "b", Type: 36},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 36},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Count"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "start", Type: 36},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 36,
	}
	result.Methods["DivMod"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "a", Type: 36},
			{Name: "b", Type: 36},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "quot", Type: 36},
			{Name: "rem", Type: 36},
			{Name: "err", Type: 65},
		},
	}
	result.Methods["GenError"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Mul"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "nested", Type: 67},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 36},
			{Name: "", Type: 65},
		},
	}
	result.Methods["QuoteAny"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "a", Type: 68},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 68},
			{Name: "", Type: 65},
		},
	}
	result.Methods["StreamingAdd"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "total", Type: 36},
			{Name: "err", Type: 65},
		},
		InStream:  36,
		OutStream: 36,
	}
	result.Methods["Sub"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "args", Type: 66},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 36},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, __wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x24, Name: "A"},
				__wiretype.FieldType{Type: 0x24, Name: "B"},
			},
			"veyron.io/veyron/veyron2/vdl/testdata/base.Args", []string(nil)},
		__wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x42, Name: "Args"},
			},
			"veyron.io/veyron/veyron2/vdl/testdata/base.NestedArgs", []string(nil)},
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}}

	return result, nil
}

// ArithCountServerStream is the server stream for Arith.Count.
type ArithCountServerStream interface {
	// SendStream returns the send side of the Arith.Count server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item int32) error
	}
}

// ArithCountContext represents the context passed to Arith.Count.
type ArithCountContext interface {
	__ipc.ServerContext
	ArithCountServerStream
}

// ArithCountContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements ArithCountContext.
type ArithCountContextStub struct {
	__ipc.ServerCall
}

// Init initializes ArithCountContextStub from ipc.ServerCall.
func (s *ArithCountContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the Arith.Count server stream.
func (s *ArithCountContextStub) SendStream() interface {
	Send(item int32) error
} {
	return implArithCountContextSend{s}
}

type implArithCountContextSend struct {
	s *ArithCountContextStub
}

func (s implArithCountContextSend) Send(item int32) error {
	return s.s.Send(item)
}

// ArithStreamingAddServerStream is the server stream for Arith.StreamingAdd.
type ArithStreamingAddServerStream interface {
	// RecvStream returns the receiver side of the Arith.StreamingAdd server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() int32
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Arith.StreamingAdd server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item int32) error
	}
}

// ArithStreamingAddContext represents the context passed to Arith.StreamingAdd.
type ArithStreamingAddContext interface {
	__ipc.ServerContext
	ArithStreamingAddServerStream
}

// ArithStreamingAddContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements ArithStreamingAddContext.
type ArithStreamingAddContextStub struct {
	__ipc.ServerCall
	valRecv int32
	errRecv error
}

// Init initializes ArithStreamingAddContextStub from ipc.ServerCall.
func (s *ArithStreamingAddContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// RecvStream returns the receiver side of the Arith.StreamingAdd server stream.
func (s *ArithStreamingAddContextStub) RecvStream() interface {
	Advance() bool
	Value() int32
	Err() error
} {
	return implArithStreamingAddContextRecv{s}
}

type implArithStreamingAddContextRecv struct {
	s *ArithStreamingAddContextStub
}

func (s implArithStreamingAddContextRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implArithStreamingAddContextRecv) Value() int32 {
	return s.s.valRecv
}
func (s implArithStreamingAddContextRecv) Err() error {
	if s.s.errRecv == __io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Arith.StreamingAdd server stream.
func (s *ArithStreamingAddContextStub) SendStream() interface {
	Send(item int32) error
} {
	return implArithStreamingAddContextSend{s}
}

type implArithStreamingAddContextSend struct {
	s *ArithStreamingAddContextStub
}

func (s implArithStreamingAddContextSend) Send(item int32) error {
	return s.s.Send(item)
}

// CalculatorClientMethods is the client interface
// containing Calculator methods.
type CalculatorClientMethods interface {
	// Arith is an example of an interface definition for an arithmetic service.
	// Things to note:
	//   * There must be at least 1 out-arg, and the last out-arg must be error.
	ArithClientMethods
	// AdvancedMath is an interface for more advanced math than arith.  It embeds
	// interfaces defined both in the same file and in an external package; and in
	// turn it is embedded by arith.Calculator (which is in the same package but
	// different file) to verify that embedding works in all these scenarios.
	AdvancedMathClientMethods
	On(__context.T, ...__ipc.CallOpt) error  // On turns the calculator on.
	Off(__context.T, ...__ipc.CallOpt) error // Off turns the calculator off.
}

// CalculatorClientStub adds universal methods to CalculatorClientMethods.
type CalculatorClientStub interface {
	CalculatorClientMethods
	__ipc.UniversalServiceMethods
}

// CalculatorClient returns a client stub for Calculator.
func CalculatorClient(name string, opts ...__ipc.BindOpt) CalculatorClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implCalculatorClientStub{name, client, ArithClient(name, client), AdvancedMathClient(name, client)}
}

type implCalculatorClientStub struct {
	name   string
	client __ipc.Client

	ArithClientStub
	AdvancedMathClientStub
}

func (c implCalculatorClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implCalculatorClientStub) On(ctx __context.T, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "On", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implCalculatorClientStub) Off(ctx __context.T, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Off", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implCalculatorClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// CalculatorServerMethods is the interface a server writer
// implements for Calculator.
type CalculatorServerMethods interface {
	// Arith is an example of an interface definition for an arithmetic service.
	// Things to note:
	//   * There must be at least 1 out-arg, and the last out-arg must be error.
	ArithServerMethods
	// AdvancedMath is an interface for more advanced math than arith.  It embeds
	// interfaces defined both in the same file and in an external package; and in
	// turn it is embedded by arith.Calculator (which is in the same package but
	// different file) to verify that embedding works in all these scenarios.
	AdvancedMathServerMethods
	On(__ipc.ServerContext) error  // On turns the calculator on.
	Off(__ipc.ServerContext) error // Off turns the calculator off.
}

// CalculatorServerStubMethods is the server interface containing
// Calculator methods, as expected by ipc.Server.
// The only difference between this interface and CalculatorServerMethods
// is the streaming methods.
type CalculatorServerStubMethods interface {
	// Arith is an example of an interface definition for an arithmetic service.
	// Things to note:
	//   * There must be at least 1 out-arg, and the last out-arg must be error.
	ArithServerStubMethods
	// AdvancedMath is an interface for more advanced math than arith.  It embeds
	// interfaces defined both in the same file and in an external package; and in
	// turn it is embedded by arith.Calculator (which is in the same package but
	// different file) to verify that embedding works in all these scenarios.
	AdvancedMathServerStubMethods
	On(__ipc.ServerContext) error  // On turns the calculator on.
	Off(__ipc.ServerContext) error // Off turns the calculator off.
}

// CalculatorServerStub adds universal methods to CalculatorServerStubMethods.
type CalculatorServerStub interface {
	CalculatorServerStubMethods
	// Describe the Calculator interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// CalculatorServer returns a server stub for Calculator.
// It converts an implementation of CalculatorServerMethods into
// an object that may be used by ipc.Server.
func CalculatorServer(impl CalculatorServerMethods) CalculatorServerStub {
	stub := implCalculatorServerStub{
		impl:                   impl,
		ArithServerStub:        ArithServer(impl),
		AdvancedMathServerStub: AdvancedMathServer(impl),
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implCalculatorServerStub struct {
	impl CalculatorServerMethods
	ArithServerStub
	AdvancedMathServerStub
	gs *__ipc.GlobState
}

func (s implCalculatorServerStub) On(ctx __ipc.ServerContext) error {
	return s.impl.On(ctx)
}

func (s implCalculatorServerStub) Off(ctx __ipc.ServerContext) error {
	return s.impl.Off(ctx)
}

func (s implCalculatorServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implCalculatorServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{CalculatorDesc, ArithDesc, AdvancedMathDesc, TrigonometryDesc, exp.ExpDesc}
}

// CalculatorDesc describes the Calculator interface.
var CalculatorDesc __ipc.InterfaceDesc = descCalculator

// descCalculator hides the desc to keep godoc clean.
var descCalculator = __ipc.InterfaceDesc{
	Name:    "Calculator",
	PkgPath: "veyron.io/veyron/veyron2/vdl/testdata/arith",
	Embeds: []__ipc.EmbedDesc{
		{"Arith", "veyron.io/veyron/veyron2/vdl/testdata/arith", "// Arith is an example of an interface definition for an arithmetic service.\n// Things to note:\n//   * There must be at least 1 out-arg, and the last out-arg must be error."},
		{"AdvancedMath", "veyron.io/veyron/veyron2/vdl/testdata/arith", "// AdvancedMath is an interface for more advanced math than arith.  It embeds\n// interfaces defined both in the same file and in an external package; and in\n// turn it is embedded by arith.Calculator (which is in the same package but\n// different file) to verify that embedding works in all these scenarios."},
	},
	Methods: []__ipc.MethodDesc{
		{
			Name: "On",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
		{
			Name: "Off",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{"offtag"},
		},
	},
}

func (s implCalculatorServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Off"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["On"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}
	var ss __ipc.ServiceSignature
	var firstAdded int
	ss, _ = s.ArithServerStub.Signature(ctx)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= __wiretype.TypeIDFirst {
				v.InArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= __wiretype.TypeIDFirst {
				v.OutArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= __wiretype.TypeIDFirst {
			v.InStream += __wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= __wiretype.TypeIDFirst {
			v.OutStream += __wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case __wiretype.SliceType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.ArrayType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.MapType:
			if wt.Key >= __wiretype.TypeIDFirst {
				wt.Key += __wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= __wiretype.TypeIDFirst {
					wt.Fields[i].Type += __wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}
	ss, _ = s.AdvancedMathServerStub.Signature(ctx)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= __wiretype.TypeIDFirst {
				v.InArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= __wiretype.TypeIDFirst {
				v.OutArgs[i].Type += __wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= __wiretype.TypeIDFirst {
			v.InStream += __wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= __wiretype.TypeIDFirst {
			v.OutStream += __wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case __wiretype.SliceType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.ArrayType:
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.MapType:
			if wt.Key >= __wiretype.TypeIDFirst {
				wt.Key += __wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= __wiretype.TypeIDFirst {
				wt.Elem += __wiretype.TypeID(firstAdded)
			}
			d = wt
		case __wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= __wiretype.TypeIDFirst {
					wt.Fields[i].Type += __wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}
