package simplechat

import (
	"fmt"
	"gorm.io/gorm"
)

type AddUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// FindUserByID finds a user by their id.
func FindUserByID(id uint) (_ *User, resp *gorm.DB, err error) {
	user := User{}
	resp = db.First(&user, id)
	if resp.Error != nil {
		err = fmt.Errorf("unable to find user in DB with ID=%d; %w", id, resp.Error)
	}
	return &user, resp, err
}

// FindUserByEmail finds a user by their email.
func FindUserByEmail(email string) (_ *User, resp *gorm.DB, err error) {
	var user User
	resp = db.Where("email = ?", email).First(&user)
	if resp.Error != nil {
		err = fmt.Errorf("unable to find user '%s' in DB; %w", email, resp.Error)
	}
	return &user, resp, err
}

// FetchAllUsers fetches all users
func FetchAllUsers() (_ *[]User, resp *gorm.DB, err error) {
	var users []User
	resp = db.Find(&users)
	if resp.Error != nil {
		err = fmt.Errorf("unable to retrieve user from in DB; %w", resp.Error)
	}
	return &users, resp, err
}

// DeleteUserByID deletes one user given their id
func DeleteUserByID(id uint) (err error) {
	user, _, err := FindUserByID(id)
	if err != nil {
		goto end
	}
	err = user.DeleteUser()
	if err != nil {
		goto end
	}
end:
	if err != nil {
		err = fmt.Errorf("unable to delete user ID='%d' from DB; %w", id, err)
	}
	return err
}

// UpdateUserByID updates user's email and name.
func UpdateUserByID(id uint, email string, name string) (err error) {
	user, _, err := FindUserByID(id)
	if err != nil {
		goto end
	}
	err = user.Update(email, name)
end:
	if err != nil {
		err = fmt.Errorf("unable to update user for ID='%d'; %w",
			id, err)
	}
	return err
}

// AddUser adds a new user object.
func AddUser(email string, name string, password string) (_ *User, err error) {
	var pwdHash string

	user := User{Email: email, Name: name}
	resp := db.Save(&user)
	if resp.Error != nil {
		err = fmt.Errorf("unable to save user in DB; %w", resp.Error)
	}

	// Encrypt password and save it in user_creds table.
	pwdHash, err = HashPwd(password)
	if err != nil {
		goto end
	}

	// Save hashed password to DB.
	err = user.SaveCreds(pwdHash)
	if err != nil {
		goto end
	}

end:
	if err != nil {
		err = fmt.Errorf("unable to add user for email='%s'; %w", email, err)
		goto end
	}

	return &user, err
}
