package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type testData struct {
	ApiKey      string `json:"apiKey"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	CountryCode int    `json:"CountryCode"`
}

var data testData

func TestMain(m *testing.M) {
	buf, err := ioutil.ReadFile("test_data.json")
	if err != nil {
		exitWithMessage("Error occurred while setting up test data", 1)
	}

	if err := json.Unmarshal(buf, &data); err != nil {
		exitWithMessage("Unable to parse test data", 1)
	}

	if data.ApiKey == "" {
		exitWithMessage("API key is required to execute tests", 1)
	}

	if data.CountryCode == 0 || data.Email == "" || data.PhoneNumber == "" {
		exitWithMessage("OneTouch API requires a user with registered device", 1)
	}

	os.Exit(m.Run())
}

func exitWithMessage(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
