package faceitgo

import (
	"fmt"
	"net/url"
	"strconv"
)

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

func (c *RESTClient) GetOrganizerByName(organizer_name string) (*Organizer, error) {
	var organizer Organizer

	if err := c.getJSON(fmt.Sprintf("/organizers?name=%s", organizer_name), &organizer); err != nil {
		return nil, err
	}

	return &organizer, nil
}

func (c *RESTClient) GetOrganizer(organizer_id string) (*Organizer, error) {
	var organizer Organizer

	if err := c.getJSON(fmt.Sprintf("/organizers/%s", organizer_id), &organizer); err != nil {
		return nil, err
	}

	return &organizer, nil
}

func (c *RESTClient) GetOrganizersChampionships(organizer_id string, offset int, limit int) (*OrganizerChampionships, error) {
	var championships OrganizerChampionships

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/organizers/%s/championships?%s", organizer_id, params.Encode()), &championships); err != nil {
		return nil, err
	}

	return &championships, nil
}

func (c *RESTClient) GetOrganizersGames(organizer_id string) ([]string, error) {
	var games []string

	if err := c.getJSON(fmt.Sprintf("/organizers/%s/games", organizer_id), &games); err != nil {
		return nil, err
	}

	return games, nil
}

func (c *RESTClient) GetOrganizersHubs(organizer_id string, offset int, limit int) (*OrganizerHubs, error) {
	var hubs OrganizerHubs

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/organizers/%s/hubs?%s", organizer_id, params.Encode()), &hubs); err != nil {
		return nil, err
	}

	return &hubs, nil
}

func (c *RESTClient) GetOrganizersTournaments(organizer_id string, game_type string, offset int, limit int) (*OrganizerTournaments, error) {
	var tournaments OrganizerTournaments

	params := url.Values{}
	params.Add("type", game_type)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if err := c.getJSON(fmt.Sprintf("/organizers/%s/tournaments?%s", organizer_id, params.Encode()), &tournaments); err != nil {
		return nil, err
	}

	return &tournaments, nil
}
