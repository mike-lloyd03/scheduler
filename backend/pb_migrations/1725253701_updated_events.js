/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("19uyx278kmcel93")

  collection.name = "event_templates"

  // remove
  collection.schema.removeField("og1cictc")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("19uyx278kmcel93")

  collection.name = "events"

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "og1cictc",
    "name": "parent_event",
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
