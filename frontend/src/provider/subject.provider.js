import { SUBJECT } from "../core/constant"; // (Pastikan 'SUBJECT' ada di constant.js Anda, misal: 'subject')
import ApiHandler from "./api.handler";


export const getPaginatedSubjects = async (limit, offset, searchQuery) => {

  const params = new URLSearchParams();
  params.append('limit', limit);
  params.append('offset', offset);


  if (searchQuery) {
    params.append('search', searchQuery); 
  }
  

  const res = await ApiHandler.get(`/${SUBJECT}/?${params.toString()}`);

  return res.data; 
};



export const createSubject = async (data) => {

  const res = await ApiHandler.post(`/${SUBJECT}/`, data); 
  return res.data;
};

export const getSubjectById = async (id) => {
  const res = await ApiHandler.get(`/${SUBJECT}/id?id=${id}`);
  return res.data;
};

export const getSubjectByCode = async (code) => {
  const res = await ApiHandler.get(`/${SUBJECT}/code?code=${code}`);
  return res.data;
};

export const updateSubject = async (id, data) => {
  const res = await ApiHandler.put(`/${SUBJECT}/?id=${id}`, data);
  return res.data;
};

export const deleteSubject = async (id) => {
  const res = await ApiHandler.delete(`/${SUBJECT}/${id}`);
  return res.data;
};