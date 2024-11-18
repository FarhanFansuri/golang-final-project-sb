# <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> Golang - Backend Development Bootcamp (Batch 61) 

Selamat datang di repository saya! Ini adalah proyek akhir yang saya kerjakan selama mengikuti bootcamp **"Golang - Backend Development"** (Batch 61) yang diselenggarakan oleh [Sanbercode](https://sanbercode.com).

## Deskripsi

Bootcamp ini berfokus pada pengembangan backend dengan menggunakan bahasa pemrograman **Golang**. Materi yang dipelajari mencakup konsep-konsep dasar Golang, pengembangan API, pengelolaan database, praktik pengujian, serta penerapan prinsip-prinsip RESTful.

## Final Project (Restaurant API)

Restaurant API adalah layanan backend yang dirancang untuk membantu manajemen operasional restoran, seperti mengelola menu, pesanan, reservasi meja, dan pembayaran.
### link project
- PPT : [Canva](https://www.canva.com/design/DAGWzfrL1S8/qvmoevX80Tm09b5e_o_hFw/edit?utm_content=DAGWzfrL1S8&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)
- API : [Railway]()
- Youtube : [Presentasi]()

## Fitur Utama

- **Manajemen Menu**  
  - CRUD (Create, Read, Update, Delete) menu makanan dan minuman.  
  - Kategori menu (makanan pembuka, utama, penutup).  
  - Informasi harga dan ketersediaan.

- **Manajemen Pesanan**  
  - Membuat pesanan baru.  
  - Mengupdate status pesanan (diproses, selesai, dibatalkan).  
  - Melihat riwayat pesanan pelanggan.

- **Manajemen Meja**  
  - Melihat daftar meja yang tersedia.  
  - Reservasi meja oleh pelanggan.

- **Manajemen Pembayaran**  
  - Mendukung berbagai metode pembayaran (tunai/kartu/digital).  
  - Melihat status pembayaran pesanan.

- **Authentication & Authorization**  
  - Sistem login untuk admin restoran.  
  - Akses fitur sesuai peran pengguna (admin/pelanggan).

---
## struktur folder
berikut adalah struktur folder yang saya gunakan
```
restaurant-api/
├── controllers/ # Logika endpoint
├── models/ # Definisi model data
├── routes/ # Routing API
├── services/ # Logika bisnis
├── middleware/ # Middleware seperti autentikasi
├── database/ # Konfigurasi database
└── main.go # Entry point aplikasi
```
