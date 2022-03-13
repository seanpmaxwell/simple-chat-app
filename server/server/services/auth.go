package services

import (
	"errors"
	"simple-chat-app/server/server/models"
	"simple-chat-app/server/server/repos"
	"simple-chat-app/server/server/util"
	"time"
)

// **** Vals **** //

const (
	checkPwdFailed = "password verification failed"
)


/**** Types ****/

// AuthService layer
type AuthService struct {
	UserRepo *repos.UserRepo
	PwdUtil  *util.PwdUtil
}


/**** Functions ****/

// Wire AuthService
func WireAuthService(userRepo *repos.UserRepo, pwdUtil *util.PwdUtil) *AuthService {
	return &AuthService{userRepo, pwdUtil}
}

// Verify user credentials
func (a *AuthService) VerifyAndFetchUser(email string, password string) (*models.User, error) {
	// Search for the user
	user, err := a.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	// Fetch the pwd hash
	pwdHash, err := a.UserRepo.GetPwdHash(user.ID)
	if err != nil {
		return nil, err
	}
	// Compare the password to the hash. Wait 500 milliseconds if it failed as a security measure.
	passed := a.PwdUtil.Verify(pwdHash, password)
	if !passed {
		time.Sleep(time.Millisecond * 500)
		return nil, errors.New(checkPwdFailed)
	}
	// Return
	return user, nil
}
