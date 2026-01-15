package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - go to another file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
	
}
type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
	
}

//fake database
var courses []Course

//middleware , helper - file
//to check courseId and courseName is not empty
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" 
}


func main(){
	fmt.Println("welcome to build api")
	r := mux.NewRouter()
	r.HandleFunc("/",serveHome)
}

//controllers - file
// serve home route 
func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to api schema<h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/type")
	params := mux.Vars(r)
	//loop through courses , find matching id 
	for _,course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course available with provided id")
	return 
}

func createOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("adding the course")
	w.Header().Set("Content-Type","application/json")
	//what if: body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	//what about data something like this -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty(){
		json.NewEncoder(w).Encode("No data inside the body")
		return
	}
	//generate a new unique id, string
	//append course into courses
	
	course.CourseId = strconv.FormatInt(time.Now().UnixNano(),10)
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("update the course")
	w.Header().Set("Content-Type","application/json")
	//first grad id from request 
	params := mux.Vars(r)
	//loop ,id , remove then add operation with myID(same id as we get)
	for index,course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)
			return

		}
	}
	//TODO : send a responce where id is not found
}