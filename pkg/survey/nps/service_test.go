package nps_test

import (
	"testing"

	"github.com/jarssin/nps-back/pkg/survey/nps"
)

type fakeRepo struct{}

func (f *fakeRepo) CreateSurvey(nps.DTO) error { return nil }

func TestCreateSurvey_Success(t *testing.T) {
	svc := nps.NewSurveyService(&fakeRepo{})
	payload := nps.ToCreateDTO{Comment: "test", Score: 10, VisitorId: "abc"}
	err := svc.CreateSurvey(payload)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestCreateSurvey_InvalidPayloadType(t *testing.T) {
	svc := nps.NewSurveyService(&fakeRepo{})
	err := svc.CreateSurvey("invalid type")
	if err == nil || err.Error() != "invalid NPS payload type" {
		t.Errorf("expected invalid NPS payload type error, got %v", err)
	}
}

func TestCreateSurvey_ValidationError(t *testing.T) {
	svc := nps.NewSurveyService(&fakeRepo{})
	payload := nps.ToCreateDTO{Comment: "", Score: 0, VisitorId: ""} // inválido
	err := svc.CreateSurvey(payload)
	if err == nil || err.Error()[:17] != "validation error:" {
		t.Errorf("expected validation error, got %v", err)
	}
}
