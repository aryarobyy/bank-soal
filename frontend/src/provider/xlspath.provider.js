// src/provider/xlspath.provider.js

import { XLSPATH } from "../core/constant"; // Anda perlu menambahkan XLSPATH = 'xlspath' ke file constant.js Anda
import ApiHandler from "./api.handler";

/**
 * Mengambil semua data file excel (mendukung paginasi)
 * Sesuai: GET /xlspath/
 */
export const getAllXlsPaths = async (limit = 10, offset = 0) => {
  const params = new URLSearchParams();
  params.append('limit', limit);
  params.append('offset', offset);
  
  const res = await ApiHandler.get(`/${XLSPATH}/?${params.toString()}`);
  
  // Asumsi backend mengembalikan { data: { data: [...], total: ... } }
  // seperti endpoint 'getUsersByRole'
  return res.data.data; 
};

/**
 * Mengambil data file excel berdasarkan ID
 * Sesuai: GET /xlspath/id?id=...
 */
export const getXlsPathById = async (id) => {
  const res = await ApiHandler.get(`/${XLSPATH}/id?id=${id}`);
  return res.data.data; 
};

/**
 * Menghapus data file excel berdasarkan ID
 * Sesuai: DELETE /xlspath?id=...
 */
export const deleteXlsPath = async (id) => {
  const res = await ApiHandler.delete(`/${XLSPATH}/?id=${id}`);
  return res.data;
};

export const downloadXlsFile = async (id) => {
  const res = await ApiHandler.get(
    `/${XLSPATH}/download?id=${id}`,
    {
      responseType: 'blob',
    }
  );
  return res; 
};