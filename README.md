# Series Tracker Backend

Backend del proyecto **Series Tracker**, desarrollado en **Go** y conectado a **PostgreSQL**.  
Expone una **API REST** que responde en **JSON** y permite gestionar series con operaciones CRUD, búsqueda, ordenamiento, paginación, documentación OpenAPI/Swagger y subida de imágenes.

---

## Descripción general

Este backend fue construido como un servidor HTTP independiente.  
No genera HTML ni conoce la interfaz del cliente. Su responsabilidad es:

- exponer endpoints REST
- validar datos en el servidor
- consultar y modificar la base de datos
- responder en formato JSON
- servir la documentación de la API
- manejar subida de imágenes

La arquitectura del proyecto está separada por responsabilidades, usando carpetas como:

- `cmd/server` → punto de entrada del servidor
- `internal/db` → conexión a PostgreSQL
- `internal/models` → estructuras de datos
- `internal/repository` → acceso a base de datos
- `internal/service` → lógica y validaciones
- `internal/handlers` → manejo de requests y responses
- `internal/middleware` → middleware como CORS
- `internal/utils` → helpers para respuestas JSON
- `docs` → Swagger UI
- `uploads` → imágenes subidas

---

## Enlaces

- **API publicada:** https://series-tracker-backend-kh31.onrender.com
- **Swagger UI publicada:** https://series-tracker-backend-kh31.onrender.com/docs/
- **Repositorio del frontend:** https://github.com/Martinho777/series-tracker-frontend 
- **Frontend publicado:** https://69eae19b24172729b3d9f709--timely-narwhal-27f5ff.netlify.app/ 

---

## Screenshot

![Swagger UI](./screenshots/swagger.png)
![JSON](./screenshots/json.png)

---

## Tecnologías utilizadas

- **Go**
- **PostgreSQL**
- **Neon** para la base de datos publicada
- **Render** para el deploy del backend
- **OpenAPI**
- **Swagger UI**

---

## Instalación y ejecución local

### 1. Clonar el repositorio

```bash
git clone https://github.com/Martinho777/series-tracker-backend
cd series-tracker-backend
```

### 2. Verificar que Go esté instalado

```bash
go version
```

Si el comando no funciona, instala Go y vuelve a intentarlo.

### 3. Instalar dependencias

Este proyecto usa el driver de PostgreSQL para Go.

```bash
go mod tidy
```

Si fuera necesario:

```bash
go get github.com/lib/pq
```

### 4. Configurar la variable de entorno

Este backend no guarda credenciales dentro del código.  
La conexión a la base de datos se lee desde una variable de entorno llamada:

```env
DATABASE_URL
```

Ejemplo:

```env
DATABASE_URL=postgresql://usuario:password@host/base_de_datos?sslmode=require
```

También se incluye un archivo `.env.example` con el formato esperado.

#### En PowerShell (Windows)

```powershell
$env:DATABASE_URL="postgresql://usuario:password@host/base_de_datos?sslmode=require"
```

### 5. Ejecutar el servidor

Desde la raíz del proyecto:

```bash
go run ./cmd/server/main.go
```

Si todo está bien, el servidor arrancará en:

```text
http://localhost:8080
```

---

## Documentación de la API

Una vez corriendo localmente, puedes abrir:

- **API base:** `http://localhost:8080/`
- **Lista de series:** `http://localhost:8080/series`
- **Swagger UI:** `http://localhost:8080/docs`
- **Spec OpenAPI YAML:** `http://localhost:8080/openapi.yaml`

---

## Variables de entorno

El proyecto usa la siguiente variable:

### `DATABASE_URL`
Cadena completa de conexión a PostgreSQL.

Ejemplo:

```env
DATABASE_URL=postgresql://usuario:password@host/base_de_datos?sslmode=require
```

### `PORT`
No es necesario definirla manualmente en local.  
Si no existe, el servidor usa `8080`.  
En producción, Render la asigna automáticamente.

---

## Estructura del proyecto

```text
series-tracker-backend/
│
├── cmd/
│   └── server/
│       └── main.go
├── docs/
│   └── index.html
├── internal/
│   ├── db/
│   │   └── postgres.go
│   ├── handlers/
│   │   └── series_handler.go
│   ├── middleware/
│   │   └── cors.go
│   ├── models/
│   │   ├── series.go
│   │   └── series_filters.go
│   ├── repository/
│   │   └── series_repository.go
│   ├── service/
│   │   └── series_service.go
│   └── utils/
│       └── response.go
├── uploads/
├── screenshots/
│   └── backend-swagger.png
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── openapi.yaml
└── README.md
```

---

## Endpoints principales

### `GET /series`
Obtiene la lista de series.

Soporta parámetros opcionales:

- `q` → búsqueda por título
- `sort` → campo de ordenamiento
- `order` → `asc` o `desc`
- `page` → número de página
- `limit` → cantidad de resultados por página

Ejemplo:

```text
/series?q=dark&sort=anio&order=desc&page=1&limit=4
```

### `GET /series/:id`
Obtiene una serie específica por ID.

### `POST /series`
Crea una nueva serie.

### `PUT /series/:id`
Actualiza una serie existente.

### `DELETE /series/:id`
Elimina una serie existente.

### `POST /upload`
Recibe un archivo de imagen, lo guarda en `uploads/` y devuelve una URL pública.

### `GET /docs`
Muestra Swagger UI.

### `GET /openapi.yaml`
Devuelve la especificación OpenAPI en YAML.

---

## Formato de datos

### Ejemplo de serie

```json
{
  "id": 1,
  "titulo": "Dark",
  "genero": "Ciencia ficción",
  "anio": 2017,
  "temporadas": 3,
  "imagen_url": "https://ejemplo.com/dark.jpg",
  "descripcion": "Serie alemana de misterio y viajes en el tiempo."
}
```

### Ejemplo de error

```json
{
  "error": "Serie no encontrada"
}
```

---

## Códigos HTTP utilizados

El backend usa códigos HTTP apropiados según la operación:

- `200 OK` → consulta o actualización exitosa
- `201 Created` → creación o upload exitoso
- `204 No Content` → eliminación exitosa
- `400 Bad Request` → datos inválidos
- `404 Not Found` → recurso no encontrado
- `405 Method Not Allowed` → método no permitido
- `500 Internal Server Error` → error interno del servidor

---

## CORS

El proyecto incluye configuración de **CORS** porque el frontend y el backend corren en orígenes distintos.

Se configuró para permitir, durante desarrollo e integración:

- métodos: `GET, POST, PUT, DELETE, OPTIONS`
- header: `Content-Type`
- origen: `*`

Esto permite que el frontend consuma la API mediante `fetch()` sin que el navegador bloquee las peticiones.

---

## Base de datos

La persistencia se realiza en **PostgreSQL**.

La tabla principal utilizada es `series`, con campos como:

- `id`
- `titulo`
- `genero`
- `anio`
- `temporadas`
- `imagen_url`
- `descripcion`
- `created_at`
- `updated_at`

Además, se aplicaron validaciones como:

- `anio >= 1900`
- `temporadas >= 1`

---

## Requisitos cumplidos

- Backend corriendo como servidor HTTP independiente
- El servidor responde JSON y no genera HTML
- API REST con endpoints CRUD completos:
  - `GET /series`
  - `GET /series/:id`
  - `POST /series`
  - `PUT /series/:id`
  - `DELETE /series/:id`
- Persistencia en una base de datos real
- Soporte para imágenes
- Separación clara entre cliente y servidor
- CORS configurado
- Repositorio separado para backend
- Proyecto publicado en internet y funcionando
- README con instrucciones claras, screenshot, challenges y reflexión

---

## Challenges implementados

- Spec de **OpenAPI/Swagger** escrita y precisa
- **Swagger UI** corriendo y servido desde el backend
- **Códigos HTTP correctos** en la API
- **Validación server-side** con respuestas descriptivas en JSON
- **Paginación** con `?page=` y `?limit=`
- **Búsqueda por nombre** con `?q=`
- **Ordenamiento** con `?sort=` y `?order=`
- **Subida de imágenes**
- Organización del código por responsabilidades
- Historial de commits progresivo

---

## Pruebas manuales recomendadas

Para probar rápidamente la API puedes abrir:

- `GET /series`
- `GET /series/1`
- `GET /series/999`
- `GET /series/hola`
- `GET /docs`

También se recomienda probar:

- `POST /series` con JSON válido
- `POST /series` con datos inválidos
- `PUT /series/:id`
- `DELETE /series/:id`
- `POST /upload` con imagen válida
- `POST /upload` con archivo inválido

---

## Deploy

El backend fue publicado en **Render** y se conecta a una base PostgreSQL publicada en **Neon**.

Para el deploy se realizaron estos ajustes:

- uso de `DATABASE_URL`
- lectura del puerto desde la variable `PORT`
- publicación de Swagger UI
- publicación del archivo `openapi.yaml`

---

## Reflexión

Para este proyecto elegí Go porque quería trabajar un backend donde el manejo de rutas, respuestas HTTP y conexión a base de datos fuera explícito y fácil de entender. En lugar de concentrar toda la lógica en un solo archivo, preferí separar el proyecto en capas como `handlers`, `service`, `repository`, `models`, `middleware` y `db`, lo cual hizo el código más ordenado y más cercano a una aplicación real. Gracias a esa estructura fue más sencillo ir agregando funcionalidades como búsqueda, ordenamiento, paginación, documentación Swagger y subida de imágenes sin desorganizar el proyecto.

También considero que usar PostgreSQL, Neon y Render fue una buena decisión porque permitió trabajar un flujo mucho más realista que uno puramente local. Aprendí no solo a construir la API, sino también a prepararla para producción usando variables de entorno, deploy público y documentación formal con OpenAPI. Si volviera a hacer un proyecto similar, sí volvería a usar este stack, especialmente cuando quiera priorizar claridad, control del backend y una separación limpia entre cliente y servidor.
