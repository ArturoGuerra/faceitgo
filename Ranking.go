package faceitgo

type (
	Ranking struct {
		Country        string `json:"country"`
		FaceitElo      int    `json:"faceit_elo"`
		GameSkillLevel int    `json:"game_skill_level"`
		Nickname       string `json:"nickname"`
		PlayerID       string `json:"player_id"`
		Position       int    `json:"position"`
	}

	Rankings struct {
		Start int       `json:"start"`
		End   int       `json:"end"`
		Items []Ranking `json:"items"`
	}
)
