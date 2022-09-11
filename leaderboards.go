package faceitgo

import (
	"fmt"
	"net/url"
	"strconv"
)

type (
	Leaderboard struct {
		CompetitionID   string `json:"competition_id"`
		CompetitionType string `json:"competition_type"`
		EndData         int    `json:"end_date"`
		GameID          string `json:"game_id"`
		Group           int    `json:"group"`
		LeaderboardID   string `json:"leaderboard_id"`
		LeaderboardMode string `json:"leaderboard_mode"`
		LeaderboardType string `json:"leaderboard_type"`
		MinMatches      int    `json:"min_matches"`
		PointsPerDraw   int    `json:"points_per_draw"`
		PointsPerLoss   int    `json:"points_per_loss"`
		PointsPerWin    int    `json:"points_per_win"`
		PointsType      string `json:"points_type"`
		RankingBoost    int    `json:"ranking_boost"`
		RankingType     string `json:"ranking_type"`
		Region          string `json:"region"`
		Round           int    `json:"round"`
		Season          int    `json:"season"`
		StartDate       int    `json:"start_date"`
		StartingPoints  int    `json:"starting_points"`
		Status          string `json:"status"`
	}

	Leaderboards struct {
		Start int           `json:"start"`
		End   int           `json:"end"`
		Items []Leaderboard `json:"items"`
	}

	LeaderboardRanking struct {
		CurrentStreak int `json:"current_streak"`
		Draw          int `json:"draw"`
		Lost          int `json:"lost"`
		Played        int `json:"played"`
		Player        struct {
			Avatar         string   `json:"avatar"`
			Country        string   `json:"country"`
			FaceitUrl      string   `json:"faceit_url"`
			MembershipType string   `json:"membership_type"`
			Memberships    []string `json:"memberships"`
			Nickname       string   `json:"nickname"`
			SkillLevel     int      `json:"skill_level"`
			UserID         string   `json:"user_id"`
		} `json:"player"`
		Points   int     `json:"points"`
		Position int     `json:"position"`
		WinRate  float64 `json:"win_rate"`
		Won      int     `json:"won"`
	}
	LeaderboardRankings struct {
		Start       int                  `json:"start"`
		End         int                  `json:"end"`
		Items       []LeaderboardRanking `json:"items"`
		Leaderboard Leaderboard          `json:"leaderboard"`
	}
)

func (c *RESTClient) GetChampionshipLeaderboards(championship_id string, offset int, limit int) (*Leaderboards, error) {
	var leaderboards Leaderboards

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	u := fmt.Sprintf("/leaderboards/championships/%s?%s", championship_id, params.Encode())
	if err := c.getJSON(u, &leaderboards); err != nil {
		return nil, err
	}

	return &leaderboards, nil
}

func (c *RESTClient) GetChampionshipsGroupRanking(championsip_id string, group int, offset int, limit int) (*LeaderboardRankings, error) {
	var leaderboardRankings LeaderboardRankings

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	u := fmt.Sprintf("/leaderboards/championships/%s/group/%d?%s", championsip_id, group, params.Encode())
	if err := c.getJSON(u, &leaderboardRankings); err != nil {
		return nil, err
	}

	return &leaderboardRankings, nil
}

func (c *RESTClient) GetHubsLeaderboards(hub_id string, offset string, limit string) (*Leaderboards, error) {
	var leaderboards Leaderboards

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	u := fmt.Sprintf("/leaderboards/hubs/%s?%s", hub_id, params.Encode())
	if err := c.getJSON(u, &leaderboards); err != nil {
		return nil, err
	}

	return &leaderboards, nil
}

func (c *RESTClient) GetHubRankings(hub_id string, offset int, limit int) (*LeaderboardRankings, error) {
	var leaderboardRankings LeaderboardRankings

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	u := fmt.Sprintf("/leaderboards/hubs/%s/general?%s", hub_id, params.Encode())
	if err := c.getJSON(u, &leaderboardRankings); err != nil {
		return nil, err
	}

	return &leaderboardRankings, nil
}

func (c *RESTClient) GetLeaderboardRankings(leaderboard_id string, offset int, limit int) (*LeaderboardRankings, error) {
	var leaderboard LeaderboardRankings

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	u := fmt.Sprintf("/leaderboards/%s?%s", leaderboard_id, params.Encode())
	if err := c.getJSON(u, &leaderboard); err != nil {
		return nil, err
	}

	return &leaderboard, nil
}
