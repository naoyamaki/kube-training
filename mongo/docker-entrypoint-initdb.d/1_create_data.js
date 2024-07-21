db = db.getSiblingDB('yourDatabaseName');

db.createUser({
  user: 'yourUsername',
  pwd: 'yourPassword',
  roles: [{ role: 'readWrite', db: 'yourDatabaseName' }]
});

db.yourCollectionName.insertMany([
  { name: 'Example', value: 1 },
  { name: 'Another Example', value: 2 }
]);
