# 🧠 Bank Soal Backend

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/github/go-mod/go-version/aryarobyy/bank-soal)
![Gin](https://img.shields.io/badge/Gin-Framework-blue)
![Status](https://img.shields.io/badge/Status-Active-success)

Backend API untuk aplikasi **Bank Soal**, dibuat menggunakan **Golang (Gin + GORM)** dan **MySQL**.  
Proyek ini dirancang untuk mendukung sistem manajemen soal dan ujian, dengan fitur autentikasi, manajemen user, serta middleware keamanan.

---

## ⚙️ Tech Stack

- **Language:** Go 1.22+
- **Framework:** Gin Gonic
- **ORM:** GORM
- **Database:** MySQL (atau PostgreSQL)
- **Auth:** JWT (Access & Refresh Token)
- **Other:** dotenv, rate limiting, middleware validation

---

# 🚀 Setup Project

## 1. Download & Install Go

**Windows/Mac:**
- Download dari: https://go.dev/dl/
- Install seperti biasa
- Verifikasi: `go version`

**Linux:**
```bash
sudo apt install golang-go
```

---

## 2. Clone Project

```bash
git clone https://github.com/aryarobyy/bank-soal
cd bank-soal/backend
```

---

## 3. Install Dependencies

```bash
go mod download
```

atau

```bash
go mod tidy
```

---

## 4. Setup Environment

### Buat file `.env`

```env
SERVER_PORT=8080
JWT_SECRET=
JWT_EXPIRED=

DB_USERNAME=root
DB_PASSWORD=
DB_NAME=
DB_HOST=
DB_PORT=

DB_MAX_IDLE_CONNS=
DB_MAX_OPEN_CONNS=
DB_CONN_MAX_LIFETIME=
```

---

## 5. Run Application

```bash
go run main.go
```

---

## 🎯 Done!

Server berjalan di: `http://localhost:8080`

---

## 📝 Common Commands

```bash
go mod tidy          # Install/update dependencies
go run main.go       # Run aplikasi
```

## 📜 License
This project is licensed under the **MIT License** — see the [LICENSE](./LICENSE) file for details.

© 2025 Roby Aryanata