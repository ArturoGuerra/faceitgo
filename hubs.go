package faceitgo

import (
	"fmt"
	"net/url"
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

	HubMatch struct {
		BestOf                  int      `json:"best_of"`
		BroadcastStartTime      int      `json:"broadcast_start_time"`
		BroadcastStartTimeLabel string   `json:"broadcast_start_time_label"`
		CalculateElo            bool     `json:"calculate_elo"`
		ChatRoomID              string   `json:"chat_room_id"`
		CompetitionID           string   `json:"competition_id"`
		CompetitionName         string   `json:"competition_name"`
		CompetitionType         string   `json:"competition_type"`
		ConfiguredAt            int      `json:"configured_at"`
		DemoUrl                 []string `json:"demo_url"`
		FaceitUrl               string   `json:"faceit_url"`
		FinishedAt              int      `json:"finished_at"`
		Game                    string   `json:"game"`
		Group                   int      `json:"group"`
		MatchID                 string   `json:"match_id"`
		OrganizerID             string   `json:"organizer_id"`
		Region                  string   `json:"region"`
		Results                 struct {
			Score map[string]int `json:"score"`
		} `json:"results"`
		Round       int    `json:"round"`
		ScheduledAt int    `json:"scheduled_at"`
		StartedAt   int    `json:"started_at"`
		Status      string `json:"status"`
		Teams       map[string]struct {
			Avatar    string `json:"avatar"`
			FactionID string `json:"faction_id"`
			Leader    string `json:"leader"`
			Name      string `json:"name"`
			Roster    []struct {
				AnticheatRequired bool   `json:"anticheat_required"`
				Avatar            string `json:"avatar"`
				GamePlayerID      string `json:"game_player_id"`
				GamePlayerName    string `json:"game_player_name"`
				GameSkillLevel    int    `json:"game_skill_level"`
				Membership        string `json:"membership"`
				Nickname          string `json:"nickname"`
				PlayerID          string `json:"player_id"`
			} `json:"roster"`
			RosterV1    map[string]interface{} `json:"roster_v1"`
			Substituted bool                   `json:"substituted"`
			Type        string                 `json:"type"`
		} `json:"teams"`
		Version int                    `json:"version"`
		Voting  map[string]interface{} `json:"voting"`
	}

	HubMatches struct {
		Start int      `json:"start"`
		Limit int      `json:"limit"`
		Items HubMatch `json:"items"`
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
)

func (c *RESTClient) GetHub(id string) (*Hub, error) {
	var hub Hub

	if err := c.getJSON(fmt.Sprintf("/hubs/%s", id), &hub); err != nil {
		return nil, err
	}

	return &hub, nil
}

func (c *RESTClient) GetHubMatches(id string, gtype string, offset string, limit string) (*HubMatches, error) {
	var hubMatches HubMatches

	params := url.Values{}
	params.Add("type", gtype)
	params.Add("offset", offset)
	params.Add("limit", limit)

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/matches?%s", id, params.Encode()), &hubMatches); err != nil {
		return nil, err
	}

	return &hubMatches, nil
}

func (c *RESTClient) GetHubMembers(id string, offset string, limit string) (*HubMembers, error) {
	var hubMembers HubMembers

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/members?%s", id, params.Encode()), &hubMembers); err != nil {
		return nil, err
	}

	return &hubMembers, nil
}

func (c *RESTClient) GetHubRoles(id string, offset string, limit string) (*HubRoles, error) {
	var hubRoles HubRoles

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/roles?%s", id, params.Encode()), &hubRoles); err != nil {
		return nil, err
	}

	return &hubRoles, nil
}

func (c *RESTClient) GetHubRules(id string, offset string, limit string) (*HubRules, error) {
	var hubRules HubRules

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/rules?%s", id, params.Encode()), &hubRules); err != nil {
		return nil, err
	}

	return &hubRules, nil
}

func (c *RESTClient) GetHubStats(id string, offset string, limit string) (*HubStats, error) {
	var hubStats HubStats

	params := url.Values{}
	params.Add("offset", offset)
	params.Add("limit", limit)

	if err := c.getJSON(fmt.Sprintf("/hubs/%s/stats?%s", id, params.Encode()), &hubStats); err != nil {
		return nil, err
	}

	return &hubStats, nil
}
