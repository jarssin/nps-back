package ordersync

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jarssin/nps-back/internal/infra/providers/cardapioweb"
	"github.com/jarssin/nps-back/pkg/order"
	"github.com/jarssin/nps-back/pkg/person"
)

type Service struct {
	client ClientI
	pR     person.RepositoryI
	oR     order.RepositoryI
	oSyncR OrderSyncRepositoryI
}

func NewService(client ClientI, pR person.RepositoryI, oR order.RepositoryI, oSyncR OrderSyncRepositoryI) ServiceI {
	return &Service{client, pR, oR, oSyncR}
}

func (s *Service) SendSurveyToLastOrders(ctx context.Context) error {
	lastSyncDate, err := s.oSyncR.FindOrderSyncLastDate(ctx)
	if err != nil {
		fmt.Printf("error finding last sync date: %v\n", err)
	}

	if lastSyncDate == "" {
		lastSyncDate = time.Now().Add(-23 * time.Hour).Format(time.RFC3339)
		s.oSyncR.UpsertOrderSyncLastDate(ctx, lastSyncDate)
	}

	orders, err := s.client.PullOrders(ctx, lastSyncDate)
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
				fmt.Printf("error retrieving order details for order ID %d: %v\n", orderID, err)
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
	var ordersToInsert []order.DTO

	for orderDetails := range orderDetailsChan {
		person := person.DTO{
			Name:  orderDetails.Customer.Name,
			Phone: orderDetails.Customer.Phone,
		}

		if person.Phone != "" {
			peopleToInsert = append(peopleToInsert, person)
		}

		order := order.DTO{
			ID:            orderDetails.ID,
			Status:        orderDetails.Status,
			OrderType:     orderDetails.OrderType,
			OderTiming:    orderDetails.OderTiming,
			SalesChannel:  orderDetails.SalesChannel,
			CreatedAt:     orderDetails.CreatedAt,
			UpdatedAt:     orderDetails.UpdatedAt,
			CustomerName:  orderDetails.Customer.Name,
			CustomerPhone: orderDetails.Customer.Phone,
			Items: func() []order.ItemDTO {
				var items []order.ItemDTO
				for _, item := range orderDetails.Items {
					items = append(items, item.ToOrderItem())
				}
				return items
			}(),
		}

		ordersToInsert = append(ordersToInsert, order)
	}

	err = s.pR.BulkCreate(ctx, peopleToInsert)
	if err != nil {
		fmt.Printf("error inserting people: %v\n", err)
	}
	err = s.oR.BulkCreate(ctx, ordersToInsert)
	if err != nil {
		fmt.Printf("error inserting orders: %v\n", err)
	}
	return nil
}
