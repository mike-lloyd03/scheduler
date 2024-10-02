package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

// Org Admin can list, view, create, update, delete event_templates in all groups own org
// Group admin can list, view, create, update, delete event_templates in own group
// Group manager can list, view, create, update, delete event_templates in own group
// Group member can list, view own event_templates
func TestEventTemplates(t *testing.T) {
	app := generateTestApp(t)
	authHeaders.Init(app)

	org1group1, err := app.Dao().FindFirstRecordByFilter("groups", "name='org1group1'")
	if err != nil {
		t.Fatal(err)
	}

	org2group1, err := app.Dao().FindFirstRecordByFilter("groups", "name='org2group1'")
	if err != nil {
		t.Fatal(err)
	}

	o1g1et1, err := app.Dao().FindFirstRecordByFilter("event_templates", "name='o1g1et1'")
	if err != nil {
		t.Fatal(err)
	}

	o1g2et1, err := app.Dao().FindFirstRecordByFilter("event_templates", "name='o1g2et1'")
	if err != nil {
		t.Fatal(err)
	}

	o2g1et1, err := app.Dao().FindFirstRecordByFilter("event_templates", "name='o2g1et1'")
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		// -----------
		// unauth
		// -----------
		{
			Name:            "unauth cannot list event templates",
			Method:          http.MethodGet,
			Url:             "/api/collections/event_templates/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view event templates",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create event templates",
			Method:          http.MethodPost,
			Url:             "/api/collections/event_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newET", "recurrence": "daily", "group":"%s"}`, org1group1.Id)),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update event templates",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete event templates",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", org1group1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		// -----------
		// org admin
		// -----------
		{
			Name:           "org admins can list own org's event_templates",
			Method:         http.MethodGet,
			Url:            "/api/collections/event_templates/records",
			RequestHeaders: authHeaders.Org1Admin,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"o1g1et1"`,
				`"name":"o1g1et2"`,
				`"name":"o1g2et1"`,
				`"name":"o1g2et2"`,
				`"totalItems":4`,
			},
			NotExpectedContent: []string{`"name":"o2g1et1"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:            "org admin can view own org's event_templates",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"o1g1et1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can create event_template in group in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/event_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newET", "recurrence": "daily", "group":"%s"}`, org1group1.Id)),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"newET"`},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot create event_template in group in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/event_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newET", "recurrence": "daily", "group":"%s"}`, org2group1.Id)),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update event_template in group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedET"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update event_template in other group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g2et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedET"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot update event_template in group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o2g1et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete event_template in group in own org",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			RequestHeaders: authHeaders.Org1Admin,
			ExpectedStatus: 204,
			ExpectedEvents: map[string]int{
				OnModelAfterDelete:          2,
				OnModelBeforeDelete:         2,
				OnRecordAfterDeleteRequest:  1,
				OnRecordBeforeDeleteRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot delete event_templates in groups in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o2g1et1.Id),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		// -----------
		// group admin
		// -----------
		{
			Name:           "group admins can list event_templates in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/event_templates/records",
			RequestHeaders: authHeaders.Org1Group1Admin,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"o1g1et1"`,
				`"name":"o1g1et2"`,
				`"totalItems":2`,
			},
			NotExpectedContent: []string{
				`"name":"o2g1et1"`,
				`"name":"o1g2et1"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admins can view event_templates in own group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"o1g1et1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view event_templates in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g2et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view event_templates in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o2g1et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin can update event_templates in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedET"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin cannot update event_templates in other group in same org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g2et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot update event_templates in other group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o2g1et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "group admin can delete event_templates in own group",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			RequestHeaders: authHeaders.Org1Group1Admin,
			ExpectedStatus: 204,
			ExpectedEvents: map[string]int{
				OnModelAfterDelete:          2,
				OnModelBeforeDelete:         2,
				OnRecordAfterDeleteRequest:  1,
				OnRecordBeforeDeleteRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin cannot delete event_templates in other group in same org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g2et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete event_templates in other group in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o2g1et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		// -----------
		// group member
		// -----------
		{
			Name:           "group member can list event_templates in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/event_templates/records",
			RequestHeaders: authHeaders.Org1Group1Member,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"name":"o1g1et1"`,
				`"name":"o1g1et2"`,
				`"totalItems":2`,
			},
			NotExpectedContent: []string{
				`"name":"o2g1et1"`,
				`"name":"o1g2et1"`,
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members can view event_templates in own group",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g1et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"o1g1et1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view event_templates in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g2et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view event_templates in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o2g1et1.Id),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot create event_templates",
			Method:          http.MethodPost,
			Url:             "/api/collections/event_templates/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"name": "newET", "recurrence": "daily", "group":"%s"}`, org1group1.Id)),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update event_templates in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", o1g2et1.Id),
			Body:            strings.NewReader(`{"name": "updatedET"}`),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot delete an event_template",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/event_templates/records/%s", org1group1.Id),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
