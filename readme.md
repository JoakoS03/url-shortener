# URL Shortener

Este es un proyecto de acortador de URLs escrito en Go utilizando el framework Gin y GORM para la base de datos.

## Estructura del Proyecto

## Instalación

1. Clona el repositorio:

   ```sh
   git clone <URL_DEL_REPOSITORIO>
   cd url_shortner
   ```

2. Instala las dependencias:

   ```sh
   go mod tidy
   ```

3. Configura la base de datos en si es necesario.

## Uso

1. Inicia el servidor:

   ```sh
   go run cmd/main.go
   ```

2. El servidor estará corriendo en .

## Endpoints

### POST /shorten

Acorta una URL.

- **URL:** `/shorten`
- **Método:** `POST`
- **Cuerpo de la solicitud:**
  ```json
  {
    "url": "https://example.com"
  }
  ```
- **Respuesta exitosa:**
  ```json
  {
    "Content-Type": "application/json",
    "data": {
      "short_url": "http://localhost:8080/r/abc123"
    }
  }
  ```

### GET /r

Redirige a la URL original.

- **URL:** `/r/{shortcode}`
- **Método:** `GET`
- **Respuesta exitosa:** Redirige a la URL original.

## Archivos Principales

- : Punto de entrada del servidor.
- : Configuración y migración de la base de datos.
- : Controladores para los endpoints.
- : Modelo para el cuerpo de la solicitud.
- : Modelo para la URL.
- : Utilidades para generar códigos cortos.
