package controller

import (
	"errors"
	"quiz-design/pkg/model"
	"quiz-design/pkg/service"
)

type CourseController struct {
	courseService service.CourseService
}

func CreateCourseController() *CourseController {
	return &CourseController{
		courseService: *service.CreateCourseService(),
	}
}

func (cc *CourseController) CreateCourse(userId string, course model.Course) {
	// TODO: authenticate here to check if admin is creating
	cc.courseService.InsertCourse(course)
}

func (cc *CourseController) CreateQuiz(userID string, courseID string, quiz model.Quiz) error {
	// TODO: authenticate here to check if admin is creating
	course, err := cc.courseService.GetCourse(courseID)
	if err != nil {
		return errors.New("cannot find course with the id")
	}
	course.InsertQuiz(quiz)
	return nil
}

func (cc *CourseController) AddQuestion(courseID string, quizID string, question model.Question) error {
	course, err := cc.courseService.GetCourse(courseID)
	if err != nil {
		return errors.New("cannot find course with the id")
	}
	quiz, err := course.GetQuiz(quizID)
	if err != nil {
		return errors.New("cannot find quiz with the id")
	}
	quiz.InsertQuestion(question)
	return nil
}

func (cc *CourseController) SubmitResponse(courseID string, userID string, quizID string, questionID string, response model.Options) error {
	course, err := cc.courseService.GetCourse(courseID)
	if err != nil {
		return errors.New("cannot find course with the id")
	}
	quiz, err := course.GetQuiz(quizID)
	if err != nil {
		return errors.New("cannot find quiz with the id")
	}
	quiz.SubmitAnswerToQuiz(questionID, userID, response)
	return nil
}

func (cc *CourseController) GetScore(courseID string, quizID string) (map[string]int, error) {
	course, err := cc.courseService.GetCourse(courseID)
	if err != nil {
		return map[string]int{}, errors.New("cannot find course with the id")
	}
	quiz, err := course.GetQuiz(quizID)
	if err != nil {
		return map[string]int{}, errors.New("cannot find quiz with the id")
	}
	return quiz.Scores, nil
}
