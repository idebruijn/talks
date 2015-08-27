import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
)

// START OMIT

func Test_create_topic_with_a_name_greater_then_255_characters(t *testing.T) {
	//setup
	topicName := randStr(256, "alphanum")

	//create topic
	r := boqs.PUT(t, withBoqsConfig(fmt.Sprintf("/v1/topics/%s", topicName)))
	code := forest.JSONPath(t, r, ".code")   // HL
	msg := forest.JSONPath(t, r, ".message") // HL

	//expect status no content
	forest.ExpectStatus(t, r, http.StatusNoContent)
}

// END OMIT
