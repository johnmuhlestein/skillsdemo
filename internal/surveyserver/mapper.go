package surveyserver

import (
	pb "skillsdemo/rpc/survey"
	"skillsdemo/ent"

)

func mapDbFeedback(feedbackEnt *ent.Feedback) *pb.FeedbackResp {
	feedbackResp := pb.FeedbackResp{}
	feedbackResp.Id = feedbackEnt.ID.String()
	feedbackResp.Status = feedbackEnt.Status

	for _, v := range feedbackEnt.Edges.Responses {
		respStatus := "new"
		if &v.AnsweredTime != nil {
			respStatus = "complete"
		}
		feedbackResp.Responses = append(feedbackResp.Responses, &pb.FbResponses{
			Id: v.ID.String(), ParsedTemplate: v.ParsedTemplate, SurveyIndex: int32(v.PromptIndex), Status: respStatus,
		    })
	} 

	return &feedbackResp
}

