# Belajar Golang RESTful API - Category Service

Repository ini berisi proyek latihan implementasi RESTful API menggunakan **Golang** dengan menerapkan prinsip **Clean Architecture**. Proyek ini mencakup manajemen data `Category` lengkap dengan validasi data, manajemen basis data transaksional, penanganan error terpusat, pengujian unit (*unit testing*), serta pengamanan jalur menggunakan *middleware*.

## 🏗️ Struktur Arsitektur & Folder

Proyek ini dipisahkan menjadi beberapa layer (*decoupled*) menggunakan konsep *interface and implementation* untuk mempermudah pengujian dan pemeliharaan kode:

* **`app/`**: Pengaturan inti aplikasi seperti konfigurasi database dan inisialisasi router.
* **`controller/`**: Layer entry-point yang memproses HTTP Request dan mengembalikan HTTP Response.
* **`service/`**: Layer logika bisnis inti, tempat validasi data masukan dijalankan.
* **`repository/`**: Layer yang berkomunikasi langsung dengan MySQL Database menggunakan transaksi database (`sql.Tx`).
* **`model/`**:
    * `domain/`: Cetak biru objek database murni (Entity).
    * `web/`: Data Transfer Object (DTO) untuk request dan response JSON.
* **`middleware/`**: Menyaring *request* masuk (misalnya memeriksa kecocokan API Key sebelum diteruskan ke router).
* **`exception/`**: Penanganan eror terpusat (*Centralized Error Handler*) memanfaatkan fitur *recovery* dari *panic server*.
* **`helper/`**: Fungsi pembantu seperti konversi JSON dan otomatisasi database commit/rollback.

---

## 🛠️ Tech Stack & Libraries

* **Language**: Go (Golang)
* **Database**: MySQL
* **Libraries / Drivers**:
    * `julienschmidt/httprouter` - HTTP Router performa tinggi.
    * `go-sql-driver/mysql` - Driver MySQL untuk database/sql Go.
    * `go-playground/validator/v10` - Library validasi objek dan struct.
    * `stretchr/testify` - Libary pembantu proses Unit Testing (Assertion).

---

## 🔑 Fitur Utama yang Dipelajari

1.  **Dependency Inversion Principle**: Memisahkan kontrak fungsi (*Interface*) dengan blok kodenya (*Implementation*) pada layer Controller, Service, dan Repository untuk kemudahan *mocking* saat testing.
2.  **Database Transaction Management**: Pengamanan proses manipulasi data menggunakan mekanisme `defer helper.CommitOrRollback(tx)` guna menghindari data korup saat query eror di tengah jalan.
3.  **Strict Data Validation**: Menggunakan tag validator khusus seperti `required`, `gt` (greater than), dan `lt` (less than) untuk menyaring request bodu sebelum diproses database.
4.  **Auth Middleware**: Menyaring akses API secara global lewat validasi header `X-API-Key`.

---

## 🚦 Jalur API (API Endpoints)

Dokumentasi spesifikasi API ini dirancang menggunakan standar OpenAPI 3.0:

| Method | Endpoint | Keterangan | Proteksi Auth |
| :--- | :--- | :--- | :--- |
| **GET** | `/api/categories` | Mengambil semua data kategori | Ya (`X-API-Key`) |
| **POST** | `/api/categories` | Membuat kategori baru | Ya (`X-API-Key`) |
| **GET** | `/api/categories/{categoryId}` | Mengambil kategori spesifik berdasarkan ID | Ya (`X-API-Key`) |
| **PUT** | `/api/categories/{categoryId}` | Mengubah data kategori berdasarkan ID | Ya (`X-API-Key`) |
| **DELETE** | `/api/categories/{categoryId}` | Menghapus kategori berdasarkan ID | Ya (`X-API-Key`) |

*Catatan: Seluruh request wajib menyertakan header `X-API-Key: RAHASIA` dan `Content-Type: application/json`.*

---

## 🏃 Cara Menjalankan Proyek

### 1. Kloning Repository
```bash
git clone [https://github.com/abdul452/belajar-golang-restful-api.git](https://github.com/abdul452/belajar-golang-restful-api.git)
cd belajar-golang-restful-api
```
### 2. Atur Database
Buat database baru di MySQL bernama `belajar_golang_restful_api` (atau sesuaikan dengan konfigurasi di file `app/database.go` kamu). Jalankan migrasi tabel `category`:
```SQL
CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL
);
```

### 3. Unduh Dependencies & Jalankan Server
```Bash
go mod tidy
go run main.go
```

Server akan berjalan di port `http://localhost:3000`.

## 🧪 Cara Menjalankan Unit Test
Pengujian skenario sukses maupun gagal untuk fungsionalitas HTTP Handler telah diotomatisasi di dalam folder `test/`. Untuk mengeksekusi semua unit test, jalankan perintah berikut di terminal:

```Bash
go test -v ./test/...
```
