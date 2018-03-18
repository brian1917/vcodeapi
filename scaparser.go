package vcodeapi

import (
	"bytes"
	"encoding/xml"
)

// SoftwareCompositionAnalysis represents the SCA section of the detailed report
type SoftwareCompositionAnalysis struct {
	ComponentsViolatedPolicy string      `xml:"components_violated_policy,attr"`
	ThirdPartyComponents     string      `xml:"third_party_components,attr"`
	ViolatePolicy            string      `xml:"violate_policy,attr"`
	VulnerableComponents     []Component `xml:"vulnerable_components"`
}

// Component is a third-party library identifed by SCA
type Component struct {
	AddedDate                        string              `xml:"added_date,attr"`
	ComponentAffectsPolicyCompliance string              `xml:"component_affects_policy_compliance,attr"`
	Description                      string              `xml:"description,attr"`
	FileName                         string              `xml:"file_name,attr"`
	Library                          string              `xml:"library,attr"`
	MaxCvssScore                     string              `xml:"max_cvss_score,attr"`
	New                              string              `xml:"new,attr"`
	Sha1                             string              `xml:"sha1,attr"`
	Vendor                           string              `xml:"vendor,attr"`
	Version                          string              `xml:"version,attr"`
	FilePaths                        FilePaths           `xml:"file_paths"`
	Licenses                         Licenses            `xml:"licenses"`
	ViolatedPolicyRules              ViolatedPolicyRules `xml:"violated_policy_rules"`
	Vulnerabilities                  Vulnerabilities     `xml:"vulnerabilities"`
}

// FilePaths is an array of filepaths
type FilePaths struct {
	FilePath []FilePath `xml:"file_path"`
}

// FilePath is the filepath of the third-party component
type FilePath struct {
	Value string `xml:"value,attr"`
}

//Licenses is an array of licenses
type Licenses struct {
	License []License `xml:"license"`
}

// License is the license associated with a third-party component identified by SCA
type License struct {
	LicenseURL string `xml:"license_url,attr"`
	Name       string `xml:"name,attr"`
	RiskRating string `xml:"risk_rating,attr"`
	SpdxID     string `xml:"spdx_id,attr"`
}

// Vulnerabilities is an array of vulnerabilities
type Vulnerabilities struct {
	Vulnerability []Vulnerability `xml:"vulnerability"`
}

// Vulnerability is a CVE associated with a third-party component identified by SCA
type Vulnerability struct {
	CveID                                string `xml:"cve_id,attr"`
	CveSummary                           string `xml:"cve_summary,attr"`
	CvssScore                            string `xml:"cvss_score,attr"`
	CweID                                string `xml:"cwe_id,attr"`
	Mitigation                           string `xml:"mitigation,attr"`
	Severity                             string `xml:"severity,attr"`
	SeverityDesc                         string `xml:"severity_desc,attr"`
	VulnerabilityAffectsPolicyCompliance string `xml:"vulnerability_affects_policy_compliance,attr"`
}

// ViolatedPolicyRules is an array of rules violating by the third-party component identifed by SCA
type ViolatedPolicyRules struct {
	PolicyRule []PolicyRule `xml:"Policy_rule"`
}

//PolicyRule is a rule violated by a third-party component identifed by SCA
type PolicyRule struct {
	Desc  string `xml:" desc,attr"`
	Type  string `xml:" type,attr"`
	Value string `xml:" value,attr" `
}

// ParseSCAReport parses the detailedreport.do API and returns a SoftwareCompositionAnalysis struct
func ParseSCAReport(credsFile, buildID string) (SoftwareCompositionAnalysis, error) {

	var SCA SoftwareCompositionAnalysis

	detailedReportAPI, err := detailedReport(credsFile, buildID)
	if err != nil {
		return SCA, err
	}

	//Create the SCA object
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
			if se.Name.Local == "software_composition_analysis" {
				decoder.DecodeElement(&SCA, &se)
			}
		}
	}

	return SCA, nil

}
