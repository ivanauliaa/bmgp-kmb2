module.exports = {
  async up(db, client) {
    // TODO write your migration here.
    // See https://github.com/seppevs/migrate-mongo/#creating-a-new-migration-script
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      db.createCollection('users'),
      db.createCollection('products'),
      db.createCollection('product_descriptions'),
      db.createCollection('product_types'),
      db.createCollection('operators'),
      db.createCollection('payment_methods'),
      db.createCollection('transactions'),
      db.createCollection('transaction_details'),
    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});

    return Promise.all([
      db.collection('users').drop(),
      db.collection('products').drop(),
      db.collection('product_descriptions').drop(),
      db.collection('product_types').drop(),
      db.collection('operators').drop(),
      db.collection('payment_methods').drop(),
      db.collection('transactions').drop(),
      db.collection('transaction_details').drop(),
    ]);
  }
};
