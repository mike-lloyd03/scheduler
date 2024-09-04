/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "31sibofialm3u1d",
    "created": "2024-09-02 04:43:42.576Z",
    "updated": "2024-09-02 04:43:42.576Z",
    "name": "org_groups",
    "type": "base",
    "system": false,
    "schema": [
      {
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
      },
      {
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
  const collection = dao.findCollectionByNameOrId("31sibofialm3u1d");

  return dao.deleteCollection(collection);
})
