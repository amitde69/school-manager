
package handlers
import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"net/http"
	"github.com/amitde69/school-manager/classes/api/resources"
    "github.com/gin-gonic/gin"
)

// func ListClasses(clientset client.Client, namespace string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
//      	classes, err := resources.ListClasses(clientset, namespace)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		c.IndentedJSON(http.StatusOK, classes)
// 	}
// }

// func GetClass(clientset client.Client, namespace string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		className := c.Param("name")
//      	class, err := resources.GetClass(clientset, namespace, className)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		if class.Name == "" {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
// 		} else {
// 			c.IndentedJSON(http.StatusOK, class)
// 		}
// 	}
// }

func AddStudent(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		newStudent := resources.StudentAdd{}
        c.BindJSON(&newStudent)
		class, err := resources.AddStudent(clientset, namespace, newStudent.StudentName, newStudent.ClassName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, class)
		}
	}
}

func RemoveStudent(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		newStudent := resources.StudentAdd{}
        c.BindJSON(&newStudent)
		class, err := resources.RemoveStudent(clientset, namespace, newStudent.StudentName, newStudent.ClassName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, class)
		}
	}
}

func ChangeTeacher(clientset client.Client, namespace string) gin.HandlerFunc {
	return func(c *gin.Context) {
		newTeacher := resources.TeacherChange{}
        c.BindJSON(&newTeacher)
		class, err := resources.ChangeTeacher(clientset, namespace, newTeacher.TeacherName, newTeacher.ClassName)
		if err != nil {
			panic(err.Error())
		}
		c.IndentedJSON(http.StatusOK, class)
	}
}

// func DeleteClass(clientset client.Client, namespace string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		className := c.Param("name")
//      	class, err := resources.DeleteClass(clientset, namespace, className)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		if class.Name == "" {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
// 		} else {
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "Class deleted"})
// 		}
// 	}
// }
