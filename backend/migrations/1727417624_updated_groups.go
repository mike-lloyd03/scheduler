package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("khdv6l9mcukmovw")
		if err != nil {
			return err
		}

		collection.UpdateRule = types.Pointer("@collection.permissions.user ?= @request.auth.id && ((@collection.permissions.group ?= id && \"admin,manager\" ~ @collection.permissions.role) || (@collection.permissions.org ?= org && @collection.permissions.role ?= \"admin\"))")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("khdv6l9mcukmovw")
		if err != nil {
			return err
		}

		collection.UpdateRule = types.Pointer("@collection.permissions.user ?= @request.auth.id && ((@collection.permissions.group ?= id && \"admin,manager\" ~ @collection.permissions.role) || (@collection.permissions.org = org && @collection.permissions.role ?= \"admin\"))")

		return dao.SaveCollection(collection)
	})
}
