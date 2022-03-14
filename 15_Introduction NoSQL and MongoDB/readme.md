- relational db (rdb) memiliki benefit: dirancang untuk segala keperluan, memiliki standar yang jelas, dan memiliki banyak dukungan tool
- rdb mengikuti prinsip ACID (atomicity, consistency, isolation, durabililty)
- noql menyediakan mekanisme yang lebih fleksibel dibanding rdb
- kelebihan nosql: schema less, fast development, support big size file, support cluster
- kapan perlu menggunakan nosql: membutuhkan skema fleksibel, acid tidak diperlukan, terdistribusi, data loggin, data sementara
- kapan tidak direkomendasikan menggunakan nosql: data finansial, data transaksi, business critical, acid compliant
- tipe-tipe nosql: key/value (redis), column family (cassandra), graph (neo4j), document based (mongodb), dsb

**mongodb**
- create collection: db.createCollection('namaCollection')
- insert one: db.coll.insertOne(data)
- insert many: db.coll.insertMany(arrayOfData)
- find: db.coll.find(query)
- update one: db.coll.updateOne(query, options)
- update many: db.coll.updateMany(query, options)
- delete one: db.coll.deleteOne(query)
- delete many: db.coll.deleteMany(query)
- count (find): db.coll.find(query).count()
- sort (find): db.coll.find(query).sort(sort)
- limit (find): db.coll.find(query).limit(number)