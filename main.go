package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Faculty struct {
	ID          string   `json:"id"`
	FacultyName string   `json:"facultyName"`
	SubjectIDs  []string `json:"subjectIDs"`
}

type Course struct {
	CourseCode    string  `json:"courseCode"`
	CourseName    string  `json:"courseName"`
	Prerequisites string  `json:"prerequisites"`
	TutorIDs      []int64 `json:"tutorIDs"`
}

type Tutor struct {
	StudentID int64   `json:"studentID"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	GPA       float32 `json:"gpa"`
}

type Class struct {
	ID         string          `json:"_id"`
	CourseCode string          `json:"courseCode"`
	TutorID    string          `json:"tutorID"`
	Timetable  []ClassSchedule `json:"timetable"`
}

type ClassSchedule struct {
	ID        string `json:"_id"`
	DayOfWeek string `json:"dayOfWeek"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

var faculties []Faculty
var courses []Course
var classes []Class
var schedules []ClassSchedule
var tutors []Tutor

func getFaculties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(faculties)
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getTutors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tutors)
}

func getClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range courses {
		if item.CourseCode == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getTutorById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tutors {
		if strconv.Itoa(int(item.StudentID)) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getClassesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range classes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getSchedules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func createTutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTutor Tutor
	_ = json.NewDecoder(r.Body).Decode(&newTutor)
	tutors = append(tutors, newTutor)
	json.NewEncoder(w).Encode(newTutor)
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
}

func createClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClass Class
	_ = json.NewDecoder(r.Body).Decode(&newClass)
	classes = append(classes, newClass)
	json.NewEncoder(w).Encode(newClass)
}

func deleteTutor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tutors {
		if strconv.Itoa(int(item.StudentID)) == params["id"] {
			tutors = append(tutors[:index], tutors[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tutors)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseCode == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}

func deleteClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range classes {
		if item.ID == params["id"] {
			classes = append(classes[:index], classes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(classes)
}

func main() {
	r := mux.NewRouter()

	// Add initial data
	tutors = append(tutors, Tutor{
		StudentID: 6411271,
		FirstName: "Lynn Thit",
		LastName:  "Nyi Nyi",
		Email:     "lnyinyi22@gmail.com",
		GPA:       3.99,
	})

	tutors = append(tutors, Tutor{
		StudentID: 6411325,
		FirstName: "Chammy",
		LastName:  "Aung",
		Email:     "chammy95@gmail.com",
		GPA:       3.70,
	})

	courses = append(courses, Course{
		CourseCode:    "CSX1001",
		CourseName:    "Fundamentals of Programming",
		Prerequisites: "None",
		TutorIDs:      []int64{6411271},
	})

	courses = append(courses, Course{
		CourseCode:    "CSX3001",
		CourseName:    "Object-Oriented Programming",
		Prerequisites: "CSX1001",
		TutorIDs:      []int64{6411271},
	})

	courses = append(courses, Course{
		CourseCode:    "BBA1001",
		CourseName:    "Business Exploration",
		Prerequisites: "None",
		TutorIDs:      []int64{6411271, 6411325},
	})

	faculties = append(faculties, Faculty{
		ID:          "1",
		FacultyName: "VMS",
		SubjectIDs:  []string{"CSX001"},
	})

	faculties = append(faculties, Faculty{
		ID:          "2",
		FacultyName: "BBA",
		SubjectIDs:  []string{"BBA001"},
	})

	classes = append(classes, Class{
		ID:         "001",
		CourseCode: "CSX001",
		TutorID:    "6411271",
		Timetable: []ClassSchedule{
			{
				ID:        "0001",
				DayOfWeek: "Monday",
				StartTime: "09:00",
				EndTime:   "11:00",
			},
			{
				ID:        "0002",
				DayOfWeek: "Tuesday",
				StartTime: "09:00",
				EndTime:   "11:00",
			},
		},
	})

	classes = append(classes, Class{
		ID:         "002",
		CourseCode: "BBA001",
		TutorID:    "6411325",
		Timetable: []ClassSchedule{
			{
				ID:        "0003",
				DayOfWeek: "Wednesday",
				StartTime: "13:00",
				EndTime:   "14:00",
			},
			{
				ID:        "0004",
				DayOfWeek: "Thursday",
				StartTime: "13:00",
				EndTime:   "14:00",
			},
		},
	})

	// Get all Faculties, Courses, Tutors, Classes, and Schedules
	r.HandleFunc("/faculties", getFaculties).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/tutors", getTutors).Methods("GET")
	r.HandleFunc("/classes", getClasses).Methods("GET")
	r.HandleFunc("/schedules", getSchedules).Methods("GET")

	// Get Courses, Tutors, and Classes by Id
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/tutors/{id}", getTutorById).Methods("GET")
	r.HandleFunc("/classes/{id}", getClassesById).Methods("GET")

	// Create Tutor, Course, and Class
	r.HandleFunc("/tutors", createTutor).Methods("POST")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/classes", createClass).Methods("POST")

	// Delete Tutors, Courses, and Classes
	r.HandleFunc("/tutors", deleteTutor).Methods("DELETE")
	r.HandleFunc("/courses", deleteCourse).Methods("DELETE")
	r.HandleFunc("/classes", deleteClass).Methods("DELETE")

	// Update Faculties, Tutors, Courses, and Classes
	// todo: ...

	// // Server and Ports
	// fmt.Println("Server running on port 8080!")
	// log.Fatal(http.ListenAndServe(":8080", r))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default to port 8000 if PORT environment variable is not set
	}
	fmt.Println("Server running on" + port + "!")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
