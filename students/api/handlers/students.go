
package handlers
import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"net/http"
	"github.com/amitde69/school-manager/api/resources"
    "github.com/gin-gonic/gin"
)

func ListUsers(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
     	users, err := resources.ListUsers(clientset, namespace)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, users)
	}
}

func GetUser(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Param("name")
     	user, err := resources.GetUser(clientset, namespace, userName)
		if err != nil {
			panic(err.Error())
		}
		if user.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}
}

func CreateUser(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		newUser := resources.UserRes{}
        c.BindJSON(&newUser)
		user, err := resources.CreateUser(clientset, namespace, newUser)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, user)
	}
}

func DeleteUser(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Param("name")
     	user, err := resources.DeleteUser(clientset, namespace, userName)
		if err != nil {
			panic(err.Error())
		}
		if user.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})
		}
	}
}
