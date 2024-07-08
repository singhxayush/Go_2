package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/// Models

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

/// Fake DB

var courses []Course

func main() {
	fmt.Printf("-------\n| API |\n-------\n")

	router := mux.NewRouter()

	// seeding
	courses = append(courses, Course{ CourseId: "1023940239", CourseName: "Golang", CoursePrice: 123, Author: &Author{Fullname: "Ayush", Website: "go.dev"}})
	courses = append(courses, Course{ CourseId: "1023120239", CourseName: "Dont Go lang", CoursePrice: 321, Author: &Author{Fullname: "Who?", Website: "dontgo.dev"}})

	// routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")

	// listen to a port
	log.Fatal(http.ListenAndServe(":8000", router))

}

/// Routes

// Serve Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// find the course with id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Invalid ID")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST One Course")
	w.Header().Set("Conetent-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Request")
	}

	// handling {} request -> Empty JSON
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty JSON Request")
		return
	}

	// genrata UID, convert it to string
	nanoseconds := time.Now().UnixNano()
	course.CourseId = strconv.Itoa(int(nanoseconds))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE One Course")
	w.Header().Set("Conetent-Type", "application/json")

	// first - get the ID from the request
	params := mux.Vars(r) // Getting key -> value pair from the params

	// loog through the Object(ID) and find the one to be updated and update
	for index, course := range courses {
		if course.CourseId == params["id"] { // exracting the value against the key "id" which contains the courseId of the course to be updated
			courses = append(courses[:index], courses[index+1:]...) // Deleteing that courseID
			var course Course // Defining a new course as a reference
			_ = json.NewDecoder(r.Body).Decode(&course) // Decoding the request body to get information to be updated into the created reference
			course.CourseId = params["id"] // Updating the reference with the x
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

/// MiddleWare

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
