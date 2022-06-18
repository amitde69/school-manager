
package handlers
import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"net/http"
	"github.com/amitde69/school-manager/students/api/resources"
    "github.com/gin-gonic/gin"
)

func ListStudents(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
     	students, err := resources.ListStudents(clientset, namespace)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, students)
	}
}

func GetStudent(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentName := c.Param("name")
     	student, err := resources.GetStudent(clientset, namespace, studentName)
		if err != nil {
			panic(err.Error())
		}
		if student.FirstName == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.IndentedJSON(http.StatusOK, student)
		}
	}
}

func CreateStudent(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		newStudent := resources.StudentRes{}
        c.BindJSON(&newStudent)
		student, err := resources.CreateStudent(clientset, namespace, newStudent)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, student)
	}
}

func DeleteStudent(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentName := c.Param("name")
     	student, err := resources.DeleteStudent(clientset, namespace, studentName)
		if err != nil {
			panic(err.Error())
		}
		if student.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Student deleted"})
		}
	}
}
