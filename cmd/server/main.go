package main

import (
	"net/http"
	"skillsdemo/internal/surveyserver"
	"skillsdemo/rpc/survey"
)

func main() {
	server := &surveyserver.Server{}
	twirpHandler := survey.NewSurveyServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}