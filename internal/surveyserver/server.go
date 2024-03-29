package surveyserver

import (
	"bytes"
	"context"
	"errors"
	"log"
	"text/template"
	"time"

	"skillsdemo/ent"
	"skillsdemo/ent/appointment"
	"skillsdemo/ent/schema"
	"skillsdemo/ent/survey"
	"skillsdemo/internal/database"
	pb "skillsdemo/rpc/survey"

	"github.com/google/uuid"
	"github.com/twitchtv/twirp"
)

type Server struct {}

func (s *Server) CreateSurvey (ctx context.Context, req *pb.CreateSurveyReq) (*pb.SurveyResp, error) {

	survey, err := wrapCreateSurveyReq(ctx,req)
	if err != nil {
		log.Printf("Error creating survey with title %s - %v\n", req.Title, err)
		return nil, twirp.InternalError("An error occured attempting to create the new survey")
	}
	resp := wrapSurvey(ctx, survey)
/*
	resp := &pb.SurveyResp{
		Id: uuid.New().String(),
		Title: req.Title,
		Description: req.Description,
		Status: req.Status,
	}
	for _, p  := range req.Prompts {
		resp.Prompts = append(resp.Prompts, &pb.Prompt{
			Id: uuid.New().String(),
			Index: p.Index,
			Title: p.Title,
			Description: p.Description,
			ResponseType: p.ResponseType,
			AllowAdditionalFeedback: p.AllowAdditionalFeedback,
		})
	}
*/
	return resp,nil
}

func (s *Server) UpdateSurveyStatus(ctx context.Context, req *pb.SurveyStatusReq) (*pb.SurveyResp, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, twirp.InvalidArgument.Errorf("The survey id [%s] is not a valid uuid", req.Id)
	}
	qSurvey, err := database.EntClient.Survey.Get(ctx,uuid)
	if errors.Is(err, &ent.NotFoundError{}) {
		return nil, twirp.NotFound.Errorf("The survey id [%s] was not found", req.Id)
	}
	switch req.Status {
		case pb.SurveyStatus_active: 
			qSurvey, err = qSurvey.Update().SetStatus(survey.Status(req.Status.String())).SetActiveTime(time.Now()).Save(ctx)
		case pb.SurveyStatus_archived:
			qSurvey, err = qSurvey.Update().SetStatus(survey.Status(req.Status.String())).SetArchiveTime(time.Now()).Save(ctx)
		case pb.SurveyStatus_unpublished:
			return nil, twirp.InvalidArgument.Errorf("You can not transition from any other status back to 'unpublished'")
	}
	if err != nil {
		log.Printf("Failed to update Survey %s to a new status %s - %v\n", req.Id, req.Status.String(), err)
		return nil, twirp.InternalError("Unable to update status")
	}
	return wrapSurvey(ctx, qSurvey), nil
}

func (s *Server) UpdateAppointmentStatus(ctx context.Context, req *pb.AppointmentStatusReq) (*pb.FeedbackResp, error) {
	uuid, err := uuid.Parse(req.Id)
	feedbackResp := pb.FeedbackResp{}
	if err != nil {
		return nil, twirp.InvalidArgument.Errorf("The survey id [%s] is not a valid uuid", req.Id)
	}
	qAppt, err := database.EntClient.Appointment.Get(ctx,uuid)
	if errors.Is(err, &ent.NotFoundError{}) {
		return nil, twirp.NotFound.Errorf("The survey id [%s] was not found", req.Id)
	}
	_, err = qAppt.Update().SetStatus(req.Status).Save(ctx)
		if err != nil {
		log.Printf("Failed to update Appointment %s to a new status %s - %v\n", req.Id, req.Status, err)
		return nil, twirp.InternalError("Unable to update status")
	}

	//eager fetch to populate the relationships
	eagerAppt, err := database.EntClient.Appointment.Query().Where(appointment.IDEQ(uuid)).WithDiagnoses().WithPatient().WithProvider().First(ctx)
		if err != nil {
		log.Printf("Failed to retrieve the eager Appointment %s to a new status %s - %v\n", req.Id, req.Status, err)
		return nil, twirp.InternalError("Unable to update status")
	}
	theSurvey, err := database.EntClient.Appointment.Query().Where(appointment.IDEQ(uuid)).QuerySurvey().WithPrompts().First(ctx)
	if err != nil {
		log.Printf("Failed to retrieve the survey associated with the appointment %s to a new status %s - %v\n", req.Id, req.Status, err)
		return nil, twirp.InternalError("Unable to update status")
	}
	parsedMap := make(map[int]string)
	for _, i := range theSurvey.Edges.Prompts {
		p := parseTemplate(i.Description, eagerAppt)
		parsedMap[i.SortOrder] = p
	}

	feedbackEnt, err := database.EntClient.Feedback.Create().SetStatus("created").SetSurvey(theSurvey).Save(ctx)
	if err != nil {
		log.Printf("Failed to create new Feedback for Appointment %s - %v\n", req.Id, err)
		return nil, twirp.InternalError("Unable to update status")
	}

	feedbackResp.Id = feedbackEnt.ID.String()
	feedbackResp.Status = feedbackEnt.Status

	for i, v := range parsedMap {
		resp, err := database.EntClient.PromptResponse.Create().SetParsedTemplate(v).SetPromptIndex(i).SetFeedback(feedbackEnt).Save(ctx)
		if err != nil {
			log.Printf("Failed to create new Prompt Response for Feedback %s - %v\n", feedbackEnt.ID.String(), err)
			return nil, twirp.InternalError("Unable to update status")
		}
		feedbackResp.Responses = append(feedbackResp.Responses, &pb.FbResponses{Id: resp.ID.String(), ParsedTemplate: v, SurveyIndex: int32(i), Status: "new"})
	} 

	return &feedbackResp, nil
}

func parseTemplate(tmplt string, appt *ent.Appointment) string {
	tmpl := template.New("prompt")
	tmpl, err := tmpl.Parse(tmplt)
	if err != nil {
		log.Print("Unable to parse the prompt template %s - %v\n", tmplt, err)
		return tmplt
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, appt)
	if err != nil {
		log.Print("Unable to execute the template binding the prompt template %s - %v\n", tmplt, err)
		return tmplt
	}
	return buf.String()
}

func wrapSurvey(ctx context.Context, s *ent.Survey) *pb.SurveyResp {
	resp := pb.SurveyResp{}
	resp.Id = s.ID.String()
	num := pb.SurveyStatus_value[s.Status.String()]
	resp.Status = pb.SurveyStatus(num)
	resp.Title = s.Title
	resp.Description = &s.Description
	prompts, err := s.QueryPrompts().All(ctx)
	if !errors.Is(err,&ent.NotFoundError{}) {
		resp.Prompts = wrapPrompts(prompts)
	}
	return &resp
}

// This would be better transactional
func wrapCreateSurveyReq(ctx context.Context,r *pb.CreateSurveyReq) (*ent.Survey, error) {
//	newPrompts := make([]*ent.Prompt, len(r.Prompts))
	newPromptIds := make([]uuid.UUID, 0)
	for _, p := range r.Prompts {
		promptCreate := database.EntClient.Prompt.Create()
		promptCreate.SetTitle(p.Title).SetDescription(p.Description).SetSortOrder(int(p.Index))
		rt := schema.Measure{Type: p.ResponseType.Type.String()}
		switch p.ResponseType.Type.String() {
		case "range":
			rt.HighRange = int(p.ResponseType.Range.HighRange)
			rt.LowRange = int(p.ResponseType.Range.LowRange)
		case "boolean":
			rt.BooleanFalse = p.ResponseType.Boolean.BoolFalse
			rt.BooleanTrue = p.ResponseType.Boolean.BoolTrue
		case "enumeration":
			for _, e := range p.ResponseType.Enumeration {
				rt.Enumeration = append(rt.Enumeration, schema.MeasureEnum{Value: int(e.Value), Label: e.Label})
			}
		case "multiflag":
			for _,f := range p.ResponseType.Flags {
				rt.Flags = append(rt.Flags, f.Label)
			}
		}
		
		np, err := promptCreate.SetResponseType(rt).Save(ctx)
		if err != nil {
			return nil, err
		}
		newPromptIds = append(newPromptIds, np.ID)
	}
	log.Printf("%d prompts created\n", len(newPromptIds))
	
	newSurvey, err := database.EntClient.Survey.Create().SetTitle(r.Title).SetDescription(*r.Description).SetStatus(survey.Status(r.GetStatus().String())).AddPromptIDs(newPromptIds...).Save(ctx)
	if err != nil {
		return nil,err
	}
	return newSurvey,nil
}

func wrapPrompts(p []*ent.Prompt) []*pb.Prompt {
	newPrompts := make([]*pb.Prompt,0)
	for _, x := range p {
		n := pb.Prompt{Id : x.ID.String(), Title: x.Title, Description: x.Description, Index: int32(x.SortOrder), AllowAdditionalFeedback: x.AdditionalFeedback}
		y := x.ResponseType
		n.ResponseType = &pb.Measure{Type: pb.MeasureType(pb.MeasureType_value[y.Type])}
		switch y.Type {
		case "range":
			n.ResponseType.Range = &pb.MeasureRange{LowRange: int32(x.ResponseType.LowRange), HighRange: int32(x.ResponseType.HighRange)}
		case "boolean":
			n.ResponseType.Boolean = &pb.MeasureBoolean{BoolTrue: x.ResponseType.BooleanTrue, BoolFalse: x.ResponseType.BooleanFalse}
		case "enumeration":
			for _, e := range x.ResponseType.Enumeration {
				n.ResponseType.Enumeration = append(n.ResponseType.Enumeration, &pb.MeasureEnum{Value: int32(e.Value), Label: e.Label}) 
			}
		case "multiflag":
			for _, f := range x.ResponseType.Flags {
				n.ResponseType.Flags = append(n.ResponseType.Flags, &pb.MeasureFlags{Label: f})
			}
		}
		newPrompts = append(newPrompts, &n)
	}
	return newPrompts
}


