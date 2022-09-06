package faceitgo

type (
	Organizer struct {
		Avatar         string `json:"avatar"`
		Cover          string `json:"cover"`
		Description    string `json:"description"`
		Facebook       string `json:"facebook"`
		FaceitUrl      string `json:"faceit_url"`
		FollowersCount int    `json:"followers_count"`
		Name           string `json:"name"`
		OrganizerID    string `json:"organizer_id"`
		Twitch         string `json:"twitch"`
		Twitter        string `json:"twitter"`
		Type           string `json:"type"`
		VK             string `json:"vk"`
		Website        string `json:"website"`
		YouTube        string `json:"youtube"`
	}

	OrganizerChampionships struct {
		Start int            `json:"start"`
		End   int            `json:"end"`
		Items []Championship `json:"items"`
	}

	OrganizerHubs struct {
		Start int   `json:"start"`
		End   int   `json:"end"`
		Items []Hub `json:"items"`
	}

	OrganizerTournaments struct {
		Start int          `json:"start"`
		End   int          `json:"end"`
		Items []Tournament `json:"items"`
	}
)
