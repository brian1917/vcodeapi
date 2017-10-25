package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
)

//TeamInfo represents the Team Information for a Veracode Team
type TeamInfo struct {
	Test     string `xml:"xmlns,attr"`
	TeamName string `xml:"team_name,attr"`
	Users    []User `xml:"user"`
	Apps     []App  `xml:"application"`
}

//User represents a User in the Veracode Platform
type User struct {
	Username  string `xml:"username,attr"`
	FirstName string `xml:"first_name,attr"`
	LastName  string `xml:"last_name,attr"`
	Email     string `xml:"email_address,attr"`
}

// App Struct is declared in applistparser.go

// ParseTeamInfo calls the Veracode getteaminfo.do API and returns a TeamInfo struct
func ParseTeamInfo(credsFile, teamID string, includeUsers, includeApplications bool) (TeamInfo, error) {
	var errMsg error
	var team TeamInfo

	teamInfoAPI, err := teamInfo(credsFile, teamID, includeUsers, includeApplications)
	if err != nil {
		log.Fatal(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(teamInfoAPI))
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
			if se.Name.Local == "teaminfo" {
				decoder.DecodeElement(&team, &se)
			}
			if se.Name.Local == "error" {
				errMsg = errors.New("api for GetTeamInfo returned with an error element")
			}
		}
	}
	return team, errMsg
}
