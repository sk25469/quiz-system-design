package model

import (
	"errors"

	"github.com/google/uuid"
)

type Quiz struct {
	CourseID  []string            `json:"course_id"`
	QuizID    string              `json:"quiz_id"`
	Questions map[string]Question `json:"questions"`
	Scores    map[string]int      // check the score of a user for this quiz
}

func CreateNewQuiz(courseID []string) *Quiz {
	return &Quiz{
		CourseID:  courseID,
		QuizID:    uuid.New().String(),
		Questions: make(map[string]Question),
		Scores:    make(map[string]int),
	}
}

func (quiz *Quiz) GetQuestion(questionID string) (*Question, error) {
	ques, ok := quiz.Questions[questionID]
	if !ok {
		return &Question{}, errors.New("cannot find the question with given id")
	}
	return &ques, nil
}

func (quiz *Quiz) InsertQuestion(question Question) {
	quiz.Questions[question.QuestionID] = question
}

func (quiz *Quiz) GetScoreForUser(userID string) (int, error) {
	score, ok := quiz.Scores[userID]
	if !ok {
		return 0, errors.New("cannot find the user for this quiz")
	}
	return score, nil
}

func (quiz *Quiz) AssignQuiz(courseID string) {
	quiz.CourseID = append(quiz.CourseID, courseID)
}

func (quiz *Quiz) SubmitAnswerToQuiz(questionID string, userID string, response Options) error {
	ques, err := quiz.GetQuestion(questionID)
	if err != nil {
		return errors.New("cannot find ques with given id")
	}
	// TODO: check if the answer is correct{}
	currScore := quiz.Scores[userID] + ques.Score
	quiz.Scores[userID] = currScore

	ques.SubmitResponse(userID, response)
	return nil
}
