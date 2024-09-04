/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("khdv6l9mcukmovw")

  collection.name = "groups"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("khdv6l9mcukmovw")

  collection.name = "group"

  return dao.saveCollection(collection)
})
