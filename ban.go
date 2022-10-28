package faceitgo

import (
	"encoding/json"
)

/*
CREATE BAN:
POST: https://api.faceit.com/hubs/v1/hub/<hubId>/ban/<userId>
Headers:
Authorization: Bearer <token>
Payload:
{
	"hubId": "string",
	"userId": "string",
	"reason": "string",
}

DELETE BAN:
DELETE: https://api.faceit.com/hubs/v1/hub/<hubId>/ban/<userId>
Headers:
Authorization: Bearer <token>

*/

type (
	BanRequest struct {
		HubId  string `json:"hubId"`
		UserId string `json:"userId"`
		Reason string `json:"reason"`
	}
)

func (c *RESTClient) CreateBan(hubId, userId, reason string) error {
	req := BanRequest{
		HubId:  hubId,
		UserId: userId,
		Reason: reason,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	_, err = c.postPriv("/hubs/v1/hub/"+hubId+"/ban/"+userId, payload)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTClient) DeleteBan(hubId, userId string) error {
	_, err := c.deletePriv("/hubs/v1/hub/" + hubId + "/ban/" + userId)
	if err != nil {
		return err
	}

	return nil
}
