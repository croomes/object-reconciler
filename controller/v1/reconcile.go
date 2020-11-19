package v1

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Reconcile implements the object controller reconciliation.
func (r ObjectReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	controller := r.C
	result := ctrl.Result{}

	// Initialize the reconciler.
	controller.InitReconcile(ctx, req)
	if err := controller.FetchInstance(); err != nil {
		return result, client.IgnoreNotFound(err)
	}

	// Run the operation.
	result, event, err := r.C.Operate()

	// Record an event if the operation returned one.
	if event != nil {
		event.Record(r.Recorder)
	}

	return result, err
}
