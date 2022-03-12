package util

import "golang.org/x/crypto/bcrypt"

/**** Types ****/

// PwdUtil layer
type PwdUtil struct{}


/**** Functions ****/

// Wire PwdUtil
func WirePwdUtil() *PwdUtil {
	return &PwdUtil{}
}

// Generate a hash from a password.
func (p *PwdUtil) Hash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Check password against hash.
func (p *PwdUtil) Verify(pwdHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password))
	return err == nil
}
