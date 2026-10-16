# Architecture

## Overview

layanan-service menggunakan struktur package Go yang memisahkan
komponen transport, business logic, domain, dan akses data.

## Directory Responsibilities

### cmd/api

Entry point aplikasi.

Bertanggung jawab untuk:
- Inisialisasi konfigurasi.
- Inisialisasi dependency.
- Menjalankan HTTP server.
- Mengelola lifecycle aplikasi.

### internal/api

Komponen yang berkaitan dengan penyusunan atau registrasi API.

Tanggung jawab detail mengikuti implementasi aktual.

### internal/web

Komponen web yang digunakan bersama oleh berbagai domain,
misalnya middleware, response helper, atau penanganan error HTTP,
jika memang ditempatkan di package ini.

### internal/layanan/domain

Representasi konsep bisnis dan aturan domain layanan.

### internal/layanan/repository

Akses data dan persistensi untuk domain layanan.

### internal/layanan/service

Business logic dan orkestrasi proses domain layanan.

### internal/layanan/web

HTTP handler atau komponen web khusus domain layanan.

### migrations

File migrasi skema database.

### remote

Komponen integrasi dengan sistem eksternal, jika itu fungsi
yang diterapkan dalam project ini.

### bin

Artefak atau executable hasil build, sesuai penggunaan direktori ini.

