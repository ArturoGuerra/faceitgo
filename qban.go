package faceitgo

import (
	"encoding/json"
	"io"
)

/*
create qban:
POST: https://api.faceit.com/queue/v1/ban
Headers:
Authorization: Bearer <token>
Payload:
{
	"queueId: "string",
	"userId": "string",
	"reason": "string",
	"banStart": "int",
	"banEnd": "int"
}

Response:
{
	"payload": {
		"id": "string",
		"version": "int",
		"createdAt": "datetime",
		"banStart": "datetime",
		"banEnd": "datetime",
		"reason": "string",
		"type": "string",
		"expired": "boolean",
		"queueId": "string",
		"userId": "string",
		"entityId": "string",
		"entityType": "string"
	},
}



delete qban:
DELETE: https://api.faceit.com/queue/v1/ban/<banId>
Headers:
Authorization: Bearer <token>

*/

type (
	QbanRequest struct {
		QueueId  string `json:"queueId"`
		UserId   string `json:"userId"`
		Reason   string `json:"reason"`
		BanStart int    `json:"banStart"`
		BanEnd   int    `json:"banEnd"`
	}

	QbanResponse struct {
		ID         string `json:"id"`
		Version    int    `json:"version"`
		CreatedAt  string `json:"createdAt"`
		BanStart   string `json:"banStart"`
		BanEnd     string `json:"banEnd"`
		Reason     string `json:"reason"`
		Type       string `json:"type"`
		Expired    bool   `json:"expired"`
		QueueId    string `json:"queueId"`
		UserId     string `json:"userId"`
		EntityId   string `json:"entityId"`
		EntityType string `json:"entityType"`
	}
)

func (c *RESTClient) CreateQban(queueId, userId, reason string, banStart, banEnd int) (*QbanResponse, error) {
	req := QbanRequest{
		QueueId:  queueId,
		UserId:   userId,
		Reason:   reason,
		BanStart: banStart,
		BanEnd:   banEnd,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.postPriv("/queue/v1/ban", payload)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var qban QbanResponse

	data := map[string]interface{}{
		"payload": &qban,
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &qban, nil
}

func (c *RESTClient) DeleteQban(banId string) error {
	_, err := c.deletePriv("/queue/v1/ban/" + banId)
	if err != nil {
		return err
	}

	return nil
}
