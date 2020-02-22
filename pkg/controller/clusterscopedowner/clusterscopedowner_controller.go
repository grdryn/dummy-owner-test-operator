package clusterscopedowner

import (
	"context"

	appv1alpha1 "github.com/grdryn/dummy-owner-test-operator/pkg/apis/app/v1alpha1"

	consolev1 "github.com/openshift/api/console/v1"

	oauthv1 "github.com/openshift/api/oauth/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_clusterscopedowner")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ClusterScopedOwner Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileClusterScopedOwner{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("clusterscopedowner-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ClusterScopedOwner
	err = c.Watch(&source.Kind{Type: &appv1alpha1.ClusterScopedOwner{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ClusterScopedOwner
	err = c.Watch(&source.Kind{Type: &corev1.Namespace{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appv1alpha1.ClusterScopedOwner{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileClusterScopedOwner implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileClusterScopedOwner{}

// ReconcileClusterScopedOwner reconciles a ClusterScopedOwner object
type ReconcileClusterScopedOwner struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ClusterScopedOwner object and makes changes based on the state read
// and what is in the ClusterScopedOwner.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileClusterScopedOwner) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ClusterScopedOwner")

	// Fetch the ClusterScopedOwner instance
	instance := &appv1alpha1.ClusterScopedOwner{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	consoleLink := &consolev1.ConsoleLink{ObjectMeta: metav1.ObjectMeta{Name: instance.Name}}
	_, err = controllerutil.CreateOrUpdate(context.TODO(), r.client, consoleLink, func() error {
		consoleLink.Spec.Href = "https://test.com"
		consoleLink.Spec.Location = consolev1.HelpMenu
		consoleLink.Spec.Text = "TEST"

		return controllerutil.SetControllerReference(instance, consoleLink, r.scheme)
	})
	if err != nil {
		return reconcile.Result{}, err
	}

	oauthClient := &oauthv1.OAuthClient{ObjectMeta: metav1.ObjectMeta{Name: instance.Name}}
	_, err = controllerutil.CreateOrUpdate(context.TODO(), r.client, oauthClient, func() error {
		oauthClient.GrantMethod = oauthv1.GrantHandlerAuto
		oauthClient.Secret = "test"
		oauthClient.RedirectURIs = []string{"http://test.com"}

		return controllerutil.SetControllerReference(instance, oauthClient, r.scheme)
	})
	if err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("Successfully reconciled ClusterScopedOwner")
	return reconcile.Result{}, nil
}
