# 3. Codificación y Resolución de Problemas
- Problema. Desarrolle un endpoint de API RESTful en un
  lenguaje de su elección (Python, Java, Node.js, etc.)
  que gestione lo siguiente:
  - Reciba un payload JSON con una lista de
  productos con sus precios y cantidades.
  - Calcule el costo total del pedido, incluyendo el
  cargo por envío a domicilio.
  - Aplique un descuento basado en el monto total
  del pedido, si corresponde.
  - Devuelva una respuesta JSON con el costo
  total y el descuento aplicado, si corresponde.
  - Considere el sistema de "estrato" en Colombia
  y si el cargo por envío podría variar según este
  valor.
- Requerimientos:
  - Escriba un código limpio y bien documentado.
  - Incluya pruebas unitarias para verificar la
  funcionalidad de su código.
  - Demuestre su comprensión de las mejores prácticas
  para el desarrollo del API.


### Prácticas y tópicos usados en la API
- [Compile-time variables](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)
- Graceful shutdown
- Structured logs
- Health checks 
- [Hexagonal pattern](https://alistair.cockburn.us/hexagonal-architecture/)
- Design patterns ([Bridge](https://refactoring.guru/design-patterns/bridge), [Decorator](https://refactoring.guru/design-patterns/decorator))

### Documentación
[OpenAPI specification](docs/OpenAPI.json)

[Required environment variables](.env.example)

###### ¿Cómo ejecutar el servidor HTTP?
```shell
make run
```
###### ¿Cómo ejecutar los test unitarios?
```shell
make test
```
###### ¿Cómo probar?
```shell
curl -X "POST" "http://localhost:8080/v1/orders" -H "Content-Type: application/json" -d '{"stratum": 1, "products": [{"price": 10, "quantity": 2}]}'
```
### Decisiones de arquitectura
###### Go project layout standard
Decidí seguir [Go project layout standard](https://github.com/golang-standards/project-layout) que es un standard no oficial bastante adoptado por la comunidad de Go.

###### Árbol de directorios
Construí el árbol de directorios siguiendo los conceptos de [hexagonal architecture pattern](https://alistair.cockburn.us/hexagonal-architecture/).
```
.
├── cmd
├── internal
│   ├── app
│   │   ├── business (Use cases, business rules, data models and ports)
│   │   ├── input    (Everything related to "drive" adapters)
│   │   └── output   (Everything related to "driven" adapters)
│   └── container (DI container)
└── pkg (Public and global code, potencially libraries)
```