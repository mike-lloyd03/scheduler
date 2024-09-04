/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("31sibofialm3u1d")

  // remove
  collection.schema.removeField("ouuhminh")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "k2cvh2ga",
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
  const collection = dao.findCollectionByNameOrId("31sibofialm3u1d")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ouuhminh",
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
  collection.schema.removeField("k2cvh2ga")

  return dao.saveCollection(collection)
})
