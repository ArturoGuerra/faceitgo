package faceitgo

import (
	"fmt"
	"net/url"
)

type (
	Team struct {
		Avatar      string `json:"avatar"`
		ChatRoomID  string `json:"chat_room_id"`
		CoverImage  string `json:"cover_image"`
		Description string `json:"description"`
		Facebook    string `json:"facebook"`
		FaceitUrl   string `json:"faceit_url"`
		Game        string `json:"game"`
		Leader      string `json:"leader"`
		Members     []struct {
			Avatar         string   `json:"avatar"`
			Country        string   `json:"country"`
			FaceitURL      string   `json:"faceit_url"`
			MembershipType string   `json:"membership_type"`
			Memberships    []string `json:"memberships"`
			Nickname       string   `json:"nickname"`
			SkillLevel     int      `json:"skill_level"`
			UserID         string   `json:"user_id"`
		} `json:"members"`
		Name     string `json:"name"`
		Nickname string `json:"nickname"`
		TeamID   string `json:"team_id"`
		TeamType string `json:"team_type"`
		Twitter  string `json:"twitter"`
		Website  string `json:"website"`
		YouTube  string `json:"youtube"`
	}

	TeamStats struct {
		GameID   string                   `json:"game_id"`
		Lifetime map[string]interface{}   `json:"lifetime"`
		Segments []map[string]interface{} `json:"segments"`
		TeamID   string                   `json:"team_id"`
	}
)

func (c *RESTClient) GetTeam(team_id string) (*Team, error) {
	var team Team

	if err := c.getJSON(fmt.Sprintf("/teams/%s", team_id), &team, nil); err != nil {
		return nil, err
	}

	return &team, nil
}

func (c *RESTClient) GetTeamStats(team_id string, game_id string) (*TeamStats, error) {
	var teamStats TeamStats

	if err := c.getJSON(fmt.Sprintf("/teams/%s/stats/%s", team_id, game_id), &teamStats, nil); err != nil {
		return nil, err
	}

	return &teamStats, nil
}

func (c *RESTClient) GetTeamTournaments(team_id string, offset string, limit string) (*Tournaments, error) {
	var tournaments Tournaments

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	if err := c.getJSON(fmt.Sprintf("/teams/%s/tournaments", team_id), &tournaments, params); err != nil {
		return nil, err
	}

	return &tournaments, nil
}
