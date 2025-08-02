package mrrobot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type MrRobotClient struct {
	httpClient *http.Client
	BaseUrl    string
}

func NewMrRobotClient(baseUrl string) *MrRobotClient {
	return &MrRobotClient{
		httpClient: &http.Client{},
		BaseUrl:    baseUrl,
	}
}

func (c *MrRobotClient) SendMessage(body BodyToSend) error {
	url := c.BaseUrl + "send"
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
