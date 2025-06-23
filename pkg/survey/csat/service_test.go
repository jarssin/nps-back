package csat_test

import (
	"testing"

	"github.com/jarssin/nps-back/pkg/survey/csat"
)

type fakeRepo struct{}

func (f *fakeRepo) CreateSurvey(csat.DTO) error { return nil }

func TestCreateSurvey_Success(t *testing.T) {
	svc := csat.NewSurveyService(&fakeRepo{})
	payload := csat.ToCreateDTO{Comment: "test", VisitorId: "abc"}
	if err := svc.CreateSurvey(payload); err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestCreateSurvey_InvalidPayloadType(t *testing.T) {
	svc := csat.NewSurveyService(&fakeRepo{})
	if err := svc.CreateSurvey("invalid type"); err == nil || err.Error() != "invalid CSAT payload type" {
		t.Errorf("expected invalid CSAT payload type error, got %v", err)
	}
}

func TestCreateSurvey_ValidationError(t *testing.T) {
	svc := csat.NewSurveyService(&fakeRepo{})
	payload := csat.ToCreateDTO{Comment: "", VisitorId: ""} // inválido
	if err := svc.CreateSurvey(payload); err == nil || err.Error()[:17] != "validation error:" {
		t.Errorf("expected validation error, got %v", err)
	}
}
