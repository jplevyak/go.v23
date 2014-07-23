// This file was auto-generated by the veyron vdl tool.
// Source: advanced.vdl

package arith

import (
	"veyron2/vdl/test_arith/exp"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Trigonometry is an interface that specifies a couple trigonometric functions.
// Trigonometry is the interface the client binds and uses.
// Trigonometry_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Trigonometry_ExcludingUniversal interface {
	Sine(ctx _gen_context.T, angle float64, opts ..._gen_ipc.CallOpt) (reply float64, err error)
	Cosine(ctx _gen_context.T, angle float64, opts ..._gen_ipc.CallOpt) (reply float64, err error)
}
type Trigonometry interface {
	_gen_ipc.UniversalServiceMethods
	Trigonometry_ExcludingUniversal
}

// TrigonometryService is the interface the server implements.
type TrigonometryService interface {
	Sine(context _gen_ipc.ServerContext, angle float64) (reply float64, err error)
	Cosine(context _gen_ipc.ServerContext, angle float64) (reply float64, err error)
}

// BindTrigonometry returns the client stub implementing the Trigonometry
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindTrigonometry(name string, opts ..._gen_ipc.BindOpt) (Trigonometry, error) {
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
	stub := &clientStubTrigonometry{client: client, name: name}

	return stub, nil
}

// NewServerTrigonometry creates a new server stub.
//
// It takes a regular server implementing the TrigonometryService
// interface, and returns a new server stub.
func NewServerTrigonometry(server TrigonometryService) interface{} {
	return &ServerStubTrigonometry{
		service: server,
	}
}

// clientStubTrigonometry implements Trigonometry.
type clientStubTrigonometry struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubTrigonometry) Sine(ctx _gen_context.T, angle float64, opts ..._gen_ipc.CallOpt) (reply float64, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Sine", []interface{}{angle}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTrigonometry) Cosine(ctx _gen_context.T, angle float64, opts ..._gen_ipc.CallOpt) (reply float64, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Cosine", []interface{}{angle}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTrigonometry) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTrigonometry) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubTrigonometry) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubTrigonometry wraps a server that implements
// TrigonometryService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubTrigonometry struct {
	service TrigonometryService
}

func (__gen_s *ServerStubTrigonometry) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Sine":
		return []interface{}{}, nil
	case "Cosine":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubTrigonometry) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Cosine"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "angle", Type: 26},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 26},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Sine"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "angle", Type: 26},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 26},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubTrigonometry) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubTrigonometry) Sine(call _gen_ipc.ServerCall, angle float64) (reply float64, err error) {
	reply, err = __gen_s.service.Sine(call, angle)
	return
}

func (__gen_s *ServerStubTrigonometry) Cosine(call _gen_ipc.ServerCall, angle float64) (reply float64, err error) {
	reply, err = __gen_s.service.Cosine(call, angle)
	return
}

// AdvancedMath is an interface for more advanced math than arith.  It embeds
// interfaces defined both in the same file and in an external package; and in
// turn it is embedded by arith.Calculator (which is in the same package but
// different file) to verify that embedding works in all these scenarios.
// AdvancedMath is the interface the client binds and uses.
// AdvancedMath_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type AdvancedMath_ExcludingUniversal interface {
	// Trigonometry is an interface that specifies a couple trigonometric functions.
	Trigonometry_ExcludingUniversal
	exp.Exp_ExcludingUniversal
}
type AdvancedMath interface {
	_gen_ipc.UniversalServiceMethods
	AdvancedMath_ExcludingUniversal
}

// AdvancedMathService is the interface the server implements.
type AdvancedMathService interface {

	// Trigonometry is an interface that specifies a couple trigonometric functions.
	TrigonometryService
	exp.ExpService
}

// BindAdvancedMath returns the client stub implementing the AdvancedMath
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindAdvancedMath(name string, opts ..._gen_ipc.BindOpt) (AdvancedMath, error) {
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
	stub := &clientStubAdvancedMath{client: client, name: name}
	stub.Trigonometry_ExcludingUniversal, _ = BindTrigonometry(name, client)
	stub.Exp_ExcludingUniversal, _ = exp.BindExp(name, client)

	return stub, nil
}

// NewServerAdvancedMath creates a new server stub.
//
// It takes a regular server implementing the AdvancedMathService
// interface, and returns a new server stub.
func NewServerAdvancedMath(server AdvancedMathService) interface{} {
	return &ServerStubAdvancedMath{
		ServerStubTrigonometry: *NewServerTrigonometry(server).(*ServerStubTrigonometry),
		ServerStubExp:          *exp.NewServerExp(server).(*exp.ServerStubExp),
		service:                server,
	}
}

// clientStubAdvancedMath implements AdvancedMath.
type clientStubAdvancedMath struct {
	Trigonometry_ExcludingUniversal
	exp.Exp_ExcludingUniversal

	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubAdvancedMath) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAdvancedMath) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubAdvancedMath) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubAdvancedMath wraps a server that implements
// AdvancedMathService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubAdvancedMath struct {
	ServerStubTrigonometry
	exp.ServerStubExp

	service AdvancedMathService
}

func (__gen_s *ServerStubAdvancedMath) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	if resp, err := __gen_s.ServerStubTrigonometry.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	if resp, err := __gen_s.ServerStubExp.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	return nil, nil
}

func (__gen_s *ServerStubAdvancedMath) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}

	result.TypeDefs = []_gen_vdlutil.Any{}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = __gen_s.ServerStubTrigonometry.Signature(call)
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
	ss, _ = __gen_s.ServerStubExp.Signature(call)
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

func (__gen_s *ServerStubAdvancedMath) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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
