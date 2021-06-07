package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var err error
var defaultOrgs []string = []string{"Google", "Oracle", "Deloitte", "Microsoft", "Amazon"}
var responseHeaders = map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Credentials": "true"}

//addCompanies : add additional companies to the default list of companies.
func addCompanies(addedCompanies []interface{}) []string {
	var newOrgs []string = defaultOrgs

	//convert to title case
	for _, c := range addedCompanies {
		s := strings.Title(strings.ToLower(c.(string)))
		newOrgs = append(newOrgs, s)
	}
	//remove duplicate values from list
	seen := make(map[string]bool)
	for _, c := range newOrgs {
		//if value is not present, add
		if _, value := seen[c]; !value {
			seen[c] = true
			newOrgs = append(newOrgs, c)
		}

	}
	// use the new list
	return newOrgs
}

//build regex to find companies in the input string
func regexBuilder(orgs []string) string {
	var regexStr []string
	for _, c := range orgs {
		regexStr = append(regexStr, `\b`+c+`\b`)
	}
	// eg (\bGoogle\b|\bAmazon\b|\bOracle\b)
	return "(" + strings.Join(regexStr, "|") + ")"

}

//find and replace all company names
func computeString(str string, regexStr string) string {
	r := regexp.MustCompile(regexStr)
	s := r.ReplaceAllString(str, `$1©`)
	return s
}

// utility function to return error response
func returnErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	response := map[string]interface{}{"data": nil, "error": err.Error()}
	responseBody, err := json.Marshal(response)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: string(responseBody), Headers: responseHeaders}, err
}

//Handler : main entrypoint handler function for Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	inputReq := make(map[string]interface{})
	//get request body
	err = json.Unmarshal([]byte(request.Body), &inputReq)
	if err != nil {
		return returnErrorResponse(err)
	}

	// get the input string
	text, ok := inputReq["text"].(string)
	if !ok {
		err := errors.New(fmt.Sprintf("got data of type %T but wanted string", inputReq["text"]))
		return returnErrorResponse(err)
	}

	var orgs []string = defaultOrgs
	// get additional companies if provided
	addedCompanies := inputReq["add_organisation"]
	if addedCompanies != nil {
		if _, ok := addedCompanies.([]interface{}); !ok {
			err := errors.New(fmt.Sprintf("got data of type %T but wanted []interface{}", inputReq["add_organisation"]))
			return returnErrorResponse(err)
		}
		orgs = addCompanies(addedCompanies.([]interface{}))
	}

	//find and replace
	regexStr := regexBuilder(orgs)
	outputText := computeString(text, regexStr)

	//send response back
	response := map[string]interface{}{"data": outputText, "error": nil}
	responseBody, err := json.Marshal(response)
	if err != nil {
		return returnErrorResponse(err)
	}
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody), Headers: responseHeaders}, nil
}

func main() {
	lambda.Start(Handler)
}
