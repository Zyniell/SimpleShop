# 🛒 SimpleShop API

API untuk aplikasi e-commerce sederhana (SimpleShop) yang dibangun menggunakan **Go (Golang)** dengan rancangan arsitektur yang bersih (*clean architecture* / *layered architecture*). 

## 🚀 Fitur Utama

- 🔐 **Autentikasi (JWT)**: Registrasi dan Login pengguna dengan token JWT yang aman.
- 📂 **Manajemen Kategori**: Endpoint CRUD standar untuk mengelola kategori produk (dilindungi *middleware* autentikasi untuk memodifikasi data, terbuka untuk membaca secara masif).
- 🛍️ **Manajemen Produk**: Endpoint CRUD untuk daftar produk beserta detailnya (terintegrasi dengan relasi kategori).
- 📦 **Pesanan (*Orders*)**: Fungsionalitas *checkout* belanja dan melihat riwayat pesanan eksklusif (*My Orders*) yang spesifik untuk tiap *user*.

## 🛠️ Stack Teknologi

- **[Go (Golang) 1.25.0](https://go.dev/)**: Bahasa pemrograman utama yang sangat tangguh, efisien, dan cepat.
- **[Gin Web Framework](https://gin-gonic.com/)**: *Router* micro-framework utama penanganan HTTP/API tingkat lanjut.
- **[PostgreSQL](https://www.postgresql.org/)**: Sistem manajemen *database* relasional SQL handal (dengan *driver* murni `github.com/lib/pq`).
- **Data Validator**: `go-playground/validator/v10` untuk validasi keamanan masukan tipe JSON (*request body payload*).
- **Konfigurasi Lingkungan**: Pengelolaan rahasia *environment* menggunakan `joho/godotenv`.
- **Autentikasi**: *JSON Web Tokens* standar keamanan tinggi via `golang-jwt/jwt/v5`.
- **Migrasi**: `rubenv/sql-migrate` untuk kontrol versi skema dan relasi pangkalan data.
