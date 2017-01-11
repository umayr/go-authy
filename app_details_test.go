package authy

import (
	"testing"
	"net/url"
)

func TestAppDetails(t *testing.T) {
	api := NewSandboxAuthyAPI("d57d919d11e6b221c9bf6f7c882028f9")
	appDetails, err := api.AppDetails(url.Values{})
	if err != nil {
		t.Error("External error found", err)
	}

	if appDetails.Name == "" {
		t.Error("Name field is empty")
	}

	if appDetails.Plan == "" {
		t.Error("Plan field is empty")
	}
}