# Praktikum

Login

![](./screenshots/Screenshot%202022-04-09%20192529.jpg)

Create user without token

![](./screenshots/Screenshot%202022-04-09%20192543.jpg)

Create user with token

![](./screenshots/Screenshot%202022-04-09%20192557.jpg)

# Resume

- clean architecture: pemisahan tanggung jawab (separation of concerns)
- tujuan: untuk membuat aplikasi yang modular, scalable, maintainable, dan testable
- hexagonal architecture: API layer (http, message broker), domain layer (business logic), SPI layer (service provider)
- constraint: independent of frameworks, testable, independent of UI, independent of database, independent of any external
- benefit: faster development jangka panjang, mudah jika ingin mengganti suatu service, standard structure
- CA layer: entities (business object), use case - domain layer (business logic), controller - presentation layer (handler), driver - data layer (data management)
- domain driven design (DDD): pengembangan yang berorientasi pada domain (business logic)
- CA adalah sebuah software architecture, sedangkan DDD adalah software design technique
- JWT merupakan auth middleware yang diletakkan pada domain di clean architecture