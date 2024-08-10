package model

import (
	"errors"
	"quiz-design/pkg/utils"

	"github.com/google/uuid"
)

type Options struct {
	Option []string `json:"options"`
}

type Question struct {
	CourseID     string             `json:"course_id"`
	QuizID       string             `json:"quiz_id"`
	QuestionID   string             `json:"question_id"`
	QuestionText string             `json:"text"`
	Answers      Options            `json:"answers"`
	Response     map[string]Options `json:"response"` // {user_id : responses}
	Score        int                `json:"score"`
}

func CreateQuestion(answers Options, courseID, quizID, questionText string) *Question {
	return &Question{
		QuestionID:   uuid.New().String(),
		CourseID:     courseID,
		QuizID:       quizID,
		QuestionText: questionText,
		Answers:      answers,
	}
}

func (q *Question) SubmitResponse(userID string, response Options) error {
	if _, ok := q.Response[userID]; !ok {
		return errors.New("user already submitted response for this quetsion")
	}
	q.Response[userID] = response
	return nil
}

func (q *Question) CheckAnswer(response Options) bool {
	return utils.CompareSlice(response.Option, q.Answers.Option)
}
