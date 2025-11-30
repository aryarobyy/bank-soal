import ApiHandler from "./api.handler";

const USER_ANSWER = "user-answer";


export const submitUserAnswer = async (payload) => {
  const res = await ApiHandler.post(`/${USER_ANSWER}/`, payload);
  return res.data;
};



export const getUserAnswersBySession = async (sessionId) => {
  try {
    const res = await ApiHandler.get(`/${USER_ANSWER}/session?session_id=${sessionId}`);
    
    if (res.data && res.data.data && Array.isArray(res.data.data.data)) return res.data.data.data;
    if (res.data && Array.isArray(res.data.data)) return res.data.data;
    if (Array.isArray(res.data)) return res.data;
    return [];

  } catch (error) {
    
    if (error.response && error.response.status === 404) {
        return [];
    }
    console.error("Gagal mengambil jawaban user:", error);
    return [];
  }
};