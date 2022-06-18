/*
Copyright 2022.

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
	// "fmt"
	// corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/apimachinery/pkg/types"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	schoolmanageriov1alpha1 "github.com/amitde69/school-manager/api/v1alpha1"
)

// StudentReconciler reconciles a Student object
type StudentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=schoolmanager.io,resources=students,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=schoolmanager.io,resources=students/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=schoolmanager.io,resources=students/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Student object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *StudentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling Students")

	student := &schoolmanageriov1alpha1.Student{}

	err := r.Get(ctx, req.NamespacedName, student)
	if err != nil {
		// if the resource is not found, then just return (might look useless as this usually happens in case of Delete events)
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Error occurred while fetching the Student resource")
		return ctrl.Result{}, err
	}

	// lbls := labels.Set{
	// 	"app": "school-manager",
	// 	"name": student.Name,
	// }
	// existingPods := &corev1.PodList{}
	// err = r.Client.List(ctx, existingPods,
	// 	&client.ListOptions{
	// 		Namespace:     student.Namespace,
	// 		LabelSelector: labels.SelectorFromSet(lbls),
	// 	})
	// if err != nil {
	// 	logger.Error(err, "Error occurred while listing pods under the user resource")
	// 	return reconcile.Result{}, err
	// }

	// existingPodNames := []string{}

	// for _, pod := range existingPods.Items {
	// 	if pod.GetObjectMeta().GetDeletionTimestamp() != nil {
	// 		continue
	// 	}
	// 	if pod.Status.Phase == corev1.PodPending || pod.Status.Phase == corev1.PodRunning {
	// 		existingPodNames = append(existingPodNames, pod.Name)
	// 	}
	// 	logger.Info("Found existing pod: " + pod.Name)
	// }

	// // if len(existingPodNames) == 0 {
	// // 	logger.Info("No existing pods found")
	// // }

	// logger.Info("Assesing number of pods against desired number of pods")

	status := schoolmanageriov1alpha1.StudentStatus{
		CurrenctClass: "-",
		Presence:      false,
	}

	if !reflect.DeepEqual(student.Status, status) {
		student.Status = status
		err := r.Client.Status().Update(ctx, student)
		if err != nil {
			logger.Error(err, "Error occurred while updating the user resource")
			return reconcile.Result{}, err
		}
	}

	// if int32(len(existingPodNames)) > student.Spec.Size {
	// 	logger.Info("Deleting a pod in the user", "expected size", student.Spec.Size, "Pod.Name", existingPods.Items[0].Name)
	// 	pod := existingPods.Items[0]
	// 	err = r.Client.Delete(ctx, &pod)
	// 	if err != nil {
	// 		logger.Error(err, "Error occurred while deleting the pod")
	// 		return reconcile.Result{}, err
	// 	}
	// }

	// if int32(len(existingPodNames)) < student.Spec.Size {
	// 	logger.Info("Adding a pod in the user", "expected size", student.Spec.Size, "Pod.Names", existingPodNames)
	// 	pod := newPodForUser(student)
	// 	if err := controllerutil.SetControllerReference(student, pod, r.Scheme); err != nil {
	// 		logger.Error(err, "unable to set owner reference on new pod")
	// 		return reconcile.Result{}, err
	// 	}
	// 	err = r.Client.Create(ctx, pod)
	// 	if err != nil {
	// 		logger.Error(err, "Error occurred while creating the pod")
	// 		return reconcile.Result{}, err
	// 	}
	// }

	return reconcile.Result{Requeue: true}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StudentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&schoolmanageriov1alpha1.Student{}).
		Complete(r)
}
