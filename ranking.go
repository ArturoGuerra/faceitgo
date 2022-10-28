package faceitgo

import (
	"fmt"
	"net/url"
)

type (
	Ranking struct {
		Country        string `json:"country"`
		FaceitElo      int    `json:"faceit_elo"`
		GameSkillLevel int    `json:"game_skill_level"`
		Nickname       string `json:"nickname"`
		PlayerID       string `json:"player_id"`
		Position       int    `json:"position"`
	}

	Rankings struct {
		Start int       `json:"start"`
		End   int       `json:"end"`
		Items []Ranking `json:"items"`
	}
)

func (c *RESTClient) GetGameRankings(game_id string, region string, country string, offset string, limit string) (*Rankings, error) {
	var rankings Rankings

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)
	params.Add("country", country)

	u := fmt.Sprintf("/rankings/games/%s/regions/%s", game_id, region)
	if err := c.getJSON(u, &rankings, params); err != nil {
		return nil, err
	}

	return &rankings, nil
}

func (c *RESTClient) GetPlayerRankings(game_id string, region string, player_id string, country string, limit string) (*Rankings, error) {
	var rankings Rankings

	params := url.Values{}
	params.Add("limit", limit)
	params.Add("country", country)

	u := fmt.Sprintf("/rankings/games/%s/regions/%s/players/%s", game_id, region, player_id)
	if err := c.getJSON(u, &rankings, params); err != nil {
		return nil, err
	}

	return &rankings, nil
}
