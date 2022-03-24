# Praktikum

Create new user

![](./screenshots/Screenshot%202022-03-24%20135009.jpg)

Get all users

![](./screenshots/Screenshot%202022-03-24%20135033.jpg)

Get user by id

![](./screenshots/Screenshot%202022-03-24%20135110.jpg)

Update user by id

![](./screenshots/Screenshot%202022-03-24%20135130.jpg)

Delete user by id

![](./screenshots/Screenshot%202022-03-24%20135206.jpg)

# Resume

- echo merupakan salah satu library http server di go
- untuk membuat handler function echo, cukup dengan membuat function dengan 1 parameter dengan tipe echo.Context dan mengembalikan error
- untuk menjalankan echo, kita perlu memanggil fungsi New(), mendaftarkan route beserta handlernya, dan memanggil fungsi Start dengan melewatkan port
- untuk mengembalikan response, kita bisa memanggil fungsi JSON() dari echo.Context lalu melewatkan status code dan data
- untuk mengambil path parameter bisa memanggil fungsi Param() dari echo.Context lalu melewatkan nama path parameter
- untuk mengambil query parameter bisa memanggil fungsi QueryParam() dari echo.Context lalu melewatkan nama query paramter
- untuk mengambil form value (payload) bisa memanggil fungsi FormValue() dari echo.Context lalu melewatkan nama form value
- alternatif mengambil form value, kita bisa melakukan binding ke sebuah struct (JSON tagged) dengan menyiapkan data struct kosong, lalu memanggil fungsi Bind() dari echo.Context lalu melewatkan reference dari struct yg tadi
