import { QUESTION } from "../core/constant";
import ApiHandler from "./api.handler";

export const getAllQuestions = async () => {
  const res = await ApiHandler.get(`/${QUESTION}/`);
  return res.data;
};

export const createQuestionWithOptions = async (data) => {
  const res = await ApiHandler.post(`${QUESTION}/options`, data);
  return res.data;
};

export const deleteQuestion = async (id) => {
  const res = await ApiHandler.delete(`/${QUESTION}/id?id=${id}`);
  return res.data;
};

export const getQuestionById = async (id) => {
  const res = await ApiHandler.get(`/${QUESTION}/id?id=${id}`);
  return res.data;
};

export const updateQuestion = async (id, data) => {
  const res = await ApiHandler.put(`${QUESTION}/${id}`, data);
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
  return res.data;
};

export const getQuestionsByDiff = async (difficulty) => {
  const res = await ApiHandler.get(`${QUESTION}/diff?diff=${difficulty}`);
  return res.data;
};

export const getQuestionsByCreator = async (creatorId) => {
  const res = await ApiHandler.get(`${QUESTION}/creator?creator_id=${creatorId}`);
  return res.data;
};
