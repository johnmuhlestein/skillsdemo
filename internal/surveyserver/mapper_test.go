package surveyserver

import (
	"skillsdemo/ent"
	"skillsdemo/ent/schema"
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

func TestMapDbMeasureEnumeration(t *testing.T) {
	dbEnums := []schema.MeasureEnum{{Label: "FirstThing", Value:1},{Label:"SecondThing",Value:2}}
	pbEnums := mapDbMeasureEnumeration(dbEnums)
	if len(pbEnums) != len(dbEnums) {
		t.Errorf("enum mapper did not return the array of the correct lenght expected %d got %d", len(dbEnums), len(pbEnums))
	}
}

func TestMapDbMeasureFlags(t *testing.T) {
	dbFlags := []string{"yellow","green","blue","red","purple"}
	pbFlags := mapDbMeasureFlags(dbFlags)
	if len(dbFlags) != len(pbFlags) {
		t.Errorf("flag mapper did not return the array with tht correct lenght. expected %d got %d", len(dbFlags), len(pbFlags))
	}
}