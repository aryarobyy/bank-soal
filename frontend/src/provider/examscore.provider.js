// src/provider/examscore.provider.js
import ApiHandler from "./api.handler";

const BASE_URL = "exam-score";

/**
 * Mendapatkan daftar nilai ujian
 * Endpoint: GET /exam-score/?limit=...&offset=...
 */
export const getExamScores = async (limit = 100, offset = 0) => {
  try {
    // Kita set limit agak besar karena ini laporan
    const res = await ApiHandler.get(`/${BASE_URL}/?limit=${limit}&offset=${offset}`);
    
    // Sesuai pola backend Anda sebelumnya, data biasanya ada di res.data.data
    // Tapi kita return res.data dulu biar fleksibel (biasanya { data: [], total: ... })
    return res.data;
  } catch (error) {
    console.error("Gagal mengambil data nilai ujian:", error);
    throw error;
  }
};