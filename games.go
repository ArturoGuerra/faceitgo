package faceitgo

import (
	"encoding/json"
	"io/ioutil"
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
	resp, err := c.Get("games")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &games)
	if err != nil {
		return nil, err
	}

	return &games, err
}

func (c *RESTClient) GetGame(gameID string) (*Game, error) {
	var game Game
	resp, err := c.Get("games/" + gameID)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &game)
	if err != nil {
		return nil, err
	}

	return &game, err
}

func (c *RESTClient) GetGameParent(gameID string) (*Game, error) {
	var game Game
	resp, err := c.Get("games/" + gameID + "/parent")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &game)
	if err != nil {
		return nil, err
	}

	return &game, err
}
