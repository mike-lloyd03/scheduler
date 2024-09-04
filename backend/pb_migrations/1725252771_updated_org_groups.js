/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("31sibofialm3u1d")

  // remove
  collection.schema.removeField("6sejmgcj")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "zu2v72ri",
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
  const collection = dao.findCollectionByNameOrId("31sibofialm3u1d")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "6sejmgcj",
    "name": "org_id",
    "type": "number",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "noDecimal": false
    }
  }))

  // remove
  collection.schema.removeField("zu2v72ri")

  return dao.saveCollection(collection)
})
