package main

import (
	"fmt"
	"os"
	// v1 "k8s.io/api/apps/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	// "k8s.io/client-go/kubernetes"
	"github.com/amitde69/user-controller/user-api/handlers"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	schoolv1alpha1 "github.com/amitde69/user-controller/api/v1alpha1"
    "github.com/gin-gonic/gin"
)

func main() {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = schoolv1alpha1.AddToScheme(scheme)
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
    router.GET("/users", handlers.ListUsers(clientset, namespace))
    router.GET("/users/:name", handlers.GetUser(clientset, namespace))
    router.POST("/users", handlers.CreateUser(clientset, namespace))
    router.DELETE("/users/:name", handlers.DeleteUser(clientset, namespace))

    router.Run(":8080")
}