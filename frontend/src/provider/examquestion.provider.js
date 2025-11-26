import ApiHandler from "./api.handler";

export const getExamQuestions = async (examId) => {
  try {
    const res = await ApiHandler.get(`/question/exam?exam_id=${examId}`);

    // BACKEND RESPONSE = { data: [ { ... } ] }
    // Jadi yang benar:
    return res.data.data;
  } catch (error) {
    console.error("❌ Gagal memuat soal ujian:", error);
    throw error;
  }
};

/**
 * Compat alias function lama
 */
export const getExamQuestionsByExamId = async (examId) => {
  return await getExamQuestions(examId);
};
/**
 * 2. Tambah soal ke ujian (SESUAI DOKUMENTASI BARU)
 * Mengirim exam_id di dalam BODY
 */
export const addExamQuestions = async (examId, questionIds) => {
  try {
    // PERHATIKAN: Tambahkan '/' sebelum '?'
    // Dari: /exam-question?exam_id=...
    // Jadi: /exam-question/?exam_id=...
    
    const res = await ApiHandler.post(`/exam-question/?exam_id=${examId}`, {
      question_ids: questionIds,
    });

    return res.data.data; // SATUKAN FORMAT RESPONSE
  } catch (error) {
    console.error("❌ Gagal menambahkan soal ke ujian:", error);
    throw error;
  }
};

/**
 * Hapus soal dari ujian
 */
// src/provider/examquestion.provider.js

export const deleteExamQuestions = async (examId, questionIdsArray) => {
  try {
    // PERBAIKAN:
    // 1. Pindahkan examId ke URL (?exam_id=...)
    // 2. Body hanya berisi question_ids
    
    const res = await ApiHandler.delete(`/exam-question/?exam_id=${examId}`, {
      data: {
        question_ids: questionIdsArray,
      },
    });

    return res.data.data; // Format sama
  } catch (error) {
    console.error("❌ Gagal menghapus soal ujian:", error);
    throw error;
  }
};

export default {
  getExamQuestions,
  getExamQuestionsByExamId,
  addExamQuestions,
  deleteExamQuestions,
};
