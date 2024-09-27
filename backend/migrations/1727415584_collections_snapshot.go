package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2024-09-02 04:37:22.680Z",
				"updated": "2024-09-21 05:43:43.240Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "tjtjykz4",
						"name": "role",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"super_admin",
								"org_admin",
								"group_admin",
								"member"
							]
						}
					},
					{
						"system": false,
						"id": "9khchabx",
						"name": "orgs",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "kliajhrdw3c7vde",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "wniiwj6t",
						"name": "groups",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "khdv6l9mcukmovw",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id || @request.auth.role = \"super_admin\"",
				"viewRule": "id = @request.auth.id || @request.auth.role = \"super_admin\"",
				"createRule": "@request.auth.role = \"super_admin\"",
				"updateRule": "id = @request.auth.id || @request.auth.role = \"super_admin\"",
				"deleteRule": "@request.auth.role = \"super_admin\"",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": false,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": "@request.auth.role = \"super_admin\"",
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "19uyx278kmcel93",
				"created": "2024-09-02 04:41:32.369Z",
				"updated": "2024-09-27 05:10:28.473Z",
				"name": "event_templates",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "b9xxjz7w",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "0jwl9eld",
						"name": "recurrence",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"daily",
								"weekly",
								"monthly"
							]
						}
					},
					{
						"system": false,
						"id": "kfrvw8bt",
						"name": "group",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "khdv6l9mcukmovw",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "(group.permissions_via_group.user.id ?= @request.auth.id && \"admin,manager\" ~ group.permissions_via_group.role)",
				"viewRule": "group.permissions_via_group.user.id ?= @request.auth.id && \"admin,manager\" ~ group.permissions_via_group.role",
				"createRule": "group.permissions_via_group.user.id ?= @request.auth.id && \"admin,manager\" ~ group.permissions_via_group.role",
				"updateRule": "group.permissions_via_group.user.id ?= @request.auth.id && \"admin,manager\" ~ group.permissions_via_group.role",
				"deleteRule": "group.permissions_via_group.user.id ?= @request.auth.id && \"admin,manager\" ~ group.permissions_via_group.role",
				"options": {}
			},
			{
				"id": "kliajhrdw3c7vde",
				"created": "2024-09-02 04:41:44.906Z",
				"updated": "2024-09-27 05:26:09.259Z",
				"name": "orgs",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "pqbs8dqo",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.orgs:each ?= id",
				"viewRule": "@request.auth.orgs:each ?= id",
				"createRule": "@request.auth.role = \"super_admin\"",
				"updateRule": "@collection.permissions.user ?= @request.auth.id && @collection.permissions.org ?= id && \"admin,manager\" ~ @collection.permissions.role",
				"deleteRule": "@request.auth.role = \"super_admin\"",
				"options": {}
			},
			{
				"id": "khdv6l9mcukmovw",
				"created": "2024-09-02 04:42:00.310Z",
				"updated": "2024-09-03 05:15:47.969Z",
				"name": "groups",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ibu83gr8",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "rdgojxzm",
						"name": "org",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "kliajhrdw3c7vde",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = \"super_admin\"",
				"viewRule": "@request.auth.role = \"super_admin\"",
				"createRule": "@request.auth.role = \"super_admin\"",
				"updateRule": "@request.auth.role = \"super_admin\"",
				"deleteRule": "@request.auth.role = \"super_admin\"",
				"options": {}
			},
			{
				"id": "vgjqljis1csnkxw",
				"created": "2024-09-02 04:47:53.398Z",
				"updated": "2024-09-05 03:56:15.422Z",
				"name": "role_templates",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "gck5ybgu",
						"name": "event_template",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "19uyx278kmcel93",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "mrnnadfd",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "i9z90t4i",
						"name": "description",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = \"super_admin\"",
				"viewRule": "@request.auth.role = \"super_admin\"",
				"createRule": "@request.auth.role = \"super_admin\"",
				"updateRule": "@request.auth.role = \"super_admin\"",
				"deleteRule": "@request.auth.role = \"super_admin\"",
				"options": {}
			},
			{
				"id": "f4i59kb5utfv4p2",
				"created": "2024-09-02 05:11:40.332Z",
				"updated": "2024-09-07 06:15:49.650Z",
				"name": "events",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ctzxbbdf",
						"name": "event_template",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "19uyx278kmcel93",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "gnetzh0b",
						"name": "datetime",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = \"super_admin\"",
				"viewRule": "@request.auth.role = \"super_admin\"",
				"createRule": "@request.auth.role = \"super_admin\"",
				"updateRule": "@request.auth.role = \"super_admin\"",
				"deleteRule": "@request.auth.role = \"super_admin\"",
				"options": {}
			},
			{
				"id": "sovlfrd5setvbbl",
				"created": "2024-09-02 05:13:09.274Z",
				"updated": "2024-09-14 07:00:50.786Z",
				"name": "roles",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "udoijbrx",
						"name": "role_template",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "vgjqljis1csnkxw",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "tfmtkkep",
						"name": "event",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "f4i59kb5utfv4p2",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "meogdjps",
						"name": "assigned_to",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.role = \"super_admin\"",
				"viewRule": "@request.auth.role = \"super_admin\"",
				"createRule": "@request.auth.role = \"super_admin\"",
				"updateRule": "@request.auth.role = \"super_admin\"",
				"deleteRule": "@request.auth.role = \"super_admin\"",
				"options": {}
			},
			{
				"id": "77t5i3stea8r1hb",
				"created": "2024-09-26 04:56:02.029Z",
				"updated": "2024-09-26 04:56:02.029Z",
				"name": "permissions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zxidosc4",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "5s5gxyn1",
						"name": "org",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "kliajhrdw3c7vde",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ydzf9flk",
						"name": "group",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "khdv6l9mcukmovw",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "zdvj3qci",
						"name": "role",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"admin",
								"manager"
							]
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
