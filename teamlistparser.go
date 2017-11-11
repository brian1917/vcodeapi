package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
)

//Team represents a Veracode team of users
type Team struct {
	TeamID       string `xml:"team_id,attr"`
	TeamName     string `xml:"team_name,attr"`
	CreationDate string `xml:"creation_date,attr"`
}

//ParseTeamList calls the getteamlist.do API and returns an array of teams
func ParseTeamList(credsFile string) ([]Team, error) {
	var teams []Team

	teamListAPI, err := teamList(credsFile)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(bytes.NewReader(teamListAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for flaw
			if se.Name.Local == "team" {
				var team Team
				decoder.DecodeElement(&team, &se)
				teams = append(teams, team)
			}
			if se.Name.Local == "error" {
				return nil, errors.New("api for GetTeamList returned with an error element")
			}
		}
	}
	return teams, nil
}
