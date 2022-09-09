package faceitgo

import (
	"fmt"
	"net/url"
)

type (
	GameAssets struct {
		Cover        string `json:"cover"`
		FeaturedImgL string `json:"featured_img_l"`
		FeaturedImgM string `json:"featured_img_m"`
		FeaturedImgS string `json:"featured_img_s"`
		FlagImgIcon  string `json:"flag_img_icon"`
		FlagImgL     string `json:"flag_img_l"`
		FlagImgM     string `json:"flag_img_m"`
		FlagImgS     string `json:"flag_img_s"`
		LandingPage  string `json:"landing_page"`
	}

	Game struct {
		Assets       GameAssets `json:"assets"`
		GameID       string     `json:"game_id"`
		LongLabel    string     `json:"long_label"`
		Order        int        `json:"order"`
		ParentGameID string     `json:"parent_game_id"`
		Platforms    []string   `json:"platforms"`
		Regions      []string   `json:"regions"`
		ShortLabel   string     `json:"short_label"`
	}

	Games struct {
		Start int    `json:"start"`
		End   int    `json:"end"`
		Items []Game `json:"items"`
	}
)

func (c *RESTClient) GetGames(offset string, limit string) (*Games, error) {
	var games Games

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	u := fmt.Sprintf("/games?%s", params.Encode())
	if err := c.getJSON(u, &games); err != nil {
		return nil, err
	}

	return &games, nil
}

func (c *RESTClient) GetGame(gameID string) (*Game, error) {
	var game Game

	if err := c.getJSON(fmt.Sprintf("/games/%s", gameID), &game); err != nil {
		return nil, err
	}

	return &game, nil
}

func (c *RESTClient) GetGameParent(gameID string) (*Game, error) {
	var game Game

	if err := c.getJSON(fmt.Sprintf("/games/%s/parent", gameID), &game); err != nil {
		return nil, err
	}

	return &game, nil
}
