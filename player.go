package faceitgo

import (
	"fmt"
	"net/url"
)

type (
	Player struct {
		Avatar             string   `json:"avatar"`
		Country            string   `json:"country"`
		CoverFeaturedImage string   `json:"cover_featured_image"`
		CoverImage         string   `json:"cover_image"`
		FaceitUrl          string   `json:"faceit_url"`
		FriendsIds         []string `json:"friends_ids"`
		Games              map[string]struct {
			FaceitElo       int                    `json:"faceit_elo"`
			GamePlayerID    string                 `json:"game_player_id"`
			GamePlayerName  string                 `json:"game_player_name"`
			GameProfileID   string                 `json:"game_profile_id"`
			Region          string                 `json:"region"`
			Regions         map[string]interface{} `json:"regions"`
			SkillLevel      int                    `json:"skill_level"`
			SkillLevelLabel string                 `json:"skill_level_label"`
		} `json:"games"`
		Infractions    map[string]interface{} `json:"infractions"`
		MembershipType string                 `json:"membership_type"`
		Memberships    []string               `json:"memberships"`
		NewSteamID     string                 `json:"new_steam_id"`
		Nickname       string                 `json:"nickname"`
		Platforms      map[string]string      `json:"platforms"`
		PlayerID       string                 `json:"player_id"`
		Settings       struct {
			Language string `json:"language"`
		} `json:"settings"`
		SteamID64     string `json:"steam_id_64"`
		SteamNickname string `json:"steam_nickname"`
	}

	PlayerMatch struct {
		CompetitionID   string   `json:"competition_id"`
		CompetitionName string   `json:"competition_name"`
		CompetitionType string   `json:"competition_type"`
		FaceitUrl       string   `json:"faceit_url"`
		FinishedAt      int      `json:"finished_at"`
		GameID          string   `json:"game_id"`
		GameMode        string   `json:"game_mode"`
		MatchID         string   `json:"match_id"`
		MatchType       string   `json:"match_type"`
		MaxPlayers      int      `json:"max_players"`
		OrganizerID     string   `json:"organizer_id"`
		PlayingPlayers  []string `json:"playing_players"`
		Region          string   `json:"region"`
		Results         struct {
			Score  map[string]int `json:"score"`
			Winner string         `json:"winner"`
		} `json:"results"`
		StartedAt int    `json:"started_at"`
		Status    string `json:"status"`
		Teams     map[string]struct {
			Avatar   string `json:"avatar"`
			Nickname string `json:"nickname"`
			Players  []struct {
				Avatar         string `json:"avatar"`
				FaceitUrl      string `json:"faceit_url"`
				GamePlayerID   string `json:"game_player_id"`
				GamePlayerName string `json:"game_player_name"`
				Nickname       string `json:"nickname"`
				PlayerID       string `json:"player_id"`
				SkillLevel     int    `json:"skill_level"`
			} `json:"players"`
		} `json:"teams"`
		TeamSize int `json:"team_size"`
	}

	PlayerMatches struct {
		Start int           `json:"start"`
		End   int           `json:"end"`
		To    int           `json:"to"`
		Items []PlayerMatch `json:"items"`
	}

	PlayerStats struct {
		GameID   string                   `json:"game_id"`
		Lifetime map[string]interface{}   `json:"lifetime"`
		PlayerID string                   `json:"player_id"`
		Segments []map[string]interface{} `json:"segments"`
	}
)

func (c *RESTClient) GetPlayer(nickname string, game string, game_player_id string) (*Player, error) {
	var player Player

	params := url.Values{}
	params.Add("nickname", nickname)
	params.Add("game", game)
	params.Add("game_player_id", game_player_id)

	err := c.getJSON("/players", &player, params)

	return &player, err
}

func (c *RESTClient) GetPlayerByID(player_id string) (*Player, error) {
	var player Player

	err := c.getJSON(fmt.Sprintf("/players/%s", player_id), &player, nil)

	return &player, err
}

func (c *RESTClient) GetPlayerMatchHistory(player_id string, offset string, limit string) (*PlayerMatches, error) {
	var matches PlayerMatches

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	err := c.getJSON(fmt.Sprintf("/players/%s/history", player_id), &matches, params)

	return &matches, err
}

func (c *RESTClient) GetPlayerHubs(player_id string, offset string, limit string) (*Hubs, error) {
	var hubs Hubs

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	err := c.getJSON(fmt.Sprintf("/players/%s/hubs", player_id), &hubs, params)

	return &hubs, err
}

func (c *RESTClient) GetPlayerStats(player_id string, game_id string) (*PlayerStats, error) {
	var stats PlayerStats

	err := c.getJSON(fmt.Sprintf("/players/%s/stats/%s", player_id, game_id), &stats, nil)

	return &stats, err
}

func (c *RESTClient) GetPlayerTournaments(player_id string, offset string, limit string) (*Tournaments, error) {
	var tournaments Tournaments

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	err := c.getJSON(fmt.Sprintf("/players/%s/tournaments", player_id), &tournaments, params)

	return &tournaments, err
}
