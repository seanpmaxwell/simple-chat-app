package simplechat

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrCheckPwdFailed = errors.New("password verification failed")
)

// VerifyAndFetchUser fetches user and verifies against email and password
func VerifyAndFetchUser(email string, password string) (*User, error) {
	var pwdHash string
	var passed bool

	// Search for the user
	user, _, err := FindUserByEmail(email)
	if err != nil {
		goto end
	}

	// Fetch the pwd hash
	pwdHash, err = user.GetPwdHash()
	if err != nil {
		goto end
	}

	// Compare the password to the hash.
	passed = CheckPwd(pwdHash, password)
	if !passed {
		// Wait 500 milliseconds if it failed as a security measure.
		time.Sleep(time.Millisecond * 500)
		err = ErrCheckPwdFailed
		goto end
	}

end:
	if err != nil {
		err = fmt.Errorf("unable to validate and/or fetch user '%s'; %w", email, err)
	}
	return user, err
}

// HashPwd generates a hash from a password.
func HashPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("unable to hash password; %w", err)
	}
	return string(hash), err
}

// CheckPwd checks password against hash.
func CheckPwd(pwdHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password))
	if err != nil {
		err = fmt.Errorf("password does not match; %w", err)
	}
	// TODO Consider convert this function to returning an error
	//      and provide more information on the reasons for failure.
	return err == nil
}
