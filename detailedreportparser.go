package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
)

type Flaw struct {
	Issueid                   string `xml:"issueid,attr"`
	CweName                   string `xml:"categoryname,attr"`
	CategoryID                string `xml:"categoryid,attr"`
	CategoryName              string
	Cweid                     string      `xml:"cweid,attr"`
	Remediation_status        string      `xml:"remediation_status,attr"`
	Mitigation_status         string      `xml:"mitigation_status,attr"`
	Affects_policy_compliance string      `xml:"affects_policy_compliance,attr"`
	Date_first_occurrence     string      `xml:"date_first_occurrence,attr"`
	Severity                  string      `xml:"severity,attr"`
	ExploitLevel              string      `xml:"exploitLevel,attr"`
	Module                    string      `xml:"module,attr"`
	Sourcefile                string      `xml:"sourcefile,attr"`
	Line                      string      `xml:"line,attr"`
	Description               string      `xml:"description,attr"`
	Mitigations               Mitigations `xml:"mitigations"`
	Annotations               Annotations `xml:"annotations"`
}

type Mitigations struct {
	Mitigation []Mitigation `xml:"mitigation"`
}

type Mitigation struct {
	Action      string `xml:"action,attr"`
	Description string `xml:"description,attr"`
	User        string `xml:"user,attr"`
	Date        string `xml:"date,attr"`
}

type Annotations struct {
	Annotation []Annotation `xml:"annotation"`
}

type Annotation struct {
	Action      string `xml:"action,attr"`
	Description string `xml:"description,attr"`
	User        string `xml:"user,attr"`
	Date        string `xml:"date,attr"`
}

type CustomField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func ParseDetailedReport(username, password, build_id string) ([]Flaw, []CustomField, error) {
	var flaws []Flaw
	var customFields []CustomField
	var errMsg error = nil

	detailedReportAPI, err := detailedReport(username, password, build_id)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(detailedReportAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for errors, flaws, and custom field
			if se.Name.Local == "error" {
				errMsg = errors.New("api for GetDetailedReport returned with an error element")
			}
			if se.Name.Local == "flaw" {
				var flaw Flaw
				decoder.DecodeElement(&flaw, &se)
				flaw.CategoryName = categoryMap[flaw.CategoryID]
				flaws = append(flaws, flaw)
			}
			if se.Name.Local == "customfield" {
				var cField CustomField
				decoder.DecodeElement(&cField, &se)
				customFields = append(customFields, cField)
			}
		}
	}
	return flaws, customFields, errMsg

}
