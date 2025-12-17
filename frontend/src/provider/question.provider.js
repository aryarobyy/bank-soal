import { QUESTION } from "../core/constant";
import ApiHandler from "./api.handler";

const extractData = (res) => {
  let items = [];
  let total = 0;
  if (res?.data?.data?.data && Array.isArray(res.data.data.data)) {
    items = res.data.data.data;
    total = res.data.data.total || 0;
  } else if (res?.data?.data && Array.isArray(res.data.data)) {
    items = res.data.data;
    total = res.data.total || items.length;
  } else if (res?.data && Array.isArray(res.data)) {
    items = res.data;
    total = items.length;
  }
  return { items, total };
};

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
export const getRandomQuestions = async (total, subjectId) => {
  const res = await ApiHandler.get(`/${QUESTION}/random?total=${total}&subject_id=${subjectId}`);
  return res.data.data;
};

export const getQuestionsByExam = async (examId, limit = 10, offset = 0) => {
  try {
    const res = await ApiHandler.get(
      `${QUESTION}/exam?exam_id=${examId}&limit=${limit}&offset=${offset}`
    );
    const responseBody = res.data;
    if (responseBody.data && Array.isArray(responseBody.data.data)) {
        return { 
            data: responseBody.data.data, 
            total: responseBody.data.total || 0 
        };
    }
    if (Array.isArray(responseBody.data)) {
        return { 
            data: responseBody.data, 
            total: responseBody.total || responseBody.data.length 
        };
    }
    if (Array.isArray(responseBody)) {
        return { data: responseBody, total: responseBody.length };
    }
    return { data: [], total: 0 };
  } catch (error) {
    console.error("Gagal load questions:", error);
    return { data: [], total: 0 };
  }
};


export const getQuestionsByDiff = async (difficulty) => {
  const res = await ApiHandler.get(`${QUESTION}/diff?diff=${difficulty}`);
  return res.data.data;
};

export const getQuestionsByCreator = async (creatorId, limit, offset) => {
  const res = await ApiHandler.get(`/${QUESTION}/creator`, {
    params: { creator_id: creatorId, limit, offset }
  });
  return extractData(res);
};
export const getQuestionsBySubject = async (subjectId, limit, offset) => {
  
  const params = new URLSearchParams();
  params.append('subject_id', subjectId);
  params.append('limit', limit);
  params.append('offset', offset);
  
  const res = await ApiHandler.get(`/${QUESTION}/subject?${params.toString()}`);
  

  return res.data.data; 
};
export const getQuestionsByCreatorAndSubject = async (creatorId, subjectId, limit, offset) => {
  const res = await ApiHandler.get(`/${QUESTION}/creator/subject`, {
    params: { 
      creator_id: creatorId, 
      subject_id: subjectId, 
      limit, 
      offset 
    }
  });
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

const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));


export const getAllQuestionsForExamDo = async (examId) => {
  let allData = [];
  let offset = 0;
  const BATCH_LIMIT = 20;
  let keepFetching = true;

  try {
    while (keepFetching) {
      const res = await ApiHandler.get(
        `${QUESTION}/exam?exam_id=${examId}&limit=${BATCH_LIMIT}&offset=${offset}`
      );
      let chunk = [];
      const responseBody = res.data;
      if (responseBody.data && Array.isArray(responseBody.data.data)) {
          chunk = responseBody.data.data;
      } else if (Array.isArray(responseBody.data)) {
          chunk = responseBody.data;
      }

      if (chunk.length > 0) {
        allData.push(...chunk);
        offset += BATCH_LIMIT;
      }

      if (chunk.length < BATCH_LIMIT) {
        keepFetching = false;
      } else {
        await delay(1500); 
      }
    }
    
    return allData; 

  } catch (error) {
    console.error("Gagal mengambil full soal:", error);
    return allData; 
  }
};