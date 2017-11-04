# Veracode API Package
[![GoDoc](https://godoc.org/github.com/brian1917/vcodeapi?status.svg)](https://godoc.org/github.com/brian1917/vcodeapi)

## Package Documentation
See here: https://godoc.org/github.com/brian1917/vcodeapi

## Description
Go package that provides easy access to the Veracode APIs. Each API typically has two files: one for making the http request and one for parsing the response.
For example, `detailedreport.go` calls the Veracode API and returns a `[byte]` and `detailedreportparser.go` parses the
XML response and returns usable objects such as flaws.

## Credentials File
Must be structured like the following:
```
[DEFAULT]
veracode_api_key_id = ID HERE
veracode_api_key_secret = SECRET HERE
```

## Included APIs
1. Get App List (`/api/5.0/getapplist.do`)
2. Get Build List (`/api/5.0/getbuildlist.do`)
3. Get Sandbox List (`/api/5.0/getsandboxlist.do`)
4. Get Detailed Report (`/api/5.0/detailedreport.do`)
5. Get Team Info (`api/3.0/getteaminfo.do`)
6. Updated Mitigation Info (`api/updatemitigationinfo.do`)
