package surveyserver

import (
	"skillsdemo/ent"
	"testing"

	"github.com/google/uuid"
)

func TestMapDbFeedback(t *testing.T) {
	feedbackEnt := ent.Feedback{ID: uuid.New(), Status: "created",}
	feedbackResp := mapDbFeedback(&feedbackEnt)
	if feedbackResp.Status != "created" {
		t.Errorf("got %s, wanted %s", feedbackResp.Status, "created")
	} 
}