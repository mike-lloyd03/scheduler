package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestRolesUnauth(t *testing.T) {
	initTest(t)

	scenarios := []tests.ApiScenario{

		{
			Name:            "unauth cannot list roles",
			Method:          http.MethodGet,
			Url:             "/api/collections/roles/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{`"items":[]`},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view roles",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create roles",
			Method:          http.MethodPost,
			Url:             "/api/collections/roles/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"role_template": "%s", "event":"%s"}`, testData.o1g1et1rt1.Id, testData.o1g1e1.Id)),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update roles",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1et1rt1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g2e1.Id)),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete roles",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestRolesOrgAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "org admins can list own org's roles",
			Method:         http.MethodGet,
			Url:            "/api/collections/roles/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
				fmt.Sprintf(`"event":"%s"`, testData.o1g2e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g2et1rt1.Id),
				`"totalItems":2`,
			},
			NotExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o2g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o2g1et1rt1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can view own org's roles",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can create role in group in own org",
			Method:         http.MethodPost,
			Url:            "/api/collections/roles/records",
			Body:           strings.NewReader(fmt.Sprintf(`{"role_template":"%s",  "event":"%s"}`, testData.o1g1et1rt1.Id, testData.o1g1e1.Id)),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
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
			Name:           "org admin can create role in other group in own org",
			Method:         http.MethodPost,
			Url:            "/api/collections/roles/records",
			Body:           strings.NewReader(fmt.Sprintf(`{"role_template":"%s",  "event":"%s"}`, testData.o1g2et1rt1.Id, testData.o1g2e1.Id)),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g2et1rt1.Id),
				fmt.Sprintf(`"event":"%s"`, testData.o1g2e1.Id),
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
			Name:            "org admin cannot create role in group in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/roles/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"role_template":"%s",  "event":"%s"}`, testData.o2g1et1rt1.Id, testData.o2g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update role in group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g2e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{fmt.Sprintf(`"event":"%s"`, testData.o1g2e1.Id)},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin can update role in other group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g2e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id)},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot update role in group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o2g1e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete role in group in own org",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
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
			Name:            "org admin cannot delete role in groups in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o2g1e1.Id),
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

func TestRolesGroupAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "group admins can list roles in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/roles/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
				`"totalItems":1`,
			},
			NotExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g2e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g2et1rt1.Id),
				fmt.Sprintf(`"event":"%s"`, testData.o2g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o2g1et1rt1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "group admins can view roles in own group",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admins cannot view roles in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g2e1r1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view roles in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o2g1e1r1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "group admin can create role in group in own org",
			Method:         http.MethodPost,
			Url:            "/api/collections/roles/records",
			Body:           strings.NewReader(fmt.Sprintf(`{"role_template":"%s",  "event":"%s"}`, testData.o1g1et1rt1.Id, testData.o1g1e1.Id)),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
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
			Name:            "group admin cannot create role in other group",
			Method:          http.MethodPost,
			Url:             "/api/collections/roles/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"role_template":"%s",  "event":"%s"}`, testData.o1g2et1rt1.Id, testData.o2g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot create role in group in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/roles/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"role_template":"%s",  "event":"%s"}`, testData.o2g1et1rt1.Id, testData.o2g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin can update roles in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g2e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{fmt.Sprintf(`"event":"%s"`, testData.o1g2e1.Id)},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin cannot update roles in other group in same org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g2e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot update roles in other group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o2g1e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "group admin can delete roles in own group",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
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
			Name:            "group admin cannot delete roles in other group in same org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g2e1r1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete roles in other group in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o2g1e1r1.Id),
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

func TestRolesGroupMember(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Member

	scenarios := []tests.ApiScenario{
		{
			Name:           "group member can list roles in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/roles/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
				`"totalItems":1`,
			},
			NotExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g2e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g2et1rt1.Id),
				fmt.Sprintf(`"event":"%s"`, testData.o2g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o2g1et1rt1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "group members can view roles in own group",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"event":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"role_template":"%s"`, testData.o1g1et1rt1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members cannot view roles in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g2e1r1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view roles in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o2g1e1r1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot create roles",
			Method:          http.MethodPost,
			Url:             "/api/collections/roles/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"role_template": "%s", "event":"%s"}`, testData.o1g1et1rt1.Id, testData.o1g1e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update roles in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1e1r1.Id),
			Body:            strings.NewReader(fmt.Sprintf(`{"event": "%s"}`, testData.o1g2e1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot delete a role",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/roles/records/%s", testData.o1g1et1rt1.Id),
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
