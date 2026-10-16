# Contributing Guide

## Development Principles

- Utamakan kode yang sederhana dan mudah dipelihara.
- Pisahkan business logic dari HTTP handling.
- Pisahkan domain dari request dan response DTO.
- Hindari abstraksi yang tidak diperlukan.
- Tangani error secara eksplisit.

## Code Style

- Ikuti format kode Go menggunakan gofmt.
- Gunakan penamaan yang mengikuti konvensi Go.
- Hindari duplikasi business logic.
- Gunakan context.Context untuk operasi yang memerlukan pembatalan atau deadline.

## Architecture

- Handler menangani komunikasi HTTP.
- Service menangani business logic.
- Repository menangani akses data.
- Domain merepresentasikan konsep bisnis.

## Testing

- Tambahkan unit test untuk business logic.
- Tambahkan integration test untuk akses database jika diperlukan.
- Pastikan test yang relevan lulus sebelum perubahan digabungkan.

## Git Workflow

Jelaskan aturan branch, commit, dan pull request.

## Pull Request Checklist

- [ ] Kode sudah diformat.
- [ ] Test sudah dijalankan.
- [ ] Tidak ada credential yang ikut di-commit.
- [ ] Dokumentasi diperbarui jika diperlukan.

