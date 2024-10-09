package API

import (
	"encoding/json"
	"fmt"
	"log"

	// "math/rand"
	"net/http"
	"strconv"

	// "time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseName string  `json:"coursename"`
	CourseId   int     `json:"courseid"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"authorname"`
	Websites string `json:"website"`
}

// db courses slice
var courses []Course

// isEmpty checks if a Course struct has empty fields
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

// homepage setup
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my courses API. This is my homepage. </h1>"))
}

// var lastCourseId = 3 // becoz already 3 raw data provided -> continuing from there

// add a new course
func AddCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding new course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("plss send some data")
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid request payload")
		return
	}

	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data")
		return
	}

	// // randomly generate unique id
	// // append course into courses
	// rand.Seed(time.Now().UnixNano())
	// course.CourseId = rand.Intn(100)

	// lastCourseId++
	// course.CourseId = lastCourseId

	//assigning next course id
	maxCourseId := 0
	for _, existingCourse := range courses {
		if existingCourse.CourseId > maxCourseId {
			maxCourseId = existingCourse.CourseId
		}
	}

	// Assign the next available ID (maxCourseId + 1)
	course.CourseId = maxCourseId + 1

	// Add the course to the databases
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	// return // redundant code ->no need becoz it will already return
}

// Get all courses
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get one course by ID
func GetOneCourseById(w http.ResponseWriter, r *http.Request) {
	// Extract course ID from the URL path parameters
	fmt.Println("getting one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	// if err != nil {
	// 	json.NewEncoder(w).Encode("Course not found")
	// 	return
	// }

	// Find course by ID
	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	// return  // redundant code
}

// Update an existing course by ID
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	// Extract course ID from the URL
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// Find the course by ID
	for i, existingCourse := range courses {
		if existingCourse.CourseId == id {
			// Decode the incoming course data
			var updatedCourse Course
			err := json.NewDecoder(r.Body).Decode(&updatedCourse)
			if err != nil {
				json.NewEncoder(w).Encode("Invalid request payload")
				return
			}

			// Update fields only if they are provided, otherwise retain the old values
			if updatedCourse.CourseName != "" {
				existingCourse.CourseName = updatedCourse.CourseName
			}
			if updatedCourse.Price != 0 {
				existingCourse.Price = updatedCourse.Price
			}
			if updatedCourse.Author != nil {
				if updatedCourse.Author.FullName != "" {
					existingCourse.Author.FullName = updatedCourse.Author.FullName
				}
				if updatedCourse.Author.Websites != "" {
					existingCourse.Author.Websites = updatedCourse.Author.Websites
				}
			}

			// Update the course in the "database"
			courses[i] = existingCourse

			// Respond with the updated course
			json.NewEncoder(w).Encode(existingCourse)
			return
		}
	}

	// If the course is not found
	json.NewEncoder(w).Encode("Course ID not found")
}

// Delete a course by ID
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// Extract course ID from the URL path parameters
	fmt.Println("deleting courses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("id donot exist")
		return
	}

	// Find and delete the course
	for i, course := range courses {
		if course.CourseId == id {
			// Remove the course
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode("course deleted")
			return

		}
	}
}

// BuildApi sets up the routes and middleware
func TestApi() {
	// Add some sample data to courses
	courses = append(courses, Course{
		CourseName: "Go Programming",
		CourseId:   1,
		Price:      100,
		Author: &Author{
			FullName: "John Doe",
			Websites: "https://johndoe.com",
		},
	})

	courses = append(courses, Course{
		CourseName: "Web Development with React",
		CourseId:   2,
		Price:      200,
		Author: &Author{
			FullName: "Jane Smith",
			Websites: "https://janesmith.dev",
		},
	})

	courses = append(courses, Course{
		// CourseName: "Django",
		CourseId: 3,
		Price:    299,
		Author: &Author{
			FullName: "Rani Smiriti",
			Websites: "https://ranismriti.dev",
		},
	})

	// Create the mux router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/course", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourseById).Methods("GET")      // Get one course by ID
	r.HandleFunc("/create", AddCourse).Methods("POST")                 // Create a new course
	r.HandleFunc("/updatecourse/{id}", UpdateCourse).Methods("PUT")    // Update a course by ID
	r.HandleFunc("/deletecourse/{id}", DeleteCourse).Methods("DELETE") // Delete a course by ID

	// Start the server
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
