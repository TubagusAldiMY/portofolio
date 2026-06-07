# Portofolio Full-Stack

Repository ini berisi website portofolio full-stack untuk menampilkan profil, proyek, produk/layanan, pengalaman, form kontak, chat berbasis AI, dan panel admin untuk mengelola konten.

## Ringkasan

- **Frontend utama:** SvelteKit, Svelte 5, TypeScript, Tailwind CSS, Three.js, GSAP.
- **Backend API:** Go, Gin, GORM, MySQL, JWT authentication, rate limiting, upload file.
- **Frontend legacy:** Vue 3 + Vite, disimpan di `frontend-vue-legacy/` sebagai versi lama.
- **Dokumentasi pendukung:** `Knowledge/`, `Journal/`, dan `Templates/`.

## Fitur

- Landing page portofolio responsif.
- Halaman proyek, detail proyek, produk/layanan, pengalaman, kontak, dan chat.
- API publik untuk projects, products, experiences, contact, dan chat AI.
- Panel admin dengan login JWT untuk CRUD konten dan melihat pesan kontak.
- Upload file statis dari backend melalui `/uploads`.
- Seed data awal untuk proyek dan produk.

## Struktur Project

```text
.
├── backend/              # REST API Go + Gin + GORM
│   ├── cmd/              # Entry point server dan seed command
│   ├── config/           # Koneksi database
│   ├── controllers/      # HTTP handlers
│   ├── middleware/       # Auth dan rate limiter
│   ├── models/           # Model database
│   └── services/         # Knowledge base service
├── frontend/             # Frontend utama SvelteKit
│   ├── src/lib/          # API client, auth, components, theme
│   └── src/routes/       # Halaman aplikasi
├── frontend-vue-legacy/  # Versi lama berbasis Vue
├── Knowledge/            # Catatan teknis dan referensi engineering
├── Journal/              # Template dan catatan kerja
└── Templates/            # Template ADR, RFC, sprint, bug report, dll.
```

## Prasyarat

- Go 1.24 atau lebih baru.
- Bun atau Node.js untuk frontend.
- MySQL 8 atau kompatibel.
- Gemini API key jika ingin mengaktifkan endpoint chat.

## Konfigurasi Environment

Backend:

```bash
cd backend
cp .env.example .env
```

Isi nilai berikut sesuai environment lokal:

```env
DB_USER=portfolio_user
DB_PASSWORD=change_me
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=portfolio_db
JWT_SECRET=change_me_to_a_64_char_hex_secret
GEMINI_API_KEY=your_gemini_api_key_here
FRONTEND_ORIGINS=http://localhost:5173
PORT=8000
```

Frontend:

```bash
cd frontend
cp .env.example .env
```

Pastikan `PUBLIC_API_URL` mengarah ke backend:

```env
PUBLIC_API_URL=http://localhost:8000
```

## Menjalankan Lokal

Siapkan database MySQL:

```sql
CREATE DATABASE portfolio_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'portfolio_user'@'localhost' IDENTIFIED BY 'change_me';
GRANT ALL PRIVILEGES ON portfolio_db.* TO 'portfolio_user'@'localhost';
FLUSH PRIVILEGES;
```

Jalankan backend:

```bash
cd backend
go mod download
go run ./cmd
```

Backend berjalan di `http://localhost:8000` dan menyediakan health check di `GET /health`.

Seed data awal:

```bash
cd backend
go run ./cmd/seed
```

Buat admin pertama:

```bash
cd backend
ADMIN_USERNAME=admin ADMIN_PASSWORD='change-this-password' go run ./cmd/seed-admin
```

Jalankan frontend utama:

```bash
cd frontend
bun install
bun run dev
```

Frontend berjalan di `http://localhost:5173`.

## Script Penting

Backend:

```bash
go run ./cmd              # menjalankan API server
go run ./cmd/seed         # seed projects dan products
go run ./cmd/seed-admin   # membuat user admin pertama
go test ./...             # menjalankan test Go
```

Frontend SvelteKit:

```bash
bun run dev       # development server
bun run build     # production build
bun run preview   # preview build lokal
bun run check     # type check SvelteKit
bun run lint      # lint dan prettier check
```

Frontend Vue legacy:

```bash
cd frontend-vue-legacy
npm install
npm run dev
```

## API Utama

```text
GET    /health
POST   /api/auth/login
POST   /api/contact
GET    /api/products
GET    /api/projects
GET    /api/projects/:id
GET    /api/experiences
POST   /api/chat
GET    /api/admin/messages
POST   /api/admin/upload
POST   /api/admin/projects
PUT    /api/admin/projects/:id
DELETE /api/admin/projects/:id
POST   /api/admin/products
PUT    /api/admin/products/:id
DELETE /api/admin/products/:id
POST   /api/admin/experiences
PUT    /api/admin/experiences/:id
DELETE /api/admin/experiences/:id
```

Endpoint `/api/admin/*` membutuhkan token JWT dari login admin.

## Build dan Deploy

Backend memiliki `Dockerfile` di `backend/` dan mengekspos port `8000`.

Frontend utama memakai SvelteKit adapter Vercel. Konfigurasi deploy tersedia di `frontend/vercel.json`.

Untuk production, pastikan:

- `APP_ENV=production`
- `JWT_SECRET` kuat dan tidak pernah commit ke Git.
- `FRONTEND_ORIGINS` hanya berisi domain frontend production.
- `PUBLIC_API_URL` mengarah ke domain API production.
- `ALLOW_REGISTER=false` setelah admin dibuat.

## Keamanan

- File `.env` dan secret lokal diabaikan oleh Git.
- Password admin disimpan menggunakan bcrypt.
- Endpoint admin diproteksi JWT.
- Endpoint login, contact, dan chat memiliki rate limit.
- Upload file dibatasi ukuran request dan disajikan dari direktori runtime.

## Lisensi

Project ini menggunakan lisensi Apache-2.0. Lihat `LICENSE` untuk detail.
