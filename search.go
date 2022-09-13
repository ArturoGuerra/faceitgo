package faceitgo

type (
	SearchChampionship struct {
		CompetitionID   string `json:"competition_id"`
		CompetitionName string `json:"competition_name"`
		Game            string `json:"game"`
		Name            string `json:"name"`
		NumberOfMembers int    `json:"number_of_members"`
		OrganizerID     string `json:"organizer_id"`
		OrganizerName   string `json:"organizer_name"`
		OrganizerType   string `json:"organizer_type"`
		PlayersCheckin  int    `json:"players_checkin"`
		PlayersJoined   int    `json:"players_joined"`
		Region          string `json:"region"`
		Slots           int    `json:"slots"`
		StartedAt       string `json:"started_at"`
		Status          string `json:"status"`
		TotalPrize      int    `json:"total_prize"`
	}

	SearchChampionships struct {
		Start int                  `json:"start"`
		End   int                  `json:"end"`
		Items []SearchChampionship `json:"items"`
	}

	SearchHub struct {
		CompetitionID   string `json:"competition_id"`
		CompetitionName string `json:"competition_name"`
		Game            string `json:"game"`
		Name            string `json:"name"`
		NumberOfMembers int    `json:"number_of_members"`
		OrganizerID     string `json:"organizer_id"`
		OrganizerName   string `json:"organizer_name"`
		OrganizerType   string `json:"organizer_type"`
		PlayersCheckin  int    `json:"players_checkin"`
		PlayersJoined   int    `json:"players_joined"`
		Region          string `json:"region"`
		Slots           int    `json:"slots"`
		StartedAt       string `json:"started_at"`
		Status          string `json:"status"`
		TotalPrize      int    `json:"total_prize"`
	}

	SearchHubs struct {
		Start int         `json:"start"`
		End   int         `json:"end"`
		Items []SearchHub `json:"items"`
	}

	SearchOrganizer struct {
		Active      bool     `json:"active"`
		Avatar      string   `json:"avatar"`
		Countries   []string `json:"countries"`
		Games       []string `json:"games"`
		Name        string   `json:"name"`
		OrganizerID string   `json:"organizer_id"`
		Partner     bool     `json:"partner"`
		Regions     []string `json:"regions"`
	}

	SearchOrganizers struct {
		Start int               `json:"start"`
		End   int               `json:"end"`
		Items []SearchOrganizer `json:"items"`
	}

	SearchPlayer struct {
		Avatar  string `json:"avatar"`
		Country string `json:"country"`
		Games   []struct {
			Name       string `json:"name"`
			SkillLevel string `json:"skill_level"`
		}
		Nickname string `json:"nickname"`
		PlayerID string `json:"player_id"`
		Status   string `json:"status"`
		Verified bool   `json:"verified"`
	}

	SearchPlayers struct {
		Start int            `json:"start"`
		End   int            `json:"end"`
		Items []SearchPlayer `json:"items"`
	}

	SearchTeam struct {
		Avatar     string `json:"avatar"`
		ChatRoomID string `json:"chat_room_id"`
		FaceitUrl  string `json:"faceit_url"`
		Game       string `json:"game"`
		Name       string `json:"name"`
		TeamID     string `json:"team_id"`
		Verified   bool   `json:"verified"`
	}

	SearchTeams struct {
		Start int          `json:"start"`
		End   int          `json:"end"`
		Items []SearchTeam `json:"items"`
	}

	SearchTournament struct {
		CompetitionID   string `json:"competition_id"`
		CompetitionName string `json:"competition_name"`
		Game            string `json:"game"`
		Name            string `json:"name"`
		NumberOfMembers int    `json:"number_of_members"`
		OrganizerID     string `json:"organizer_id"`
		OrganizerName   string `json:"organizer_name"`
		OrganizerType   string `json:"organizer_type"`
		PlayersCheckin  int    `json:"players_checkin"`
		PlayersJoined   int    `json:"players_joined"`
		Region          string `json:"region"`
		Slots           int    `json:"slots"`
		StartedAt       string `json:"started_at"`
		Status          string `json:"status"`
		TotalPrize      int    `json:"total_prize"`
	}
	SearchTournaments struct {
		Start int                `json:"start"`
		End   int                `json:"end"`
		Items []SearchTournament `json:"items"`
	}
)
