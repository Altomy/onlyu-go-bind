package onlyugobind

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name" `
	Phone        string `json:"phone"`
	Country      string `json:"country"`
	MobileNumber string `json:"mobileNumber"`
	CallingCode  string `json:"callingCode"`
	Avatar       string `json:"avatar"`
	ServerID     string `json:"serverID"`
	Token        string `json:"token"`
}

func StoreUser(payload string) (string, error) {
	//
	var oldUser User
	checkHaveError := db.First(&oldUser).Error
	if checkHaveError == nil {
		err := db.Delete(&oldUser).Error
		if err != nil {
			return "Error", err
		}
	}

	var user User
	json.Unmarshal([]byte(payload), &user)

	// user := User{Name: "ahmed", Phone: "+96278", Country: "Jordan", MobileNumber: "7878", Avatar: "", ServerID: "324"}

	insertError := db.Create(&user).Error
	if insertError != nil {
		return "Error", insertError
	}

	return "Created", nil
}

func CheckUser(payload string) (string, error) {
	var users User

	fetchError := db.First(&users).Error
	if fetchError != nil {
		return "Error", fetchError
	}

	result, _ := json.Marshal(users)

	return string(result), nil
}
