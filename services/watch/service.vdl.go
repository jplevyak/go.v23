// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// Package watch defines an API for watching updates that match a query.
//
// API Overview
//
// Watcher service allows a client to watch for updates that match a
// query.  For each watched query, the client will receive a reliable
// stream of watch events without re-ordering.
//
// The watching is done by starting a streaming RPC. The argument to
// the RPC contains the query. The result stream consists of a
// never-ending sequence of Change messages (until the call fails or
// is cancelled).
//
// Root Entity
//
// The Object name that receives the Watch RPC is called the root
// entity.  The root entity is the parent of all entities that the
// client cares about.  Therefore, the query is confined to children
// of the root entity, and the names in the Change messages are all
// relative to the root entity.
//
// Watch Request
//
// When a client makes a watch request, it can indicate whether it
// wants to receive the initial states of the entities that match the
// query, just new changes to the entities, or resume watching from a
// particular point in a previous watch stream.  On receiving a watch
// request, the server sends one or more messages to the client. The
// first message informs the client that the server has registered the
// client's request; the instant of time when the client receives the
// event is referred to as the client's "watch point" for that query.
//
// Atomic Delivery
//
// The response stream consists of a sequence of Change messages. Each
// Change message contains an optional continued bit
// (default=false). A sub-sequence of Change messages with
// continued=true followed by a Change message with continued=false
// forms an "atomic group". Systems that support multi-entity atomic
// updates may guarantee that all changes resulting from a single
// atomic update are delivered in the same "atomic group". It is up to
// the documentation of a particular system that implements the Watch
// API to document whether or not it supports such grouping. We expect
// that most callers will ignore the notion of atomic delivery and the
// continued bit, i.e., they will just process each Change message as
// it is received.
//
// Initial State
//
// The first atomic group delivered by a watch call is special. It is
// delivered as soon as possible and contains the initial state of the
// entities being watched.  The client should consider itself caught up
// after processing this first atomic group.  The messages in this first
// atomic group depend on the value of ResumeMarker.
//
//   (1) ResumeMarker is "" or not specified: For every entity P that
//       matches the query and exists, there will be at least one message
//       delivered with entity == P and the last such message will contain
//       the current state of P.  For every entity Q (including the entity
//       itself) that matches the query but does not exist, either no
//       message will be delivered, or the last message for Q will have
//       state == DOES_NOT_EXIST. At least one message for entity="" will
//       be delivered.
//
//   (2) ResumeMarker == "now": there will be exactly one message with
//       entity = "" and state INITIAL_STATE_SKIPPED.  The client cannot
//       assume whether or not the entity exists after receiving this
//       message.
//
//   (3) ResumeMarker has a value R from a preceding watch call on this
//       entity: The same messages as described in (1) will be delivered
//       to the client except that any information implied by messages
//       received on the preceding call up to and including R may not be
//       delivered. The expectation is that the client will start with
//       state it had built up from the preceding watch call, apply the
//       changes received from this call and build an up-to-date view of
//       the entities without having to fetch a potentially large amount
//       of information that has not changed.  Note that some information
//       that had already been delivered by the preceding call might be
//       delivered again.
//
// Ordering and Reliability
//
// The Change messages that apply to a particular element of the
// entity will be delivered eventually in order without loss for the
// duration of the RPC. Note however that if multiple Changes apply to
// the same element, the implementation is free to suppress them and
// deliver just the last one.  The underlying system must provide the
// guarantee that any relevant update received for an entity E after a
// client's watch point for E MUST be delivered to that client.
//
// These tight guarantees allow for the following simplifications in
// the client:
//
//   (1) The client does not need to have a separate polling loop to
//       make up for missed updates.
//
//   (2) The client does not need to manage timestamps/versions
//       manually; the last update delivered corresponds to the
//       eventual state of the entity.
package watch

import (
	"v.io/core/veyron2/services/security/access"

	"v.io/core/veyron2/services/watch/types"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "v.io/core/veyron2"
	__context "v.io/core/veyron2/context"
	__ipc "v.io/core/veyron2/ipc"
	__vdlutil "v.io/core/veyron2/vdl/vdlutil"
)

// GlobWatcherClientMethods is the client interface
// containing GlobWatcher methods.
//
// GlobWatcher allows a client to receive updates for changes to objects
// that match a pattern.  See the package comments for details.
type GlobWatcherClientMethods interface {
	// WatchGlob returns a stream of changes that match a pattern.
	WatchGlob(ctx *__context.T, req types.GlobRequest, opts ...__ipc.CallOpt) (GlobWatcherWatchGlobCall, error)
}

// GlobWatcherClientStub adds universal methods to GlobWatcherClientMethods.
type GlobWatcherClientStub interface {
	GlobWatcherClientMethods
	__ipc.UniversalServiceMethods
}

// GlobWatcherClient returns a client stub for GlobWatcher.
func GlobWatcherClient(name string, opts ...__ipc.BindOpt) GlobWatcherClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implGlobWatcherClientStub{name, client}
}

type implGlobWatcherClientStub struct {
	name   string
	client __ipc.Client
}

func (c implGlobWatcherClientStub) c(ctx *__context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.GetClient(ctx)
}

func (c implGlobWatcherClientStub) WatchGlob(ctx *__context.T, i0 types.GlobRequest, opts ...__ipc.CallOpt) (ocall GlobWatcherWatchGlobCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "WatchGlob", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implGlobWatcherWatchGlobCall{Call: call}
	return
}

// GlobWatcherWatchGlobClientStream is the client stream for GlobWatcher.WatchGlob.
type GlobWatcherWatchGlobClientStream interface {
	// RecvStream returns the receiver side of the GlobWatcher.WatchGlob client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() types.Change
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// GlobWatcherWatchGlobCall represents the call returned from GlobWatcher.WatchGlob.
type GlobWatcherWatchGlobCall interface {
	GlobWatcherWatchGlobClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implGlobWatcherWatchGlobCall struct {
	__ipc.Call
	valRecv types.Change
	errRecv error
}

func (c *implGlobWatcherWatchGlobCall) RecvStream() interface {
	Advance() bool
	Value() types.Change
	Err() error
} {
	return implGlobWatcherWatchGlobCallRecv{c}
}

type implGlobWatcherWatchGlobCallRecv struct {
	c *implGlobWatcherWatchGlobCall
}

func (c implGlobWatcherWatchGlobCallRecv) Advance() bool {
	c.c.valRecv = types.Change{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implGlobWatcherWatchGlobCallRecv) Value() types.Change {
	return c.c.valRecv
}
func (c implGlobWatcherWatchGlobCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implGlobWatcherWatchGlobCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// GlobWatcherServerMethods is the interface a server writer
// implements for GlobWatcher.
//
// GlobWatcher allows a client to receive updates for changes to objects
// that match a pattern.  See the package comments for details.
type GlobWatcherServerMethods interface {
	// WatchGlob returns a stream of changes that match a pattern.
	WatchGlob(ctx GlobWatcherWatchGlobContext, req types.GlobRequest) error
}

// GlobWatcherServerStubMethods is the server interface containing
// GlobWatcher methods, as expected by ipc.Server.
// The only difference between this interface and GlobWatcherServerMethods
// is the streaming methods.
type GlobWatcherServerStubMethods interface {
	// WatchGlob returns a stream of changes that match a pattern.
	WatchGlob(ctx *GlobWatcherWatchGlobContextStub, req types.GlobRequest) error
}

// GlobWatcherServerStub adds universal methods to GlobWatcherServerStubMethods.
type GlobWatcherServerStub interface {
	GlobWatcherServerStubMethods
	// Describe the GlobWatcher interfaces.
	Describe__() []__ipc.InterfaceDesc
}

// GlobWatcherServer returns a server stub for GlobWatcher.
// It converts an implementation of GlobWatcherServerMethods into
// an object that may be used by ipc.Server.
func GlobWatcherServer(impl GlobWatcherServerMethods) GlobWatcherServerStub {
	stub := implGlobWatcherServerStub{
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

type implGlobWatcherServerStub struct {
	impl GlobWatcherServerMethods
	gs   *__ipc.GlobState
}

func (s implGlobWatcherServerStub) WatchGlob(ctx *GlobWatcherWatchGlobContextStub, i0 types.GlobRequest) error {
	return s.impl.WatchGlob(ctx, i0)
}

func (s implGlobWatcherServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implGlobWatcherServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{GlobWatcherDesc}
}

// GlobWatcherDesc describes the GlobWatcher interface.
var GlobWatcherDesc __ipc.InterfaceDesc = descGlobWatcher

// descGlobWatcher hides the desc to keep godoc clean.
var descGlobWatcher = __ipc.InterfaceDesc{
	Name:    "GlobWatcher",
	PkgPath: "v.io/core/veyron2/services/watch",
	Doc:     "// GlobWatcher allows a client to receive updates for changes to objects\n// that match a pattern.  See the package comments for details.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "WatchGlob",
			Doc:  "// WatchGlob returns a stream of changes that match a pattern.",
			InArgs: []__ipc.ArgDesc{
				{"req", ``}, // types.GlobRequest
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{access.Tag("Resolve")},
		},
	},
}

// GlobWatcherWatchGlobServerStream is the server stream for GlobWatcher.WatchGlob.
type GlobWatcherWatchGlobServerStream interface {
	// SendStream returns the send side of the GlobWatcher.WatchGlob server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item types.Change) error
	}
}

// GlobWatcherWatchGlobContext represents the context passed to GlobWatcher.WatchGlob.
type GlobWatcherWatchGlobContext interface {
	__ipc.ServerContext
	GlobWatcherWatchGlobServerStream
}

// GlobWatcherWatchGlobContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements GlobWatcherWatchGlobContext.
type GlobWatcherWatchGlobContextStub struct {
	__ipc.ServerCall
}

// Init initializes GlobWatcherWatchGlobContextStub from ipc.ServerCall.
func (s *GlobWatcherWatchGlobContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the GlobWatcher.WatchGlob server stream.
func (s *GlobWatcherWatchGlobContextStub) SendStream() interface {
	Send(item types.Change) error
} {
	return implGlobWatcherWatchGlobContextSend{s}
}

type implGlobWatcherWatchGlobContextSend struct {
	s *GlobWatcherWatchGlobContextStub
}

func (s implGlobWatcherWatchGlobContextSend) Send(item types.Change) error {
	return s.s.Send(item)
}
