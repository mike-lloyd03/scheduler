/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("31sibofialm3u1d");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "31sibofialm3u1d",
    "created": "2024-09-02 04:43:42.576Z",
    "updated": "2024-09-02 04:52:51.756Z",
    "name": "org_groups",
    "type": "base",
    "system": false,
    "schema": [
      {
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
      },
      {
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
})
