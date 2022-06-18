package resources

import (
	"context"
	// v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	schoolmanageriov1alpha1 "github.com/amitde69/school-manager/classes/controller/api/v1alpha1"
	// "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
)

// list class resource using schoolmanageriov1alpha1
func ListClasses(clientset client.Client, namespace string) ([]ClassRes, error) {
	list := &schoolmanageriov1alpha1.ClassList{}
	err := clientset.List(context.Background(), list, client.InNamespace(namespace))
	if err != nil {
		return nil, err
	}
	classesRes := make([]ClassRes, len(list.Items))
	for i, class := range list.Items {
		classesRes[i].FirstName = class.Spec.FirstName
		classesRes[i].LastName = class.Spec.LastName
		classesRes[i].Age = class.Spec.Age
		classesRes[i].Id = class.Spec.Id
		classesRes[i].Classes = class.Spec.Classes
	}
	// for _, class := range classes.Items {
	// 	fmt.Printf("name: %s\nage: %s\n", class.Spec.Name, class.Spec.Age)
	// }
	return classesRes, err
}

// get class resource using schoolmanageriov1alpha1
func GetClass(clientset client.Client, namespace string, className string) (ClassRes, error) {
	// class := &schoolmanageriov1alpha1.Class{}
	classRes := ClassRes{}
	classExists := &schoolmanageriov1alpha1.ClassList{}
	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "class-controller", "name": className}))
	if err != nil {
		return classRes, err
	}
	if len(classExists.Items) == 0 {
		return classRes, err
	}
	classRes.FirstName = classExists.Items[0].Spec.FirstName
	classRes.LastName = classExists.Items[0].Spec.LastName
	classRes.Age = classExists.Items[0].Spec.Age
	classRes.Id = classExists.Items[0].Spec.Id
	classRes.Classes = classExists.Items[0].Spec.Classes
	return classRes, err
}

// delete class resource using schoolmanageriov1alpha1
func DeleteClass(clientset client.Client, namespace string, className string) (schoolmanageriov1alpha1.Class, error) {
	class := &schoolmanageriov1alpha1.Class{}
	classExists := &schoolmanageriov1alpha1.ClassList{}
	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "class-controller", "name": className}))
	if err != nil {
		return *class, err
	}
	if len(classExists.Items) == 0 {
		return *class, err
	}
	class = &classExists.Items[0]
	err = clientset.Delete(context.Background(), class)
	return *class, err
}

// create class resource using schoolmanageriov1alpha1
func CreateClass(clientset client.Client, namespace string, newClass ClassRes) (ClassRes, error) {
	class := newClassCR(namespace, &newClass)
	classExists := &schoolmanageriov1alpha1.ClassList{}
	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "class-controller", "name": newClass.FirstName}))
	if len(classExists.Items) == 0 {
		err = clientset.Create(context.Background(), class)
	} else {
		class.ResourceVersion = classExists.Items[0].ResourceVersion
		err = clientset.Update(context.Background(), class)
	}
	if err != nil {
		return ClassRes{}, err
	}
	return newClass, err
}


func newClassCR(namespace string, class *ClassRes) *schoolmanageriov1alpha1.Class {
	labels := map[string]string{
		"app": "class-controller",
		"name": class.FirstName,
	}
	return &schoolmanageriov1alpha1.Class{
		ObjectMeta: metav1.ObjectMeta{
			Name: class.FirstName,
			Namespace:    namespace,
			Labels:       labels,
		},
		Spec: schoolmanageriov1alpha1.ClassSpec{
			FirstName: class.FirstName,
			LastName: class.LastName,
			Age: class.Age,
			Id: class.Id,
			Classes: class.Classes,
		},
	}
}

type ClassRes struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Age  int32 `json:"age"`
	Id int32 `json:"id"`
	Classes []string `json:"classes"`
}