package service

import (
	"errors"
	"quiz-design/pkg/model"
)

type CourseService struct {
	Courses map[string]model.Course
}

func CreateCourseService() *CourseService {
	return &CourseService{
		Courses: make(map[string]model.Course),
	}
}

func (cs *CourseService) GetCourse(courseID string) (*model.Course, error) {
	course, ok := cs.Courses[courseID]
	if !ok {
		return &model.Course{}, errors.New("cannot find course with that id")
	}

	return &course, nil
}

func (cs *CourseService) InsertCourse(course model.Course) {
	cs.Courses[course.CourseID] = course
}
