package faceitgo

type (
	Tournament struct {
		AnticheatRequired           bool     `json:"anticheat_required"`
		Custom                      bool     `json:"custom"`
		FaceitUrl                   string   `json:"faceit_url"`
		FeaturedImage               string   `json:"featured_image"`
		GameID                      string   `json:"game_id"`
		InviteType                  string   `json:"invite_type"`
		MatchType                   string   `json:"match_type"`
		MaxSkill                    int      `json:"max_skill"`
		MembershipType              string   `json:"membership_type"`
		MinSkill                    int      `json:"min_skill"`
		Name                        string   `json:"name"`
		NumberOfPlayers             int      `json:"number_of_players"`
		NumberOfPlayersCheckedin    int      `json:"number_of_players_checkedin"`
		NumberOfPlayersJoined       int      `json:"number_of_players_joined"`
		NumberOfPlayersParticipants int      `json:"number_of_players_participants"`
		OrganizerID                 string   `json:"organizer_id"`
		PrizeType                   string   `json:"prize_type"`
		Region                      string   `json:"region"`
		StartedAt                   string   `json:"started_at"`
		SubscriptionsCount          int      `json:"subscriptions_count"`
		TeamSize                    int      `json:"team_size"`
		TotalPrize                  int      `json:"total_prize"`
		TournamentID                string   `json:"tournament_id"`
		WhitelistCountries          []string `json:"whitelist_countries"`
	}

	Tournaments struct {
		Start int          `json:"start"`
		End   int          `json:"end"`
		Items []Tournament `json:"items"`
	}

	TournamentDetails struct {
		AnticheapRequired           bool          `json:"anticheat_required"`
		BestOf                      int           `json:"best_of"`
		CalculateElo                bool          `json:"calculate_elo"`
		CompetitionID               string        `json:"competition_id"`
		CoverImage                  string        `json:"cover_image"`
		Custom                      bool          `json:"custom"`
		Description                 string        `json:"description"`
		FaceitUrl                   string        `json:"faceit_url"`
		FeaturedImage               string        `json:"featured_image"`
		GameData                    Game          `json:"game_data"`
		GameID                      string        `json:"game_id"`
		InviteType                  string        `json:"invite_type"`
		MatchType                   string        `json:"match_type"`
		MaxSkill                    int           `json:"max_skill"`
		MembershipType              string        `json:"membership_type"`
		MinSkill                    int           `json:"min_skill"`
		Name                        string        `json:"name"`
		NumberOfPlayers             int           `json:"number_of_players"`
		NumberOfPlayersCheckedin    int           `json:"number_of_players_checkedin"`
		NumberOfPlayersJoined       int           `json:"number_of_players_joined"`
		NumberOfPlayersParticipants int           `json:"number_of_players_participants"`
		OrganizerID                 string        `json:"organizer_id"`
		OrganizerData               Organizer     `json:"organizer_data"`
		PrizeType                   string        `json:"prize_type"`
		Region                      string        `json:"region"`
		Rounds                      []interface{} `json:"rounds"`
		Rule                        string        `json:"rule"`
		StartedAt                   string        `json:"started_at"`
		Status                      string        `json:"status"`
		SubstitutesAllowed          int           `json:"substitutes_allowed"`
		SubstitutionsAllowed        int           `json:"substitutions_allowed"`
		TeamSize                    int           `json:"team_size"`
		TotalPrize                  interface{}   `json:"total_prize"`
		TournamentID                string        `json:"tournament_id"`
		Voting                      interface{}   `json:"voting"`
		WhitelistCountries          []string      `json:"whitelist_countries"`
	}

	TournamentBracket struct {
		Game    string `json:"game"`
		Matches []struct {
			FaceitUrl string `json:"faceit_url"`
			MatchID   string `json:"match_id"`
			Position  int    `json:"position"`
			Results   struct {
				Score  map[string]int `json:"score"`
				Winner string         `json:"winner"`
			} `json:"results"`
			Round int    `json:"round"`
			Stats string `json:"stats"`
			Teams map[string]struct {
				Avatar   string `json:"avatar"`
				Nickname string `json:"nickname"`
				TeamID   string `json:"team_id"`
			} `json:"teams"`
		} `json:"matches"`
		Name   string `json:"name"`
		Rounds []struct {
			BestOf               int    `json:"best_of"`
			Label                string `json:"label"`
			Matches              int    `json:"matches"`
			Round                int    `json:"round"`
			StartTime            int    `json:"start_time"`
			StartsAsap           bool   `json:"starts_asap"`
			SubstitutionName     int    `json:"substitution_name"`
			SubstitutionsAllowed bool   `json:"substitutions_allowed"`
		} `json:"rounds"`
		Status string `json:"status"`
	}

	TournamentMatches struct {
		Start int     `json:"start"`
		End   int     `json:"end"`
		Items []Match `json:"items"`
	}

	TournamentTeam struct {
		Nickname   string `json:"nickname"`
		Skilllevel int    `json:"skilllevel"`
		SubsDone   int    `json:"subs_done"`
		TeamID     string `json:"team_id"`
		TeamLeader string `json:"team_leader"`
		TeamType   string `json:"team_type"`
	}

	TournamentTeams struct {
		CheckedIn []TournamentTeam `json:"checked_in"`
		Finished  []TournamentTeam `json:"finished"`
		Joined    []TournamentTeam `json:"joined"`
		Started   []TournamentTeam `json:"started"`
	}
)
