import { EXAM } from "../core/constant";
import ApiHandler from "./api.handler";

// ✅ Get all exams
export const getAllExam = async (limit, offset) => {
  const res = await ApiHandler.get(`/${EXAM}?limit=${limit}&offset=${offset}`);
  return res.data;
};

// ✅ Get exam by ID
export const getExamById = async (id) => {
  const res = await ApiHandler.get(`/${EXAM}/id?id=${id}/`);
  return res.data;
};

// ✅ Create exam
export const createExam = async (data) => {
  const res = await ApiHandler.post(`/${EXAM}/`, data);
  return res.data;
};

// ✅ Update exam
export const updateExam = async (id, data) => {
  const res = await ApiHandler.put(`/${EXAM}/id?id=${id}/`, data);
  return res.data;
};

// ✅ Delete exam
export const deleteExam = async (id) => {
  const res = await ApiHandler.delete(`/${EXAM}/id?id=${id}/`);
  return res.data;
};
