# Portofolio Pribadi (Vue.js & Tailwind CSS)

Ini adalah proyek website portofolio pribadi yang dibangun untuk menampilkan profil, proyek, pengalaman, dan cara untuk terhubung. Proyek ini dibuat dengan teknologi web modern yang berfokus pada kecepatan, kemudahan pengembangan, dan desain yang responsif.

## 📜 Tentang Proyek
Website ini berfungsi sebagai CV digital interaktif. Tujuannya adalah untuk memberikan gambaran lengkap tentang kemampuan teknis dan perjalanan karier kepada calon perekrut atau klien.

### Teknologi Utama yang Digunakan:
- **Vue.js 3**: Kerangka kerja JavaScript progresif untuk membangun antarmuka pengguna (UI).
- **Vite**: Build tool modern yang memberikan pengalaman pengembangan yang sangat cepat.
- **Tailwind CSS**: Kerangka kerja CSS utility-first untuk membuat desain kustom dengan cepat tanpa meninggalkan HTML.

## ✨ Fitur Utama
- **Desain Responsif**: Tampilan yang optimal di berbagai perangkat, mulai dari desktop hingga mobile.
- **Komponen Terstruktur**: Setiap bagian website (Navbar, Hero, About, dll.) dipecah menjadi komponen Vue yang terpisah untuk kemudahan pengelolaan.
- **Single Page Application (SPA)**: Navigasi yang mulus dan cepat tanpa perlu me-reload halaman.
- **Formulir Kontak Fungsional**: Terintegrasi dengan layanan Formspree untuk menerima pesan langsung ke email.
- **Animasi Halus**: Animasi transisi saat menggulir halaman untuk pengalaman pengguna yang lebih menarik.

## 🚀 Instalasi & Menjalankan Proyek
Berikut adalah langkah-langkah untuk menjalankan proyek ini di komputermu.

### Prasyarat:
- **Node.js** (versi 18 atau lebih tinggi direkomendasikan)
- **NPM** atau **Yarn** (terinstal bersama Node.js)

### Langkah-langkah:
1. Clone repositori ini:
   ```bash
   git clone https://github.com/TubagusAldiMY/my-portofolio.git
   ```

2. Masuk ke direktori proyek:
   ```bash
   cd my-portofolio/portofolio-vuejs
   ```

3. Instal semua dependensi:
   ```bash
   npm install
   ```

4. Jalankan server pengembangan:
   ```bash
   npm run dev
   ```
   Buka browser dan akses `http://localhost:5173` (atau port lain yang ditampilkan di terminal).

5. Build untuk produksi:
   ```bash
   npm run build
   ```
   Hasilnya akan tersedia di dalam folder `dist/`.

## 📁 Struktur Proyek
Berikut adalah penjelasan singkat tentang file dan folder penting dalam proyek ini:

```
/portofolio-vuejs
├── public/                # Aset statis (gambar, CV.pdf, favicon, dll.)
├── src/
│   ├── assets/            # Aset yang diproses oleh Vite (gambar komponen)
│   ├── components/        # Komponen-komponen Vue
│   │   ├── AboutSection.vue
│   │   ├── ContactSection.vue
│   │   ├── ExperienceSection.vue
│   │   ├── Footer.vue
│   │   ├── HeroSection.vue
│   │   ├── Navbar.vue
│   │   └── ProjectsSection.vue
│   ├── directives/        # (Opsional) Custom directive seperti scroll-animation
│   ├── App.vue            # Komponen root aplikasi
│   ├── main.js            # Titik masuk aplikasi
│   └── style.css          # File CSS utama (impor Tailwind)
├── .gitignore             # File/folder yang diabaikan oleh Git
├── index.html             # Template HTML utama
├── package.json           # Daftar dependensi dan skrip proyek
└── vite.config.js         # File konfigurasi untuk Vite
```

## 🧩 Penjelasan Komponen (src/components/)
- **Navbar.vue**: Mengelola bilah navigasi atas, termasuk link-link halaman dan menu hamburger untuk tampilan mobile.
- **HeroSection.vue**: Bagian pembuka website yang berisi nama, jabatan, deskripsi singkat, dan tombol call-to-action.
- **AboutSection.vue**: Berisi foto, biografi singkat, dan daftar teknologi yang dikuasai.
- **ProjectsSection.vue**: Menampilkan galeri proyek dalam bentuk kartu (card). Data proyek dikelola di dalam komponen ini.
- **ExperienceSection.vue**: Menampilkan riwayat pengalaman kerja atau studi dalam format linimasa (timeline).
- **ContactSection.vue**: Berisi informasi kontak (email, LinkedIn, GitHub) dan formulir kontak.
- **Footer.vue**: Bagian kaki website yang berisi link sosial media dan informasi hak cipta.

## 🔧 Kustomisasi
Konten website ini dirancang agar mudah diubah. Berikut adalah cara mengubah data utama:

- **Mengubah Info Pribadi**: Buka `HeroSection.vue` dan `AboutSection.vue` untuk mengubah teks deskripsi dan bio.
- **Mengubah Daftar Proyek**: Buka `src/components/ProjectsSection.vue` dan ubah array `projects` di dalam bagian `<script setup>`.
- **Mengubah Pengalaman**: Buka `src/components/ExperienceSection.vue` dan ubah array `experiences`.
- **Mengubah Link Sosial Media**: Buka `src/components/ContactSection.vue` dan `src/components/Footer.vue` untuk memperbarui array `socialLinks`.
- **Mengganti Endpoint Form**: Jangan lupa mengganti URL `action` di dalam tag `<form>` pada `ContactSection.vue` dengan URL Formspree milikmu.
- **Mengganti CV**: Ganti file CV di folder `public/` dan pastikan nama filenya sesuai dengan yang ada di link pada `HeroSection.vue`.
