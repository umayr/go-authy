package authy

import (
	"net/url"
	"testing"
)

func Test_TransformParams(t *testing.T) {
	str := transformParams(url.Values{
		"foo": []string{"unicorn", "rainbow"},
		"bar": []string{"pony"},
	})

	if str != "bar=pony&foo[]=unicorn&foo[]=rainbow" {
		t.Error("Unable to propery parse url params into a sorted string")
	}
}
