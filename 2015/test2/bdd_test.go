package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
)

//// flag variables
//var server = flag.String("server", "localhost", "the server against which the tests are run")
//var port = flag.String("port", "49011", "the port against which the tests are run")
//var boqs *forest.APITesting

//func init() {
//	flag.Parse()
//	boqs = forest.NewClient(fmt.Sprintf("http://%s:%s", *server, *port), new(http.Client))
//}

//suite setup
var boqs = forest.NewClient(("http://localhost:8060"), new(http.Client))

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
	forest.ExpectStatus(t, r, http.StatusBadRequest)
	ExpectErrorCode(t, r, 3016)
}

// END OMIT
