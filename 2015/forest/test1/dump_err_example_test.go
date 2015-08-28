package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
)

// START OMIT

//suite setup
var boqs = forest.NewClient("http://localhost:9999", new(http.Client))

func withBoqsConfig(staticPath string) *forest.RequestConfig {
	return forest.NewConfig(staticPath).
		Header("Content-Type", "application/json").
		Header("Accept", "application/json")
}

func Test_create_topic(t *testing.T) {
	//setup
	topicName := "golang"

	//create topic
	r := boqs.PUT(t, withBoqsConfig(fmt.Sprintf("/v1/topics/%s", topicName))) // HL

	//expect status no content
	forest.ExpectStatus(t, r, http.StatusNoContent) // HL
}

// END OMIT
