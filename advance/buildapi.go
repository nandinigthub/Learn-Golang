package advance

import (
	"fmt"
)

type Course struct {
	CourseName string `json:"coursename"`
	CourseId   int    `json:"courseid"`
	Price      int    `json:"price"`
	// pasing custom type as pointer
	Author *Author `json:"author"`
}

type Author struct {
	FullName string `json:"authorname"`
	Websites string `json:"website"`
}

// fake db

var courses []Course

// // creating a helper which will help apply condition
// func (c *Course) isEmpty() bool {
// 	return c.CourseName == "" && c.CourseId == 0
// }

func BuildApi() {
	fmt.Println("my first api")

	// Add some initial data to the courses slice -> seeding
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
		// CourseId:   3,
		Price: 299,
		Author: &Author{
			FullName: "Rani Smiriti",
			Websites: "https://ranismriti.dev",
		},
	})

}
