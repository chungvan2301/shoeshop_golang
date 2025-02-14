# ShoesShop Backend API

## Overview
The **ShoesShop Backend API** is a RESTful API built using **Gin (Golang)** for managing an online shoe store. It provides authentication, product management, and order processing with **MongoDB** as the database and **Cloudinary** for image storage.

## Tech Stack
- **Gin** – Web framework for Golang
- **MongoDB** – NoSQL database
- **Cloudinary** – Image storage and management
- **JWT** – Authentication and authorization

## Features

### 1. **User Authentication & Authorization**
- Google OAuth and local authentication
- JWT-based authentication
- Role-based access control (User, Admin)

### 2. **Product Management**
- Create, update, and delete products (Admin only)
- Upload and manage product images via **Cloudinary**

## Installation & Setup

### Prerequisites
Ensure you have the following installed:
- **Go 1.19+**
- **MongoDB**

### Clone the Repository
```sh
git clone https://github.com/yourusername/shoesshop.git
cd shoesshop
```


