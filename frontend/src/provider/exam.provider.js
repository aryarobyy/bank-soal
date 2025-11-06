import { EXAM } from "../core/constant";
import ApiHandler from "./api.handler";

export const getAllExam = async () => {
  const res = await ApiHandler.get(`/${EXAM}/`);
  return res.data;
};

export const getExamById = async (id) => {
  const res = await ApiHandler.get(`/${EXAM}/id?id=${id}`);
  return res.data;
};

export const createExam = async (data) => {
  const res = await ApiHandler.post(`/${EXAM}/`, data);
  return res.data;
};

export const updateExam = async (id, data) => {
  const res = await ApiHandler.put(`/${EXAM}/id?id=${id}`, data);
  return res.data;
};

export const deleteExam = async (id) => {
  const res = await ApiHandler.delete(`/${EXAM}/id?id=${id}`);
  return res.data;
};