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

package vcodeapi </br>
import "github.com/brian1917/vcodeapi"


### FUNCTIONS

func AppList(username, password string) ([]byte, error)</br>
func BuildList(username, password, app_id string) ([]byte, error)</br>
func GetAppList(username, password string) ([]App, error)</br>
func GetBuildList(username, password, app_id string) ([]string, error)</br>
func ParseDetailedReport(username, password, build_id string) ([]Flaw, []CustomField, error)</br>
func SandboxList(username, password, appID string) ([]byte, error)</br>

### TYPES

type Annotation struct {</br>
    Action      string `xml:"action,attr"`</br>
    Description string `xml:"description,attr"`</br>
    User        string `xml:"user,attr"`</br>
    Date        string `xml:"date,attr"`</br>
}

type Annotations struct {</br>
    Annotation Annotation `xml:"annotation"`</br>
}

type App struct {</br>
    AppID   string `xml:"app_id,attr"`</br>
    AppName string `xml:"app_name,attr"`</br>
}

type Build struct {</br>
    BuildID string `xml:"build_id,attr"`</br>
}

type CustomField struct {</br>
    Name  string `xml:"name,attr"`</br>
    Value string `xml:"value,attr"`</br>
}

type Flaw struct {</br>
    Issueid                   string      `xml:"issueid,attr"`</br>
    CategoryName              string      `xml:"categoryname,attr"`</br>
    Cweid                     string      `xml:"cweid,attr"`</br>
    Remediation_status        string      `xml:"remediation_status,attr"`</br>
    Mitigation_status         string      `xml:"mitigation_status,attr"`</br>
    Affects_policy_compliance string      `xml:"affects_policy_compliance,attr"`</br>
    Date_first_occurrence     string      `xml:"date_first_occurrence,attr"`</br>
    Severity                  string      `xml:"severity,attr"`</br>
    ExploitLevel              string      `xml:"exploitLevel,attr"`</br>
    Module                    string      `xml:"module,attr"`</br>
    Sourcefile                string      `xml:"sourcefile,attr"`</br>
    Line                      string      `xml:"line,attr"`</br>
    Description               string      `xml:"description,attr"`</br>
    Mitigations               Mitigations `xml:"mitigations"`</br>
    Annotations               Annotations `xml:"annotations"`</br>
}

type Mitigation struct {</br>
    Action      string `xml:"action,attr"`</br>
    Description string `xml:"description,attr"`</br>
    User        string `xml:"user,attr"`</br>
    Date        string `xml:"date,attr"`</br>
}

type Mitigations struct {</br>
    Mitigation Mitigation `xml:"mitigation"`</br>
}

