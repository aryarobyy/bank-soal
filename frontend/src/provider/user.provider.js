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
    headers: { "Content-Type": "multipart/form-data" },
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

// ## PERBAIKAN 1: Menggunakan Path Parameter ( /user/:id ) ##
export const deleteUser = async (id) => {
  // Menggunakan format path parameter: /user/11
  const res = await ApiHandler.delete(`${USER}/${id}`);
  return res.data;
};

/**
 * provider baru
 */
export const getUsersByRole = async (role) => {
  const res = await ApiHandler.get(`${USER}/role?role=${role}`);
  return res.data;
};

export const changeRole = async (userId, adminId, role) => {
  const res = await ApiHandler.put(`${USER}/role?id=${userId}`, {
    admin_id: adminId,
    role: role,
  });
  return res.data;
};

/**
 * ## PERBAIKAN 2: Menambahkan admin_id untuk melewati 'old_password' ##
 * @param {number} id - ID user (masuk ke query param)
 * @param {string} newPassword - Password baru (masuk ke body)
 * @param {number} adminId - ID Admin (masuk ke body)
 */
export const changePassword = async (id, newPassword, adminId) => {
  // 1. Tetap mengirim sebagai JSON
  // 2. Mengganti key 'password' menjadi 'new_password'
  const res = await ApiHandler.put(`${USER}/password?id=${id}`, {
    new_password: newPassword, // <-- INI PERBAIKANNYA
    admin_id: adminId, 
  });
  return res.data;
};

export const getUserByEmail = async (email) => {
  const res = await ApiHandler.get(`${USER}/email?email=${email}`);
  return res.data;
};

export const getUserByNim = async (nim) => {
  const res = await ApiHandler.get(`${USER}/nim?nim=${nim}`);
  return res.data;
};

export const checkUser = async (id) => {
  const res = await ApiHandler.get(`${USER}/check?id=${id}`);
  return res.data;
};