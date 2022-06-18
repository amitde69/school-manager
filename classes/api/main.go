package main

import (
	"fmt"
	"os"
	// v1 "k8s.io/api/apps/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	// "k8s.io/client-go/kubernetes"
	"github.com/amitde69/school-manager/classes/api/handlers"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	classesapi "github.com/amitde69/school-manager/classes/controller/api/v1alpha1"
	studentsapi "github.com/amitde69/school-manager/students/controller/api/v1alpha1"
    "github.com/gin-gonic/gin"
)

func main() {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = classesapi.AddToScheme(scheme)
	_ = studentsapi.AddToScheme(scheme)
	kubeconfig := ctrl.GetConfigOrDie()
	clientset, err := client.New(kubeconfig, client.Options{ Scheme: scheme })
	if err != nil {
        fmt.Printf("error: %v", err)
	}	
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		namespace = "default"
	}

    router := gin.Default()
    router.GET("/classes", handlers.ListClasses(clientset, namespace))
    router.GET("/classes/:name", handlers.GetClass(clientset, namespace))
    router.POST("/classes", handlers.CreateClass(clientset, namespace))
    router.POST("/classes/student", handlers.AddStudent(clientset, namespace))
    router.POST("/classes/teacher", handlers.ChangeTeacher(clientset, namespace))
    router.DELETE("/classes/:name", handlers.DeleteClass(clientset, namespace))

    router.Run(":8080")
}