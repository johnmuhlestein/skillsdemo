package surveyserver

import (
	"context"
	"errors"
	"log"
	"time"

	"skillsdemo/ent"
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
	qSurvey, err := database.EntClient.Survey.Query().Where(survey.IDEQ(uuid)).First(ctx)
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


