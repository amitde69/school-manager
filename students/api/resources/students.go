package resources

import (
	"context"
	// v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	schoolmanageriov1alpha1 "github.com/amitde69/school-manager/students/controller/api/v1alpha1"
	// "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
)

// list student resource using schoolmanageriov1alpha1
func ListStudents(clientset client.Client, namespace string) ([]StudentRes, error) {
	list := &schoolmanageriov1alpha1.StudentList{}
	err := clientset.List(context.Background(), list, client.InNamespace(namespace))
	if err != nil {
		return nil, err
	}
	studentsRes := make([]StudentRes, len(list.Items))
	for i, student := range list.Items {
		studentsRes[i].Name = student.Spec.Name
		studentsRes[i].Age = student.Spec.Age
		studentsRes[i].Size = student.Spec.Size
	}
	// for _, student := range students.Items {
	// 	fmt.Printf("name: %s\nage: %s\n", student.Spec.Name, student.Spec.Age)
	// }
	return studentsRes, err
}

// get student resource using schoolmanageriov1alpha1
func GetStudent(clientset client.Client, namespace string, studentName string) (StudentRes, error) {
	// student := &schoolmanageriov1alpha1.Student{}
	studentRes := StudentRes{}
	studentExists := &schoolmanageriov1alpha1.StudentList{}
	err := clientset.List(context.Background(), studentExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "student-controller", "name": studentName}))
	if err != nil {
		return studentRes, err
	}
	if len(studentExists.Items) == 0 {
		return studentRes, err
	}
	studentRes.Name = studentExists.Items[0].Spec.Name
	studentRes.Age = studentExists.Items[0].Spec.Age
	studentRes.Size = studentExists.Items[0].Spec.Size
	return studentRes, err
}

// delete student resource using schoolmanageriov1alpha1
func DeleteStudent(clientset client.Client, namespace string, studentName string) (schoolmanageriov1alpha1.Student, error) {
	student := &schoolmanageriov1alpha1.Student{}
	studentExists := &schoolmanageriov1alpha1.StudentList{}
	err := clientset.List(context.Background(), studentExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "student-controller", "name": studentName}))
	if err != nil {
		return *student, err
	}
	if len(studentExists.Items) == 0 {
		return *student, err
	}
	student = &studentExists.Items[0]
	err = clientset.Delete(context.Background(), student)
	return *student, err
}

// create student resource using schoolmanageriov1alpha1
func CreateStudent(clientset client.Client, namespace string, newStudent StudentRes) (StudentRes, error) {
	student := newStudentCR(namespace, &newStudent)
	studentExists := &schoolmanageriov1alpha1.StudentList{}
	err := clientset.List(context.Background(), studentExists, client.InNamespace(namespace), 
		client.MatchingLabels(labels.Set{"app": "student-controller", "name": newStudent.Name}))
	if len(studentExists.Items) == 0 {
		err = clientset.Create(context.Background(), student)
	} else {
		student.ResourceVersion = studentExists.Items[0].ResourceVersion
		err = clientset.Update(context.Background(), student)
	}
	if err != nil {
		return StudentRes{}, err
	}
	return newStudent, err
}


func newStudentCR(namespace string, student *StudentRes) *schoolmanageriov1alpha1.Student {
	labels := map[string]string{
		"app": "student-controller",
		"name": student.Name,
	}
	return &schoolmanageriov1alpha1.Student{
		ObjectMeta: metav1.ObjectMeta{
			Name: student.Name,
			Namespace:    namespace,
			Labels:       labels,
		},
		Spec: schoolmanageriov1alpha1.StudentSpec{
			Name: student.Name,
			Age: student.Age,
			Size: student.Size,
		},
	}
}

type StudentRes struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Size int32 `json:"size"`
}