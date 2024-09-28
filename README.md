# BasicTrade API

BasicTrade API is a backend service where an admin (seller) can create products, upload product images, and manage product variants. This application supports full authentication, CRUD operations, and file/image uploads using Cloudinary.

## Features

- **Authentication**: Register and Login functionalities for admin users.
- **CRUD Operations**:
  - Admins can create, update, delete, and view products.
  - Admins can create, update, delete, and view product variants.
- **File/Image Upload**: Admins can upload product images to Cloudinary with support for JPG, JPEG, PNG, and SVG formats.
- **Authorization**: All endpoints that modify data (Create, Update, Delete) are secured with JWT authorization.
- **Search**: Implemented search by product name and variant name.
- **Pagination**: Both product and variant listing endpoints support pagination for better data management.
- **UUID Support**: The project uses UUIDs for all records to ensure unique and secure identifiers across the system.

## Requirements

To run and deploy this project, you will need:

- **Go**: The primary programming language used to build this backend services.
- **MySQL**: The relational database used for storing application data. The project interacts directly with MySQL without using an ORM (Object-Relational Mapping).
- **Cloudinary**: A cloud-based service for managing and hosting images, used to store product images uploaded by admins.
- **Railway**: A cloud deployment platform used for hosting and deploying the application.

## Libraries and Frameworks Used

The following libraries and frameworks are used in this project:

- [**Gin Gonic**](github.com/gin-gonic/gin): A high-performance HTTP web framework in Go.
- [**Golang JWT**](github.com/golang-jwt/jwt/v5): A Go library used for implementing JSON Web Tokens (JWT) for authentication and authorization.
- [**Go Validator**](github.com/go-playground/validator/v10): A library for struct validation, used to validate input fields.
- [**Google UUID**](github.com/google/uuid): A library to generate universally unique identifiers (UUID) for records.
- [**GoDotEnv**](github.com/joho/godotenv): A library for managing environment variables via a `.env` file.
- [**Cloudinary Go**](github.com/cloudinary/cloudinary-go/v2): A library for uploading and managing media files using Cloudinary.
- [**MySQL Driver**](github.com/go-sql-driver/mysql): A MySQL driver for Go, used for database operations.
- [**Golang Crypto**](golang.org/x/crypto/bcrypt): A collection of cryptography packages in Go, used for password hashing and encryption.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/alifmaulidanar/go-final-project.git
cd go-final-project
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up your environment variables by creating a `.env` file in the project root and adding the following:

```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=

JWT_SECRET=

CLOUDINARY_CLOUD_NAME=
CLOUDINARY_API_KEY=
CLOUDINARY_API_SECRET=
CLOUDINARY_URL=
```

4. Create your database tables for `admins`, `products`, and `variants`.
   - You can use the provided MySQL dump file located in the `/database` directory to quickly set up the necessary tables.
   - Simply import the SQL dump into your MySQL database using the following command:

```bash
mysql -u [username] -p [database_name] < /database/go-final-project-dump.sql
```

## Running the Project

You can run the server manually or using `air` for live-reloading.

### Run the Server Manually

```bash
go run main.go
```

### Run the Server with Air (Hot Reload)

Install Air if you haven't:

```bash
go install github.com/air-verse/air@latest
```

Then, run the server with Air:

```bash
air
```

## Endpoints

| Method | URL                                                | Description         | Auth |
| ------ | -------------------------------------------------- | ------------------- | ---- |
| POST   | {{url}}/auth/register                              | Register Admin      | ❌   |
| POST   | {{url}}/auth/login                                 | Login Admin         | ❌   |
| GET    | {{url}}/products?limit=10&offset=0&search=         | Get All Products    | ❌   |
| POST   | {{url}}/products                                   | Create Product      | ✅   |
| PUT    | {{url}}/products/:productUUID                      | Update Product      | ✅   |
| DELETE | {{url}}/products/:productUUID                      | Delete Product      | ✅   |
| GET    | {{url}}/products/:productUUID                      | Get Product by UUID | ❌   |
| GET    | {{url}}/products/variant?limit=10&offset=0&search= | Get All Variants    | ❌   |
| POST   | {{url}}/products/variants                          | Create Variant      | ✅   |
| PUT    | {{url}}/products/variants/:variantUUID             | Update Variant      | ✅   |
| DELETE | {{url}}/products/variants/:variantUUID             | Delete Variant      | ✅   |
| GET    | {{url}}/products/variant/:variantUUID              | Get Variant by UUID | ❌   |

## Deployment / Live Demo API

The project has been successfully deployed on Railway and can be accessed through the following link:

[Live Demo API](https://your-railway-url.com)

Feel free to test the API endpoints using the live demo.

## Postman Collection

You can find the Postman collection to test the API in the repository under `Base Trade API - Alif Maulidanar.json`.
