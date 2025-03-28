package db

import (
	"encoding/json"
	"os"
)

func LoadCourses() (c map[string]Course, err error) {
	data, err := os.ReadFile("courses.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return
	}

	for k, course := range c {
		for i, a := range course.Articles {
			c[k].Articles[i].Markdown, err = a.RenderHTML()
			if err != nil {
				return
			}
		}
	}

	return
}

type Course struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Articles []Article `json:"articles"`
	Lessons  []Lesson  `json:"lessons"`
}
