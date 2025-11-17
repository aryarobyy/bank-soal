import { USER } from "../core/constant";
import ApiHandler from "./api.handler";

export const register = async (data) => {
  const res = await ApiHandler.post(`${USER}/register`, data);
  return res.data;
};

export const login = async (data) => {
  const res = await ApiHandler.post(`${USER}/login`, data);
  return res;
};

export const refreshToken = async () => {
  const res = await ApiHandler.post(`${USER}/refresh`, {});
  return res;
};

// src/provider/user.provider.js

// ... (fungsi register, login, refreshToken tetap sama) ...

export const updateUser = async (data, id) => {
  const formData = new FormData();
  
  // ## AWAL PERUBAHAN ##
  // Kita ubah cara memasukkan data ke FormData
  for (const key in data) {
    const value = data[key];

    if (value === null) {
      // Jika nilainya null, kirim sebagai string KOSONG ("")
      // Ini akan mencegah 'null' berubah menjadi string "null"
      formData.append(key, "");
    } else if (value !== undefined) {
      // Jika nilainya bukan undefined (bisa jadi string, angka, boolean, atau file),
      // baru kita tambahkan.
      formData.append(key, value);
    }
    // Jika nilainya 'undefined', kita tidak melakukan apa-apa (field tidak akan dikirim)
  }
  // ## AKHIR PERUBAHAN ##

  const res = await ApiHandler.put(`${USER}/${id}`, formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });
  return res.data;
};

// ... (sisa file user.provider.js tetap sama) ...

export const getUserById = async (id) => {
  const res = await ApiHandler.get(`${USER}/id?id=${id}`);
  return res.data;
};

export const getUsers = async (limit = 10, offset = 0) => {
  // 1. Tambahkan parameter limit & offset ke URL
  const res = await ApiHandler.get(
    `${USER}/?limit=${limit}&offset=${offset}`
  );
  
  // 2. Kembalikan res.data.data
  // Berdasarkan backend Go Anda (GetMany) dan struktur Exam, 
  // ini seharusnya mengembalikan objek: { data: [...], total: ... }
  return res.data.data;
};

export const getUserByName = async (name) => {
  const res = await ApiHandler.get(`${USER}/name?name=${name}`);
  return res;
};

// ## PERBAIKAN 1: Menggunakan Path Parameter ( /user/:id ) ##
export const deleteUser = async (id) => {
  // Menggunakan format path parameter: /user/11
  const res = await ApiHandler.delete(`${USER}/${id}`);
  return res.data;
};

/**
 * provider baru
 */
export const getUsersByRole = async (role, limit = 10, offset = 0) => {
  // 1. Buat query parameters
  const params = new URLSearchParams();
  params.append('role', role);
  params.append('limit', limit);
  params.append('offset', offset);

  // 2. Panggil endpoint
  const res = await ApiHandler.get(`${USER}/role?${params.toString()}`);
  
  // 3. Kembalikan objek { data: [...], total: ... }
  // (Berdasarkan provider Anda yang lain, 'res.data.data' adalah format yang benar)
  return res.data.data;
};

export const changeRole = async (userId, adminId, role) => {
  const res = await ApiHandler.put(`${USER}/role?id=${userId}`, {
    admin_id: adminId,
    role: role,
  });
  return res.data;
};

/**
 * ## PERBAIKAN 2: Menambahkan admin_id untuk melewati 'old_password' ##
 * @param {number} id - ID user (masuk ke query param)
 * @param {string} newPassword - Password baru (masuk ke body)
 * @param {number} adminId - ID Admin (masuk ke body)
 */
export const changePassword = async (id, newPassword, adminId) => {
  // 1. Tetap mengirim sebagai JSON
  // 2. Mengganti key 'password' menjadi 'new_password'
  const res = await ApiHandler.put(`${USER}/password?id=${id}`, {
    new_password: newPassword, // <-- INI PERBAIKANNYA
    admin_id: adminId, 
  });
  return res.data;
};

export const getUserByEmail = async (email) => {
  const res = await ApiHandler.get(`${USER}/email?email=${email}`);
  return res.data;
};

export const getUserByNim = async (nim) => {
  const res = await ApiHandler.get(`${USER}/nim?nim=${nim}`);
  return res.data;
};

export const checkUser = async (id) => {
  const res = await ApiHandler.get(`${USER}/check?id=${id}`);
  return res.data;
};

// src/provider/user.provider.js
// ... (fungsi-fungsi lain) ...

export const generateUsers = async (prefix, start, end, academic_year) => {
 
  // 1. Buat query params (untuk data di URL)
  const params = new URLSearchParams();
  params.append('prefix', String(prefix));
  params.append('start', String(start));
  params.append('end', String(end));

  // 2. Buat request body (untuk data di BODY)
  // Sesuai 'user_srv.go', body hanya berisi 'academic_year'
  const requestBody = {
    academic_year: String(academic_year),
  };

  // 3. Panggil POST dengan KEDUA-DUANYA
  const res = await ApiHandler.post(
    `${USER}/generate?${params.toString()}`, // <-- Data di URL
    requestBody                             // <-- Data di Body
    // <-- 'responseType: blob' DIHAPUS (sudah benar)
  );

  // 4. Kembalikan data JSON
  return res.data;
};