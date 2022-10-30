package faceitgo

import (
	"fmt"
	"net/url"
	"strconv"
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

func (c *RESTClient) GetGames(offset, limit *int) (*Games, error) {
	var games Games

	params := url.Values{}
	if offset != nil {
		params.Add("offset", strconv.Itoa(*offset))
	}

	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	u := fmt.Sprintf("/games?%s", params.Encode())
	if err := c.getJSON(u, &games, nil); err != nil {
		return nil, err
	}

	return &games, nil
}

func (c *RESTClient) GetGame(game_id string) (*Game, error) {
	var game Game

	if err := c.getJSON(fmt.Sprintf("/games/%s", game_id), &game, nil); err != nil {
		return nil, err
	}

	return &game, nil
}

func (c *RESTClient) GetGameParent(game_id string) (*Game, error) {
	var game Game

	if err := c.getJSON(fmt.Sprintf("/games/%s/parent", game_id), &game, nil); err != nil {
		return nil, err
	}

	return &game, nil
}
