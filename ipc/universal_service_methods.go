package ipc

// UniversalServiceMethods defines the set of methods that are implemented on
// all services.
type UniversalServiceMethods interface {
	// TODO(bprosnitz) Remove GetMethodTags and fetch the method tags from
	// signature instead.
	// GetMethodTags returns the tags associated with the given method.
	GetMethodTags(method string, opts ...ClientCallOpt) ([]interface{}, error)
	// Signature returns a description of the service.
	Signature(opts ...ClientCallOpt) (ServiceSignature, error)
	// UnresolveStep returns the names for the remote service, rooted at the
	// service's immediate namespace ancestor.
	UnresolveStep(opts ...ClientCallOpt) ([]string, error)
}
