package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
)

//suite setup
var boqs = forest.NewClient("http://localhost:9999", new(http.Client))

//default headers
func withBoqsConfig(staticPath string) *forest.RequestConfig {
	return forest.NewConfig(staticPath).
		Header("Content-Type", "application/json").
		Header("Accept", "application/json")
}

// START OMIT

func Test_create_topic_with_a_name_greater_then_255_characters(t *testing.T) {
	Given(t, "I want to create a new topic with a name > 255 characters")
	nameWith256Characters := randStr(256, "alphanum")

	When(t, "I create a new topic with that name")
	r := boqs.PUT(t, withBoqsConfig(fmt.Sprintf("/v1/topics/%s", nameWith256Characters)))

	Then(t, "I expect the correct error code and status")
	forest.ExpectStatus(t, r, http.StatusOK)
	ExpectErrorCode(t, r, 3016)
}

// END OMIT
