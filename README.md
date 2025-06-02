# uala-microblog-challenge

## 🚀 Setup

### Requisitos

- Go 1.20+ (o compatible)
- [swag](https://github.com/swaggo/swag) para generar la documentación Swagger (`go install github.com/swaggo/swag/cmd/swag@latest`)
- Docker (opcional)

### Instalación y ejecución local

1. Instalar dependencias:    `go mod tidy`

2. Generar la documentación Swagger:    `make swagger`

3. Compilar el proyecto:    `make build`

4. Ejecutar el servidor:    `make run`

El servidor estará disponible en http://localhost:8080/api/v1/swagger/index.html

---

## 🐳 Docker

1. Construir la imagen:    `make docker-build`

2. Ejecutar el contenedor:    `make docker-run`

---

## 📖 Documentación Swagger

La API expone su documentación interactiva utilizando Swagger:

- http://localhost:8080/api/v1/swagger/index.html

Desde ahí podes probar los endpoints, ver los parámetros y las respuestas esperadas.

---

## 🛠️ Tests

- `make test` — Ejecuta los unit tests.
- `make test-integration` — Ejecuta los tests de integración (es un script que crea usuarios, tweets y sigue a usuarios).

---

## 📦 Comandos útiles

- make build — Compila el binario principal.
- make run — Ejecuta el servidor localmente.
- make swagger — Genera la documentación Swagger.
- make docker — Construye y ejecuta el proyecto en Docker.

---

## ✍️ Autor

Carlos Santiago Yanzon  
https://github.com/bizk 