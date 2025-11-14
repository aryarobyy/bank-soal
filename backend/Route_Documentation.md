# Dokumentasi Route Aplikasi Bank Soal

Dokumentasi ini menjelaskan semua route yang tersedia dalam aplikasi bank soal ini. Setiap route memiliki penjelasan terkait fungsinya, metode HTTP yang digunakan, parameter yang dibutuhkan, dan siapa yang dapat mengaksesnya.

## Daftar Isi
- [Route User](#route-user)
- [Route Exam](#route-exam)
- [Route Subject](#route-subject)
- [Route Question](#route-question)
- [Route Option](#route-option)
- [Route Exam Question](#route-exam-question)
- [Route Exam Session](#route-exam-session)
- [Route Exam Score](#route-exam-score)
- [Route User Answer](#route-user-answer)
- [Route Xls Path](#route-xls-path)

## Route User

Route untuk mengelola data pengguna dalam sistem.

### POST /user/register
- **Deskripsi**: Mendaftarkan pengguna baru ke dalam sistem
- **Metode HTTP**: POST
- **Akses**: Umum (tidak memerlukan otentikasi)
- **Parameter Request (JSON)**:
  - `name` (string, wajib): Nama pengguna
  - `email` (string, wajib): Email pengguna
  - `password` (string, wajib): Password pengguna
  - `major` (string, wajib): Jurusan pengguna
  - `faculty` (string, wajib): Fakultas pengguna
- **Validasi**: Semua field wajib diisi, email harus dalam format yang valid
- **Contoh Request**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "major": "Informatika",
  "faculty": "Ilmu Komputer"
}
```

### POST /user/login
- **Deskripsi**: Login pengguna ke dalam sistem
- **Metode HTTP**: POST
- **Akses**: Umum (tidak memerlukan otentikasi)
- **Parameter Request (JSON)**:
  - `email` (string, wajib): Email pengguna
  - `password` (string, wajib): Password pengguna
- **Validasi**: Email harus dalam format yang valid
- **Contoh Request**:
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

### POST /user/refresh
- **Deskripsi**: Memperbarui access token menggunakan refresh token
- **Metode HTTP**: POST
- **Akses**: Umum (mengambil refresh token dari cookie)
- **Parameter**: Refresh token dari cookie

### GET /user/id
- **Deskripsi**: Mendapatkan data pengguna berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID pengguna

### GET /user/email
- **Deskripsi**: Mendapatkan data pengguna berdasarkan email
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `email` (string, wajib): Email pengguna

### GET /user/nim
- **Deskripsi**: Mendapatkan data pengguna berdasarkan NIM
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `nim` (string, wajib): NIM pengguna

### GET /user/name
- **Deskripsi**: Mendapatkan data pengguna berdasarkan nama
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `name` (string, opsional): Nama pengguna
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /user/role
- **Deskripsi**: Mendapatkan data pengguna berdasarkan peran
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `role` (string, wajib): Peran pengguna ("admin", "user", "lecturer")
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /user/
- **Deskripsi**: Mendapatkan data banyak pengguna
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### PUT /user/:id
- **Deskripsi**: Mengupdate data pengguna
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID pengguna yang akan diperbarui
- **Parameter Request (Form)**:
  - `name` (string, opsional): Nama pengguna
  - `email` (string, opsional): Email pengguna
  - `nim` (string, opsional): NIM pengguna
  - `nip` (string, opsional): NIP pengguna
  - `nidn` (string, opsional): NIDN pengguna
  - `major` (string, opsional): Jurusan pengguna
  - `faculty` (string, opsional): Fakultas pengguna
  - `status` (string, opsional): Status pengguna
  - `academic_year` (string, opsional): Tahun akademik
  - `image` (file, opsional): Gambar profil pengguna

### PUT /user/password
- **Deskripsi**: Mengganti password pengguna
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin)
- **Parameter Query**:
  - `id` (integer, wajib): ID pengguna
- **Parameter Request (JSON)**:
  - `new_password` (string, wajib): Password baru

### DELETE /user/:id
- **Deskripsi**: Menghapus pengguna
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID pengguna yang akan dihapus

### PUT /user/role
- **Deskripsi**: Mengubah peran pengguna
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Super Admin)
- **Parameter Query**:
  - `id` (integer, wajib): ID pengguna
- **Parameter Request (JSON)**:
  - `role` (string, wajib): Peran baru ("admin", "user", "lecturer")

### POST /user/generate
- **Deskripsi**: Membuat pengguna dalam jumlah besar
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin)
- **Parameter Query**:
  - `prefix` (string, wajib): Awalan NIM (harus 2 karakter)
  - `start` (integer, wajib): Angka awal
  - `end` (integer, wajib): Angka akhir
- **Output**: File Excel berisi NIM dan password pengguna

## Route Exam

Route untuk mengelola data ujian dalam sistem.

### POST /exam/
- **Deskripsi**: Membuat ujian baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `title` (string, wajib): Judul ujian
  - `creator_id` (integer, wajib): ID pembuat ujian
  - `long_time` (integer, wajib): Durasi ujian dalam menit
  - `started_at` (string, wajib): Tanggal mulai ujian (format: 2006-01-02 15:04:05)
  - `finished_at` (string, wajib): Tanggal selesai ujian (format: 2006-01-02 15:04:05)
- **Contoh Request**:
```json
{
  "title": "Ujian Tengah Semester",
  "creator_id": 1,
  "long_time": 90,
  "started_at": "2023-11-15 08:00:00",
  "finished_at": "2023-11-15 09:30:00"
}
```

### GET /exam/
- **Deskripsi**: Mendapatkan data banyak ujian
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /exam/id
- **Deskripsi**: Mendapatkan data ujian berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID ujian

### PUT /exam/:id
- **Deskripsi**: Mengupdate data ujian
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID ujian yang akan diperbarui
- **Parameter Request (JSON)**:
  - `title` (string, opsional): Judul ujian
  - `creator_id` (integer, opsional): ID pembuat ujian
  - `long_time` (integer, opsional): Durasi ujian dalam menit
  - `started_at` (string, opsional): Tanggal mulai ujian
  - `finished_at` (string, opsional): Tanggal selesai ujian

### DELETE /exam/:id
- **Deskripsi**: Menghapus ujian
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID ujian yang akan dihapus

## Route Subject

Route untuk mengelola data mata kuliah dalam sistem.

### POST /subject/
- **Deskripsi**: Membuat mata kuliah baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `title` (string, wajib): Nama mata kuliah
  - `code` (string, wajib): Kode mata kuliah
- **Contoh Request**:
```json
{
  "title": "Algoritma dan Pemrograman",
  "code": "IF101"
}
```

### GET /subject/
- **Deskripsi**: Mendapatkan data banyak mata kuliah
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /subject/id
- **Deskripsi**: Mendapatkan data mata kuliah berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID mata kuliah

### GET /subject/code
- **Deskripsi**: Mendapatkan data mata kuliah berdasarkan kode
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `code` (string, wajib): Kode mata kuliah

### PUT /subject/
- **Deskripsi**: Mengupdate data mata kuliah
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `id` (integer, wajib): ID mata kuliah
  - `title` (string, opsional): Nama mata kuliah
  - `code` (string, opsional): Kode mata kuliah

### DELETE /subject/
- **Deskripsi**: Menghapus mata kuliah
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `id` (integer, wajib): ID mata kuliah yang akan dihapus

## Route Question

Route untuk mengelola data soal dalam sistem.

### POST /question/
- **Deskripsi**: Membuat soal baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `creator_id` (integer, wajib): ID pembuat soal
  - `question_text` (string, wajib): Teks soal
- **Contoh Request**:
```json
{
  "creator_id": 1,
  "question_text": "Berapa hasil dari 2 + 2?"
}
```

### POST /question/options
- **Deskripsi**: Membuat soal baru beserta pilihan jawabannya
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `creator_id` (integer, wajib): ID pembuat soal
  - `question_text` (string, wajib): Teks soal

### POST /question/json
- **Deskripsi**: Membuat soal dari format JSON
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)

### GET /question/
- **Deskripsi**: Mendapatkan data banyak soal
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /question/id
- **Deskripsi**: Mendapatkan data soal berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID soal

### GET /question/exam
- **Deskripsi**: Mendapatkan soal berdasarkan ujian
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID ujian

### GET /question/diff
- **Deskripsi**: Mendapatkan soal berdasarkan tingkat kesulitan
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `difficulty` (string, wajib): Tingkat kesulitan soal

### GET /question/creator
- **Deskripsi**: Mendapatkan soal berdasarkan pembuat
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID pembuat soal

### GET /question/subject
- **Deskripsi**: Mendapatkan soal berdasarkan mata kuliah
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID mata kuliah

### PUT /question/:id
- **Deskripsi**: Mengupdate data soal
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID soal yang akan diperbarui
- **Parameter Request (JSON)**:
  - `question_text` (string, opsional): Teks soal
  - `difficulty` (string, opsional): Tingkat kesulitan
  - `subject_id` (integer, opsional): ID mata kuliah

### DELETE /question/:id
- **Deskripsi**: Menghapus soal
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID soal yang akan dihapus

## Route Option

Route untuk mengelola pilihan jawaban dalam sistem.

### POST /option/
- **Deskripsi**: Membuat pilihan jawaban baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `question_id` (integer, wajib): ID soal
  - `option_label` (string, wajib): Label pilihan (A, B, C, D, E)
  - `option_text` (string, wajib): Teks pilihan
  - `is_correct` (boolean, wajib): Apakah pilihan ini benar
- **Contoh Request**:
```json
{
  "question_id": 1,
  "option_label": "A",
  "option_text": "Jawaban A",
  "is_correct": true
}
```

### GET /option/
- **Deskripsi**: Mendapatkan data banyak pilihan jawaban
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /option/id
- **Deskripsi**: Mendapatkan data pilihan jawaban berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID pilihan jawaban

### PUT /option/:id
- **Deskripsi**: Mengupdate data pilihan jawaban
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID pilihan jawaban yang akan diperbarui
- **Parameter Request (JSON)**:
  - `option_text` (string, opsional): Teks pilihan
  - `is_correct` (boolean, opsional): Apakah pilihan ini benar

### DELETE /option/:id
- **Deskripsi**: Menghapus pilihan jawaban
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID pilihan jawaban yang akan dihapus

## Route Exam Question

Route untuk mengelola relasi antara ujian dan soal.

### POST /exam-question/
- **Deskripsi**: Menambahkan soal ke dalam ujian
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `exam_id` (integer, wajib): ID ujian
  - `question_ids` (array integer, wajib): Array ID soal

### PUT /exam-question/
- **Deskripsi**: Mengupdate soal dalam ujian
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `exam_id` (integer, wajib): ID ujian
  - `question_ids` (array integer, wajib): Array ID soal baru

### DELETE /exam-question/
- **Deskripsi**: Menghapus soal dari ujian
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `exam_id` (integer, wajib): ID ujian
  - `question_ids` (array integer, wajib): Array ID soal yang akan dihapus

## Route Exam Session

Route untuk mengelola sesi ujian dalam sistem.

### POST /exam-session/
- **Deskripsi**: Membuat sesi ujian baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi
- **Parameter Request (JSON)**:
  - `exam_id` (integer, wajib): ID ujian
- **Contoh Request**:
```json
{
  "exam_id": 1
}
```

### GET /exam-session/
- **Deskripsi**: Mendapatkan data banyak sesi ujian
- **Metode HTTP**: GET
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /exam-session/id
- **Deskripsi**: Mendapatkan data sesi ujian berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID sesi ujian

### PUT /exam-session/:id
- **Deskripsi**: Mengupdate data sesi ujian
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID sesi ujian yang akan diperbarui
- **Parameter Request (JSON)**:
  - `status` (string, opsional): Status sesi ujian

### PUT /exam-session/:id/no
- **Deskripsi**: Mengupdate nomor soal saat ini dalam sesi ujian
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID sesi ujian
- **Parameter Request (JSON)**:
  - `current_no` (integer, wajib): Nomor soal saat ini

### PUT /exam-session/:id/finish
- **Deskripsi**: Menyelesaikan sesi ujian
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID sesi ujian

### DELETE /exam-session/:id
- **Deskripsi**: Menghapus sesi ujian
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID sesi ujian yang akan dihapus

## Route Exam Score

Route untuk mengelola nilai ujian dalam sistem.

### POST /exam-score/
- **Deskripsi**: Membuat nilai ujian baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Request (JSON)**:
  - `exam_id` (integer, wajib): ID ujian
  - `user_id` (integer, wajib): ID pengguna
  - `status` (string, wajib): Status pengerjaan ("done", "not-started", "in-progress")
- **Contoh Request**:
```json
{
  "exam_id": 1,
  "user_id": 1,
  "status": "done",
  "score": 85.5
}
```

### GET /exam-score/
- **Deskripsi**: Mendapatkan data banyak nilai ujian
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /exam-score/id
- **Deskripsi**: Mendapatkan data nilai ujian berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID nilai ujian

### PUT /exam-score/:id
- **Deskripsi**: Mengupdate data nilai ujian
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID nilai ujian yang akan diperbarui
- **Parameter Request (JSON)**:
  - `score` (float, opsional): Nilai ujian
  - `status` (string, opsional): Status pengerjaan

### DELETE /exam-score/:id
- **Deskripsi**: Menghapus nilai ujian
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID nilai ujian yang akan dihapus

## Route User Answer

Route untuk mengelola jawaban pengguna dalam sistem.

### POST /user-answer/
- **Deskripsi**: Membuat jawaban pengguna baru
- **Metode HTTP**: POST
- **Akses**: Terotentikasi
- **Parameter Request (JSON)**:
  - `exam_session_id` (integer, wajib): ID sesi ujian
  - `user_id` (integer, wajib): ID pengguna
  - `question_id` (integer, wajib): ID soal
  - `answer` (string, wajib): Jawaban yang dipilih
- **Contoh Request**:
```json
{
  "exam_session_id": 1,
  "user_id": 1,
  "question_id": 1,
  "answer": "A"
}
```

### GET /user-answer/
- **Deskripsi**: Mendapatkan data banyak jawaban pengguna
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /user-answer/id
- **Deskripsi**: Mendapatkan data jawaban pengguna berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID jawaban pengguna

### GET /user-answer/session
- **Deskripsi**: Mendapatkan jawaban pengguna berdasarkan sesi ujian
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID sesi ujian

### GET /user-answer/question
- **Deskripsi**: Mendapatkan jawaban pengguna berdasarkan soal
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `id` (integer, wajib): ID soal

### GET /user-answer/user
- **Deskripsi**: Mendapatkan jawaban pengguna
- **Metode HTTP**: GET
- **Akses**: Terotentikasi
- **Parameter Query**:
  - `exam_session_id` (integer, wajib): ID sesi ujian

### PUT /user-answer/:id
- **Deskripsi**: Mengupdate data jawaban pengguna
- **Metode HTTP**: PUT
- **Akses**: Terotentikasi
- **Parameter Path**:
  - `id` (integer, wajib): ID jawaban pengguna yang akan diperbarui
- **Parameter Request (JSON)**:
  - `answer` (string, opsional): Jawaban yang dipilih

### DELETE /user-answer/:id
- **Deskripsi**: Menghapus jawaban pengguna
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin atau Dosen)
- **Parameter Path**:
  - `id` (integer, wajib): ID jawaban pengguna yang akan dihapus

## Route Xls Path

Route untuk mengelola file Excel dalam sistem.

### GET /xlspath/
- **Deskripsi**: Mendapatkan data banyak file Excel
- **Metode HTTP**: GET
- **Akses**: Terotentikasi (Role Admin)
- **Parameter Query**:
  - `limit` (integer, opsional): Jumlah data maksimal (default: 20)
  - `offset` (integer, opsional): Offset data (default: 0)

### GET /xlspath/id
- **Deskripsi**: Mendapatkan data file Excel berdasarkan ID
- **Metode HTTP**: GET
- **Akses**: Terotentikasi (Role Admin)
- **Parameter Query**:
  - `id` (integer, wajib): ID file Excel

### DELETE /xlspath/
- **Deskripsi**: Menghapus file Excel
- **Metode HTTP**: DELETE
- **Akses**: Terotentikasi (Role Admin)
- **Parameter Request (JSON)**:
  - `id` (integer, wajib): ID file Excel yang akan dihapus