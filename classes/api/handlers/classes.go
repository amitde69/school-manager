
package handlers
import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"net/http"
	"github.com/amitde69/school-manager/classes/api/resources"
    "github.com/gin-gonic/gin"
)

func ListClasses(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
     	classes, err := resources.ListClasses(clientset, namespace)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, classes)
	}
}

func GetClass(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		className := c.Param("name")
     	class, err := resources.GetClass(clientset, namespace, className)
		if err != nil {
			panic(err.Error())
		}
		if class.FirstName == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.IndentedJSON(http.StatusOK, class)
		}
	}
}

func CreateClass(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		newClass := resources.ClassRes{}
        c.BindJSON(&newClass)
		class, err := resources.CreateClass(clientset, namespace, newClass)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, class)
	}
}

func DeleteClass(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		className := c.Param("name")
     	class, err := resources.DeleteClass(clientset, namespace, className)
		if err != nil {
			panic(err.Error())
		}
		if class.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Class deleted"})
		}
	}
}
