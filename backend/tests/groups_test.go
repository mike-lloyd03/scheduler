package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestGroupsUnauth(t *testing.T) {
	initTest(t)

	scenarios := []tests.ApiScenario{
		{
			Name:            "unauth cannot list groups",
			Method:          http.MethodGet,
			Url:             "/api/collections/groups/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create group",
			Method:          http.MethodPost,
			Url:             "/api/collections/groups/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newOrg", "org":"%s"}`, testData.org1.Id)),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete group",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGroupsOrgAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "org admins can list own org's groups",
			Method:         http.MethodGet,
			Url:            "/api/collections/groups/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"org1group1"`,
				`"name":"org1group2"`,
				`"totalItems":2`},
			NotExpectedContent: []string{`"name":"org2group1"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "org admin can view own org's groups",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1group1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can create group in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/groups/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newOrg", "org":"%s"}`, testData.org1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"newOrg"`},
			ExpectedEvents:  map[string]int{OnModelAfterCreate: 1, OnModelBeforeCreate: 1, OnRecordAfterCreateRequest: 1, OnRecordBeforeCreateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot create group in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/groups/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newOrg", "org":"%s"}`, testData.org2.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update own org's groups",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedGroup"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot update other org's groups",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org2group1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete own org's groups",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			RequestHeaders: auth,
			ExpectedStatus: 204,
			ExpectedEvents: map[string]int{
				OnModelAfterDelete:          testData.onDeleteGroup,
				OnModelBeforeDelete:         testData.onDeleteGroup,
				OnModelAfterUpdate:          4,
				OnModelBeforeUpdate:         4,
				OnRecordAfterDeleteRequest:  1,
				OnRecordBeforeDeleteRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot delete other org's groups",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org2group1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGroupsGroupAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:               "group admins can list own groups",
			Method:             http.MethodGet,
			Url:                "/api/collections/groups/records",
			RequestHeaders:     auth,
			ExpectedStatus:     200,
			ExpectedContent:    []string{`"name":"org1group1"`, `"totalItems":1`},
			NotExpectedContent: []string{`"name":"org2group1"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "group admins can view own group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1group1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot create group",
			Method:          http.MethodPost,
			Url:             "/api/collections/groups/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newOrg", "org":"%s"}`, testData.org2.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin can update own groups",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedGroup"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete own groups",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGroupsOrgMember(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Member

	scenarios := []tests.ApiScenario{
		{
			Name:            "org member cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/groups/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newOrg", "org":"%s"}`, testData.org1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org members cannot delete an org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGroupsGroupMember(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Member

	scenarios := []tests.ApiScenario{
		{
			Name:               "group members can list own groups",
			Method:             http.MethodGet,
			Url:                "/api/collections/groups/records",
			RequestHeaders:     auth,
			ExpectedStatus:     200,
			ExpectedContent:    []string{`"name":"org1group1"`, `"totalItems":1`},
			NotExpectedContent: []string{`"name":"org2group1"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "group member can view own group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1group1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot create group",
			Method:          http.MethodPost,
			Url:             "/api/collections/groups/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newOrg", "org":"%s"}`, testData.org1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org1group1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group member cannot update other group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/groups/records/%s", testData.org2group1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
