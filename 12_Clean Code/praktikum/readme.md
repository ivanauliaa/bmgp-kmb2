**Problem 1**

Banyak kekurangan: 5

Bagian mana saja:
- Naming class user
- Naming class userservice
- Naming fungsi getallusers
- Naming fungsi getuserbyid
- Naming parameter userid di fungsi getuserbyid

Alasan:
- Naming class seharusnya menggunakan UpperCamelCase, sehingga seharusnya "class user" diubah menjadi "class User", dan "class userservice" diubah menjadi "class UserService"
- Naming function/method seharusnya menggunakan lowerCamelCase, sehingga seharusnya fungsi "getallusers" dibuah menjadi "getAllUsers", dan fungsi "getuserbyid" diubah menjadi "getUserByID"
- Naming parameter "userid" di fungsi "getuserbyid" seharusnya diubah menjadi "id", karena dari nama fungsinya sudah merepresentasikan untuk mendapatkan data user by id (id user)

**Problem 2**

``` java
class Kendaraan {
  var totalRoda  = 0;
  var kecepatanPerJam = 0;
}

class Mobil extends Kendaraan {
  void berjalan() {
    tambahKecepatan(10);
  }

  tambahKecepatan(var kecepatanBaru) {
    kecepatanPerJam = kecepatanPerJam + kecepatanBaru;
  }
}

void main() {
  mobilCepat = new Mobil();

  mobilCepat.berjalan();
  mobilCepat.berjalan();
  mobilCepat.berjalan();

  mobilLamban = new Mobil();

  mobilLamban.berjalan();
}
```