package surveyserver

import (
	"skillsdemo/ent"
	"skillsdemo/ent/schema"
	pb "skillsdemo/rpc/survey"
)

func mapDbFeedback(feedbackEnt *ent.Feedback) *pb.FeedbackResp {
	feedbackResp := pb.FeedbackResp{Id: feedbackEnt.ID.String(), Status: feedbackEnt.Status}

	for _, v := range feedbackEnt.Edges.Responses {
		respStatus := "new"
		if v.AnsweredTime != nil {
			respStatus = "complete"
		}
		promptResponse := pb.FbResponse{Id: v.ID.String(), ParsedTemplate: v.ParsedTemplate, SurveyIndex: int32(v.PromptIndex), Status: respStatus}
		
		for _, p := range feedbackEnt.Edges.Survey.Edges.Prompts {
			if p.SortOrder == v.PromptIndex {
				switch p.ResponseType.Type {
				case "freeform":
					promptResponse.Type = pb.MeasureType_freeform
					if (v.FreeformValue == nil) {
						promptResponse.FreeformValue = ""
					} else {
						promptResponse.FreeformValue = *v.FreeformValue
					}
				case "range":
					promptResponse.Type = pb.MeasureType_range
					if (v.RangeValue == nil) {
						promptResponse.RangeValue = 0
					} else {
						promptResponse.RangeValue = int32(*v.RangeValue)
					}
					promptResponse.Range = &pb.MeasureRange{HighRange: int32(p.ResponseType.HighRange), LowRange: int32(p.ResponseType.LowRange)}
				case "boolean":
					promptResponse.Type = pb.MeasureType_boolean
					if (v.BoolValue == nil) {
						promptResponse.BooleanValue = ""
					} else {
						promptResponse.BooleanValue = *v.BoolValue
					}
					promptResponse.Boolean = &pb.MeasureBoolean{BoolTrue: p.ResponseType.BooleanTrue, BoolFalse: p.ResponseType.BooleanFalse}
				case "enumeration":
					promptResponse.Type = pb.MeasureType_enumeration
					promptResponse.EnumerationLabelValue = v.EnumValue.Label
					promptResponse.EnumerationValue = int32(v.EnumValue.Value)
					promptResponse.Enumeration = mapDbMeasureEnumeration(p.ResponseType.Enumeration)
				case "multiflag":
					promptResponse.Type = pb.MeasureType_multiflag
					promptResponse.FlagValues = v.LabelValues
					promptResponse.Flags = mapDbMeasureFlags(p.ResponseType.Flags)
				}			
			}
		}
		feedbackResp.Responses = append(feedbackResp.Responses, &promptResponse)
	} 

	return &feedbackResp
}

func mapDbFeedbackPrompt(fbResp *ent.PromptResponse, sPrompt *ent.Prompt) *pb.FbResponse {
	respStatus := "new"
	if fbResp.AnsweredTime != nil {
		respStatus = "complete"
	}

	promptResponse := pb.FbResponse{Id: fbResp.ID.String(), 
	ParsedTemplate: fbResp.ParsedTemplate, 
	SurveyIndex: int32(fbResp.PromptIndex), Status: respStatus}

	switch sPrompt.ResponseType.Type {
	case "freeform":
		promptResponse.Type = pb.MeasureType_freeform
		if (fbResp.FreeformValue == nil) {
			promptResponse.FreeformValue = ""
		} else {
			promptResponse.FreeformValue = *fbResp.FreeformValue
		}
	case "range":
		promptResponse.Type = pb.MeasureType_range
		if (fbResp.RangeValue == nil) {
			promptResponse.RangeValue = 0
		} else {
			promptResponse.RangeValue = int32(*fbResp.RangeValue)
		}
		promptResponse.Range = &pb.MeasureRange{HighRange: int32(sPrompt.ResponseType.HighRange), LowRange: int32(sPrompt.ResponseType.LowRange)}
	case "boolean":
		promptResponse.Type = pb.MeasureType_boolean
		if (fbResp.BoolValue == nil) {
			promptResponse.BooleanValue = ""
		} else {
			promptResponse.BooleanValue = *fbResp.BoolValue
		}
		promptResponse.Boolean = &pb.MeasureBoolean{BoolTrue: sPrompt.ResponseType.BooleanTrue, BoolFalse: sPrompt.ResponseType.BooleanFalse}
	case "enumeration":
		promptResponse.Type = pb.MeasureType_enumeration
		promptResponse.EnumerationLabelValue = fbResp.EnumValue.Label
		promptResponse.EnumerationValue = int32(fbResp.EnumValue.Value)
		promptResponse.Enumeration = mapDbMeasureEnumeration(sPrompt.ResponseType.Enumeration)
	case "multiflag":
		promptResponse.Type = pb.MeasureType_multiflag
		promptResponse.FlagValues = fbResp.LabelValues
		promptResponse.Flags = mapDbMeasureFlags(sPrompt.ResponseType.Flags)
	}

	return &promptResponse

}

func mapDbMeasureEnumeration(dbEnums []schema.MeasureEnum) []*pb.MeasureEnum {
	pbEnums := []*pb.MeasureEnum{}
	for _, e := range dbEnums {
		pbEnums = append(pbEnums, &pb.MeasureEnum{Label: e.Label, Value: int32(e.Value)})
	}
	return pbEnums
}

func mapDbMeasureFlags(dbFlags []string) []*pb.MeasureFlags {
	pbFlags := []*pb.MeasureFlags{}
	for _, e := range dbFlags {
		pbFlags = append(pbFlags, &pb.MeasureFlags{Label: e})
	}
	return pbFlags
}

