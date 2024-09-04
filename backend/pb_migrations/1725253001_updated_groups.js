/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("khdv6l9mcukmovw")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "rdgojxzm",
    "name": "org",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "kliajhrdw3c7vde",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("khdv6l9mcukmovw")

  // remove
  collection.schema.removeField("rdgojxzm")

  return dao.saveCollection(collection)
})
