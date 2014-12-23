package ipc

import (
	"veyron.io/veyron/veyron2/verror"
	"veyron.io/veyron/veyron2/verror2"
	"veyron.io/veyron/veyron2/vom"

	// Ensure all standard vdl types are registered.
	_ "veyron.io/veyron/veyron2/vdl/vdlroot"
)

func init() {
	// The verror.Standard type is used by the ipc package to marshal errors.
	// TODO(toddw): Remove this after the vom2 transition.
	vom.Register(verror.Standard{})
	vom.Register(verror2.Standard{})
}
