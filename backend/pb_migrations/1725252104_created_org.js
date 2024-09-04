/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "kliajhrdw3c7vde",
    "created": "2024-09-02 04:41:44.906Z",
    "updated": "2024-09-02 04:41:44.906Z",
    "name": "org",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "pqbs8dqo",
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
  const collection = dao.findCollectionByNameOrId("kliajhrdw3c7vde");

  return dao.deleteCollection(collection);
})
