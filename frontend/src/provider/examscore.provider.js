import ApiHandler from "./api.handler";

const BASE_URL = "exam-score";


export const getExamScores = async (limit = 100, offset = 0, examId) => {
  try {
    const params = new URLSearchParams();
    params.append('limit', limit);
    params.append('offset', offset);
    

    if (examId) {
        params.append('exam_id', examId);
    }

    const res = await ApiHandler.get(`/${BASE_URL}/?${params.toString()}`);
    
   
    if (res.data && Array.isArray(res.data.data)) {
        return res.data.data;
    } else if (Array.isArray(res.data)) {
        return res.data;
    }
    
    return [];
  } catch (error) {
    console.error("Gagal mengambil data nilai ujian:", error);
    return [];
  }
};