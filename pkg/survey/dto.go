package survey

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DTO struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Comment   string             `json:"comment" bson:"comment"`
	Score     int                `json:"score" bson:"score" validate:"required,min=1,max=10"`
	VisitorId string             `json:"visitorId" bson:"visitorId" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

var validate *validator.Validate

func (s *DTO) Validate() error {
	validate = validator.New()
	return validate.Struct(s)
}
