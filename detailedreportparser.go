package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
)

// DetReport represents the detailed report returned for a build
type DetReport struct {
	AppName                string          `xml:"app_name,attr"`
	AppID                  string          `xml:"app_id,attr"`
	PolicyName             string          `xml:"policy_name,attr"`
	PolicyComplianceStatus string          `xml:"policy_compliance_status,attr"`
	PolicyRulesStatus      string          `xml:"policy_rules_status,attr"`
	GracePeriodExpired     string          `xml:"grace_period_expired,attr"`
	BusinessUnit           string          `xml:"business_unit,attr"`
	StaticAnalysis         StaticAnalysis  `xml:"static-analysis"`
	DynamicAnalysis        DynamicAnalysis `xml:"dynamic-analysis"`
	ManualAnalysis         ManualAnalysis  `xml:"manual-analysis"`
}

// StaticAnalysis represents a static scan from Veracode
type StaticAnalysis struct {
	AnalysisSize  string  `xml:"analysis_size_bytes,attr"`
	EngineVersion string  `xml:"engine_version,attr"`
	PublishedDate string  `xml:"published_date,attr"`
	Rating        string  `xml:"rating,attr"`
	Score         string  `xml:"score,attr"`
	SubmittedDate string  `xml:"submitted_date,attr"`
	Version       string  `xml:"version,attr"`
	Modules       Modules `xml:"modules"`
}

// Modules is an array of module
type Modules struct {
	Module []Module `xml:"module"`
}

// Module represents a scannable module in Veracode
type Module struct {
	Architecture string `xml:"architecture,attr"`
	Compiler     string `xml:"compiler,attr"`
	Domain       string `xml:"domain,attr"`
	Loc          string `xml:"loc,attr"`
	Name         string `xml:"name,attr"`
	Numflawssev0 string `xml:"numflawssev0,attr"`
	Numflawssev1 string `xml:"numflawssev1,attr"`
	Numflawssev2 string `xml:"numflawssev2,attr"`
	Numflawssev3 string `xml:"numflawssev3,attr"`
	Numflawssev4 string `xml:"numflawssev4,attr"`
	Numflawssev5 string `xml:"numflawssev5,attr"`
	Os           string `xml:"os,attr"`
	Score        string `xml:"score,attr"`
	TargetURL    string `xml:"target_url,attr"`
}

// DynamicAnalysis represents a dynamic scan from Veracode
type DynamicAnalysis struct {
	DynamicScanType    string  `xml:"dynamic_scan_type,attr"`
	PublishedDate      string  `xml:"published_date,attr"`
	Rating             string  `xml:"rating,attr"`
	ScanExitStatusDesc string  `xml:"scan_exit_status_desc,attr"`
	ScanExitStatusID   string  `xml:"scan_exit_status_id,attr"`
	Score              string  `xml:"score,attr"`
	SubmittedDate      string  `xml:"submitted_date,attr"`
	Version            string  `xml:"version,attr"`
	Modules            Modules `xml:"modules"`
}

// ManualAnalysis represents a manual assessment from Veracode
type ManualAnalysis struct {
	PublishedDate string  `xml:"published_date,attr"`
	Rating        string  `xml:"rating,attr"`
	Score         string  `xml:"score,attr"`
	SubmittedDate string  `xml:"submitted_date,attr"`
	Version       string  `xml:"version,attr"`
	Modules       Modules `xml:"modules"`
}

// Flaw represents a finding from a Veracode test (static, dynamic, or MPT)
type Flaw struct {
	Issueid                 string `xml:"issueid,attr"`
	CweName                 string `xml:"categoryname,attr"`
	CategoryID              string `xml:"categoryid,attr"`
	CategoryName            string
	Cweid                   string `xml:"cweid,attr"`
	RemediationStatus       string `xml:"remediation_status,attr"`
	MitigationStatus        string `xml:"mitigation_status,attr"`
	AffectsPolicyCompliance string `xml:"affects_policy_compliance,attr"`
	PolicyName              string
	DateFirstOccurrence     string      `xml:"date_first_occurrence,attr"`
	Severity                string      `xml:"severity,attr"`
	ExploitLevel            string      `xml:"exploitLevel,attr"`
	Module                  string      `xml:"module,attr"`
	Sourcefile              string      `xml:"sourcefile,attr"`
	Line                    string      `xml:"line,attr"`
	Description             string      `xml:"description,attr"`
	FlawURL                 string      `xml:"url,attr"`
	VulnParameter           string      `xml:"vuln_parameter,attr"`
	Mitigations             Mitigations `xml:"mitigations"`
	Annotations             Annotations `xml:"annotations"`
}

// Mitigations are an array individual mitigations
type Mitigations struct {
	Mitigation []Mitigation `xml:"mitigation"`
}

// Mitigation is an individual documentation of a compensating control or reason a policy-violating flaw will not be addressed.
type Mitigation struct {
	Action      string `xml:"action,attr"`
	Description string `xml:"description,attr"`
	User        string `xml:"user,attr"`
	Date        string `xml:"date,attr"`
}

// Annotations are an array of individual annotations (comments)
type Annotations struct {
	Annotation []Annotation `xml:"annotation"`
}

// Annotation is a comment on a flaw (separate from comments attached to mitigation actions)
type Annotation struct {
	Action      string `xml:"action,attr"`
	Description string `xml:"description,attr"`
	User        string `xml:"user,attr"`
	Date        string `xml:"date,attr"`
}

// CustomField is metadata for an application profile (extracted from detailed report API)
type CustomField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// ParseDetailedReport parses the detailedreport.do API and returns an DetailedReport struct, an array of Flaws, and an array of Custom Fields.
func ParseDetailedReport(credsFile, buildID string) (DetReport, []Flaw, []CustomField, error) {
	var flaws []Flaw
	var customFields []CustomField
	var detRep DetReport

	detailedReportAPI, err := detailedReport(credsFile, buildID)
	if err != nil {
		return detRep, nil, nil, err
	}

	/**
	Two decoders isn't ideal, but we are avoiding putting entire document into structs.
	We are decoding the entire document in decoder1 to get metadeta around the build.
	We decode flaws and custom fields in decoder 2.
	**/

	//Create the detailed report object
	detailedReportDecoder := xml.NewDecoder(bytes.NewReader(detailedReportAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := detailedReportDecoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "detailedreport" {
				detailedReportDecoder.DecodeElement(&detRep, &se)
			}
		}
	}

	//Create arrays of flaws and custom fields
	flawAndCustomDecoder := xml.NewDecoder(bytes.NewReader(detailedReportAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := flawAndCustomDecoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for errors, flaws, and custom field
			if se.Name.Local == "error" {
				return detRep, nil, nil, errors.New("api for GetDetailedReport returned with an error element")
			}
			if se.Name.Local == "flaw" {
				var flaw Flaw
				flawAndCustomDecoder.DecodeElement(&flaw, &se)
				flaw.CategoryName = categoryMap[flaw.CategoryID]
				flaw.PolicyName = detRep.PolicyName
				flaws = append(flaws, flaw)
			}
			if se.Name.Local == "customfield" {
				var cField CustomField
				flawAndCustomDecoder.DecodeElement(&cField, &se)
				customFields = append(customFields, cField)
			}
		}
	}
	return detRep, flaws, customFields, nil

}
