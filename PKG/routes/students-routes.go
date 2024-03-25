package routes
import (
"github.com/gorilla/mux"
"github.com/20WilliamPanjaitan/go_students_crud_mysql/pkg/controllers"
)
var RegisterStudentsRoutes = func(router *mux.Router) {
router.HandleFunc("/student/",
controllers.CreateStudent).Methods("POST")
router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")
router.HandleFunc("/student/{studentId}",
controllers.GetStudentById).Methods("GET")
router.HandleFunc("/student/{studentId}",
controllers.UpdateStudent).Methods("PUT")
router.HandleFunc("/student/{studentId}",
controllers.DeleteStudent).Methods("DELETE")
}
