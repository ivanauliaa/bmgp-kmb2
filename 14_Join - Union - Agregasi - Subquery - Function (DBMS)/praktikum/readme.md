## Todo First

- Melakukan migrate migrate-up.sql

``` mysql -u root alta_online_shop < migrate-up.sql ```

![migration](../screenshots/Screenshot%202022-03-12%20072134.jpg)

## Problem 1

- **1. Insert**

a. 5 operators

![insert](../screenshots/Screenshot%202022-03-12%20073505.jpg)

b. 3 product types

![insert](../screenshots/Screenshot%202022-03-12%20073522.jpg)

c. 2 product dg product type 1 dan operator 3

d. 3 product dg product type 2 dan operator 1

e. 3 product dg product type 3 dan operator 4

![insert](../screenshots/Screenshot%202022-03-12%20073737.jpg)

f. product description setiap product

![insert](../screenshots/Screenshot%202022-03-12%20073800.jpg)

g. 3 payment methods

![insert](../screenshots/Screenshot%202022-03-12%20073818.jpg)

h. 5 users

![insert](../screenshots/Screenshot%202022-03-12%20081035.jpg)

i. 3 transactions setiap user

![insert](../screenshots/Screenshot%202022-03-12%20075359.jpg)

j. 3 product setiap transaction

![insert](../screenshots/Screenshot%202022-03-12%20075619.jpg)


- **2. Select**

a. nama user gender M

b. product 3

c. user created_at 7 hari kebelakang dan nama mengandung 'a'

d. jumlah user gender F

e. users name asc

f. 5 product

![select](../screenshots/Screenshot%202022-03-12%20082523.jpg)

- **3. Update**

a. nama product 1 -> 'product dummy'

b. qty transaction detail dg product 1 -> 3

![update](../screenshots/Screenshot%202022-03-12%20083213.jpg)

- **4. Delete**

a. product 1

b. product type 1

![delete](../screenshots/Screenshot%202022-03-12%20084313.jpg)

## Problem 2

- 1. Gabungkan data transaksi dari user id 1 dan user id 2

- 2. Tampilkan jumlah harga transaksi user id 1

- 3. Tampilkan total transaksi dengan product type 2

![1, 2, 3](../screenshots/Screenshot%202022-03-13%20092217.jpg)

- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan

![4](../screenshots/Screenshot%202022-03-13%20092900.jpg)

- 5. Tampilkan semua field table transaction, field name table product dan field name table user

![5](../screenshots/Screenshot%202022-03-13%20092554.jpg)

- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud

- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate

![function 6 7](../screenshots/Screenshot%202022-03-13%20091744.jpg)

- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query

![8](../screenshots/Screenshot%202022-03-13%20092558.jpg)