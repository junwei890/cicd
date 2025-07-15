package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey (t *testing.T) {
	header1 := http.Header{}
	header1.Set("Authorization", "ApiKey heylois")
	header2 := http.Header{}
	header2.Set("Authorization", "ApiKeybruh")
	header3 := http.Header{}
	header3.Set("Authorization", "Apikey fartsoundeffect")
	header4 := http.Header{}
	header4.Set("Content-Type", "application/json")

	testCases := []struct{
		name string
		header http.Header
		result string
		errorPresent bool
	}{
		{
			name: "Test pass",
			header: header1,
			result: "heylois",
			errorPresent: false,
		},
		{
			name: "Length of header less than 2",
			header: header2,
			result: "",
			errorPresent: true,
		},
		{
			name: "Wrong spelling in header",
			header: header3,
			result: "",
			errorPresent: true,
		},
		{
			name: "Authorization header not present",
			header: header4,
			result: "",
			errorPresent: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := GetAPIKey(testCase.header)
			if (err != nil) != testCase.errorPresent && result != testCase.result {
				t.Errorf("Test case: %s, failed. Actual results: %v, %s", testCase.name, err, result)
			}
		})
	}
}
