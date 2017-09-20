# Veracode API Package

## Package Documentation
See here: https://godoc.org/github.com/brian1917/vcodeapi

## Description
Go package that provides easy access to the Veracode APIs. Each API has two files: one for making the http request and one for parsing the response.
For example, `detailedreport.go` calls the Veracode API and returns a `[byte]` and `detailedreportparser.go` parses the
XML response and returns usable objects such as flaws.

## Included APIs
1. Get App List (`/api/5.0/getapplist.do`)
2. Get Build List (`/api/5.0/getbuildlist.do`)
3. Get Sandbox List (`/api/5.0/getsandboxlist.do`)
4. Get Detailed Report (`/api/5.0/detailedreport.do`)
