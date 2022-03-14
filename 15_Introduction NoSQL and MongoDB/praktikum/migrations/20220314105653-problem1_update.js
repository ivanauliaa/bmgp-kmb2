module.exports = {
  async up(db, client) {
    // TODO write your migration here.
    // See https://github.com/seppevs/migrate-mongo/#creating-a-new-migration-script
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      db.collection('products').updateOne({_id: 1}, {'$set': {name: 'product dummy'}}),
      db.collection('transaction_details').updateMany({product_id: 1}, {'$set': {qty: 3}}),

    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});

    return Promise.all([
      db.collection('products').updateOne({_id: 1}, {'$set': {name: 'Water Mop'}}),
      db.collection('transaction_details').updateMany({product_id: 1}, {'$set': {qty: 1}}),
    ]);
  }
};
