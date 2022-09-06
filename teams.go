package faceitgo

type (
	Team struct {
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
	}

	TeamStats struct {
		GameID   string                   `json:"game_id"`
		Lifetime map[string]interface{}   `json:"lifetime"`
		Segments []map[string]interface{} `json:"segments"`
		TeamID   string                   `json:"team_id"`
	}
)
