package faceitgo

type (
	RoomMember struct {
		IsOnline bool     `json:"is_online"`
		MemberID string   `json:"member_id"`
		Nickname string   `json:"nickname"`
		Photo    string   `json:"photo"`
		Roles    []string `json:"roles"`
		Status   string   `json:"status"`
	}

	RoomRoles struct {
		Color       string   `json:"color"`
		Displayed   bool     `json:"displayed"`
		Mentionable bool     `json:"mentionable"`
		Name        string   `json:"name"`
		Permissions []string `json:"permissions"`
		Ranking     int      `json:"ranking"`
		RoleID      string   `json:"role_id"`
	}

	Room struct {
		Members []RoomMember `json:"members"`
		Name    string       `json:"name"`
		Roles   []RoomRoles  `json:"roles"`
	}

	RoomMessage struct {
		Avatar    string `json:"avatar"`
		Body      string `json:"body"`
		From      string `json:"from"`
		ID        string `json:"id"`
		Nickname  string `json:"nickname"`
		Timestamp string `json:"timestamp"`
	}

	RoomMessages struct {
		IsLastPage bool          `json:"is_last_page"`
		Messages   []RoomMessage `json:"messages"`
	}
)
