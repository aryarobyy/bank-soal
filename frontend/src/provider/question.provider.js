import { QUESTION } from "../core/constant";
import ApiHandler from "./api.handler";

export const getAllQuestions = async () => {
  const res = await ApiHandler.get(`/${QUESTION}/`);
  return res.data.data;
};

export const getmanyQuestions = async (limit, offset) => {
  const res = await ApiHandler.get(
    `/${QUESTION}/?limit=${limit}&offset=${offset}`
  );
  return res.data.data;
};

export const createQuestionWithOptions = async (data) => {
  const formData = new FormData();

  // Loop semua data dari formatPayload
  for (const key in data) {
    const value = data[key];

    if (key === "options") {
      // Backend (sesuai Postman) mengharapkan 'options' sebagai string JSON
      formData.append(key, JSON.stringify(value));
    } else if (value !== null && value !== undefined) {
      // Ini akan menangani string, angka, boolean, dan File (gambar)
      formData.append(key, value);
    }
  }

  // Kirim sebagai multipart/form-data
  const res = await ApiHandler.post(`${QUESTION}/`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return res.data;
};

export const deleteQuestion = async (id) => {
  // Menggunakan format path parameter: /question/5
  const res = await ApiHandler.delete(`/${QUESTION}/${id}`);
  return res.data;
};

export const getQuestionById = async (id) => {
  const res = await ApiHandler.get(`/${QUESTION}/id?id=${id}`);
  return res.data;
};

export const updateQuestion = async (id, data) => {
  const formData = new FormData();

  for (const key in data) {
    const value = data[key];

    if (key === "options") {
      // Backend (sesuai Postman) mengharapkan 'options' sebagai string JSON
      formData.append(key, JSON.stringify(value));
    } else if (value !== null && value !== undefined) {
      // Ini akan menangani string, angka, boolean, dan File (gambar)
      // Jika 'image' (value) adalah null, ini akan dilewati (perilaku benar)
      formData.append(key, value);
    }
  }

  // Endpoint update (PUT /question/:id)
  const res = await ApiHandler.put(`${QUESTION}/${id}`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return res.data;
};

export const createQuestion = async (data) => {
  const res = await ApiHandler.post(`${QUESTION}`, data);
  return res.data;
};

export const createQuestionFromJson = async (data) => {
  const res = await ApiHandler.post(`${QUESTION}/json`, data);
  return res.data;
};

export const getQuestionsByExam = async (examId) => {
  // Parameter dikembalikan ke 'exam_id' agar sesuai dengan Postman
  const res = await ApiHandler.get(`${QUESTION}/exam?exam_id=${examId}`); // <-- PERUBAHAN DI SINI
  return res.data.data;
};
export const getQuestionsByDiff = async (difficulty) => {
  const res = await ApiHandler.get(`${QUESTION}/diff?diff=${difficulty}`);
  return res.data.data;
};

export const getQuestionsByCreator = async (creatorId, limit, offset) => {
  // 1. Tambahkan parameter limit & offset
  const params = new URLSearchParams();
  params.append("creator_id", creatorId);
  params.append("limit", limit);
  params.append("offset", offset);

  const res = await ApiHandler.get(`/${QUESTION}/creator?${params.toString()}`);

  return res.data.data;
};

export const getQuestionsBySubject = async (subjectId, limit, offset) => {
  // 1. Tambahkan parameter limit & offset
  const params = new URLSearchParams();
  params.append("subject_id", subjectId);
  params.append("limit", limit);
  params.append("offset", offset);

  // 2. Kirim parameter ke API
  const res = await ApiHandler.get(`/${QUESTION}/subject?${params.toString()}`);

  // 3. Kembalikan res.data (yang berisi { data: [...], total: ... })
  return res.data.data;
};
