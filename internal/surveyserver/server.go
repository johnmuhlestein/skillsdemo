package surveyserver

import (
	"context"

	pb "skillsdemo/rpc/survey"

	"github.com/google/uuid"
)

type Server struct {}

func (s *Server) MakeSurvey (ctx context.Context, req *pb.SurveyReq) (*pb.SurveyResp, error) {

	return &pb.SurveyResp{
		Id: uuid.New().String(),
		AppointmentId: req.AppointmentId,
		Title: req.Title,
		Description: req.Description,
		Status: req.Status,
		Prompts: req.Prompts,
	},nil
}