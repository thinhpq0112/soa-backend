# SOA Backend

A Golang-based backend for managing products, suppliers, and categories with structured APIs, including PDF generation and filtering.

```bash
git clone https://github.com/thinhpq0112/soa-backend.git
```

### Create a .env file in the root directory of the project, replace with your own values; or you can use the .env.example file

```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
```

or use Makefile
```bash
make create-env
```

### Run the following commands to start the project:

```bash
go run cmd/main.go
```

### SWAGGER UI
http://localhost:8080/api/swagger/index.html
