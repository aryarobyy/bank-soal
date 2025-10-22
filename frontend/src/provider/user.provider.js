import { USER } from "../core/constant";
import ApiHandler from "./api.handler";

export const register = async (data) => {
  const res = await ApiHandler.post(`${USER}/register`, data);
  return res.data;
};

export const login = async (data) => {
  const res = await ApiHandler.post(`${USER}/login`, data);
  return res;
};

export const refreshToken = async () => {
  const res = await ApiHandler.post(`${USER}/refresh`, {});
  return res;
};

export const updateUser = async (data, id) => {
  const formData = new FormData();

  for (const key in data) {
    formData.append(key, data[key]);
  }

  const res = await ApiHandler.put(`${USER}/${id}`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return res.data;
};

export const getUserById = async (id) => {
  const res = await ApiHandler.get(`${USER}/id?id=${id}`);
  return res.data;
};

export const getUsers = async () => {
  const res = await ApiHandler.get(`${USER}/`);
  return res.data;
};

export const getUserByName = async (name) => {
  const res = await ApiHandler.get(`${USER}/name?name=${name}`);
  return res;
};
