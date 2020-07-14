package main

import (
	"context"
	"time"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	runinformer "github.com/tektoncd/pipeline/pkg/client/injection/informers/pipeline/v1alpha1/run"
	runreconciler "github.com/tektoncd/pipeline/pkg/client/injection/reconciler/pipeline/v1alpha1/run"
	tkncontroller "github.com/tektoncd/pipeline/pkg/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
)

const controllerName = "group-task-controller"

func main() {
	sharedmain.Main(controllerName, newController)
}

func newController(ctx context.Context, cmw configmap.Watcher) *controller.Impl {
	c := &Reconciler{}
	impl := runreconciler.NewImpl(ctx, c, func(impl *controller.Impl) controller.Options {
		return controller.Options{
			AgentName: controllerName,
		}
	})
	c.enqueueAfter = impl.EnqueueAfter

	runinformer.Get(ctx).Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: tkncontroller.FilterRunRef("example.dev/v0", "Wait"),
		Handler:    controller.HandleAll(impl.Enqueue),
	})

	return impl
}

type Reconciler struct {
	enqueueAfter func(interface{}, time.Duration)
}

// ReconcileKind implements Interface.ReconcileKind.
func (c *Reconciler) ReconcileKind(ctx context.Context, r *v1alpha1.Run) reconciler.Event {
	logger := logging.FromContext(ctx)
	logger.Infof("Reconciling %s/%s", r.Namespace, r.Name)

	// Ignore completed runs.
	if r.IsDone() {
		logger.Info("Run is finished, done reconciling")
		return nil
	}

	// TODO: look up grouptask, create concatenated TaskRun.

	return reconciler.NewEvent(corev1.EventTypeNormal, "RunReconciled", "Run reconciled: \"%s/%s\"", r.Namespace, r.Name)
}
