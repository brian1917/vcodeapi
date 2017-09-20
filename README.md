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
