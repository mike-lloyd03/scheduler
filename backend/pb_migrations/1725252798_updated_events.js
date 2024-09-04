/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("19uyx278kmcel93")

  // remove
  collection.schema.removeField("cnubuvre")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "kfrvw8bt",
    "name": "group",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "khdv6l9mcukmovw",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("19uyx278kmcel93")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cnubuvre",
    "name": "group_id",
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
  collection.schema.removeField("kfrvw8bt")

  return dao.saveCollection(collection)
})
