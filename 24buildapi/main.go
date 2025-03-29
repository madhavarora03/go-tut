package main

// Model for course ---> file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
