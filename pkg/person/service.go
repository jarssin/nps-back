package person

import (
	"log"

	"github.com/jarssin/nps-back/pkg/mrrobot"
)

type Service struct {
	mrRobotClient *mrrobot.MrRobotClient
}

func NewService(m *mrrobot.MrRobotClient) ServiceI {
	return &Service{mrRobotClient: m}
}

func (s *Service) SendSurvey(people []DTO) {
	for _, p := range people {
		log.Printf("Sending survey to %s with phone %s", p.Name, p.Phone)
		s.mrRobotClient.SendMessage(mrrobot.BodyToSend{
			Msg:   "Olá " + p.Name + ", por favor, responda nossa pesquisa. \nhttps://nps-static-front.storage.googleapis.com/index.html?surveyType=nps&phone=" + p.Phone + "&name=" + p.Name,
			Phone: p.Phone,
		})
	}

	log.Println("All surveys sent successfully")
}
