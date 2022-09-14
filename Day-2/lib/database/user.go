package database

import (
	"Day-2/config"
	"Day-2/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func GetUserById(id uint) (models.User, error) {
	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func CreateNewUser(user models.User) error {
	// fmt.Println("Masuk db", user)
	// test := models.User{Email: "melati@falcon.com", Password: "1234"}
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUserById(user models.User, id uint) error {
	updatedUser, _ := GetUserById(id)
	updatedUser.Password = user.Password
	if err := config.DB.Save(&updatedUser).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserById(id uint) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
