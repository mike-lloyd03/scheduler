package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestOrgsUnauth(t *testing.T) {
	initTest(t)

	scenarios := []tests.ApiScenario{
		{
			Name:            "unauth cannot list orgs",
			Method:          http.MethodGet,
			Url:             "/api/collections/orgs/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/orgs/records",
			Body:            strings.NewReader(`{"name": "newOrg"}`),
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestOrgsOrgAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:               "org admins can list own orgs",
			Method:             http.MethodGet,
			Url:                "/api/collections/orgs/records",
			RequestHeaders:     auth,
			ExpectedStatus:     200,
			ExpectedContent:    []string{`"name":"org1"`, `"totalItems":1`},
			NotExpectedContent: []string{`"name":"org2"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "org admin can view own org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot view other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org2.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/orgs/records",
			Body:            strings.NewReader(`{"name": "newOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedOrg"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot update other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org2.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot delete own org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot delete other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org2.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestOrgsGroupMember(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Member

	scenarios := []tests.ApiScenario{
		{
			Name:               "org member can list own orgs",
			Method:             http.MethodGet,
			Url:                "/api/collections/orgs/records",
			RequestHeaders:     auth,
			ExpectedStatus:     200,
			ExpectedContent:    []string{`"name":"org1"`, `"totalItems":1`},
			NotExpectedContent: []string{`"name":"org2"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "org member can view own org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member cannot view other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org2.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/orgs/records",
			Body:            strings.NewReader(`{"name": "newOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member cannot update an org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member cannot update other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org2.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org members cannot delete an org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", testData.org1.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
