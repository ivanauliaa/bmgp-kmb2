# Praktikum

Create User

![](./screenshots/Screenshot%202022-04-10%20201714.jpg)

Get Users

![](./screenshots/Screenshot%202022-04-10%20201737.jpg)

Directories

![](./screenshots/Screenshot%202022-04-10%20201832.jpg)

# Resume
- clean architecture: pemisahan tanggung jawab (separation of concerns)
- tujuan: untuk membuat aplikasi yang modular, scalable, maintainable, dan testable
- hexagonal architecture: API layer (http, message broker), domain layer (business logic), SPI layer (service provider)
- constraint: independent of frameworks, testable, independent of UI, independent of database, independent of any external
- benefit: faster development jangka panjang, mudah jika ingin mengganti suatu service, standard structure
- CA layer: entities (business object), use case - domain layer (business logic), controller - presentation layer (handler), driver - data layer (data management)
- domain driven design (DDD): pengembangan yang berorientasi pada domain (business logic)
- CA adalah sebuah software architecture, sedangkan DDD adalah software design technique