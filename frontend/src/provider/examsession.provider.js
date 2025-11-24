import ApiHandler from "./api.handler";

const EXAM_SESSION = "exam-session";

// 🚀 Create exam session
export const createExamSession = async (examId) => {
  const res = await ApiHandler.post(`/${EXAM_SESSION}/`, {
    exam_id: examId,
  });
  return res.data;
};

// 📌 Get session by ID -> pakai query ?id=
export const getExamSessionById = async (sessionId) => {
  const res = await ApiHandler.get(`/${EXAM_SESSION}/id?id=${sessionId}`);
  return res.data;
};

// 📌 Get all sessions (admin/lecturer)
export const getExamSessions = async () => {
  const res = await ApiHandler.get(`/${EXAM_SESSION}/`);
  return res.data;
};

// 🔁 Update entire exam-session (rarely used)
export const updateExamSession = async (sessionId, payload) => {
  const res = await ApiHandler.put(`/${EXAM_SESSION}/${sessionId}`, payload);
  return res.data;
};

// 🔁 Update current_no -> /exam-session/:id/no
export const updateCurrentNo = async (sessionId, currentNo) => {
  const res = await ApiHandler.put(`/${EXAM_SESSION}/${sessionId}/no`, {
    current_no: currentNo,
  });
  return res.data;
};

// 🏁 Finish exam session → /finish (NO ID)
export const finishExamSession = async (sessionId) => {
  const res = await ApiHandler.put(`/exam-session/finish?id=${sessionId}`);
  return res.data;
};

// ❌ Delete session -> /exam-session/:id
export const deleteExamSession = async (sessionId) => {
  const res = await ApiHandler.delete(`/${EXAM_SESSION}/${sessionId}`);
  return res.data;
};
