package faceitgo

import (
	"net/url"
	"strconv"
)

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

func (c *RESTClient) SearchChampionships(name string, game string, region string, game_type string, offset int, limit int) (*SearchChampionships, error) {
	var searchChampionships SearchChampionships

	params := url.Values{}
	params.Add("name", name)
	params.Add("game", game)
	params.Add("region", region)
	params.Add("type", game_type)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON("/search/championships", &searchChampionships, params)

	return &searchChampionships, err
}

func (c *RESTClient) SearchHubs(name string, game string, region string, offset int, limit int) (*SearchHubs, error) {
	var searchHubs SearchHubs

	params := url.Values{}
	params.Add("name", name)
	params.Add("game", game)
	params.Add("region", region)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON("/search/hubs", &searchHubs, params)

	return &searchHubs, err
}

func (c *RESTClient) SearchOrganizers(name string, offset int, limit int) (*SearchOrganizers, error) {
	var searchOrganizers SearchOrganizers

	params := url.Values{}
	params.Add("name", name)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON("/search/organizers", &searchOrganizers, params)

	return &searchOrganizers, err
}

func (c *RESTClient) SearchPlayers(nickname string, game string, country string, offset int, limit int) (*SearchPlayers, error) {
	var searchPlayers SearchPlayers

	params := url.Values{}
	params.Add("nickname", nickname)
	params.Add("game", game)
	params.Add("country", country)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON("/search/players", &searchPlayers, params)

	return &searchPlayers, err
}

func (c *RESTClient) SearchTeams(nickname string, game string, offset int, limit int) (*SearchTeams, error) {
	var searchTeams SearchTeams

	params := url.Values{}
	params.Add("name", nickname)
	params.Add("game", game)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON("/search/teams", &searchTeams, params)

	return &searchTeams, err
}

func (c *RESTClient) SearchTournaments(name string, game string, region string, game_type string, offset int, limit int) (*SearchTournaments, error) {
	var searchTournaments SearchTournaments

	params := url.Values{}
	params.Add("name", name)
	params.Add("game", game)
	params.Add("region", region)
	params.Add("type", game_type)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	err := c.getJSON("/search/tournaments", &searchTournaments, params)

	return &searchTournaments, err
}
