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
	return
}

type Course struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Articles []Article `json:"articles"`
	Lessons  []Lesson  `json:"lessons"`
}
