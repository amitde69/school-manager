package resources

import (
	"context"
	"fmt"
	// v1 "k8s.io/api/apps/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	"errors"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"
	classesapi "github.com/amitde69/school-manager/classes/controller/api/v1alpha1"
	studentsapi "github.com/amitde69/school-manager/students/controller/api/v1alpha1"
	// "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
)

// // list class resource using classesapi
// func ListClasses(clientset client.Client, namespace string) ([]ClassRes, error) {
// 	list := &classesapi.ClassList{}
// 	err := clientset.List(context.Background(), list, client.InNamespace(namespace))
// 	if err != nil {
// 		return nil, err
// 	}
// 	classesRes := make([]ClassRes, len(list.Items))
// 	for i, class := range list.Items {
// 		classesRes[i].Name = class.Spec.Name
// 		classesRes[i].Teacher = class.Spec.Teacher
// 		classesRes[i].Students = class.Spec.Students
// 	}
// 	// for _, class := range classes.Items {
// 	// 	fmt.Printf("name: %s\nage: %s\n", class.Spec.Name, class.Spec.Age)
// 	// }
// 	return classesRes, err
// }

// // get class resource using classesapi
// func GetClass(clientset client.Client, namespace string, className string) (ClassRes, error) {
// 	// class := &classesapi.Class{}
// 	classRes := ClassRes{}
// 	classExists := &classesapi.ClassList{}
// 	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
// 		client.MatchingLabels(labels.Set{"app": "class-controller", "name": className}))
// 	if err != nil {
// 		return classRes, err
// 	}
// 	if len(classExists.Items) == 0 {
// 		return classRes, err
// 	}
// 	classRes.Name = classExists.Items[0].Spec.Name
// 	classRes.Teacher = classExists.Items[0].Spec.Teacher
// 	classRes.Students = classExists.Items[0].Spec.Students
// 	return classRes, err
// }

// // delete class resource using classesapi
// func DeleteClass(clientset client.Client, namespace string, className string) (classesapi.Class, error) {
// 	class := &classesapi.Class{}
// 	classExists := &classesapi.ClassList{}
// 	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
// 		client.MatchingLabels(labels.Set{"app": "class-controller", "name": className}))
// 	if err != nil {
// 		return *class, err
// 	}
// 	if len(classExists.Items) == 0 {
// 		return *class, err
// 	}
// 	class = &classExists.Items[0]
// 	err = clientset.Delete(context.Background(), class)
// 	return *class, err
// }


func AddStudent(clientset client.Client, namespace string, student string, class string) (ClassRes, error) {
	classExists := &classesapi.ClassList{}
	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "class-controller", "name": class}))
	if len(classExists.Items) == 0 {
		err = errors.New("class not found")
		return ClassRes{}, err
	}
	if classExists.Items[0].Status.Available == false {
		err = errors.New("class is not available")
		return ClassRes{}, err
	}	
	studentExists := &studentsapi.StudentList{}
	err = clientset.List(context.Background(), studentExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "student-controller", "name": student}))
	if len(studentExists.Items) == 0 {
		fmt.Printf("student not found")
		err = errors.New("student not found")
		return ClassRes{}, err
	}
	newClass := &classesapi.Class{}
	newClass = &classExists.Items[0]
	// newClass := &classesapi.Class{}
	// newClass.Spec.Name = classExists.Items[0].Spec.Name
	// newClass.Spec.Teacher = classExists.Items[0].Spec.Teacher
	// newClass.Spec.Students = classExists.Items[0].Spec.Students
	newClass.Spec.Students = append(newClass.Spec.Students, student)
	// newClass.ResourceVersion = classExists.Items[0].ResourceVersion
	err = clientset.Update(context.Background(), newClass)
	if err != nil {
		return ClassRes{}, err
	}
	classres := ClassRes{}
	classres.Name = newClass.Spec.Name
	classres.Teacher = newClass.Spec.Teacher
	classres.Students = newClass.Spec.Students  
	return classres, err
}

func RemoveStudent(clientset client.Client, namespace string, student string, class string) (ClassRes, error) {
	classExists := &classesapi.ClassList{}
	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "class-controller", "name": class}))
	if len(classExists.Items) == 0 {
		err = errors.New("class not found")
		return ClassRes{}, err
	}
	if classExists.Items[0].Status.Available == false {
		err = errors.New("class is not available")
		return ClassRes{}, err
	}	
	studentExists := &studentsapi.StudentList{}
	err = clientset.List(context.Background(), studentExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "student-controller", "name": student}))
	if len(studentExists.Items) == 0 {
		fmt.Printf("student not found")
		err = errors.New("student not found")
		return ClassRes{}, err
	}
	newClass := &classesapi.Class{}
	newClass = &classExists.Items[0]
	for i, liststudent := range newClass.Spec.Students { 
		if reflect.DeepEqual(liststudent, student) {
			newClass.Spec.Students = RemoveIndex(newClass.Spec.Students, i)
			break
		}
	}
	err = clientset.Update(context.Background(), newClass)
	if err != nil {
		return ClassRes{}, err
	}
	classres := ClassRes{}
	classres.Name = newClass.Spec.Name
	classres.Teacher = newClass.Spec.Teacher
	classres.Students = newClass.Spec.Students  
	return classres, err
}

// create class resource using classesapi
func ChangeTeacher(clientset client.Client, namespace string, teacher string, class string) (ClassRes, error) {
	classExists := &classesapi.ClassList{}
	err := clientset.List(context.Background(), classExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "class-controller", "name": class}))
	if len(classExists.Items) == 0 {
		fmt.Printf("class not found")
		return ClassRes{}, err
	}	
	newClass := &classesapi.Class{}
	newClass = &classExists.Items[0]
	newClass.Spec.Teacher = teacher
	err = clientset.Update(context.Background(), newClass)
	if err != nil {
		return ClassRes{}, err
	}
	classres := ClassRes{}
	classres.Name = newClass.Spec.Name
	classres.Teacher = newClass.Spec.Teacher
	classres.Students = newClass.Spec.Students  
	return classres, err
}


// func newClassCR(namespace string, class *ClassRes) *classesapi.Class {
// 	labels := map[string]string{
// 		"app": "class-controller",
// 		"name": class.Name,
// 	}
// 	return &classesapi.Class{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: class.Name,
// 			Namespace:    namespace,
// 			Labels:       labels,
// 		},
// 		Spec: classesapi.ClassSpec{
// 			Name: class.Name,
// 			Teacher: class.Teacher,
// 			Students: class.Students,
// 		},
// 	}
// }

type StudentAdd struct {
	StudentName string `json:"studentname"`
	ClassName string `json:"classname"`
}
type TeacherChange struct {
	TeacherName string `json:"teachername"`
	ClassName string `json:"classname"`
}


func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}