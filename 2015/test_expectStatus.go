package forest

import "testing"

func TestExpectStatus404(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/status/404"))
	ExpectStatus(t, r, 404)
}
