import { EXAM } from "../core/constant";
import ApiHandler from "./api.handler";


export const getAllExam = async (limit = 10, offset = 0, creatorId = null) => {
  try {

    const params = new URLSearchParams();
    params.append('limit', limit);
    params.append('offset', offset);
    
    if (creatorId) {
      params.append('creator_id', creatorId);
    }


    const res = await ApiHandler.get(
      `/${EXAM}/?${params.toString()}`
    );
    

    return res.data.data; 
    
  } catch (error) {
    console.error("Gagal mengambil data ujian:", error);
    return { data: [], total: 0 }; 
  }
};



export const getExamById = async (id) => {
  try {
    const res = await ApiHandler.get(`/${EXAM}/id?id=${id}`);
    return res.data?.data || res.data || null;
  } catch (error) {
    console.error(`Gagal mengambil data ujian ID ${id}:`, error);
    throw error;
  }
};

export const createExam = async (data) => {
  try {
    const res = await ApiHandler.post(`/${EXAM}/`, data);
    return res.data;
  } catch (error) {
    console.error("Gagal membuat ujian:", error);
    throw error;
  }
};


export const updateExam = async (id, data) => {
  try {
    const res = await ApiHandler.put(`/${EXAM}/${id}`, data);
    return res.data;
  } catch (error) {
    console.error(`Gagal memperbarui ujian ID ${id}:`, error);
    throw error;
  }
};


export const deleteExam = async (id) => {
  try {
    const examId = parseInt(id);
    const res = await ApiHandler.delete(`/${EXAM}/${examId}`);
    return res.data;
  } catch (error) {
    console.error(`Gagal menghapus ujian ID ${id}:`, error);
    throw error;
  }
};

export const addQuestions = async (id, data) => {
  console.log(data)
  const res = await ApiHandler.put(`/${EXAM}/q/add/${id}`, data)
  return res.data.data
}

export const replaceQuestions = async (id, data) => {
  const res = await ApiHandler.put(`/${EXAM}/q/replace/${id}`, data)
  return res.data.data
}

export const removeQuestions = async (examId, payload) => {
  const res = await ApiHandler.delete(`/${EXAM}/q/${examId}`, {
    data: payload 
  });
  return res.data;
};