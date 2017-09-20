# Veracode API Package

## Description
Easy access to the Veracode APIs. Each API has two files: one for making the http request and one for parsing the response.
For example, `detailedreport.go` calls the Veracode API and returns a `[byte]`. `detailedreportparser.go` parses the
XML response and returns relevant information on flaws, the app, etc.

## Included APIs
1. func AppList(username, password)
2. func BuildList(username, password, app\_id string)
3. func DetailedReport (username, password, build\_id)

All APIs return type []byte
