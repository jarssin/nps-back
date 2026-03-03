package ordersync

import (
	"context"
	"fmt"
	"sync"

	"github.com/jarssin/nps-back/internal/infra/providers/cardapioweb"
	"github.com/jarssin/nps-back/pkg/person"
)

type Service struct {
	client ClientI
	pR     person.RepositoryI
}

func NewService(client ClientI, pR person.RepositoryI) ServiceI {
	return &Service{client, pR}
}

func (s *Service) SendSurveyToLastOrders(ctx context.Context) error {
	orders, err := s.client.PullOrders(ctx, "")
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	orderDetailsChan := make(chan cardapioweb.OrderDetailsDTO)
	for _, order := range orders {
		wg.Add(1)
		go func(orderID int64) {
			defer wg.Done()
			orderDetails, err := s.client.OrderDetails(ctx, orderID)
			if err != nil {
				fmt.Printf("Error retrieving order details for order ID %d: %v\n", orderID, err)
				return
			}

			orderDetailsChan <- orderDetails
		}(order.ID)
	}

	go func() {
		wg.Wait()
		close(orderDetailsChan)
	}()

	var peopleToInsert []person.DTO

	for orderDetails := range orderDetailsChan {
		person := person.DTO{
			Name:  orderDetails.Customer.Name,
			Phone: orderDetails.Customer.Phone,
		}

		if person.Phone != "" {
			peopleToInsert = append(peopleToInsert, person)
		}
	}

	err = s.pR.BulkCreate(ctx, peopleToInsert)
	if err != nil {
		return fmt.Errorf("error inserting people: %w", err)
	}

	return nil
}
