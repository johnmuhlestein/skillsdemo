package main

import (
	"context"
	"log"
	"net/http"

	"skillsdemo/rpc/survey"
)

func main() {
	client := survey.NewSurveyProtobufClient("http://localhost:8080", &http.Client{})

	description := "This is the first servey using the RPC client to generate a dummy survery"

	survey, err := client.MakeSurvey(context.Background(), &survey.SurveyReq{
		AppointmentId: "b7b85b29-473b-422d-93f9-4106607921ed",
		Title: "First Survey Test",
		Description: &description,
		Status: survey.SurveyStatus_unpublished,
		Prompts: &survey.Prompt{
			Index: 1,
			Title: "Did you like the survey",
			Description: "On a scale of 1 - 10 with 1 indicating not like the survey and 10 indicating loving the survey how much did you enjoy the survey?",
    		ResponseType: &survey.Measure{
				Type: survey.MeasureType_range,
				MeasureOneof: &survey.Measure_Range{
					Range: &survey.MeasureRange{
						LowRange: 1,
						HighRange: 10,
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("Unable to make the survey %v/n", err)
	}
	log.Println(survey)
}