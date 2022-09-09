package faceitgo

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Championship struct {
		AnticheatRequired    bool   `json:"anticheat_required"`
		Avatar               string `json:"avatar"`
		BackgroundImage      string `json:"background_image"`
		ChampionshipID       string `json:"championship_id"`
		ChampionshipStart    int    `json:"championship_start"`
		CheckinClear         int    `json:"checkin_clear"`
		CheckinEnabled       bool   `json:"checkin_enabled"`
		CheckinStart         int    `json:"checkin_start"`
		CoverImage           string `json:"cover_image"`
		CurrentSubscriptions int    `json:"current_subscriptions"`
		Description          string `json:"description"`
		FaceitUrl            string `json:"faceit_url"`
		Featured             bool   `json:"featured"`
		Full                 bool   `json:"full"`
		GamaData             Game   `json:"game_data"`
		GameID               string `json:"game_id"`
		ID                   string `json:"id"`
		JoinChecks           struct {
			AllowedTeamTypes                []string `json:"allowed_team_types"`
			BlacklistGeoCountries           []string `json:"blacklist_geo_countries"`
			JoinPolicy                      string   `json:"join_policy"`
			MaxSkillLevel                   int      `json:"max_skill_level"`
			MembershipType                  string   `json:"membership_type"`
			MinSkillLevel                   int      `json:"min_skill_level"`
			WhitelistGeoCountries           []string `json:"whitelist_geo_countries"`
			WhitelistGeoCountriesMinPlayers int      `json:"whitelist_geo_countries_min_players"`
		} `json:"join_checks"`
		Name          string    `json:"name"`
		OrganizerData Organizer `json:"organizer_data"`
		OrganizerID   string    `json:"organizer_id"`
		Prizes        []struct {
			FaceitPoints int `json:"faceit_points"`
			Rank         int `json:"rank"`
		} `json:"prizes"`
		Region   string `json:"region"`
		RulesID  string `json:"rules_id"`
		Schedule map[string]struct {
			Data   int    `json:"data"`
			Status string `json:"status"`
		} `json:"schedule"`
		SeedingStrategy string `json:"seeding_strategy"`
		Slots           int    `json:"slots"`
		Status          string `json:"status"`
		Stream          struct {
			Active   bool   `json:"active"`
			Platform string `json:"platform"`
			Source   string `json:"source"`
			Title    string `json:"title"`
		} `json:"stream"`
		SubscriptionEnd           int  `json:"subscription_end"`
		SubscriptionStart         int  `json:"subscription_start"`
		SubscriptionLocked        bool `json:"subscription_locked"`
		SubstitutionConfiguration struct {
			MaxSubstitutions int `json:"max_substitutions"`
			MaxSubstitutes   int `json:"max_substitutes"`
		} `json:"substitution_configuration"`
		TotalGroups int    `json:"total_groups"`
		TotalPrizes int    `json:"total_prizes"`
		TotalRounds int    `json:"total_rounds"`
		Type        string `json:"type"`
	}

	Championships struct {
		Start int            `json:"start"`
		End   int            `json:"end"`
		Items []Championship `json:"items"`
	}

	ChampionshipMatch struct {
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
		Version     int    `json:"version"`
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
		} `json:"teams"`
	}

	ChampionshipMatches struct {
		Start int                 `json:"start"`
		End   int                 `json:"end"`
		Items []ChampionshipMatch `json:"items"`
	}

	ChampionshipResult struct {
		Bounds struct {
			Left  int `json:"left"`
			Right int `json:"right"`
		} `json:"bounds"`
		Placements []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"placements"`
	}

	ChampionshipResults struct {
		Start int                  `json:"start"`
		End   int                  `json:"end"`
		Items []ChampionshipResult `json:"items"`
	}

	ChampionshipSubscription struct {
		Coach       string   `json:"coach"`
		Coleader    string   `json:"coleader"`
		Group       int      `json:"group"`
		Leader      string   `json:"leader"`
		Roster      []string `json:"roster"`
		Status      string   `json:"status"`
		Substitutes []string `json:"substitutes"`
		Team        struct {
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
		} `json:"team"`
	}

	ChampionshipSubscriptions struct {
		Start int                        `json:"start"`
		End   int                        `json:"end"`
		Items []ChampionshipSubscription `json:"items"`
	}
)

func (c *RESTClient) GetChampionships(game string, gtype string, offset string, limit string) (*Championships, error) {
	var championships Championships

	if err := c.getJSON("/championships", &championships); err != nil {
		return nil, err
	}

	return &championships, nil
}

func (c *RESTClient) GetChampionship(championshipID string) (*Championship, error) {
	var championship Championship
	resp, err := c.get("/championships/" + championshipID)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &championship)
	if err != nil {
		return nil, err
	}

	return &championship, err
}

func (c *RESTClient) GetChampionshipMatches(championshipID string, gtype string, offset string, limit string) (*ChampionshipMatches, error) {
	var championshipMatches ChampionshipMatches
	resp, err := c.get("/championships/" + championshipID + "/matches")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &championshipMatches)
	if err != nil {
		return nil, err
	}

	return &championshipMatches, err
}

func (c *RESTClient) GetChampionshipResults(championshipID string, offset string, limit string) (*ChampionshipResults, error) {
	var championshipResults ChampionshipResults
	resp, err := c.get("/championships/" + championshipID + "/results")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &championshipResults)
	if err != nil {
		return nil, err
	}

	return &championshipResults, err
}

func (c *RESTClient) GetChampionshipSubscriptions(championshipID string, offset string, limit string) (*ChampionshipSubscriptions, error) {
	var championshipSubscriptions ChampionshipSubscriptions
	resp, err := c.get("/championships/" + championshipID + "/subscriptions")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &championshipSubscriptions)
	if err != nil {
		return nil, err
	}

	return &championshipSubscriptions, err
}
