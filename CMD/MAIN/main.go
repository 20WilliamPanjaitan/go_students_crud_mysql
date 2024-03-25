package main
import (
"fmt"
"log"
"net/http"
"github.com/gorilla/mux"
"github.com/jinzhu/gorm/dialects/mysql"
"github.com/rudychandra/go_students_crud_mysql/pkg/routes"
)
func main() {
r := mux.NewRouter()
routes.RegisterStudentsRoutes(r)
http.Handle("/", r)
fmt.Print("Starting Server localhost:9010")
log.Fatal(http.ListenAndServe("localhost:9010", r))
}
