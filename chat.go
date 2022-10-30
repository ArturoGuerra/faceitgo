package faceitgo

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

type (
	RoomDetail struct {
		Name    string `json:"name"`
		Members []struct {
			IsOnline bool     `json:"is_online"`
			MemberID string   `json:"member_id"`
			Nickname string   `json:"nickname"`
			Photo    string   `json:"photo"`
			Roles    []string `json:"roles"`
			Status   string   `json:"status"`
		}
		Roles []struct {
			Color       string   `json:"color"`
			Displayed   bool     `json:"displayed"`
			Mentionable bool     `json:"mentionable"`
			Permissions []string `json:"permissions"`
			Ranking     int      `json:"ranking"`
			RoleID      string   `json:"role_id"`
		}
	}
)

func (c *RESTClient) GetRoom(roomID string) (*RoomDetail, error) {
	resp, err := c.getChat("/rooms/" + roomID)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		_, err := httpErrorHandler(resp)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var room RoomDetail
	err = json.Unmarshal(body, &room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (c *RESTClient) GetRoomMessages(roomID string, before *string, limit *int) (*RoomMessages, error) {
	params := url.Values{}
	if before != nil {
		params.Add("before", *before)
	}

	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	resp, err := c.getChat("/rooms/" + roomID + "/messages?" + params.Encode())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		_, err := httpErrorHandler(resp)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var messages RoomMessages
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}

	return &messages, nil
}

func (c *RESTClient) PostRoomMessage(roomID, message string) (string, error) {
	data, err := json.Marshal(map[string]string{
		"body": message,
	})
	if err != nil {
		return "", err
	}

	resp, err := c.postChat("/rooms/"+roomID+"/messages", data)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		_, err := httpErrorHandler(resp)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var msg RoomMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return "", err
	}

	return msg.ID, nil
}
