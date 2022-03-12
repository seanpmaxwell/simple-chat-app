package repos

import (
	"simple-chat-app/server/server/models"

	"gorm.io/gorm"
)

/**** Types ****/

// UserRepo Layer
type UserRepo struct {
	Db *gorm.DB
}


/**** Functions ****/

// Wire UserRepo
func WireUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// Find a user by their id.
func (u *UserRepo) FindById(id uint) (*models.User, error) {
	user := models.User{}
	resp := u.Db.First(&user, id)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &user, nil
}

// Find a user by email.
func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	resp := u.Db.Where("email = ?", email).First(&user)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &user, nil
}

// Fetch all users.
func (u *UserRepo) FetchAll() (*[]models.User, error) {
	var users []models.User
	// pick up here, check this
	resp := u.Db.Omit("UserCreds").Find(&users)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &users, nil
}

// Add a new user.
func (u *UserRepo) Add(email string, name string) (*models.User, error) {
	newUser := models.User{Email: email, Name: name}
	resp := u.Db.Save(&newUser)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &newUser, nil
}

// Update user's email and name.
func (u *UserRepo) Update(user *models.User, email string, name string) {
	u.Db.Model(user).Updates(models.User{Email: email, Name: name})
}

// Delete one user.
func (u *UserRepo) Delete(id uint) error {
	resp := u.Db.Unscoped().Where("id = ?", id).Delete(&models.User{})
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

// Fetch a user's hashed password
func (u *UserRepo) GetPwdHash(userId uint) (string, error) {
	var userCreds models.UserCreds
	resp := u.Db.Where("user_id = ?", userId).First(&userCreds)
	if resp.Error != nil {
		return "", resp.Error
	}
	return userCreds.Pwdhash, nil
}

// Create a user credentials row to store confidentials stuff.
func (u *UserRepo) SaveUserCreds(id uint, pwdHash string) error {
	userCreds := models.UserCreds{Pwdhash: pwdHash, UserID: id}
	resp := u.Db.Save(&userCreds)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
