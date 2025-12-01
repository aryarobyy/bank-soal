import { QUESTION } from "../core/constant";
import ApiHandler from "./api.handler";

export const getAllQuestions = async () => {
  const res = await ApiHandler.get(`/${QUESTION}/`);
  return res.data.data;
};

export const getmanyQuestions = async (limit,offset) => {
  const res = await ApiHandler.get(`/${QUESTION}/?limit=${limit}&offset=${offset}`);
  return res.data.data;
};

export const createQuestionWithOptions = async (data) => {
  const formData = new FormData();

  for (const key in data) {
    const value = data[key];
    
    if (key === 'options') {

      formData.append(key, JSON.stringify(value));
    } else if (value !== null && value !== undefined) {

      formData.append(key, value);
    }
  }


  const res = await ApiHandler.post(`${QUESTION}/`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
  return res.data;
};

export const deleteQuestion = async (id) => {
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
    
    if (key === 'options') {
  
      formData.append(key, JSON.stringify(value));
    } else if (value !== null && value !== undefined) {
 
      formData.append(key, value);
    }
  }
  

  const res = await ApiHandler.put(`${QUESTION}/${id}`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
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

  const res = await ApiHandler.get(`${QUESTION}/exam?exam_id=${examId}`); 
  return res.data.data;
};
export const getQuestionsByDiff = async (difficulty) => {
  const res = await ApiHandler.get(`${QUESTION}/diff?diff=${difficulty}`);
  return res.data.data;
};

export const getQuestionsByCreator = async (creatorId, limit, offset) => {

  const params = new URLSearchParams();
  params.append('creator_id', creatorId);
  params.append('limit', limit);
  params.append('offset', offset);

  
  const res = await ApiHandler.get(`/${QUESTION}/creator?${params.toString()}`);
  
 
  return res.data.data; 
};

export const getQuestionsBySubject = async (subjectId, limit, offset) => {
  
  const params = new URLSearchParams();
  params.append('subject_id', subjectId);
  params.append('limit', limit);
  params.append('offset', offset);

  const res = await ApiHandler.get(`/${QUESTION}/subject?${params.toString()}`);
  

  return res.data.data; 
};

export const createWithJson = async () => {
  const res = await ApiHandler.post(
    `/${QUESTION}/json/`,{
      headers: {
      'Content-Type': 'application/json',
    },
  }
  )
  return res.data.data
}