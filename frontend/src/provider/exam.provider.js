import { EXAM } from "../core/constant";
import ApiHandler from "./api.handler";

// ✅ Get all exams (support pagination, aman jika BE pakai query)
export const getAllExam = async (limit = 10, offset = 0) => {
  try {
    const res = await ApiHandler.get(
      `/${EXAM}/?limit=${limit}&offset=${offset}`
    );
    // backend kamu kadang return { data: [...] }, kadang langsung array
    return res.data?.data || res.data || [];
  } catch (error) {
    console.error("Gagal mengambil data ujian:", error);
    throw error;
  }
};

// ✅ Get exam by ID (mendukung format lama dan baru)
export const getExamById = async (id) => {
  try {
    const res = await ApiHandler.get(`/${EXAM}/id?id=${id}`);
    return res.data?.data || res.data || null;
  } catch (error) {
    console.error(`Gagal mengambil data ujian ID ${id}:`, error);
    throw error;
  }
};

// ✅ Create exam (POST)
export const createExam = async (data) => {
  try {
    const res = await ApiHandler.post(`/${EXAM}/`, data);
    return res.data;
  } catch (error) {
    console.error("Gagal membuat ujian:", error);
    throw error;
  }
};

// ✅ Update exam (PUT)
export const updateExam = async (id, data) => {
  try {
    const res = await ApiHandler.put(`/${EXAM}/id?id=${id}`, data);
    return res.data;
  } catch (error) {
    console.error(`Gagal memperbarui ujian ID ${id}:`, error);
    throw error;
  }
};

// ✅ Delete exam (DELETE) — diperbaiki agar lebih fleksibel
export const deleteExam = async (id) => {
  try {
    // beberapa backend pakai /exam/:id, beberapa pakai /exam/id?id=
    // kita coba versi pertama, kalau gagal fallback ke yang lama
    try {
      const res = await ApiHandler.delete(`/${EXAM}/${id}`);
      return res.data;
    } catch {
      const res = await ApiHandler.delete(`/${EXAM}/id?id=${id}`);
      return res.data;
    }
  } catch (error) {
    console.error(`Gagal menghapus ujian ID ${id}:`, error);
    throw error;
  }
};
