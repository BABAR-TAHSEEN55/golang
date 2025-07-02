package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for Course
type Course struct {
	CourseId    string `json:"courseId"`
	CourseName  string `json:"courseName"`
	CoursePrice int    `json:"coursePrice"`
	Author      *Author
}
type Admin struct {
	Fullname string
}
type Author struct {
	Fullname string `json:"fullName"`
	Website  string `json:"website"`
}

// Fake Db
var Courses []Course

// Middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
	// Returns true if both are true
}

func main() {

	fmt.Println("Simple API in Golang")
	// Seeding
	Courses = append(Courses, Course{CourseId: "1", CourseName: "Golang", CoursePrice: 299, Author: &Author{Fullname: "Gopher", Website: "https://youtube.com"}})
	Courses = append(Courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 899, Author: &Author{Fullname: "NestJs", Website: "https://google.com"}})
	// Router
	r := mux.NewRouter()
	// Routing
	r.HandleFunc("/", ServeHome)
	r.HandleFunc("/course", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", AddCourses).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controller
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Custom API</h1>")) // It expects the byte format ( can't directly use string )
}
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses on the Course ")
	w.Header().Set("Content-Type", "application/json")
	// NewEncoder-> Creates JSON and writes it to HTTP Header and ENcode actually converts it
	json.NewEncoder(w).Encode(Courses)
}
func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Specific Course")
	params := mux.Vars(r) // Getting ID from the URL
	for _, course := range Courses {
		if course.CourseId == params["id"] { json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Courses Found")
	return
}
func AddCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding Courses...")
	// What if input is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide Input")
		return
	}
	// What if Input is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty Data")
		return
	}
	// Generating Course Id
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(Courses)
	return

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCourse ..")
	w.Header().Set("Content-Type", "application/json")

	// Grabbing Id to be Updated
	params := mux.Vars(r)
	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			var course Course
			//NOTE: : Whenever NewDecoder is used it needs &course so create a variable with suitable ref
			json.NewDecoder(r.Body).Decode(&course)
			// Overriding the ID
			// Adding New Data

			//FIX : I dont get this ? how Ik but stilL how ?
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}

	}

}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			// This is mainly used for JSON Responses
			json.NewEncoder(w).Encode(map[string]string{"message": "Deleted "})
			// This is used for Plain Text
			// w.Write([]byte("Deleted"))
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"error": "Not Found "})
	// w.Write([]byte("Not Found"))

}

//NOTE : Seeding means adding some mock data into your application when it starts for testing purposes , showcasing API Workarounds
