import (
	"fmt"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
)

// START OMIT
func Test_create_topic_with_a_name_greater_then_255_characters(t *testing.T) {
	Given(t, "I want to create a new topic with a name > 255 characters") // HL
	nameWith256Characters := randStr(256, "alphanum")

	When(t, "I create a new topic with that name") // HL
	r := boqs.PUT(t, withBoqsConfig(fmt.Sprintf("/v1/topics/%s", nameWith256Characters)))

	Then(t, "I expect the correct error code and status") // HL
	forest.ExpectStatus(t, r, http.StatusBadRequest)
}

// END OMIT
