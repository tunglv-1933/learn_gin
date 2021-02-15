package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var courses []Course

func main() {
	router := gin.Default()
	courses = append(courses, Course{ID: "1", Title: "My first course", Body: "This is the content of my first course"})

	router.GET("/courses", func(c *gin.Context) {
		c.JSON(200, courses)
	})

	router.POST("/courses", func(c *gin.Context) {
		var course Course
		if err := c.ShouldBind(&course); err == nil {
			lastIDCourse, _ := strconv.ParseInt(courses[len(courses)-1].ID, 10, 16)
			course.ID = strconv.FormatInt(lastIDCourse+1, 16)
			courses = append(courses, course)
			c.JSON(200, course)
		} else {
			c.JSON(200, gin.H{"message": "error!"})
		}
		return
	})

	router.GET("/course/:id", func(c *gin.Context) {
		found := false

		for _, item := range courses {
			if item.ID == c.Param("id") {
				found = true
				c.JSON(200, item)
			}
		}

		if !found {
			c.JSON(200, gin.H{"message": "not found!"})
		}
		return
	})

	router.PUT("/course/:id", func(c *gin.Context) {
		found := false

		for index, item := range courses {
			if item.ID == c.Param("id") {
				found = true
				c.Bind(&item)
				courses[index] = item
				c.JSON(200, item)
			}
		}

		if !found {
			c.JSON(200, gin.H{"message": "not found!"})
		}
		return
	})

	router.DELETE("/course/:id", func(c *gin.Context) {
		found := false

		for index, item := range courses {
			if item.ID == c.Param("id") {
				found = true
				courses = append(courses[:index], courses[index+1:]...)
				c.JSON(200, gin.H{"message": "delete successfull"})
			}
		}

		if !found {
			c.JSON(200, gin.H{"message": "not found!"})
		}
		return
	})

	router.Run(":8000")
}
