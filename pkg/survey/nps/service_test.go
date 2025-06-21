package nps_test

import (
	"testing"

	"github.com/jarssin/nps-back/pkg/survey/nps"
)

type fakeRepo struct{}

func (f *fakeRepo) CreateSurvey(nps.DTO) error { return nil }

func TestCreateSurvey(t *testing.T) {
	svc := nps.NewSurveyService(&fakeRepo{})
	err := svc.CreateSurvey(nps.DTO{Comment: "test", Score: 10, VisitorId: "abc"})
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}
