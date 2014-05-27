// This file was auto-generated by the veyron vdl tool.
// Source: application.vdl

// Package application defines the interface for serving application
// metadata.
//
// OVERVIEW: Applications are expected to be organized using veyron's
// hierarchical namespace. The nodes of the hierarchy are expected to
// implement: 1) the MountTable interface, to enable application
// discovery and hierarchy organization and 2) the Application
// interface, to enable application (metadata) management.
package application

import (
	"veyron2/security"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

// Envelope is a collection of metadata that describes an application.
type Envelope struct {
	// Arguments is an array of command-line arguments to be used when
	// executing the binary.
	Args []string
	// Binary is a veyron name that identifies the application binary.
	Binary string
	// Environment is a map that stores the environment variable values
	// to be used when executing the binary.
	Env []string
}

// Repository provides access to application envelopes. An
// application envelope is identified by an application name and an
// application version, which are specified through the veyron name,
// and a profile name, which is specified using a method argument.
//
// Example:
// /apps/search/v1.Match([]string{"base", "media"})
//   returns an application envelope that can be used for downloading
//   and executing the "search" application, version "v1", runnable
//   on either the "base" or "media" profile.
//
// Further, we envision that there will be special "latest" and
// "release" versions that will be symbolic links whose mapping is
// maintained by a mount table.
// Repository is the interface the client binds and uses.
// Repository_InternalNoTagGetter is the interface without the TagGetter
// and UnresolveStep methods (both framework-added, rathern than user-defined),
// to enable embedding without method collisions.  Not to be used directly by
// clients.
type Repository_InternalNoTagGetter interface {

	// Match checks if any of the given profiles contains an application
	// envelope for the given application version (specified through the
	// veyron name suffix) and if so, returns this envelope. If multiple
	// profile matches are possible, the method returns the first
	// matching profile, respecting the order of the input argument.
	Match(Profiles []string, opts ..._gen_ipc.ClientCallOpt) (reply Envelope, err error)
}
type Repository interface {
	_gen_vdl.TagGetter
	// UnresolveStep returns the names for the remote service, rooted at the
	// service's immediate namespace ancestor.
	UnresolveStep(opts ..._gen_ipc.ClientCallOpt) ([]string, error)
	Repository_InternalNoTagGetter
}

// RepositoryService is the interface the server implements.
type RepositoryService interface {

	// Match checks if any of the given profiles contains an application
	// envelope for the given application version (specified through the
	// veyron name suffix) and if so, returns this envelope. If multiple
	// profile matches are possible, the method returns the first
	// matching profile, respecting the order of the input argument.
	Match(context _gen_ipc.Context, Profiles []string) (reply Envelope, err error)
}

// BindRepository returns the client stub implementing the Repository
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindRepository(name string, opts ..._gen_ipc.BindOpt) (Repository, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdl.ErrTooManyOptionsToBind
	}
	stub := &clientStubRepository{client: client, name: name}

	return stub, nil
}

// NewServerRepository creates a new server stub.
//
// It takes a regular server implementing the RepositoryService
// interface, and returns a new server stub.
func NewServerRepository(server RepositoryService) interface{} {
	return &ServerStubRepository{
		service: server,
	}
}

// clientStubRepository implements Repository.
type clientStubRepository struct {
	client _gen_ipc.Client
	name   string
}

func (c *clientStubRepository) GetMethodTags(method string) []interface{} {
	return GetRepositoryMethodTags(method)
}

func (__gen_c *clientStubRepository) Match(Profiles []string, opts ..._gen_ipc.ClientCallOpt) (reply Envelope, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Match", []interface{}{Profiles}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *clientStubRepository) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = c.client.StartCall(c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubRepository wraps a server that implements
// RepositoryService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubRepository struct {
	service RepositoryService
}

func (s *ServerStubRepository) GetMethodTags(method string) []interface{} {
	return GetRepositoryMethodTags(method)
}

func (s *ServerStubRepository) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Match"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Profiles", Type: 61},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
			{Name: "", Type: 66},
		},
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Args"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Binary"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "Env"},
			},
			"veyron2/services/mgmt/application.Envelope", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (s *ServerStubRepository) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubRepository) Match(call _gen_ipc.ServerCall, Profiles []string) (reply Envelope, err error) {
	reply, err = __gen_s.service.Match(call, Profiles)
	return
}

func GetRepositoryMethodTags(method string) []interface{} {
	switch method {
	case "Match":
		return []interface{}{security.Label(1)}
	default:
		return nil
	}
}
