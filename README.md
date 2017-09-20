# Veracode API Package

## Description
Go package that provides easy access to the Veracode APIs. Each API has two files: one for making the http request and one for parsing the response.
For example, `detailedreport.go` calls the Veracode API and returns a `[byte]` and `detailedreportparser.go` parses the
XML response and returns usable objects such as flaws.

## Included APIs
1. getapplist.do
2. getbuildlist.do
3. getdetailedreport.do

## PACKAGE DOCUMENTATION

package vcodeapi
import "github.com/brian1917/vcodeapi"


### FUNCTIONS

func AppList(username, password string) ([]byte, error)
func BuildList(username, password, app_id string) ([]byte, error)
func GetAppList(username, password string) ([]App, error)
func GetBuildList(username, password, app_id string) ([]string, error)
func ParseDetailedReport(username, password, build_id string) ([]Flaw, []CustomField, error)
func SandboxList(username, password, appID string) ([]byte, error)

### TYPES

type Annotation struct {
    Action      string `xml:"action,attr"`
    Description string `xml:"description,attr"`
    User        string `xml:"user,attr"`
    Date        string `xml:"date,attr"`
}

type Annotations struct {
    Annotation Annotation `xml:"annotation"`
}

type App struct {
    AppID   string `xml:"app_id,attr"`
    AppName string `xml:"app_name,attr"`
}

type Build struct {
    BuildID string `xml:"build_id,attr"`
}

type CustomField struct {
    Name  string `xml:"name,attr"`
    Value string `xml:"value,attr"`
}

type Flaw struct {
    Issueid                   string      `xml:"issueid,attr"`
    CategoryName              string      `xml:"categoryname,attr"`
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

type Mitigation struct {
    Action      string `xml:"action,attr"`
    Description string `xml:"description,attr"`
    User        string `xml:"user,attr"`
    Date        string `xml:"date,attr"`
}

type Mitigations struct {
    Mitigation Mitigation `xml:"mitigation"`
}

