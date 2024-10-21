package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestPermissionsUnauth(t *testing.T) {
	initTest(t)

	scenarios := []tests.ApiScenario{
		{
			Name:            "unauth cannot list permissions",
			Method:          http.MethodGet,
			Url:             "/api/collections/permissions/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view permissions",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1AdminPermission.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create permissions",
			Method:          http.MethodPost,
			Url:             "/api/collections/permissions/records",
			Body:            strings.NewReader(`{"user":"user1","org":"org1","role":"role",}`),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update permissions",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
			Body:            strings.NewReader(`{"name":"user1"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete permissions",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestPermissionsOrgAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "org admins can list own org's users' permissions",
			Method:         http.MethodGet,
			Url:            "/api/collections/permissions/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"user":"%s"`, testData.org1Admin.Id),
				fmt.Sprintf(`"org":"%s"`, testData.org1.Id),
				`"group":""`,
				`"role":"admin"`,
				fmt.Sprintf(`"user":"%s"`, testData.org1group1Admin.Id),
				fmt.Sprintf(`"group":"%s"`, testData.org1group1.Id),
				`"org":""`,
				`"role":"admin"`,
				`"totalItems":2`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can view own org's user's permissions",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1AdminPermission.Id),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"user":"%s"`, testData.org1group1Admin.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can create org permissions for users in own org",
			Method:         http.MethodPost,
			Url:            "/api/collections/permissions/records",
			Body:           strings.NewReader(fmt.Sprintf(`{"user":"%s", "org":"%s", "role":"manager"}`, testData.org1group1Member.Id, testData.org1.Id)),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"user":"%s"`, testData.org1group1Member.Id),
				fmt.Sprintf(`"org":"%s"`, testData.org1.Id),
				`"group":""`,
				`"role":"manager"`,
			},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can create group permissions for users in own org",
			Method:         http.MethodPost,
			Url:            "/api/collections/permissions/records",
			Body:           strings.NewReader(fmt.Sprintf(`{"user":"%s", "group":"%s", "role":"manager"}`, testData.org1group1Member.Id, testData.org1group1.Id)),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"user":"%s"`, testData.org1group1Member.Id),
				fmt.Sprintf(`"group":"%s"`, testData.org1group1.Id),
				`"org":""`,
				`"role":"manager"`,
			},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot create permissions for users in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/permissions/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"user":"%s", "org":"%s", "role":"manager"}`, testData.org2group1Member.Id, testData.org1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot create permissions on other orgs for users in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/permissions/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"user":"%s", "org":"%s", "role":"manager"}`, testData.org1group1Member.Id, testData.org2.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update permission for user in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1AdminPermission.Id),
			Body:            strings.NewReader(`{"role": "manager"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"role":"manager"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot update permissions for users in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org2group1AdminPermission.Id),
			Body:            strings.NewReader(`{"role": "manager"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete permissions for users in own org",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1AdminPermission.Id),
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
			Name:            "org admin cannot delete permissions for users in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org2group1AdminPermission.Id),
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

//
// func TestPermissionsGroupAdmin(t *testing.T) {
// 	initTest(t)
// 	auth := authHeaders.Org1Group1Admin
//
// 	scenarios := []tests.ApiScenario{
// 		{
// 			Name:           "group admins can list permissions in own org",
// 			Method:         http.MethodGet,
// 			Url:            "/api/collections/permissions/records",
// 			RequestHeaders: auth,
// 			ExpectedStatus: 200,
// 			ExpectedContent: []string{
// 				`"username":"org1Admin"`,
// 				`"username":"org1Member"`,
// 				`"username":"org1group1Admin"`,
// 				`"username":"org1group1Member"`,
// 				`"username":"org1group1Manager"`,
// 				`"username":"org1group2Member"`,
// 				`"totalItems":6`,
// 			},
// 			NotExpectedContent: []string{
// 				`"username":"org2group1Member"`,
// 			},
// 			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
// 			TestAppFactory: generateTestApp,
// 		},
// 		{
// 			Name:            "group admins can view permissions in own org",
// 			Method:          http.MethodGet,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  200,
// 			ExpectedContent: []string{`"username":"org1group1Member"`},
// 			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admins can view permissions in other group in same org",
// 			Method:          http.MethodGet,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group2Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  200,
// 			ExpectedContent: []string{`"username":"org1group2Member"`},
// 			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admins cannot view permissions in other group in other org",
// 			Method:          http.MethodGet,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org2group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admin cannot create permissions in own org",
// 			Method:          http.MethodPost,
// 			Url:             "/api/collections/permissions/records",
// 			Body:            strings.NewReader(fmt.Sprintf(`{"username":"newUser", "password":"1234567890", "passwordConfirm":"1234567890", "orgs":["%s"], "role":"member"}`, testData.org1.Id)),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  400,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admin cannot create permissions in other org",
// 			Method:          http.MethodPost,
// 			Url:             "/api/collections/permissions/records",
// 			Body:            strings.NewReader(fmt.Sprintf(`{"username":"newUser", "password":"1234567890", "passwordConfirm":"1234567890", "orgs":["%s"], "role":"member"}`, testData.org2.Id)),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  400,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admin cannot update permissions in own org",
// 			Method:          http.MethodPatch,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
// 			Body:            strings.NewReader(`{"name": "newName"}`),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admin cannot update permissions in other org",
// 			Method:          http.MethodPatch,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org2group1Member.Id),
// 			Body:            strings.NewReader(`{"name": "newName"}`),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admin cannot delete permissions in own group",
// 			Method:          http.MethodDelete,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group admin cannot delete permissions in other org",
// 			Method:          http.MethodDelete,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org2group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 	}
//
// 	for _, scenario := range scenarios {
// 		scenario.Test(t)
// 	}
// }
//
// func TestPermissionsGroupMember(t *testing.T) {
// 	initTest(t)
// 	auth := authHeaders.Org1Group1Member
//
// 	scenarios := []tests.ApiScenario{
// 		{
// 			Name:           "group member can list permissions in own org",
// 			Method:         http.MethodGet,
// 			Url:            "/api/collections/permissions/records",
// 			RequestHeaders: auth,
// 			ExpectedStatus: 200,
// 			ExpectedContent: []string{
// 				`"username":"org1Admin"`,
// 				`"username":"org1Member"`,
// 				`"username":"org1group1Admin"`,
// 				`"username":"org1group1Member"`,
// 				`"username":"org1group1Manager"`,
// 				`"username":"org1group2Member"`,
// 				`"totalItems":6`,
// 			},
// 			NotExpectedContent: []string{
// 				`"username":"org2group1Member"`,
// 			},
// 			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
// 			TestAppFactory: generateTestApp,
// 		},
// 		{
// 			Name:            "group members cannot view permissions in own org",
// 			Method:          http.MethodGet,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group2Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group members cannot view permissions in other org",
// 			Method:          http.MethodGet,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org2group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group members can view own user",
// 			Method:          http.MethodGet,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  200,
// 			ExpectedContent: []string{`"username":"org1group1Member"`},
// 			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group members cannot create permissions",
// 			Method:          http.MethodPost,
// 			Url:             "/api/collections/permissions/records",
// 			Body:            strings.NewReader(`{"username":"user1","password":"pass123456","passwordConfirm":"pass123456",}`),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  400,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group members cannot update permissions in own group",
// 			Method:          http.MethodPatch,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group2Member.Id),
// 			Body:            strings.NewReader(`{"name":"user1"}`),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group members can update own user",
// 			Method:          http.MethodPatch,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
// 			Body:            strings.NewReader(`{"name":"user1"}`),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  200,
// 			ExpectedContent: []string{`"name":"user1"`},
// 			ExpectedEvents: map[string]int{
// 				OnModelAfterUpdate:          1,
// 				OnModelBeforeUpdate:         1,
// 				OnRecordAfterUpdateRequest:  1,
// 				OnRecordBeforeUpdateRequest: 1,
// 			},
// 			TestAppFactory: generateTestApp,
// 		},
// 		{
// 			Name:            "group members cannot delete a user",
// 			Method:          http.MethodDelete,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group2Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 		{
// 			Name:            "group members cannot delete own user",
// 			Method:          http.MethodDelete,
// 			Url:             fmt.Sprintf("/api/collections/permissions/records/%s", testData.org1group1Member.Id),
// 			RequestHeaders:  auth,
// 			ExpectedStatus:  404,
// 			ExpectedContent: []string{`"data":{}`},
// 			TestAppFactory:  generateTestApp,
// 		},
// 	}
//
// 	for _, scenario := range scenarios {
// 		scenario.Test(t)
// 	}
// }
