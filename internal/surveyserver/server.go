package surveyserver

import (
	"context"

	pb "skillsdemo/rpc/survey"

	"github.com/google/uuid"
)

type Server struct {}

func (s *Server) CreateSurvey (ctx context.Context, req *pb.CreateSurveyReq) (*pb.SurveyResp, error) {

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
	return resp,nil
}


