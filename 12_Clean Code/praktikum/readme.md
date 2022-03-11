**Problem 1**

Banyak kekurangan: 7

Bagian mana saja:

- Naming struct userservice
- Naming method getallusers
- Naming method getuserbyid
- Naming properti t di userservice
- Tipe receiver di method getallusers dan getuserbyid
- Naming receiver u di method getallusers dan getuserbyid
- Naming variabel r penerima return forrange di getuserbyid

Alasan:

- Naming struct dan method seharusnya menggunakan camelCase (untuk penentuan kapital huruf pertama menyesuaikan kebutuhan exported atau tidak), sehingga seharusnya struct "userservice" diubah menjadi "userService", method "getallusers" diubah menjadi "getAllUsers", dan method "getuserbyid" diubah menjadi "getUserByID"
- Naming properti t di userservice kurang representatif, seharusnya diubah menjadi users
- Tipe receiver di method getallusers dan getuserbyid seharusnya menggunakan pointer, agar lebih hemat memori dengan tidak terus menerus membuat copy dari data struct terkait
- Naming receiver u di method getallusers dan getuserbyid kurang representatif, seharusnya diubah menjadi misalnya service
- Naming variabel r penerima return forrange di getuserbyid kurang representatif, seharusnya diubah menjadi user

**Problem 2**

```go
package main

type kendaraan struct {
  totalRoda int
  kecepatanPerJam int
}

type mobil struct {
  kendaraan
}

func (m *mobil) berjalan() {
  m.tambahKecepatan(10)
}

func (m *mobil) tambahKecepatan(kecepatanBaru) {
  m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
  mobilcepat := mobil{}

  iterasi := 3
  for (i := 0; i < iterasi; i++) {
    mobilcepat.berjalan()
  }

  mobillamban := mobil{}
  mobillamban.berjalan()
}

```
