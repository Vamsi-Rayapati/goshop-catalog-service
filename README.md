# 🛍️ GoShop Catalog Service

A microservice for managing product catalogs in the **GoShop e-commerce platform**. This service exposes RESTful APIs to handle **categories**, **products**, and **product images**, with AWS S3 integration for image storage.


## ✨ Features

- 📦 **Product Management**  
  Full CRUD operations for products.

- 🏷️ **Category Management**  
  Organize products into categories for better discoverability.

- 🖼️ **Image Management**  
  Upload, update, and delete product images using **AWS S3** via pre-signed URLs.

- 🔐 **JWT Authentication**  
  Secure access to APIs using JSON Web Tokens.

- 🗄️ **Database Integration**  
  Uses **MySQL** with **GORM** for relational data persistence.

- 🚀 **High Performance**  
  Built with the **Gin** web framework for fast and lightweight HTTP handling.

- 🐳 **Docker Support**  
  Containerized for consistent deployment across environments.

## 🚀 Getting Started
1. Update .env with AWS IAM user credentials
```bash
AWS_API_KEY="XXX"
AWS_SECRET="XXXXXXX"
```
2. Install Dependencies
```bash
go mod tidy
```
3. Run the Service
```bash
go run main.go
```

## 🧰 Tech Stack

| Component         | Technology      |
|------------------|-----------------|
| Language          | Go 1.23.5        |
| Web Framework     | Gin              |
| ORM               | GORM             |
| Database          | MySQL            |
| Cloud Storage     | AWS S3 (Pre-signed URLs) |
| Authentication    | JWT              |
| Containerization  | Docker           |


## ✨ Author
Made with ❤️ by Vamsi Rayapati
