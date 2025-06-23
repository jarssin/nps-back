package csat

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTO struct {
	Id                 primitive.ObjectID `json:"_id" bson:"_id"`
	Comment            string             `json:"comment" bson:"comment"`
	VisitorId          string             `json:"visitorId" bson:"visitorId" validate:"required"`
	Name               string             `json:"name" bson:"name"`
	Phone              string             `json:"phone" bson:"phone"`
	JourneyEvaluations map[string]int     `json:"journeyEvaluations" bson:"journeyEvaluations"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
}

type ToCreateDTO struct {
	Comment            string         `json:"comment" bson:"comment"`
	VisitorId          string         `json:"visitorId" bson:"visitorId" validate:"required"`
	Name               string         `json:"name" bson:"name"`
	Phone              string         `json:"phone" bson:"phone"`
	JourneyEvaluations map[string]int `json:"journeyEvaluations" bson:"journeyEvaluations"`
}

func (c ToCreateDTO) ToDTO() DTO {
	return DTO{
		Id:                 primitive.NewObjectID(),
		Comment:            c.Comment,
		VisitorId:          c.VisitorId,
		Name:               c.Name,
		Phone:              c.Phone,
		JourneyEvaluations: c.JourneyEvaluations,
		CreatedAt:          time.Now(),
	}
}

var validate *validator.Validate

func (s *DTO) Validate() error {
	validate = validator.New()
	return validate.Struct(s)
}
