package cardapioweb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	http    *http.Client
	baseUrl string
	token   string
}

func NewClient(httpClient *http.Client, baseUrl, token string) *Client {
	return &Client{
		http:    httpClient,
		baseUrl: baseUrl,
		token:   token,
	}
}

func (c *Client) PullOrders(ctx context.Context, lastCallDate string) ([]OrderDTO, error) {
	url := c.baseUrl + "/api/partner/v1/orders"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if lastCallDate == "" {
		lastCallDate = time.Now().Add(-23 * time.Hour).Format(time.RFC3339)
	}

	req.Header.Set("X-API-KEY", c.token)
	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = fmt.Sprintf("updated_since=%s", lastCallDate)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error pulling orders: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var orders []OrderDTO
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return orders, nil
}

func (c *Client) OrderDetails(ctx context.Context, orderID int64) (OrderDetailsDTO, error) {
	url := fmt.Sprintf(c.baseUrl+"/api/partner/v1/orders/%d", orderID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return OrderDetailsDTO{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("X-API-KEY", c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return OrderDetailsDTO{}, fmt.Errorf("error fetching order details: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return OrderDetailsDTO{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var orderDetails OrderDetailsDTO
	if err := json.NewDecoder(resp.Body).Decode(&orderDetails); err != nil {
		return OrderDetailsDTO{}, fmt.Errorf("error decoding response: %w", err)
	}

	return orderDetails, nil
}
