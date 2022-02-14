package user

import (
	"fmt"
	"simple-chat-app/server/src/daos"
	"simple-chat-app/server/src/models"
)

/**
Find a user by email.
*/
func FindByEmail(email string) (*models.User, error) {
	db := daos.GetDbConn()
	var user models.User
	resp := db.Where("email = ?", email).First(&user)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &user, nil
}

/**
Fetch all users.
*/
func FetchAll() *[]models.User {
	db := daos.GetDbConn()
	var users []models.User
	result := db.Find(&models.User{})
	fmt.Println(result.Error)

	return &users
}

/**
Add a new user.
*/
func AddOne(email string, name string) (*models.User, error) {
	db := daos.GetDbConn()
	newUser := models.User{Email: email, Name: name}
	resp := db.Save(&newUser)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &newUser, nil
}

/**
Delete one user.
*/
func DeleteOne(id uint) error {
	db := daos.GetDbConn()
	resp := db.Where("id = ?", id).Delete(&models.User{})
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
