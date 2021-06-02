package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var err error
var companies []string = []string{"Google", "Oracle", "Deloitte", "Microsoft", "Amazon"}
var responseHeaders = map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Credentials": "true"}

func addCompanies(addedCompanies []interface{}) {

	for _, c := range addedCompanies {
		s := strings.Title(strings.ToLower(c.(string)))
		companies = append(companies, s)
	}
	seen := make(map[string]bool)
	var newCompanies []string
	for _, c := range companies {
		if _, value := seen[c]; !value {
			seen[c] = true
			newCompanies = append(newCompanies, c)
		}

	}
	companies = newCompanies
}

func regexBuilder() string {
	var regexStr []string
	for _, c := range companies {
		regexStr = append(regexStr, `\b`+c+`\b`)
	}
	return "(" + strings.Join(regexStr, "|") + ")"

}

func computeString(str string) string {
	regexStr := regexBuilder()
	r := regexp.MustCompile(regexStr)
	words := r.FindAllString(str, -1)
	fmt.Println(words)
	s := r.ReplaceAllString(str, `$1©`)
	return s
}

func returnErrorResponse(err error) events.APIGatewayProxyResponse {
	response := map[string]interface{}{"data": nil, "error": err.Error()}
	responseBody, err := json.Marshal(response)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: string(responseBody), Headers: responseHeaders}
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	inputReq := make(map[string]interface{})
	err = json.Unmarshal([]byte(request.Body), &inputReq)
	if err != nil {
		return returnErrorResponse(err), nil
	}
	text := inputReq["text"].(string)
	addedCompanies := inputReq["add_organisation"]
	if addedCompanies != nil {
		addCompanies(addedCompanies.([]interface{}))
	}

	outputText := computeString(text)
	response := map[string]interface{}{"data": outputText, "error": nil}
	responseBody, err := json.Marshal(response)
	if err != nil {
		return returnErrorResponse(err), nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(responseBody), Headers: responseHeaders}, nil
}
func main() {
	lambda.Start(Handler)
}
