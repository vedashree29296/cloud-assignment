package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request     events.APIGatewayProxyRequest
		expect      string
		statusCode  int
		errorString string
	}{
		{
			// Success Scenario
			request:    events.APIGatewayProxyRequest{Body: `{"text":"Google has announced new features"}`},
			expect:     `{"data":"Google© has announced new features","error":null}`,
			statusCode: 200,
		},
		{
			// Test
			// when a valid name is provided in the HTTP body
			request:    events.APIGatewayProxyRequest{Body: `{"text":123}`},
			expect:     `{"data":null,"error":"got data of type float64 but wanted string"}`,
			statusCode: 500,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		if err != nil {
			assert.Equal(t, test.errorString, err.Error())
		}
		assert.Equal(t, test.statusCode, response.StatusCode)
		assert.Equal(t, test.expect, response.Body)
	}

}
