package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AppDetails encapsulates the response from the Authy API when requesting an app details.
type AppDetails struct {
	HTTPResponse    *http.Response
	ID              int    `json:"app_id"`
	Name            string `json:"name"`
	Plan            string `json:"plan"`
	SMSEnabled      bool   `json:"sms_enabled"`
	WhiteLabel      bool   `json:"white_label"`
	OneTouchEnabled bool   `json:"onetouch_enabled"`
}

// NewAppDetails receives a http request, parses the body and return an instance of AppDetails
func NewAppDetails(response *http.Response) (*AppDetails, error) {
	appDetails := &AppDetails{HTTPResponse: response}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &struct {
		App *AppDetails `json:"app"`
	}{
		App: appDetails,
	})
	if err != nil {
		return nil, err
	}

	return appDetails, nil
}
