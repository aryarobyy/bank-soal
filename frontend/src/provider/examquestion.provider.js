import ApiHandler from "./api.handler";

/**
 * GET /question/exam?exam_id=XX
 * Mengambil semua soal yang ada di satu ujian.
 */
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
 * Tambah soal ke ujian
 */
export const addExamQuestions = async (examId, questionIds) => {
  try {
    const res = await ApiHandler.post(`/exam-question/`, {
      exam_id: examId,
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
export const deleteExamQuestions = async (examId, questionIdsArray) => {
  try {
    const res = await ApiHandler.delete(`/exam-question/`, {
      data: {
        exam_id: examId,
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
