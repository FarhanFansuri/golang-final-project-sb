# <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> Golang - Final Project Bootcamp SB (Batch 61) 

Selamat datang di repository saya! Ini adalah proyek akhir yang saya kerjakan selama mengikuti bootcamp **"Golang - Backend Development"** (Batch 61) yang diselenggarakan oleh [Sanbercode](https://sanbercode.com)

## Deskripsi

Bootcamp ini berfokus pada pengembangan backend dengan menggunakan bahasa pemrograman **Golang**. Materi yang dipelajari mencakup konsep-konsep dasar Golang, pengembangan API, pengelolaan database, praktik pengujian, serta penerapan prinsip-prinsip RESTful.

## Final Project (Dompetku API)

**Dompetku** adalah aplikasi API manajemen keuangan pribadi yang dirancang untuk membantu pengguna melacak dan mengelola pendapatan serta pengeluaran mereka dengan mudah dan efisien. Dengan Dompetku, pengguna dapat memperoleh wawasan tentang kebiasaan finansial mereka, mengatur anggaran, dan membuat keputusan keuangan yang lebih baik.

### link project
- PPT : [Canva](https://www.canva.com/design/DAGW7EPTsCU/QY6SAmKUiqZta59tuwoRRg/edit?utm_content=DAGW7EPTsCU&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)
- API : [Railway](https://golang-final-project-sb-production.up.railway.app/)
- Youtube : [Presentasi](https://youtu.be/oNCPtTnBX6w)


## 📋 Fitur Utama

### 1. Manajemen Transaksi (Pendapatan & Pengeluaran)

Dompetku memungkinkan pengguna untuk mencatat setiap transaksi keuangan mereka, baik itu pendapatan maupun pengeluaran, dengan detail yang lengkap. Fitur ini mencakup:

- **Pencatatan Transaksi:** Pengguna dapat menambahkan transaksi baru dengan informasi seperti jumlah, jenis (pendapatan atau pengeluaran), kategori, deskripsi, dan tanggal.
- **Klasifikasi Kategori:** Transaksi dapat dikategorikan ke dalam berbagai kategori seperti makanan, transportasi, hiburan, dan lainnya untuk memudahkan analisis.
- **CRUD Transaksi:** Pengguna dapat melakukan operasi Create, Read, Update, dan Delete (CRUD) pada transaksi mereka, memastikan data keuangan selalu terbarui dan akurat.

**Contoh Payload untuk Menambahkan Transaksi:**
```json
{
    "amount": 100000,
    "type": "expense",
    "category": "food",
    "description": "Dinner at restaurant"
}
```

### 2. Keamanan Akun dan Data

Keamanan adalah prioritas utama dalam Dompetku. Aplikasi ini memastikan bahwa data pengguna terlindungi dengan standar keamanan yang tinggi melalui:

- **Autentikasi Berbasis JWT:** Menggunakan JSON Web Tokens (JWT) untuk autentikasi pengguna, memastikan bahwa hanya pengguna yang berwenang yang dapat mengakses data mereka.
- **Enkripsi Password:** Password pengguna disimpan dalam bentuk terenkripsi menggunakan bcrypt, menjaga kerahasiaan informasi login.
- **Akses Terbatas:** Setiap pengguna hanya dapat mengakses data mereka sendiri, mencegah akses tidak sah ke informasi finansial pribadi.

---
## struktur folder
berikut adalah struktur folder yang saya gunakan
```
restaurant-api/
├── controllers/ # Logika endpoint
├── models/ # Definisi model data
├── routes/ # Routing API
├── middleware/ # Middleware seperti autentikasi
└── main.go # Entry point aplikasi
```
