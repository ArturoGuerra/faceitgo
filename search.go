package faceitgo

type (
	SearchChampionships struct {
		Start int            `json:"start"`
		End   int            `json:"end"`
		Items []Championship `json:"items"`
	}

	SearchHubs struct {
		Start int   `json:"start"`
		End   int   `json:"end"`
		Items []Hub `json:"items"`
	}

	SearchOrganizers struct {
		Start int         `json:"start"`
		End   int         `json:"end"`
		Items []Organizer `json:"items"`
	}

	SearchPlayers struct {
		Start int      `json:"start"`
		End   int      `json:"end"`
		Items []Player `json:"items"`
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

	SearchTournaments struct {
		Start int          `json:"start"`
		End   int          `json:"end"`
		Items []Tournament `json:"items"`
	}
)
