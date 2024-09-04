/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("vgjqljis1csnkxw")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dke3wipu",
    "name": "event_id",
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

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("vgjqljis1csnkxw")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dke3wipu",
    "name": "event_it",
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

  return dao.saveCollection(collection)
})
