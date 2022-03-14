function maleUsers(db) {
  getData(db, 'users', { gender: 'M' }, function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Male users');
    console.log(docs);
    console.log('');
  });
}

function productID3(db) {
  getData(db, 'products', { _id: 3 }, function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Product ID 3');
    console.log(docs);
    console.log('');
  });
}

function femaleUsersCount(db) {
  db.collection('users').count({ gender: 'F' }, function (err, counter) {
    if (err) {
      throw err;
    }

    console.log('Female users count');
    console.log(counter);
    console.log('');
  });
}

function usersOrderedByName(db) {
  db.collection('users').find().sort({ name: 1 }).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Users Orderded by Name');
    console.log(docs);
    console.log('');
  });
}

function productsLimitTo5(db) {
  db.collection('products').find().limit(5).toArray(function (err, docs) {
    if (err) {
      throw err;
    }

    console.log('Products Limit to 5');
    console.log(docs);
    console.log('');
  });
}

function getData(db, collection, filter, callback) {
  let arr;

  db.collection(collection).find(filter === undefined ? {} : filter).toArray(callback);

  return arr;
}

module.exports = {
  async up(db, client) {
    // TODO write your migration here.
    // See https://github.com/seppevs/migrate-mongo/#creating-a-new-migration-script
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      maleUsers(db),
      productID3(db),
      femaleUsersCount(db),
      usersOrderedByName(db),
      productsLimitTo5(db),
    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // await db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});
    return
  }
};
