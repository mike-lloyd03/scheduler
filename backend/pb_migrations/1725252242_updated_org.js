/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("kliajhrdw3c7vde")

  collection.name = "orgs"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("kliajhrdw3c7vde")

  collection.name = "org"

  return dao.saveCollection(collection)
})
