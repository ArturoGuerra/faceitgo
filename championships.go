package faceitgo

import (
	"fmt"
	"net/url"
	"strconv"
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

func (c *RESTClient) GetChampionships(game_id string, game_type *string, offset, limit *int) (*Championships, error) {
	var championships Championships

	params := url.Values{}
	params.Add("game", game_id)
	if game_type != nil {
		params.Add("type", *game_type)
	}

	if offset != nil {
		params.Add("offset", strconv.Itoa(*offset))
	}

	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	if err := c.getJSON("/championships", &championships, params); err != nil {
		return nil, err
	}

	return &championships, nil
}

func (c *RESTClient) GetChampionship(championship_id string) (*Championship, error) {
	var championship Championship

	if err := c.getJSON(fmt.Sprintf("/championships/%s", championship_id), &championship, nil); err != nil {
		return nil, err
	}

	return &championship, nil
}

func (c *RESTClient) GetChampionshipMatches(championship_id string, game_type *string, offset, limit *int) (*Matches, error) {
	var championshipMatches Matches

	params := url.Values{}
	if game_type != nil {
		params.Add("type", *game_type)
	}

	if offset != nil {
		params.Add("offset", strconv.Itoa(*offset))
	}

	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	u := fmt.Sprintf("/championships/%s/matches", championship_id)
	if err := c.getJSON(u, &championshipMatches, params); err != nil {
		return nil, err
	}

	return &championshipMatches, nil
}

func (c *RESTClient) GetChampionshipResults(championship_id string, offset, limit *int) (*ChampionshipResults, error) {
	var championshipResults ChampionshipResults

	params := url.Values{}
	if offset != nil {
		params.Add("offset", strconv.Itoa(*offset))
	}

	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	u := fmt.Sprintf("/championships/%s/results", championship_id)

	if err := c.getJSON(u, &championshipResults, params); err != nil {
		return nil, err
	}

	return &championshipResults, nil
}

func (c *RESTClient) GetChampionshipSubscriptions(championship_id string, offset, limit *int) (*ChampionshipSubscriptions, error) {
	var championshipSubscriptions ChampionshipSubscriptions

	params := url.Values{}
	if offset != nil {
		params.Add("offset", strconv.Itoa(*offset))
	}

	if limit != nil {
		params.Add("limit", strconv.Itoa(*limit))
	}

	if err := c.getJSON(fmt.Sprintf("/championships/%s/subscriptions", championship_id), &championshipSubscriptions, params); err != nil {
		return nil, err
	}

	return &championshipSubscriptions, nil
}
