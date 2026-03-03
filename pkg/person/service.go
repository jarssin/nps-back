package person

import (
	"context"
	"log"

	"github.com/jarssin/nps-back/pkg/mrrobot"
)

type Service struct {
	mrRobotClient *mrrobot.MrRobotClient
	repository    RepositoryI
}

func NewService(m *mrrobot.MrRobotClient, r RepositoryI) ServiceI {
	return &Service{mrRobotClient: m, repository: r}
}

func (s *Service) Create(ctx context.Context, person DTO) error {
	return s.repository.Create(ctx, person)
}

func (s *Service) SendSurvey(ctx context.Context, people []DTO) {
	for _, p := range people {
		person, err := s.repository.Find(ctx, p.Phone)
		if err != nil && err.Error() != "mongo: no documents in result" {
			log.Printf("Error finding person with phone %s: %v", p.Phone, err)
			continue
		}

		if person.Phone == "" {
			log.Printf("Person with phone %s not found, creating new record", p.Phone)
			if err := s.repository.Create(ctx, p); err != nil {
				log.Printf("Error creating person with phone %s: %v", p.Phone, err)
				continue
			}
		}

		log.Printf("Sending survey to %s with phone %s", person.Name, person.Phone)
		s.mrRobotClient.SendMessage(mrrobot.BodyToSend{
			Msg:   "Olá " + p.Name + ", por favor, responda nossa pesquisa. \nhttps://nps-static-front.storage.googleapis.com/index.html?surveyType=nps&phone=" + p.Phone + "&name=" + p.Name,
			Phone: p.Phone,
		})
	}

	log.Println("All surveys sent successfully")
}
