package model

import (
	"errors"

	"github.com/google/uuid"
)

type Course struct {
	CourseID   string `json:"course_id"`
	Quizes     map[string]Quiz
	CourseName string
}

func CreateNewCourse(courseName string) *Course {
	return &Course{
		CourseID:   uuid.New().String(),
		CourseName: courseName,
	}
}

func (course *Course) InsertQuiz(quiz Quiz) {
	course.Quizes[quiz.QuizID] = quiz
}

func (course *Course) GetQuiz(quizID string) (*Quiz, error) {
	quiz, ok := course.Quizes[quizID]
	if !ok {
		return &Quiz{}, errors.New("cannot find quiz with given id")
	}
	return &quiz, nil
}
