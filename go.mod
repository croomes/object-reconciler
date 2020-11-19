module github.com/croomes/object-reconciler

go 1.15

replace github.com/croomes/object-reconciler => /home/simon/go/src/github.com/croomes/object-reconciler

require (
	github.com/go-logr/logr v0.1.0
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.3
)
