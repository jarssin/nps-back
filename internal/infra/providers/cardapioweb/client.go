package cardapioweb

import (
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

func (c *Client) PullOrders(lastCallDate string) ([]OrderDTO, error) {
	url := c.baseUrl + "/api/partner/v1/orders"
	req, err := http.NewRequest("GET", url, nil)
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

func (c *Client) OrderDetails(orderID string) (OrderDetailsDTO, error) {
	// Implement the logic to get order details from CardapioWeb using c.http, c.baseUrl, and c.token
	// This might involve making a GET request to an endpoint with the orderID
	return OrderDetailsDTO{}, nil
}
