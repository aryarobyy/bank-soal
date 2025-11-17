import { SUBJECT } from "../core/constant"; // (Pastikan 'SUBJECT' ada di constant.js Anda, misal: 'subject')
import ApiHandler from "./api.handler";

/**
 * Mengambil daftar subjek yang sudah dipaginasi dari server.
 * Sesuai dengan API baru: GET /subject
 */
export const getPaginatedSubjects = async (limit, offset, searchQuery) => {
  // Buat query parameters
  const params = new URLSearchParams();
  params.append('limit', limit);
  params.append('offset', offset);

  // Tambahkan search query jika ada
  if (searchQuery) {
    params.append('search', searchQuery); 
  }
  
  // "sortBy" DIHAPUS SESUAI PERMINTAAN ANDA

  // ## PERBAIKAN ERROR 301 ADA DI SINI: Tambahkan '/' ##
  const res = await ApiHandler.get(`/${SUBJECT}/?${params.toString()}`);
  
  // Sesuai info backend "response.data.total"
  return res.data; 
};

// (Provider lain yang Anda tunjukkan di Postman)

export const createSubject = async (data) => {
  const res = await ApiHandler.post(`/${SUBJECT}`, data);
  return res.data;
};

export const getSubjectById = async (id) => {
  const res = await ApiHandler.get(`/${SUBJECT}/id?id=${id}`);
  return res.data;
};

export const getSubjectByCode = async (code) => {
  const res = await ApiHandler.get(`/${SUBJECT}/code?code=${code}`);
  return res.data;
};

export const updateSubject = async (id, data) => {
  const res = await ApiHandler.put(`/${SUBJECT}?id=${id}`, data);
  return res.data;
};

export const deleteSubject = async (id) => {
  const res = await ApiHandler.delete(`/${SUBJECT}?id=${id}`);
  return res.data;
};