import ApiHandler from "./api.handler";

const EXAM_SESSION = "exam-session";


export const createExamSession = async (examId) => {
  const res = await ApiHandler.post(`/${EXAM_SESSION}/`, {
    exam_id: examId,
  });
  return res.data;
};


export const getExamSessionById = async (sessionId) => {
  const res = await ApiHandler.get(`/${EXAM_SESSION}/id?id=${sessionId}`);
  return res.data;
};



export const updateExamSession = async (sessionId, payload) => {
  const res = await ApiHandler.put(`/${EXAM_SESSION}/${sessionId}`, payload);
  return res.data;
};


export const updateCurrentNo = async (sessionId, currentNo) => {
  const res = await ApiHandler.put(`/${EXAM_SESSION}/${sessionId}/no`, {
    current_no: currentNo,
  });
  return res.data;
};


export const finishExamSession = async (payload) => {
  const res = await ApiHandler.put(`/${EXAM_SESSION}/finish`, payload);
  return res.data;
};

export const deleteExamSession = async (sessionId) => {
  const res = await ApiHandler.delete(`/${EXAM_SESSION}/${sessionId}`);
  return res.data;
};

export const getExamSessions = async (limit = 100, offset = 0, examId = null) => {
  const params = new URLSearchParams();
  params.append('limit', limit);
  params.append('offset', offset);
  
  if (examId) {
    params.append('exam_id', examId);
  }
  try {
    const res = await ApiHandler.get(`/${EXAM_SESSION}/?${params.toString()}`);
    if (res.data && res.data.data && Array.isArray(res.data.data.data)) {
      
        return res.data.data;
    } 
    if (res.data && Array.isArray(res.data.data)) {
        return res.data.data;
    }
    return []; 
  } catch (error) {
    console.error("Gagal mengambil data sesi ujian:", error);
    return [];
  }
};

export const getExamSessionByUser = async (userId, limit = 10, offset = 0) => {
  const params = new URLSearchParams();
  params.append('user_id', userId);
  params.append('limit', limit);
  params.append('offset', offset);

  try {

    const res = await ApiHandler.get(`/${EXAM_SESSION}/user?${params.toString()}`);
    
    if (res.data && res.data.data && Array.isArray(res.data.data.data)) {
        return res.data.data.data;
    }
    
   
    if (res.data && Array.isArray(res.data.data)) {
        return res.data.data;
    }

    return [];

  } catch (error) {
    console.error("Gagal ambil session user:", error);
    return [];
  }
};