function bookAuthorID1And2(db) {
  db.collection('books').find({ authorID: { '$in': [1, 2] } }).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Books with Author ID 1 & 2')
    console.log(docs);
    console.log();
  });
}

function bookTitleAndPriceAuthorID1(db) {
  db.collection('books').find({ authorID: 1 }, { projection: { title: 1, price: 1 } }).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Book titles and prices with Author ID 1')
    console.log(docs);
    console.log();
  });
}

function bookAuthorID2TotalPages(db) {
  db.collection('books').aggregate([
    { '$match': { authorID: 2 } },
    { '$group': { totalPages: { '$sum': '$stats.page' }, _id: null } },
    { '$project': { _id: '2', totalPages: 1 } },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Book with Author ID 2 Total Pages')
    console.log(docs);
    console.log();
  });
}

function relatedAuthorBooks(db) {
  db.collection('authors').aggregate([
    { '$lookup': { from: 'books', localField: '_id', foreignField: 'authorID', as: 'books' } },
    { '$unwind': '$books' },
    { '$unwind': '$books.stats' },
    { '$unwind': '$books.category' },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Related Author and Books')
    console.log(docs);
    console.log();
  });
}

function relatedBookAuthor(db) {
  db.collection('books').aggregate([
    { '$lookup': { from: 'authors', localField: 'authorID', foreignField: '_id', as: 'author' } },
    { '$unwind': '$author' },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Related Book and Author')
    console.log(docs);
    console.log();
  });
}

function relatedBookAuthorPublisher(db) {
  db.collection('books').aggregate([
    { '$lookup': { from: 'authors', localField: 'authorID', foreignField: '_id', as: 'author' } },
    { '$unwind': '$author' },
    { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
    { '$unwind': '$publisher' },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Related Book, Author, and Publisher')
    console.log(docs);
    console.log();
  });
}

function summary(db) {
  db.collection('authors').aggregate([
    { '$lookup': { from: 'books', localField: '_id', foreignField: 'authorID', as: 'books' } },
    { '$project': { _id: { '$concat': ['$firstName', ' ', '$lastName'] }, number_of_books: { '$size': '$books' }, list_of_book: '$books.title' } },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Author books summary')
    console.log(docs);
    console.log();
  });
}

function discount(db) {
  db.collection('books').aggregate([
    { '$project': { title: 1, price: 1, promo: { '$cond': { if: { '$lt': ['$price', 60000] }, then: '1%', else: { '$cond': { if: { '$lt': ['$price', 90000] }, then: '2%', else: '3%' } } } } } },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Discount Promo')
    console.log(docs);
    console.log();
  });
}

function bookSortByPrice(db) {
  db.collection('books').aggregate([
    { '$lookup': { from: 'authors', localField: 'authorID', foreignField: '_id', as: 'author' } },
    { '$sort': { price: -1 } },
    { '$unwind': '$author' },
    { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
    { '$unwind': '$publisher' },
    { '$project': { title: 1, price: 1, author: { '$concat': ['$author.firstName', ' ', '$author.lastName'] }, publisher: '$publisher.publisherName' } },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Books Sort by Price')
    console.log(docs);
    console.log();
  });
}

function bookPricePublisher(db) {
  db.collection('books').aggregate([
    { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
    { '$unwind': '$publisher' },
    { '$project': { title: 1, price: 1, publisher: '$publisher.publisherName' } },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Book, Price, and Publisher')
    console.log(docs);
    console.log();
  });
}

function book3And4PricePublisher(db) {
  db.collection('books').aggregate([
    { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
    { '$match': { _id: { '$in': [3, 4] } } },
    { '$unwind': '$publisher' },
    { '$project': { title: 1, price: 1, publisher: '$publisher.publisherName' } },
  ]).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Book ID 3 & 4, Price, and Publisher')
    console.log(docs);
    console.log();
  });
}

module.exports = {
  async up(db, client) {
    // TODO write your migration here.
    // See https://github.com/seppevs/migrate-mongo/#creating-a-new-migration-script
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      bookAuthorID1And2(db),
      bookTitleAndPriceAuthorID1(db),
      bookAuthorID2TotalPages(db),
      relatedAuthorBooks(db),
      relatedBookAuthor(db),
      relatedBookAuthorPublisher(db),
      summary(db),
      discount(db),
      bookSortByPrice(db),
      bookPricePublisher(db),
      book3And4PricePublisher(db),
    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});

    return;
  }
};
