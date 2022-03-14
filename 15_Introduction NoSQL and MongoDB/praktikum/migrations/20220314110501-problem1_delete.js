function productObject(_id, product_type_id, operator_id, code, name, status) {
  return {
    _id,
    product_type_id,
    operator_id,
    code,
    name,
    status,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

module.exports = {
  async up(db, client) {
    // TODO write your migration here.
    // See https://github.com/seppevs/migrate-mongo/#creating-a-new-migration-script
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      db.collection('products').deleteOne({_id: 1}),
      db.collection('products').deleteMany({product_type_id: 1}),
    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});

    return Promise.all([
      db.collection('products').insertOne(productObject(1, 1, 3, 'RT1', 'product dummy', '200')),
      db.collection('products').insertOne(productObject(2, 1, 3, 'RT2', 'Rak Sepatu', '200')),
    ])
  }
};
