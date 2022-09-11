package faceitgo

import (
	"fmt"
	"net/url"
	"strconv"
)

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

	TournamentBrackets struct {
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

func (c *RESTClient) GetTournament(tournament_id string) (TournamentDetails, error) {
	var tournament TournamentDetails
	err := c.getJSON(fmt.Sprintf("/tournaments/%s", tournament_id), &tournament, nil)
	return tournament, err
}

func (c *RESTClient) GetTournamentBrackets(tournament_id string) (TournamentBrackets, error) {
	var tournament TournamentBrackets
	err := c.getJSON(fmt.Sprintf("/tournaments/%s/brackets", tournament_id), &tournament, nil)
	return tournament, err
}

func (c *RESTClient) GetTournamentMatches(tournament_id string, offset int, limit int) (TournamentMatches, error) {
	var tournament TournamentMatches

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON(fmt.Sprintf("/tournaments/%s/matches", tournament_id), &tournament, params)
	return tournament, err
}

func (c *RESTClient) GetTournamentTeams(tournament_id string, offset int, limit int) (TournamentTeams, error) {
	var tournament TournamentTeams

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON(fmt.Sprintf("/tournaments/%s/teams", tournament_id), &tournament, params)
	return tournament, err
}
