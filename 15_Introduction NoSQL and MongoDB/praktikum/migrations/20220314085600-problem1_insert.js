function operatorObject(_id, name) {
  return {
    _id,
    name,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

function productTypeObject(_id, name) {
  return {
    _id,
    name,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

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

function productDescriptionObject(_id, product_id, description) {
  return {
    _id,
    product_id,
    description,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

function paymentMethodObject(_id, name, status) {
  return {
    _id,
    name,
    status,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

function userObject(_id, name, status, dob, gender) {
  return {
    _id,
    name,
    status,
    dob,
    gender,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

function transactionObject(_id, user_id, payment_method_id, status, total_qty, total_price) {
  return {
    _id,
    user_id,
    payment_method_id,
    status,
    total_qty,
    total_price,
    created_at: Date.now(),
    updated_at: Date.now(),
  };
}

function transactionDetailObject(transaction_id, product_id, status, qty, price) {
  return {
    transaction_id,
    product_id,
    status,
    qty,
    price,
    created_at: Date.now(),
    updated_at: Date.now(),
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
    // db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: true}});

    return Promise.all([
      db.collection('operators').insertMany([
        operatorObject(1, 'Operator 1'),
        operatorObject(2, 'Operator 2'),
        operatorObject(3, 'Operator 3'),
        operatorObject(4, 'Operator 4'),
        operatorObject(5, 'Operator 5'),
      ]),
      db.collection('product_types').insertMany([
        productTypeObject(1, 'Rumah Tangga'),
        productTypeObject(2, 'Elektronik'),
        productTypeObject(3, 'Bahan Bangunan'),
      ]),
      db.collection('products').insertMany([
        productObject(1, 1, 3, 'RT1', 'Water Mop', '200'),
        productObject(2, 1, 3, 'RT2', 'Rak Sepatu', '200'),
        productObject(3, 2, 1, 'E1', 'LED TV', '404'),
        productObject(4, 2, 1, 'E2', 'Soundbar', '200'),
        productObject(5, 2, 1, 'E3', 'Mesin Cuci', '200'),
        productObject(6, 3, 4, 'BB1', 'Keramik', '404'),
        productObject(7, 3, 4, 'BB2', 'Wastafel', '404'),
        productObject(8, 3, 4, 'BB3', 'Cat Tembok', '404'),
      ]),
      db.collection('product_descriptions').insertMany([
        productDescriptionObject(1, 1, 'Lorem'),
        productDescriptionObject(2, 2, 'Ipsum'),
        productDescriptionObject(3, 3, 'Dolor'),
        productDescriptionObject(4, 4, 'Sit'),
        productDescriptionObject(5, 5, 'Amet'),
        productDescriptionObject(6, 6, 'Consectetuer'),
        productDescriptionObject(7, 7, 'Adipiscing'),
        productDescriptionObject(8, 8, 'Elit'),
      ]),
      db.collection('payment_methods').insertMany([
        paymentMethodObject(1, 'OVO', 200),
        paymentMethodObject(2, 'Gopay', 200),
        paymentMethodObject(3, 'ShopeePay', 200),
      ]),
      db.collection('users').insertMany([
        userObject(1, 'Ivan', 200, Date('2000-10-27'), 'M'),
        userObject(2, 'Aulia', 200, Date('2000-10-28'), 'M'),
        userObject(3, 'Rahman', 200, Date('2000-10-29'), 'F'),
        userObject(4, 'Karim', 200, Date('2000-10-30'), 'M'),
        userObject(5, 'Benzema', 200, Date('2000-10-31'), 'F'),
      ]),
      db.collection('transactions').insertMany([
        transactionObject(1, 1, 1, 'PENDING', 6, 60000),
        transactionObject(2, 1, 2, 'PENDING', 15, 150000),
        transactionObject(3, 1, 3, 'CANCELLED', 16, 160000),
        transactionObject(4, 2, 1, 'PENDING', 9, 90000),
        transactionObject(5, 2, 2, 'SUCCESS', 18, 180000),
        transactionObject(6, 2, 3, 'CANCELLED', 6, 60000),
        transactionObject(7, 3, 1, 'SUCCESS', 8, 80000),
        transactionObject(8, 3, 2, 'PENDING', 18, 180000),
        transactionObject(9, 3, 3, 'PENDING', 11, 110000),
        transactionObject(10, 4, 1, 'CANCELLED', 12, 120000),
        transactionObject(11, 4, 2, 'PENDING', 21, 210000),
        transactionObject(12, 4, 3, 'PENDING', 6, 60000),
        transactionObject(13, 5, 1, 'CANCELLED', 15, 150000),
        transactionObject(14, 5, 2, 'PENDING', 16, 160000),
        transactionObject(15, 5, 3, 'SUCCESS', 9, 90000),
      ]),
      db.collection('transaction_details').insertMany([
        transactionDetailObject(1, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(1, 2, 'PRE-ORDER', 2, 20000),
        transactionDetailObject(1, 3, 'AVAILABLE', 3, 30000),
        transactionDetailObject(2, 4, 'AVAILABLE', 4, 40000),
        transactionDetailObject(2, 5, 'AVAILABLE', 5, 50000),
        transactionDetailObject(2, 6, 'AVAILABLE', 6, 60000),
        transactionDetailObject(3, 7, 'AVAILABLE', 7, 70000),
        transactionDetailObject(3, 8, 'PRE-ORDER', 8, 80000),
        transactionDetailObject(3, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(4, 2, 'AVAILABLE', 2, 20000),
        transactionDetailObject(4, 3, 'AVAILABLE', 3, 30000),
        transactionDetailObject(4, 4, 'AVAILABLE', 4, 40000),
        transactionDetailObject(5, 5, 'PRE-ORDER', 5, 50000),
        transactionDetailObject(5, 6, 'AVAILABLE', 6, 60000),
        transactionDetailObject(5, 7, 'AVAILABLE', 7, 70000),
        transactionDetailObject(6, 8, 'AVAILABLE', 8, 80000),
        transactionDetailObject(6, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(6, 2, 'AVAILABLE', 2, 20000),
        transactionDetailObject(7, 3, 'PRE-ORDER', 3, 30000),
        transactionDetailObject(7, 4, 'AVAILABLE', 4, 40000),
        transactionDetailObject(7, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(8, 5, 'AVAILABLE', 5, 50000),
        transactionDetailObject(8, 6, 'AVAILABLE', 6, 60000),
        transactionDetailObject(8, 7, 'AVAILABLE', 7, 70000),
        transactionDetailObject(9, 8, 'AVAILABLE', 8, 80000),
        transactionDetailObject(9, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(9, 2, 'AVAILABLE', 2, 20000),
        transactionDetailObject(10, 3, 'AVAILABLE', 3, 30000),
        transactionDetailObject(10, 4, 'PRE-ORDER', 4, 40000),
        transactionDetailObject(10, 5, 'AVAILABLE', 5, 50000),
        transactionDetailObject(11, 6, 'AVAILABLE', 6, 60000),
        transactionDetailObject(11, 7, 'AVAILABLE', 7, 70000),
        transactionDetailObject(11, 8, 'AVAILABLE', 8, 80000),
        transactionDetailObject(12, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(12, 2, 'PRE-ORDER', 2, 20000),
        transactionDetailObject(12, 3, 'AVAILABLE', 3, 30000),
        transactionDetailObject(13, 4, 'AVAILABLE', 4, 40000),
        transactionDetailObject(13, 5, 'AVAILABLE', 5, 50000),
        transactionDetailObject(13, 6, 'AVAILABLE', 6, 60000),
        transactionDetailObject(14, 7, 'AVAILABLE', 7, 70000),
        transactionDetailObject(14, 8, 'PRE-ORDER', 8, 80000),
        transactionDetailObject(14, 1, 'AVAILABLE', 1, 10000),
        transactionDetailObject(15, 2, 'AVAILABLE', 2, 20000),
        transactionDetailObject(15, 3, 'AVAILABLE', 3, 30000),
        transactionDetailObject(15, 4, 'AVAILABLE', 4, 40000),
      ]),
    ]);
  },

  async down(db, client) {
    // TODO write the statements to rollback your migration (if possible)
    // Example:
    // db.collection('albums').updateOne({artist: 'The Beatles'}, {$set: {blacklisted: false}});

    return Promise.all([
      emptyData(db, 'operators', 5),
      emptyData(db, 'product_types', 3),
      emptyData(db, 'products', 8),
      emptyData(db, 'product_descriptions', 8),
      emptyData(db, 'payment_methods', 3),
      emptyData(db, 'users', 5),
      emptyData(db, 'transactions', 15),
      emptyData(db, 'transaction_details'),
    ]);
  }
};
