# AGENTS.md

## Project Overview

`layanan-service` adalah backend service berbasis Go.

Dokumen ini menjadi panduan utama bagi AI coding agent ketika
menganalisis, membuat, memperbaiki, atau melakukan refactoring kode
di repository ini.

Sebelum mengubah kode, pahami implementasi yang sudah ada dan
ikuti konvensi project. Jangan mengasumsikan implementasi hanya
berdasarkan nama direktori.

## Technology Stack

- Language: Go
- Dependency management: Go Modules
- Application entry point: `cmd/api`
- Database migrations: `migrations`
- Build and development commands: `Makefile`
- Containerization: `Dockerfile`

Periksa `go.mod`, `Makefile`, dan konfigurasi yang tersedia untuk
mengetahui versi serta teknologi yang benar-benar digunakan.

Jangan mengasumsikan database, HTTP router, ORM, atau library tertentu
sebelum memeriksa implementasi aktual.

## Project Structure

### `cmd/api`

Entry point aplikasi.

Periksa bagaimana aplikasi menginisialisasi konfigurasi,
dependency, HTTP server, dan koneksi database.

### `internal/api`

Package untuk komponen API yang ditempatkan di sini.

Periksa implementasi aktual sebelum menentukan tanggung jawab
dan hubungan dependency package ini dengan package lainnya.

### `internal/layanan`

Implementasi domain bisnis layanan.

#### `internal/layanan/domain`

Tempat konsep domain dan aturan bisnis yang berkaitan dengan layanan.

Utamakan agar domain tidak bergantung pada detail HTTP,
database, atau framework tertentu.

#### `internal/layanan/repository`

Tempat implementasi akses data dan persistensi domain layanan.

Hindari menempatkan business workflow di repository.

#### `internal/layanan/service`

Tempat business logic dan orkestrasi proses domain layanan.

Service dapat berkoordinasi dengan repository dan dependency
lainnya sesuai kebutuhan bisnis.

#### `internal/layanan/web`

Tempat komponen HTTP yang khusus menangani domain layanan,
sesuai implementasi aktual.

### `internal/web`

Tempat komponen web yang digunakan bersama oleh berbagai bagian
aplikasi, jika sesuai dengan implementasi aktual.

Contohnya dapat berupa middleware, helper HTTP, atau penanganan
error yang bersifat umum.

Jangan menduplikasi komponen yang sebenarnya sudah tersedia di sini.

### `migrations`

Tempat file migrasi database.

Periksa mekanisme migrasi yang digunakan sebelum menambahkan
atau mengubah skema database.

### `remote`

Tempat komponen integrasi eksternal jika itu merupakan tanggung
jawab package ini.

Periksa implementasi aktual sebelum menambahkan integrasi baru.

### `docs`

Tempat dokumentasi teknis project.

### `bin`

Tempat executable atau artefak build sesuai konvensi project.

### `Makefile`

Sumber perintah standar untuk build, test, menjalankan aplikasi,
dan pekerjaan pengembangan lainnya.

Periksa target yang tersedia sebelum menjalankan atau
mendokumentasikan perintah.

## Architecture Principles

### 1. Follow Existing Architecture

- Ikuti struktur package yang sudah ada.
- Jangan membuat layer atau abstraksi baru tanpa kebutuhan yang jelas.
- Jangan memindahkan kode antar-package tanpa alasan yang kuat.
- Periksa implementasi serupa sebelum membuat komponen baru.
- Utamakan perubahan yang konsisten dengan pola yang sudah digunakan.

### 2. Separation of Concerns

Pertahankan pemisahan tanggung jawab antara:

- Domain: konsep dan aturan bisnis.
- Repository: akses data dan persistensi.
- Service: business logic dan orkestrasi proses.
- Web: komunikasi HTTP dan pemetaan request/response.

Sesuaikan batas tanggung jawab dengan desain aktual project.
Jangan memaksakan abstraksi yang tidak diperlukan.

### 3. Domain and DTO

Pisahkan model domain dari struktur request dan response HTTP
apabila keduanya memiliki kebutuhan atau siklus perubahan berbeda.

Jangan menggunakan DTO sebagai pengganti domain secara otomatis.

Hindari membuat struktur baru jika struktur yang sudah ada
masih sesuai dan dapat digunakan kembali.

### 4. Dependency Management

- Ikuti pola dependency injection yang digunakan project.
- Hindari global state yang tidak diperlukan.
- Hindari circular dependencies.
- Jangan menambahkan dependency eksternal jika standard library
  atau dependency yang tersedia sudah memadai.

### 5. Error Handling

- Tangani error secara eksplisit.
- Jangan mengabaikan error tanpa alasan yang jelas.
- Ikuti konvensi error yang sudah digunakan project.
- Pisahkan error internal dari detail yang aman dikirim ke client.
- Jangan mengubah format error API secara tidak konsisten.

### 6. Database and Transactions

- Ikuti pola akses database yang sudah digunakan.
- Gunakan parameterized queries jika menulis SQL.
- Kelola transaksi sesuai kebutuhan konsistensi data.
- Hindari perubahan skema database tanpa migrasi yang sesuai.
- Jangan mengubah skema atau data produksi secara langsung.

### 7. HTTP API

- Ikuti pola routing, middleware, request validation,
  response, dan status code yang sudah digunakan.
- Pertahankan kompatibilitas API kecuali perubahan kontrak
  memang diperlukan.
- Jangan mengekspos detail internal atau informasi sensitif
  melalui response HTTP.

## Coding Conventions

- Ikuti idiom dan konvensi bahasa Go.
- Gunakan `gofmt` untuk memformat kode.
- Gunakan nama package, tipe, fungsi, dan variabel yang jelas.
- Utamakan implementasi sederhana dan mudah dipelihara.
- Hindari duplikasi business logic.
- Gunakan `context.Context` untuk operasi yang membutuhkan
  pembatalan atau deadline.
- Hindari perubahan kode yang tidak berkaitan dengan tugas.
- Jangan menambahkan komentar yang hanya mengulang kode.

## Testing

Sebelum mengubah implementasi:

1. Periksa test yang sudah tersedia.
2. Identifikasi perilaku yang harus dipertahankan.
3. Tambahkan atau perbarui test untuk perubahan perilaku.
4. Jalankan test yang relevan.

Gunakan perintah yang sesuai dengan project, misalnya:

    go test ./...

Jika test gagal, periksa apakah kegagalan disebabkan oleh
perubahan yang dibuat atau masalah yang sudah ada sebelumnya.

Jangan menyatakan test berhasil jika belum menjalankannya.

## Development Workflow

Sebelum melakukan perubahan:

1. Baca instruksi repository dan dokumentasi yang relevan.
2. Periksa implementasi package terkait.
3. Telusuri alur dari entry point atau HTTP handler menuju service,
   repository, dan database jika relevan.
4. Identifikasi dependency dan dampak perubahan.
5. Buat perubahan sekecil mungkin untuk memenuhi kebutuhan.

Setelah melakukan perubahan:

1. Format kode yang berubah.
2. Jalankan test yang relevan.
3. Periksa perubahan untuk menemukan regresi yang mungkin terjadi.
4. Perbarui dokumentasi jika perilaku, konfigurasi, atau arsitektur
   berubah.

## Restrictions

- Jangan mengarang fungsi, package, tabel, atau endpoint yang
  belum diverifikasi keberadaannya.
- Jangan mengganti library atau framework tanpa alasan yang jelas.
- Jangan mengubah kontrak API tanpa mempertimbangkan kompatibilitas.
- Jangan menghapus kode yang tampak tidak digunakan tanpa
  memeriksa kemungkinan penggunaannya.
- Jangan menambahkan dependency hanya untuk menyelesaikan
  masalah sederhana.
- Jangan memasukkan credential, token, atau secret ke repository.
- Jangan menjalankan operasi destruktif terhadap database
  atau lingkungan bersama tanpa otorisasi yang sesuai.

## Important Rule

Implementasi kode adalah sumber kebenaran untuk perilaku aplikasi.

Dokumentasi membantu memahami desain, tetapi jika dokumentasi
dan implementasi berbeda, identifikasi perbedaannya terlebih
dahulu dan jangan langsung mengasumsikan salah satunya benar.

Jika informasi penting belum jelas, telusuri kode yang relevan
sebelum mengambil keputusan.

