// src/provider/examquestion.provider.js
import ApiHandler from "./api.handler";

// ✅ Ambil semua soal dalam ujian tertentu
export const getExamQuestionsByExamId = async (examId) => {
  try {
    const res = await ApiHandler.get(`/exam-question?exam_id=${examId}`);
    // Backend biasanya mengembalikan { code, status, data: [...] }
    return res.data?.data || [];
  } catch (error) {
    console.error("❌ Gagal memuat soal ujian:", error);
    throw error;
  }
};

// ✅ Tambah soal ke ujian
export const addExamQuestions = async (examId, questionIds) => {
  try {
    const res = await ApiHandler.post(`/exam-question?exam_id=${examId}`, {
      question_ids: questionIds,
    });
    return res.data;
  } catch (error) {
    console.error("❌ Gagal menambahkan soal ke ujian:", error);
    throw error;
  }
};

// ✅ Hapus soal dari ujian (jika endpoint tersedia)
export const deleteExamQuestion = async (examQuestionId) => {
  try {
    const res = await ApiHandler.delete(
      `/exam-question/id?id=${examQuestionId}`
    );
    return res.data;
  } catch (error) {
    console.error("❌ Gagal menghapus soal ujian:", error);
    throw error;
  }
};
