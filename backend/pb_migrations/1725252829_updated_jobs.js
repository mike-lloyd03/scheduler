/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("vgjqljis1csnkxw")

  // remove
  collection.schema.removeField("dke3wipu")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "gck5ybgu",
    "name": "event",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "19uyx278kmcel93",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("vgjqljis1csnkxw")

  // add
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

  // remove
  collection.schema.removeField("gck5ybgu")

  return dao.saveCollection(collection)
})
