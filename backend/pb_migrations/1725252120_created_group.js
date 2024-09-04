/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "khdv6l9mcukmovw",
    "created": "2024-09-02 04:42:00.310Z",
    "updated": "2024-09-02 04:42:00.310Z",
    "name": "group",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ibu83gr8",
        "name": "name",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("khdv6l9mcukmovw");

  return dao.deleteCollection(collection);
})
