function bookObject(_id, title, authorID, publisherID, price, stats, publishedAt, category) {
  return {
    _id,
    title,
    authorID,
    publisherID,
    price,
    stats,
    publishedAt,
    category,
  };
}

function authorObject(_id, firstName, lastName) {
  return {
    _id,
    firstName,
    lastName,
  };
}

function publisherObject(_id, publisherName) {
  return {
    _id,
    publisherName,
  };
}

function emptyData(db, collection, length) {
  if (length === undefined) {
    db.collection(collection).deleteMany({});
  } else {
    db.collection(collection).deleteMany(
      {
        '_id': {
          '$in': Array.from({ length: length }, (_, i) => i + 1),
        },
      },
    );
  }
}

module.exports = {
  async up(db, client) {
    // TODO write your migration here.
    // See https://github.com/seppevs/migrate-mongo/#creating-a-new-migration-script
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      db.collection('books').insertMany([
        bookObject(1, 'Wawasan Pancasila', 1, 1, 71200, {page: 324, weight: 300}, new Date('2018-10-01'), ['social', 'politics']),
        bookObject(2, 'Mata Air Keteladanan', 1, 2, 106260, {page: 672, weight: 650}, new Date('2017-09-01'), ['social', 'politics']),
        bookObject(3, 'Revolusi Pancasila', 1, 1, 54400, {page: 220, weight: 500}, new Date('201505-01'), ['social', 'politics']),
        bookObject(4, 'Self Driving', 2, 1, 58650, {page: 286, weight: 300}, new Date('2018-05-01'), ['self-development']),
        bookObject(5, 'Self Disruption', 2, 2, 83300, {page: 400, weight: 800}, new Date('2018-05-01'), ['self-development']),
      ]),
      db.collection('authors').insertMany([
        authorObject(1, 'Yudi', 'Latif'),
        authorObject(2, 'Rhenald', 'Kasali'),
      ]),
      db.collection('publishers').insertMany([
        publisherObject(1, 'Expose'),
        publisherObject(2, 'Mizan'),
      ]),
    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});

    return Promise.all([
      emptyData(db, 'books', 5),
      emptyData(db, 'authors', 2),
      emptyData(db, 'publishers', 2),
    ]);
  }
};
