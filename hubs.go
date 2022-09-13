package faceitgo

import (
	"fmt"
	"net/url"
	"strconv"
)

type (
	HubStats struct {
		GameID  string `json:"game_id"`
		Players []struct {
			Nickname string                 `json:"nickname"`
			PlayerID string                 `json:"player_id"`
			Stats    map[string]interface{} `json:"stats"`
		} `json:"players"`
	}

	HubRules struct {
		Body      string `json:"body"`
		Game      string `json:"game"`
		Name      string `json:"name"`
		Organizer string `json:"organizer"`
		RuleID    string `json:"rule_id"`
	}

	HubRole struct {
		Color         string `json:"color"`
		Name          string `json:"name"`
		Ranking       int    `json:"ranking"`
		RoleID        string `json:"role_id"`
		VisibleOnChat bool   `json:"visible_on_chat"`
	}

	HubRoles struct {
		Start int       `json:"start"`
		End   int       `json:"end"`
		Items []HubRole `json:"items"`
	}

	HubMember struct {
		Avatar    string   `json:"avatar"`
		FaceitUrl string   `json:"faceit_url"`
		Nickname  string   `json:"nickname"`
		Roles     []string `json:"roles"`
		UserID    string   `json:"user_id"`
	}

	HubMembers struct {
		Start int         `json:"start"`
		End   int         `json:"end"`
		Items []HubMember `json:"items"`
	}

	Hub struct {
		Avatar          string    `json:"avatar"`
		BackgroundImage string    `json:"background_image"`
		ChatRoomID      string    `json:"chat_room_id"`
		CoverImage      string    `json:"cover_image"`
		Description     string    `json:"description"`
		FaceitUrl       string    `json:"faceit_url"`
		GameData        Game      `json:"game_data"`
		GameID          string    `json:"game_id"`
		HubID           string    `json:"hub_id"`
		JoinPermission  string    `json:"join_permission"`
		MaxSkillLevel   int       `json:"max_skill_level"`
		MinSkillLevel   int       `json:"min_skill_level"`
		Name            string    `json:"name"`
		OrganizerData   Organizer `json:"organizer_data"`
		OrganizerID     string    `json:"organizer_id"`
		PlayersJoined   int       `json:"players_joined"`
		Region          string    `json:"region"`
		RuleID          string    `json:"rule_id"`
	}

	Hubs struct {
		Start int   `json:"start"`
		End   int   `json:"end"`
		Items []Hub `json:"items"`
	}
)

func (c *RESTClient) GetHub(hub_id string) (*Hub, error) {
	var hub Hub

	if err := c.getJSON(fmt.Sprintf("/hubs/%s", hub_id), &hub, nil); err != nil {
		return nil, err
	}

	return &hub, nil
}

func (c *RESTClient) GetHubMatches(hub_id string, game_type string, offset int, limit int) (*Matches, error) {
	var hubMatches Matches

	params := url.Values{}
	params.Add("type", game_type)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/matches", hub_id), &hubMatches, params); err != nil {
		return nil, err
	}

	return &hubMatches, nil
}

func (c *RESTClient) GetHubMembers(hub_id string, offset int, limit int) (*HubMembers, error) {
	var hubMembers HubMembers

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/members", hub_id), &hubMembers, nil); err != nil {
		return nil, err
	}

	return &hubMembers, nil
}

func (c *RESTClient) GetHubRoles(hub_id string, offset int, limit int) (*HubRoles, error) {
	var hubRoles HubRoles

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/roles", hub_id), &hubRoles, params); err != nil {
		return nil, err
	}

	return &hubRoles, nil
}

func (c *RESTClient) GetHubRules(hub_id string, offset int, limit int) (*HubRules, error) {
	var hubRules HubRules

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/rules", hub_id), &hubRules, params); err != nil {
		return nil, err
	}

	return &hubRules, nil
}

func (c *RESTClient) GetHubStats(hub_id string, offset int, limit int) (*HubStats, error) {
	var hubStats HubStats

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/stats", hub_id), &hubStats, nil); err != nil {
		return nil, err
	}

	return &hubStats, nil
}
