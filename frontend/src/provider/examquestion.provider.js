// src/provider/examquestion.provider.js
import ApiHandler from "./api.handler";

/**
 * 1. Ambil semua soal dalam ujian tertentu
 * (Ini endpoint yang error CORS, path sudah benar)
 */
export const getExamQuestionsByExamId = async (examId) => {
  try {
    const res = await ApiHandler.put(`/exam-question?exam_id=${examId}`);
    return res.data?.data || [];
  } catch (error) {
    console.error("❌ Gagal memuat soal ujian:", error);
    throw error;
  }
};

/**
 * 2. Tambah soal ke ujian (SESUAI DOKUMENTASI BARU)
 * Mengirim exam_id di dalam BODY
 */
export const addExamQuestions = async (examId, questionIds) => {
  try {
    // Endpointnya adalah '/exam-question/'
    // Body-nya berisi 'exam_id' dan 'question_ids'
    const res = await ApiHandler.post(`/exam-question/`, {
      exam_id: examId,
      question_ids: questionIds,
    });
    return res.data;
  } catch (error) {
    console.error("❌ Gagal menambahkan soal ke ujian:", error);
    throw error;
  }
};

/**
 * 3. Hapus soal dari ujian (SESUAI DOKUMENTASI BARU)
 * Mengirim exam_id dan array question_ids di BODY
 */
export const deleteExamQuestions = async (examId, questionIdsArray) => {
  try {
    const res = await ApiHandler.delete(`/exam-question/`, {
      // DELETE bisa punya body, kita kirim data di 'data'
      data: {
        exam_id: examId,
        question_ids: questionIdsArray,
      }
    });
    return res.data;
  } catch (error) {
    console.error("❌ Gagal menghapus soal ujian:", error);
    throw error;
  }
};