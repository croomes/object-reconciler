package v1

import (
	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
)

// ObjectReconciler defines a Kubernetes object reconciler.
type ObjectReconciler struct {
	Log      logr.Logger
	C        Controller
	Recorder record.EventRecorder
}
