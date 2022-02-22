package simplechat

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `json:"email" gorm:"unique;size:255;not null"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	UserCreds UserCreds `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}

type UserCreds struct {
	gorm.Model
	Pwdhash string `gorm:"size:255;not null"`
	UserID  uint
}

// GetPwdHash fetches a user's hashed password by their ID
func (user *User) GetPwdHash() (_ string, err error) {
	resp := db.First(&user.UserCreds)
	if resp.Error != nil {
		err = fmt.Errorf("unable to retrieve user creds from DB; %w", resp.Error)
		goto end
	}
end:
	if err != nil {
		err = fmt.Errorf("unable to get password hash for user '%s'; %w", user.Name, err)
	}
	return user.UserCreds.Pwdhash, err
}

// SaveCreds creates a user credentials row to store user's password hash
func (user *User) SaveCreds(pwdHash string) (err error) {
	user.UserCreds.UserID = user.ID
	user.UserCreds.Pwdhash = pwdHash
	resp := db.Save(&user.UserCreds)
	if resp.Error != nil {
		err = fmt.Errorf("unable to save password hash for user '%s'; %w",
			user.Email, resp.Error)
	}
	return resp.Error
}

// DeleteUser deletes a user
func (user *User) DeleteUser() (err error) {
	resp := db.Model(user).Delete(&User{})
	if resp.Error != nil {
		err = fmt.Errorf("unable to delete user '%s'; %w", user.Email, resp.Error)
	}
	return err
}

// Update updates a user's email and name.
func (user *User) Update(email string, name string) (err error) {
	resp := db.Model(user).Updates(User{Email: email, Name: name})
	if resp.Error != nil {
		err = fmt.Errorf("unable to update user '%s'; %w", email, resp.Error)
	}
	return err
}
