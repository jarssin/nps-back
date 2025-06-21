package csat_test

import (
	"testing"

	"github.com/jarssin/nps-back/pkg/survey/csat"
)

type fakeRepo struct{}

func (f *fakeRepo) CreateSurvey(csat.DTO) error { return nil }

func TestCreateSurvey(t *testing.T) {
	svc := csat.NewSurveyService(&fakeRepo{})
	payload := csat.DTO{Comment: "test", VisitorId: "abc"}
	if err := svc.CreateSurvey(payload); err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}
