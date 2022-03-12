package services

import (
	"simple-chat-app/server/server/models"
	"simple-chat-app/server/server/repos"
	"simple-chat-app/server/server/util"
)

/**** Functions ****/

// UserService Layer
type UserService struct {
	UserRepo *repos.UserRepo
	PwdUtil  *util.PwdUtil
}


/**** Functions ****/

// Wire UserService
func WireUserService(userRepo *repos.UserRepo, pwdUtil *util.PwdUtil) *UserService {
	return &UserService{userRepo, pwdUtil}
}

// Fetch all users.
func (u *UserService) FetchAll() (*[]models.User, error) {
	return u.UserRepo.FetchAll()
}

// Add a new user object.
func (u *UserService) Add(email string, name string, password string) error {
	// Save the user
	user, err := u.UserRepo.Add(email, name)
	if err != nil {
		return err
	}
	// Ecrypt password and save it in user_creds table.
	pwdHash, err := u.PwdUtil.Hash(password)
	if err != nil {
		return err
	}
	err = u.UserRepo.SaveUserCreds(user.ID, pwdHash)
	if err != nil {
		return err
	}
	return nil
}

// Update user's email and name.
func (u *UserService) Update(id uint, email string, name string) error {
	user, err := u.UserRepo.FindById(id)
	if err != nil {
		return err
	}
	u.UserRepo.Update(user, email, name)
	return nil
}

// Delete one user
func (u *UserService) Delete(id uint) error {
	return u.UserRepo.Delete(id)
}
