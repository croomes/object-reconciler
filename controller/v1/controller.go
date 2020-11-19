package v1

import (
	"context"

	eventv1 "github.com/croomes/object-reconciler/event/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Controller is the controller interface that must be implemented by an
// external controller. It provides methods required for reconciling a
// external controller.
type Controller interface {
	// InitReconcile initializes the node object request.
	InitReconcile(ctx context.Context, req reconcile.Request)

	// FetchInstance retrieves the request's object.
	FetchInstance() error

	// Operate runs the core operation of the controller that ensures that
	// the child objects or the other objects and configurations in the
	// environment are in the desired state. It should be able to update any
	// existing resources or create one, if there's a configuration drift,
	// based on the type of objects.
	// The returned result is the returned reconcile result. eventMessage is a
	// message that's emitted to the primary object. It's related to the change
	// done in the operation. A controller can use this to emit event of one
	// change performed by the reconciler and return the result with requeue
	// set to true.
	Operate() (result ctrl.Result, event eventv1.ReconcilerEvent, err error)
}
