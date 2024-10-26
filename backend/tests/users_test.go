package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestUsersUnauth(t *testing.T) {
	initTest(t)

	scenarios := []tests.ApiScenario{
		{
			Name:            "unauth cannot list users",
			Method:          http.MethodGet,
			Url:             "/api/collections/users/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view users",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create users",
			Method:          http.MethodPost,
			Url:             "/api/collections/users/records",
			Body:            strings.NewReader(`{"username":"user1","password":"pass123456","passwordConfirm":"pass123456",}`),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update users",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"name":"user1"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete users",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestUsersOrgAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "org admins can list own org's users",
			Method:         http.MethodGet,
			Url:            "/api/collections/users/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"username":"org1Admin"`,
				`"username":"org1Member"`,
				`"username":"org1group1Admin"`,
				`"username":"org1group1Member"`,
				`"username":"org1group1Manager"`,
				`"username":"org1group2Member"`,
				`"totalItems":6`,
			},
			NotExpectedContent: []string{
				`"username":"org2group1Member"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can view own org's users",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"username":"org1group1Member"`,
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin can create users in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/users/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"username":"newUser", "password":"1234567890", "passwordConfirm":"1234567890", "orgs":["%s"], "role":"member"}`, testData.org1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"username":"newUser"`},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot create users in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/users/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"username":"newUser", "password":"1234567890", "passwordConfirm":"1234567890", "orgs":["%s"], "role":"member"}`, testData.org2.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update user in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"name": "newName"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"newName"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot update user in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org2group1Member.Id),
			Body:            strings.NewReader(`{"name": "newName"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete user in own org",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			RequestHeaders: auth,
			ExpectedStatus: 204,
			ExpectedEvents: map[string]int{
				OnModelAfterDelete:          1,
				OnModelBeforeDelete:         1,
				OnRecordAfterDeleteRequest:  1,
				OnRecordBeforeDeleteRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot delete users in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org2group1Member.Id),
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

func TestUsersGroupAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "group admins can list users in own org",
			Method:         http.MethodGet,
			Url:            "/api/collections/users/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"username":"org1Admin"`,
				`"username":"org1Member"`,
				`"username":"org1group1Admin"`,
				`"username":"org1group1Member"`,
				`"username":"org1group1Manager"`,
				`"username":"org1group2Member"`,
				`"totalItems":6`,
			},
			NotExpectedContent: []string{
				`"username":"org2group1Member"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admins can view users in own org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"username":"org1group1Member"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins can view users in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group2Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"username":"org1group2Member"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view users in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org2group1Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot create users in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/users/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"username":"newUser", "password":"1234567890", "passwordConfirm":"1234567890", "orgs":["%s"], "role":"member"}`, testData.org1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot create users in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/users/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"username":"newUser", "password":"1234567890", "passwordConfirm":"1234567890", "orgs":["%s"], "role":"member"}`, testData.org2.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot update users in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"name": "newName"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot update users in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org2group1Member.Id),
			Body:            strings.NewReader(`{"name": "newName"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete users in own group",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete users in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org2group1Member.Id),
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

func TestUsersGroupMember(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Member

	scenarios := []tests.ApiScenario{
		{
			Name:           "group member can list users in own org",
			Method:         http.MethodGet,
			Url:            "/api/collections/users/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"username":"org1Admin"`,
				`"username":"org1Member"`,
				`"username":"org1group1Admin"`,
				`"username":"org1group1Member"`,
				`"username":"org1group1Manager"`,
				`"username":"org1group2Member"`,
				`"totalItems":6`,
			},
			NotExpectedContent: []string{
				`"username":"org2group1Member"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members cannot view users in own org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group2Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view users in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org2group1Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members can view own user",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"username":"org1group1Member"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot create users",
			Method:          http.MethodPost,
			Url:             "/api/collections/users/records",
			Body:            strings.NewReader(`{"username":"user1","password":"pass123456","passwordConfirm":"pass123456",}`),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update users in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group2Member.Id),
			Body:            strings.NewReader(`{"name":"user1"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members can update own user",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"name":"user1"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"user1"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members can update own password",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"oldPassword":"1234567890","password":"abcdefghijk","passwordConfirm":"abcdefghijk"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"username":"org1group1Member"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members cannot update own groups",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"groups":["%s"]}`, testData.org2group1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update own orgs",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"orgs":["%s"]}`, testData.org2.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update own role",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"role":"super_admin"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot delete a user",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group2Member.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot delete own user",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/users/records/%s", testData.org1group1Member.Id),
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
