package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// UserActivityType is activity type that needs to be sent to API
type UserActivityType string

const (
	// PasswordReset is a user activity type that wraps `password_reset`
	PasswordReset UserActivityType = "password_reset"
	// Banned is a user activity type that wraps `banned`
	Banned UserActivityType = "banned"
	// Unbanned is a user activity type that wraps `unbanned`
	Unbanned UserActivityType = "unbanned"
	// CookieLogin is a user activity type that wraps `cookie_login`
	CookieLogin UserActivityType = "cookie_login"
)

// User is an Authy User
type User struct {
	HTTPResponse *http.Response
	ID           string
	UserData     struct {
		ID int `json:"id"`
	} `json:"user"`
	Errors  map[string]string `json:"errors"`
	Message string            `json:"message"`
}

// UserStatus is a user with information loaded from Authy API
type UserStatus struct {
	HTTPResponse *http.Response
	ID           string
	StatusData   struct {
		ID          int      `json:"authy_id"`
		Confirmed   bool     `json:"confirmed"`
		Registered  bool     `json:"registered"`
		Country     int      `json:"country_code"`
		PhoneNumber string   `json:"phone_number"`
		Devices     []string `json:"devices"`
	} `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// UserActivity is a user's activity with message from Authy API
type UserActivity struct {
	HTTPResponse *http.Response
	Message      string `json:"message"`
	Success      bool   `json:"success"`
}

// NewUser returns an instance of User
func NewUser(httpResponse *http.Response) (*User, error) {
	userResponse := &User{HTTPResponse: httpResponse}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return userResponse, err
	}

	err = json.Unmarshal(body, userResponse)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return userResponse, err
	}

	userResponse.ID = strconv.Itoa(userResponse.UserData.ID)
	return userResponse, nil
}

// NewStatus returns an instance of UserStatus
func NewUserStatus(httpResponse *http.Response) (*UserStatus, error) {
	statusResponse := &UserStatus{HTTPResponse: httpResponse}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		Logger.Println("Error reading from API:", err)
		return statusResponse, err
	}

	err = json.Unmarshal(body, statusResponse)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return statusResponse, err
	}

	statusResponse.ID = strconv.Itoa(statusResponse.StatusData.ID)
	return statusResponse, nil
}

// NewUserActivity returns an instance of UserActivity
func NewUserActivity(httpResponse *http.Response) (*UserActivity, error) {
	userActivityResponse := &UserActivity{HTTPResponse: httpResponse}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return userActivityResponse, err
	}

	err = json.Unmarshal(body, userActivityResponse)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return userActivityResponse, err
	}

	return userActivityResponse, nil
}

// Valid returns true if the user was created successfully
func (response *User) Valid() bool {
	if response.HTTPResponse.StatusCode != 200 {
		return false
	}

	return true
}
