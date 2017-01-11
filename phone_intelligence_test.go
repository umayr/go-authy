package authy

import (
	"testing"
	"net/url"
)

func Test_PhoneInformation(t *testing.T) {
	api := NewSandboxAuthyAPI("bf12974d70818a08199d17d5e2bae630")
	info, err := api.PhoneInformation(1, "7754615609", url.Values{})
	if err != nil {
		t.Error("Internal error occurred")
	}

	if !info.Success {
		t.Error("Couldn't fetch phone information")
	}
}