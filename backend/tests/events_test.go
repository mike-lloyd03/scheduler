package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

// Org Admin can list, view, create, update, delete events in all groups own org
// Group admin can list, view, create, update, delete events in own group
// Group manager can list, view, create, update, delete events in own group
// Group member can list, view own events
func TestEvents(t *testing.T) {
	app := generateTestApp(t)
	authHeaders.Init(app)
	testData.Init(app)

	scenarios := []tests.ApiScenario{
		// -----------
		// unauth
		// -----------
		{
			Name:            "unauth cannot list events",
			Method:          http.MethodGet,
			Url:             "/api/collections/events/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot view events",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create events",
			Method:          http.MethodPost,
			Url:             "/api/collections/events/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"event_template": "%s", "datetime": "123456",}`, testData.o1g1et1.Id)),
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot update events",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			Body:            strings.NewReader(`{"name": "updatedGroup"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot delete events",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		// -----------
		// org admin
		// -----------
		{
			Name:           "org admins can list own org's events",
			Method:         http.MethodGet,
			Url:            "/api/collections/events/records",
			RequestHeaders: authHeaders.Org1Admin,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g1e1.Id),
				fmt.Sprintf(`"id":"%s"`, testData.o1g2e1.Id),
				`"totalItems":4`,
			},
			NotExpectedContent: []string{
				fmt.Sprintf(`"event_template":"%s"`, testData.o2g1et1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "org admin can view own org's events",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			RequestHeaders: authHeaders.Org1Admin,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g1e1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin can create events in group in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/events/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"event_template": "%s", "datetime": "2024-10-01 12:00:00.000Z"}`, testData.o1g1et1.Id)),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"datetime":"2024-10-01 12:00:00.000Z"`},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot create event in group in other org",
			Method:          http.MethodPost,
			Url:             "/api/collections/events/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"event_template": "%s", "datetime": "2024-10-01 12:00:00.000Z"}`, testData.o2g1et1.Id)),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update event in group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"datetime":"2024-10-03 12:00:00.000Z"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin can update event in other group in own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g2e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"datetime":"2024-10-03 12:00:00.000Z"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "org admin cannot update event in group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o2g1e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "org admin can delete event in group in own org",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			RequestHeaders: authHeaders.Org1Admin,
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
			Name:            "org admin cannot delete events in groups in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o2g1e1.Id),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},

		// -----------
		// group admin
		// -----------
		{
			Name:           "group admins can list events in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/events/records",
			RequestHeaders: authHeaders.Org1Group1Admin,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g1e1.Id),
				`"totalItems":2`,
			},
			NotExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g2e1.Id),
				fmt.Sprintf(`"id":"%s"`, testData.o2g1e1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "group admins can view events in own group",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			RequestHeaders: authHeaders.Org1Group1Admin,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g1e1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admins cannot view events in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g2e1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admins cannot view events in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o2g1e1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin can create events in group in own org",
			Method:          http.MethodPost,
			Url:             "/api/collections/events/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"event_template": "%s", "datetime": "2024-10-01 12:00:00.000Z"}`, testData.o1g1et1.Id)),
			RequestHeaders:  authHeaders.Org1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"datetime":"2024-10-01 12:00:00.000Z"`},
			ExpectedEvents: map[string]int{
				OnModelAfterCreate:          1,
				OnModelBeforeCreate:         1,
				OnRecordAfterCreateRequest:  1,
				OnRecordBeforeCreateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin can update events in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"datetime":"2024-10-03 12:00:00.000Z"`},
			ExpectedEvents: map[string]int{
				OnModelAfterUpdate:          1,
				OnModelBeforeUpdate:         1,
				OnRecordAfterUpdateRequest:  1,
				OnRecordBeforeUpdateRequest: 1,
			},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group admin cannot update events in other group in same org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g2e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot update events in other group in other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o2g1e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:           "group admin can delete events in own group",
			Method:         http.MethodDelete,
			Url:            fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			RequestHeaders: authHeaders.Org1Group1Admin,
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
			Name:            "group admin cannot delete events in other group in same org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g2e1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group admin cannot delete events in other group in other org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o2g1e1.Id),
			RequestHeaders:  authHeaders.Org1Group1Admin,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},

		// -----------
		// group member
		// -----------
		{
			Name:           "group member can list events in own group",
			Method:         http.MethodGet,
			Url:            "/api/collections/events/records",
			RequestHeaders: authHeaders.Org1Group1Member,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g1e1.Id),
				`"totalItems":2`,
			},
			NotExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g2e1.Id),
				fmt.Sprintf(`"id":"%s"`, testData.o2g1e1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordsListRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:           "group members can view events in own group",
			Method:         http.MethodGet,
			Url:            fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			RequestHeaders: authHeaders.Org1Group1Member,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				fmt.Sprintf(`"id":"%s"`, testData.o1g1e1.Id),
			},
			ExpectedEvents: map[string]int{OnRecordViewRequest: 1},
			TestAppFactory: generateTestApp,
		},
		{
			Name:            "group members cannot view events in other group in same org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g2e1.Id),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot view events in other group in other org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o2g1e1.Id),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot create events",
			Method:          http.MethodPost,
			Url:             "/api/collections/events/records",
			Body:            strings.NewReader(fmt.Sprintf(`{"event_template": "%s", "datetime": "2024-10-01 12:00:00.000Z"}`, testData.o1g1et1.Id)),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  400,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot update events in own group",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
			Body:            strings.NewReader(`{"datetime": "2024-10-03 12:00:00.000Z"}`),
			RequestHeaders:  authHeaders.Org1Group1Member,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "group members cannot delete an event",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/events/records/%s", testData.o1g1e1.Id),
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
