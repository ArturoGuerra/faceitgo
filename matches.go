package faceitgo

type (
	Match struct {
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
		MatchID                 string   `json:"match_id"`
		OrganizerID             string   `json:"organizer_id"`
		Region                  string   `json:"region"`
		Results                 struct {
			Score map[string]int `json:"score"`
		} `json:"results"`
		Round int `json:"round"`
		Teams []struct {
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

	MatchStats struct {
		Rounds []struct {
			BestOf        int                    `json:"best_of"`
			CompetitionID string                 `json:"competition_id"`
			GameID        string                 `json:"game_id"`
			GameMode      string                 `json:"game_mode"`
			MatchID       string                 `json:"match_id"`
			MatchRound    int                    `json:"match_round"`
			Played        string                 `json:"played"`
			RoundStats    map[string]interface{} `json:"round_stats"`
			Teams         map[string]struct {
				Players map[string]struct {
					Nickname    string                 `json:"nickname"`
					PlayerID    string                 `json:"player_id"`
					PlayerStats map[string]interface{} `json:"player_stats"`
				} `json:"players"`
				Premade   bool                   `json:"premade"`
				TeamID    string                 `json:"team_id"`
				TeamStats map[string]interface{} `json:"team_stats"`
			} `json:"teams"`
		} `json:"rounds"`
	}
)
