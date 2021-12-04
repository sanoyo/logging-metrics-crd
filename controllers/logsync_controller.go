/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"cloud.google.com/go/logging/logadmin"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	loggingv1 "github.com/sanoyo/logsync-crd/api/v1"
)

var setupLog = ctrl.Log.WithName("setup")

// LogSyncReconciler reconciles a LogSync object
type LogSyncReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=logging.yo.logsync,resources=logsyncs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=logging.yo.logsync,resources=logsyncs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=logging.yo.logsync,resources=logsyncs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *LogSyncReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	instance := &loggingv1.LogSync{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	if instance.Spec.ProjectID == "" || instance.Spec.Destination == "" || instance.Spec.Filter == "" {
		setupLog.Error(err, "bad argument")
	}

	projectID := instance.Spec.ProjectID
	destination := instance.Spec.Destination
	filter := instance.Spec.Filter

	client, err := logadmin.NewClient(ctx, projectID)
	if err != nil {
		setupLog.Error(err, "Unable to initialize")
	}

	_, err = client.CreateSink(ctx, &logadmin.Sink{
		ID:          "gcs-sync",
		Destination: destination,
		Filter:      filter,
	})
	if err != nil {
		setupLog.Error(err, "Unable to create logsync")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LogSyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loggingv1.LogSync{}).
		Complete(r)
}
