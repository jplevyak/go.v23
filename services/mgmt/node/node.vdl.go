// This file was auto-generated by the veyron vdl tool.
// Source: node.vdl

// Package node supports managing a node and applications running on
// the node.
package node

import (
	"veyron2/services/mgmt/binary"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// Description enumerates the profiles that a Node supports.
type Description struct {
	// Profiles is a set of names of supported profiles.	Each name can
	// either be an object name that resolves to a Profile, or can be the
	// profile's label, e.g.:
	//   "profiles/google/cluster/diskfull"
	//   "linux-media"
	//
	// Profiles for nodes can be provided by hand, but they can also be
	// automatically derived by examining the node.
	Profiles map[string]struct{}
}

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Application can be used to manage applications on a device. The
// idea is that this interace will be invoked using an object name that
// identifies the application and its installations and instances
// where applicable.
//
// In particular, the interface methods can be divided into three
// groups based on their intended receiver:
//
// 1) Method receiver is an application:
// -- Install()
//
// 2) Method receiver is an application installation:
// -- Start()
// -- Uninstall()
// -- Update()
//
// 3) Method receiver is application installation instance:
// -- Refresh()
// -- Restart()
// -- Resume()
// -- Stop()
// -- Suspend()
//
// For groups 2) and 3), the suffix that specifies the receiver can
// optionally omit the installation and/or instance, in which case the
// operation applies to all installations and/or instances in the
// scope of the suffix.
//
// Examples:
// # Install Google Maps on the node.
// device/apps.Install("/google.com/appstore/maps") --> "google maps/0"
//
// # Start an instance of the previously installed maps application installation.
// device/apps/google maps/0.Start() --> { "0" }
//
// # Start a second instance of the previously installed maps application installation.
// device/apps/google maps/0.Start() --> { "1" }
//
// # Stop the first instance previously started.
// device/apps/google maps/0/0.Stop()
//
// # Install a second Google Maps installation.
// device/apps.Install("/google.com/appstore/maps") --> "google maps/1"
//
// # Start an instance for all maps application installations.
// device/apps/google maps.Start() --> {"0/2", "1/0"}
//
// # Refresh the state of all instances of all maps application installations.
// device/apps/google maps.Refresh()
//
// # Refresh the state of all instances of the maps application installation
// identified by the given suffix.
// device/apps/google maps/0.Refresh()
//
// # Refresh the state of the maps application installation instance identified by
// the given suffix.
// device/apps/google maps/0/2.Refresh()
//
// # Update the second maps installation to the latest version available.
// device/apps/google maps/1.Update()
//
// # Update the first maps installation to a specific version.
// device/apps/google maps/0.UpdateTo("/google.com/appstore/beta/maps")
//
// Further, the following methods complement one another:
// -- Install() and Uninstall()
// -- Start() and Stop()
// -- Suspend() and Resume()
//
// Finally, an application installation instance can be in one of
// three abstract states: 1) "does not exist", 2) "running", or 3)
// "suspended". The interface methods transition between these
// abstract states using the following state machine:
//
// apply(Start(), "does not exists") = "running"
// apply(Refresh(), "running") = "running"
// apply(Refresh(), "suspended") = "suspended"
// apply(Restart(), "running") = "running"
// apply(Restart(), "suspended") = "running"
// apply(Resume(), "suspended") = "running"
// apply(Resume(), "running") = "running"
// apply(Stop(), "running") = "does not exist"
// apply(Stop(), "suspended") = "does not exist"
// apply(Suspend(), "running") = "suspended"
// apply(Suspend(), "suspended") = "suspended"
//
// In other words, invoking any method using an existing application
// installation instance as a receiver is well-defined.
// Application is the interface the client binds and uses.
// Application_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Application_ExcludingUniversal interface {
	// Install installs the application identified by the argument and
	// returns an object name suffix that identifies the new installation.
	//
	// The argument should be an object name. The service it identifies must
	// implement repository.Application, and is expected to return either
	// the requested version (if the object name encodes a specific
	// version), or otherwise the latest available version, as appropriate.
	//
	// The returned suffix, when appended to the name used to reach the
	// receiver for Install, can be used to control the installation object.
	// The suffix will contain the title of the application as a prefix,
	// which can then be used to control all the installations of the given
	// application.
	Install(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (reply string, err error)
	// Refresh refreshes the state of application installation(s)
	// instance(s).
	Refresh(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Restart restarts execution of application installation(s)
	// instance(s).
	Restart(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Resume resumes execution of application installation(s)
	// instance(s).
	Resume(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Revert reverts application installation(s) to the most recent
	// previous installation.
	Revert(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Start starts an instance of application installation(s) and
	// returns the object name(s) that identifies/identify the new
	// instance(s).
	Start(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error)
	// Stop attempts a clean shutdown of application installation(s)
	// instance(s). If the deadline is non-zero and the instance(s) in
	// questions are still running after the given deadline, shutdown of
	// the instance(s) is enforced.
	//
	// TODO(jsimsa): Switch deadline to time.Duration when built-in types
	// are implemented.
	Stop(ctx _gen_context.T, Deadline uint64, opts ..._gen_ipc.CallOpt) (err error)
	// Suspend suspends execution of application installation(s)
	// instance(s).
	Suspend(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Uninstall uninstalls application installation(s).
	Uninstall(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Update updates the application installation(s) from the object name
	// provided during Install.  If the new application envelope contains a
	// different application title, the update does not occur, and an error
	// is returned.
	Update(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// UpdateTo updates the application installation(s) to the application
	// specified by the object name argument.  If the new application
	// envelope contains a different application title, the update does not
	// occur, and an error is returned.
	UpdateTo(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (err error)
}
type Application interface {
	_gen_ipc.UniversalServiceMethods
	Application_ExcludingUniversal
}

// ApplicationService is the interface the server implements.
type ApplicationService interface {

	// Install installs the application identified by the argument and
	// returns an object name suffix that identifies the new installation.
	//
	// The argument should be an object name. The service it identifies must
	// implement repository.Application, and is expected to return either
	// the requested version (if the object name encodes a specific
	// version), or otherwise the latest available version, as appropriate.
	//
	// The returned suffix, when appended to the name used to reach the
	// receiver for Install, can be used to control the installation object.
	// The suffix will contain the title of the application as a prefix,
	// which can then be used to control all the installations of the given
	// application.
	Install(context _gen_ipc.ServerContext, Name string) (reply string, err error)
	// Refresh refreshes the state of application installation(s)
	// instance(s).
	Refresh(context _gen_ipc.ServerContext) (err error)
	// Restart restarts execution of application installation(s)
	// instance(s).
	Restart(context _gen_ipc.ServerContext) (err error)
	// Resume resumes execution of application installation(s)
	// instance(s).
	Resume(context _gen_ipc.ServerContext) (err error)
	// Revert reverts application installation(s) to the most recent
	// previous installation.
	Revert(context _gen_ipc.ServerContext) (err error)
	// Start starts an instance of application installation(s) and
	// returns the object name(s) that identifies/identify the new
	// instance(s).
	Start(context _gen_ipc.ServerContext) (reply []string, err error)
	// Stop attempts a clean shutdown of application installation(s)
	// instance(s). If the deadline is non-zero and the instance(s) in
	// questions are still running after the given deadline, shutdown of
	// the instance(s) is enforced.
	//
	// TODO(jsimsa): Switch deadline to time.Duration when built-in types
	// are implemented.
	Stop(context _gen_ipc.ServerContext, Deadline uint64) (err error)
	// Suspend suspends execution of application installation(s)
	// instance(s).
	Suspend(context _gen_ipc.ServerContext) (err error)
	// Uninstall uninstalls application installation(s).
	Uninstall(context _gen_ipc.ServerContext) (err error)
	// Update updates the application installation(s) from the object name
	// provided during Install.  If the new application envelope contains a
	// different application title, the update does not occur, and an error
	// is returned.
	Update(context _gen_ipc.ServerContext) (err error)
	// UpdateTo updates the application installation(s) to the application
	// specified by the object name argument.  If the new application
	// envelope contains a different application title, the update does not
	// occur, and an error is returned.
	UpdateTo(context _gen_ipc.ServerContext, Name string) (err error)
}

// BindApplication returns the client stub implementing the Application
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindApplication(name string, opts ..._gen_ipc.BindOpt) (Application, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubApplication{client: client, name: name}

	return stub, nil
}

// NewServerApplication creates a new server stub.
//
// It takes a regular server implementing the ApplicationService
// interface, and returns a new server stub.
func NewServerApplication(server ApplicationService) interface{} {
	return &ServerStubApplication{
		service: server,
	}
}

// clientStubApplication implements Application.
type clientStubApplication struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubApplication) Install(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (reply string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Install", []interface{}{Name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Refresh(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Refresh", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Restart(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Restart", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Resume(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Resume", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Revert(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Revert", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Start(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Start", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Stop(ctx _gen_context.T, Deadline uint64, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Stop", []interface{}{Deadline}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Suspend(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Suspend", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Uninstall(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Uninstall", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Update(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Update", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) UpdateTo(ctx _gen_context.T, Name string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UpdateTo", []interface{}{Name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubApplication) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubApplication wraps a server that implements
// ApplicationService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubApplication struct {
	service ApplicationService
}

func (__gen_s *ServerStubApplication) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Install":
		return []interface{}{}, nil
	case "Refresh":
		return []interface{}{}, nil
	case "Restart":
		return []interface{}{}, nil
	case "Resume":
		return []interface{}{}, nil
	case "Revert":
		return []interface{}{}, nil
	case "Start":
		return []interface{}{}, nil
	case "Stop":
		return []interface{}{}, nil
	case "Suspend":
		return []interface{}{}, nil
	case "Uninstall":
		return []interface{}{}, nil
	case "Update":
		return []interface{}{}, nil
	case "UpdateTo":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubApplication) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Install"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Name", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 3},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Refresh"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Restart"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Resume"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Revert"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Start"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 61},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Stop"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Deadline", Type: 53},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Suspend"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Uninstall"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Update"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["UpdateTo"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Name", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubApplication) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
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

func (__gen_s *ServerStubApplication) Install(call _gen_ipc.ServerCall, Name string) (reply string, err error) {
	reply, err = __gen_s.service.Install(call, Name)
	return
}

func (__gen_s *ServerStubApplication) Refresh(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Refresh(call)
	return
}

func (__gen_s *ServerStubApplication) Restart(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Restart(call)
	return
}

func (__gen_s *ServerStubApplication) Resume(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Resume(call)
	return
}

func (__gen_s *ServerStubApplication) Revert(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Revert(call)
	return
}

func (__gen_s *ServerStubApplication) Start(call _gen_ipc.ServerCall) (reply []string, err error) {
	reply, err = __gen_s.service.Start(call)
	return
}

func (__gen_s *ServerStubApplication) Stop(call _gen_ipc.ServerCall, Deadline uint64) (err error) {
	err = __gen_s.service.Stop(call, Deadline)
	return
}

func (__gen_s *ServerStubApplication) Suspend(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Suspend(call)
	return
}

func (__gen_s *ServerStubApplication) Uninstall(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Uninstall(call)
	return
}

func (__gen_s *ServerStubApplication) Update(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Update(call)
	return
}

func (__gen_s *ServerStubApplication) UpdateTo(call _gen_ipc.ServerCall, Name string) (err error) {
	err = __gen_s.service.UpdateTo(call, Name)
	return
}

// Node can be used to manage a node. The idea is that this interace
// will be invoked using an object name that identifies the node.
// Node is the interface the client binds and uses.
// Node_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Node_ExcludingUniversal interface {
	// Application can be used to manage applications on a device. The
	// idea is that this interace will be invoked using an object name that
	// identifies the application and its installations and instances
	// where applicable.
	//
	// In particular, the interface methods can be divided into three
	// groups based on their intended receiver:
	//
	// 1) Method receiver is an application:
	// -- Install()
	//
	// 2) Method receiver is an application installation:
	// -- Start()
	// -- Uninstall()
	// -- Update()
	//
	// 3) Method receiver is application installation instance:
	// -- Refresh()
	// -- Restart()
	// -- Resume()
	// -- Stop()
	// -- Suspend()
	//
	// For groups 2) and 3), the suffix that specifies the receiver can
	// optionally omit the installation and/or instance, in which case the
	// operation applies to all installations and/or instances in the
	// scope of the suffix.
	//
	// Examples:
	// # Install Google Maps on the node.
	// device/apps.Install("/google.com/appstore/maps") --> "google maps/0"
	//
	// # Start an instance of the previously installed maps application installation.
	// device/apps/google maps/0.Start() --> { "0" }
	//
	// # Start a second instance of the previously installed maps application installation.
	// device/apps/google maps/0.Start() --> { "1" }
	//
	// # Stop the first instance previously started.
	// device/apps/google maps/0/0.Stop()
	//
	// # Install a second Google Maps installation.
	// device/apps.Install("/google.com/appstore/maps") --> "google maps/1"
	//
	// # Start an instance for all maps application installations.
	// device/apps/google maps.Start() --> {"0/2", "1/0"}
	//
	// # Refresh the state of all instances of all maps application installations.
	// device/apps/google maps.Refresh()
	//
	// # Refresh the state of all instances of the maps application installation
	// identified by the given suffix.
	// device/apps/google maps/0.Refresh()
	//
	// # Refresh the state of the maps application installation instance identified by
	// the given suffix.
	// device/apps/google maps/0/2.Refresh()
	//
	// # Update the second maps installation to the latest version available.
	// device/apps/google maps/1.Update()
	//
	// # Update the first maps installation to a specific version.
	// device/apps/google maps/0.UpdateTo("/google.com/appstore/beta/maps")
	//
	// Further, the following methods complement one another:
	// -- Install() and Uninstall()
	// -- Start() and Stop()
	// -- Suspend() and Resume()
	//
	// Finally, an application installation instance can be in one of
	// three abstract states: 1) "does not exist", 2) "running", or 3)
	// "suspended". The interface methods transition between these
	// abstract states using the following state machine:
	//
	// apply(Start(), "does not exists") = "running"
	// apply(Refresh(), "running") = "running"
	// apply(Refresh(), "suspended") = "suspended"
	// apply(Restart(), "running") = "running"
	// apply(Restart(), "suspended") = "running"
	// apply(Resume(), "suspended") = "running"
	// apply(Resume(), "running") = "running"
	// apply(Stop(), "running") = "does not exist"
	// apply(Stop(), "suspended") = "does not exist"
	// apply(Suspend(), "running") = "suspended"
	// apply(Suspend(), "suspended") = "suspended"
	//
	// In other words, invoking any method using an existing application
	// installation instance as a receiver is well-defined.
	Application_ExcludingUniversal
	// Describe generates a description of the node.
	Describe(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply Description, err error)
	// IsRunnable checks if the node can execute the given binary.
	IsRunnable(ctx _gen_context.T, Description binary.Description, opts ..._gen_ipc.CallOpt) (reply bool, err error)
	// Reset resets the node. If the deadline is non-zero and the node
	// in question is still running after the given deadline expired,
	// reset of the node is enforced.
	//
	// TODO(jsimsa): Switch deadline to time.Duration when built-in types
	// are implemented.
	Reset(ctx _gen_context.T, Deadline uint64, opts ..._gen_ipc.CallOpt) (err error)
}
type Node interface {
	_gen_ipc.UniversalServiceMethods
	Node_ExcludingUniversal
}

// NodeService is the interface the server implements.
type NodeService interface {

	// Application can be used to manage applications on a device. The
	// idea is that this interace will be invoked using an object name that
	// identifies the application and its installations and instances
	// where applicable.
	//
	// In particular, the interface methods can be divided into three
	// groups based on their intended receiver:
	//
	// 1) Method receiver is an application:
	// -- Install()
	//
	// 2) Method receiver is an application installation:
	// -- Start()
	// -- Uninstall()
	// -- Update()
	//
	// 3) Method receiver is application installation instance:
	// -- Refresh()
	// -- Restart()
	// -- Resume()
	// -- Stop()
	// -- Suspend()
	//
	// For groups 2) and 3), the suffix that specifies the receiver can
	// optionally omit the installation and/or instance, in which case the
	// operation applies to all installations and/or instances in the
	// scope of the suffix.
	//
	// Examples:
	// # Install Google Maps on the node.
	// device/apps.Install("/google.com/appstore/maps") --> "google maps/0"
	//
	// # Start an instance of the previously installed maps application installation.
	// device/apps/google maps/0.Start() --> { "0" }
	//
	// # Start a second instance of the previously installed maps application installation.
	// device/apps/google maps/0.Start() --> { "1" }
	//
	// # Stop the first instance previously started.
	// device/apps/google maps/0/0.Stop()
	//
	// # Install a second Google Maps installation.
	// device/apps.Install("/google.com/appstore/maps") --> "google maps/1"
	//
	// # Start an instance for all maps application installations.
	// device/apps/google maps.Start() --> {"0/2", "1/0"}
	//
	// # Refresh the state of all instances of all maps application installations.
	// device/apps/google maps.Refresh()
	//
	// # Refresh the state of all instances of the maps application installation
	// identified by the given suffix.
	// device/apps/google maps/0.Refresh()
	//
	// # Refresh the state of the maps application installation instance identified by
	// the given suffix.
	// device/apps/google maps/0/2.Refresh()
	//
	// # Update the second maps installation to the latest version available.
	// device/apps/google maps/1.Update()
	//
	// # Update the first maps installation to a specific version.
	// device/apps/google maps/0.UpdateTo("/google.com/appstore/beta/maps")
	//
	// Further, the following methods complement one another:
	// -- Install() and Uninstall()
	// -- Start() and Stop()
	// -- Suspend() and Resume()
	//
	// Finally, an application installation instance can be in one of
	// three abstract states: 1) "does not exist", 2) "running", or 3)
	// "suspended". The interface methods transition between these
	// abstract states using the following state machine:
	//
	// apply(Start(), "does not exists") = "running"
	// apply(Refresh(), "running") = "running"
	// apply(Refresh(), "suspended") = "suspended"
	// apply(Restart(), "running") = "running"
	// apply(Restart(), "suspended") = "running"
	// apply(Resume(), "suspended") = "running"
	// apply(Resume(), "running") = "running"
	// apply(Stop(), "running") = "does not exist"
	// apply(Stop(), "suspended") = "does not exist"
	// apply(Suspend(), "running") = "suspended"
	// apply(Suspend(), "suspended") = "suspended"
	//
	// In other words, invoking any method using an existing application
	// installation instance as a receiver is well-defined.
	ApplicationService
	// Describe generates a description of the node.
	Describe(context _gen_ipc.ServerContext) (reply Description, err error)
	// IsRunnable checks if the node can execute the given binary.
	IsRunnable(context _gen_ipc.ServerContext, Description binary.Description) (reply bool, err error)
	// Reset resets the node. If the deadline is non-zero and the node
	// in question is still running after the given deadline expired,
	// reset of the node is enforced.
	//
	// TODO(jsimsa): Switch deadline to time.Duration when built-in types
	// are implemented.
	Reset(context _gen_ipc.ServerContext, Deadline uint64) (err error)
}

// BindNode returns the client stub implementing the Node
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindNode(name string, opts ..._gen_ipc.BindOpt) (Node, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubNode{client: client, name: name}
	stub.Application_ExcludingUniversal, _ = BindApplication(name, client)

	return stub, nil
}

// NewServerNode creates a new server stub.
//
// It takes a regular server implementing the NodeService
// interface, and returns a new server stub.
func NewServerNode(server NodeService) interface{} {
	return &ServerStubNode{
		ServerStubApplication: *NewServerApplication(server).(*ServerStubApplication),
		service:               server,
	}
}

// clientStubNode implements Node.
type clientStubNode struct {
	Application_ExcludingUniversal

	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubNode) Describe(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply Description, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Describe", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubNode) IsRunnable(ctx _gen_context.T, Description binary.Description, opts ..._gen_ipc.CallOpt) (reply bool, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "IsRunnable", []interface{}{Description}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubNode) Reset(ctx _gen_context.T, Deadline uint64, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Reset", []interface{}{Deadline}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubNode) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubNode) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubNode) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubNode wraps a server that implements
// NodeService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubNode struct {
	ServerStubApplication

	service NodeService
}

func (__gen_s *ServerStubNode) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	if resp, err := __gen_s.ServerStubApplication.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	switch method {
	case "Describe":
		return []interface{}{}, nil
	case "IsRunnable":
		return []interface{}{}, nil
	case "Reset":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubNode) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Describe"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 66},
			{Name: "", Type: 67},
		},
	}
	result.Methods["IsRunnable"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Description", Type: 68},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 2},
			{Name: "", Type: 67},
		},
	}
	result.Methods["Reset"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Deadline", Type: 53},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 67},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.MapType{Key: 0x3, Elem: 0x2, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x41, Name: "Profiles"},
			},
			"veyron2/services/mgmt/node.Description", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x41, Name: "Profiles"},
			},
			"veyron2/services/mgmt/binary.Description", []string(nil)},
	}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = __gen_s.ServerStubApplication.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					wt.Fields[i].Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}

func (__gen_s *ServerStubNode) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
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

func (__gen_s *ServerStubNode) Describe(call _gen_ipc.ServerCall) (reply Description, err error) {
	reply, err = __gen_s.service.Describe(call)
	return
}

func (__gen_s *ServerStubNode) IsRunnable(call _gen_ipc.ServerCall, Description binary.Description) (reply bool, err error) {
	reply, err = __gen_s.service.IsRunnable(call, Description)
	return
}

func (__gen_s *ServerStubNode) Reset(call _gen_ipc.ServerCall, Deadline uint64) (err error) {
	err = __gen_s.service.Reset(call, Deadline)
	return
}
