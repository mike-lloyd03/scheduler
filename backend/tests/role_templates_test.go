package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestRoleTemplatesUnauth(t *testing.T) {
	initTest(t)

	scenarios := []tests.ApiScenario{
		{
			Name:            "unauth cannot list role templates",
			Method:          http.MethodGet,
			Url:             "/api/collections/role_templates/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view role templates",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create role templates",
			Method:          http.MethodPost,
			Url:             "/api/collections/role_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newRole", "event_template":"%s"}`, testData.o1g1et1.Id)),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update role templates",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete role templates",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestRoleTemplatesOrgAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "org admins can list own org's role_templates",
			Method:         http.MethodGet,
			Url:            "/api/collections/role_templates/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"o1g1et1rt1"`,
				`"name":"o1g2et1rt1"`,
				`"totalItems":2`,
			},
			NotExpectedContent: []string{`"name":"o2g1et1rt1"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "org admin can view own org's role_templates",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"o1g1et1rt1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can create role_template in group in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/role_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newRole",  "event_template":"%s"}`, testData.o1g1et1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"newRole"`},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin can create role_template in other group in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/role_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newRole",  "event_template":"%s"}`, testData.o1g2et1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"newRole"`},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot create role_template in group in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/role_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newRole",  "event_template":"%s"}`, testData.o2g1et1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update role_template in group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedRole"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update role_template in other group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g2et1rt1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedRole"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot update role_template in group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o2g1et1rt1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete role_template in group in own org",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			RequestHeaders: auth,
			ExpectedStatus: 204,
			ExpectedEvents: map[string]int{
				OnModelAfterDelete:          testData.onDeleteRoleTemplate,
				OnModelBeforeDelete:         testData.onDeleteRoleTemplate,
				OnRecordAfterDeleteRequest:  1,
				OnRecordBeforeDeleteRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot delete role_templates in groups in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o2g1et1.Id),
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

func TestRoleTemplatesGroupAdmin(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Admin

	scenarios := []tests.ApiScenario{
		{
			Name:           "group admins can list role_templates in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/role_templates/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"o1g1et1rt1"`,
				`"totalItems":1`,
			},
			NotExpectedContent: []string{
				`"name":"o2g1et1rt1"`,
				`"name":"o1g2et1rt1"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admins can view role_templates in own group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"o1g1et1rt1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view role_templates in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g2et1rt1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view role_templates in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o2g1et1rt1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin can update role_templates in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedRole"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin cannot update role_templates in other group in same org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g2et1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot update role_templates in other group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o2g1et1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "group admin can delete role_templates in own group",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			RequestHeaders: auth,
			ExpectedStatus: 204,
			ExpectedEvents: map[string]int{
				OnModelAfterDelete:          testData.onDeleteRoleTemplate,
				OnModelBeforeDelete:         testData.onDeleteRoleTemplate,
				OnRecordAfterDeleteRequest:  1,
				OnRecordBeforeDeleteRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin cannot delete role_templates in other group in same org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g2et1rt1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete role_templates in other group in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o2g1et1rt1.Id),
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

func TestRoleTemplatesGroupMember(t *testing.T) {
	initTest(t)
	auth := authHeaders.Org1Group1Member

	scenarios := []tests.ApiScenario{
		{
			Name:           "group member can list role_templates in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/role_templates/records",
			RequestHeaders: auth,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"o1g1et1rt1"`,
				`"totalItems":1`,
			},
			NotExpectedContent: []string{
				`"name":"o2g1et1rt1"`,
				`"name":"o1g2et1rt1"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members can view role_templates in own group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"o1g1et1rt1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view role_templates in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g2et1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view role_templates in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o2g1et1.Id),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot create role_templates",
			Method:          http.MethodPost,
			Url:             "/api/collections/role_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newRole", "event_template":"%s"}`, testData.o1g1et1.Id)),
			RequestHeaders:  auth,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update role_templates in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g2et1.Id),
			Body:            strings.NewReader(`{"name": "updatedRole"}`),
			RequestHeaders:  auth,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot delete a role_template",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/role_templates/records/%s", testData.o1g1et1rt1.Id),
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
