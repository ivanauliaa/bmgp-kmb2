# Praktikum

## User

Create new user

![](./screenshots/Screenshot%202022-03-25%20153705.jpg)

Get all users

![](./screenshots/Screenshot%202022-03-25%20153853.jpg)

Get user by id

![](./screenshots/Screenshot%202022-03-25%20153955.jpg)

Update user by id

![](./screenshots/Screenshot%202022-03-25%20154016.jpg)

Delete user by id

![](./screenshots/Screenshot%202022-03-25%20154033.jpg)

## Book

Create new book

![](./screenshots/Screenshot%202022-03-26%20144757.jpg)

Get all books

![](./screenshots/Screenshot%202022-03-26%20144849.jpg)

Get book by id

![](./screenshots/Screenshot%202022-03-26%20144908.jpg)

Update book by id

![](./screenshots/Screenshot%202022-03-26%20144931.jpg)

Delete book by id

![](./screenshots/Screenshot%202022-03-26%20144948.jpg)

# Resume

- object relational mapper (ORM) merupakan suatu teknik untuk mengubah data dari database dengan skema terkait menjadi object di suatu programming language dengan format yang sesuai
- gorm merupakan library orm
- kelebihan menggunakan orm: mengurangi perulangan statement query, langsung mengubah data menjadi siap pakai di suatu programmming language, cache query, screening data sebelum dikirim ke database
- kekurangan menggunakan orm: menambah cost dari suatu proses karena harus melewati lapisan orm terlebih dahulu, memuat relasi yg tidak diperlukan, query yang kompleks bisa sangat ribet dituliskan menggunakan orm, bisa jadi ada sql function yang tidak didukung oleh library orm terkait
- migration digunakan sebagai versioning dalam pengelolaan skema database
- sebelum melakukan proses gorm, kita harus mendapatkan koneksi ke db terkait terlebih dahulu
- kita perlu mendefinisikan skema suatu tabel ke dalam model
- kita perlu mendefinisikan routes yang bisa diterima
- kita perlu mendefinisikan handler untuk masing" routes