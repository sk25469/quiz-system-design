package service

import "quiz-design/pkg/model"

type UserService struct {
	Users map[string]*model.User
}

func CreateNewUserService() *UserService {
	return &UserService{
		Users: make(map[string]*model.User),
	}
}

func (us *UserService) InsertUser(name string, email string, userType model.UserType) {
	// TODO: check if email exists
	user := model.RegisterUser(int(userType), name, email)
	us.Users[user.ID] = user
}

func (us *UserService) RegisterWithOAuth() {

}

func (us *UserService) GetScore(userID string) int {
	score := us.Users[userID].Score
	return score
}
