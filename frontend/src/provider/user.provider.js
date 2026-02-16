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
    const value = data[key];

    if (value === null) {
  
      formData.append(key, "");
    } else if (value !== undefined) {
   
      formData.append(key, value);
    }
  
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

export const getUsers = async (limit = 10, offset = 0) => {

  const res = await ApiHandler.get(
    `${USER}/?limit=${limit}&offset=${offset}`
  );
  

  return res.data.data;
};

export const getUserByName = async (name) => {
  const res = await ApiHandler.get(`${USER}/name?name=${name}`);
  return res;
};

export const deleteUser = async (id) => {
 
  const res = await ApiHandler.delete(`${USER}/${id}`);
  return res.data;
};


export const getUsersByRole = async (role, limit = 10, offset = 0) => {

  const params = new URLSearchParams();
  params.append('role', role);
  params.append('limit', limit);
  params.append('offset', offset);


  const res = await ApiHandler.get(`${USER}/role?${params.toString()}`);

  return res.data.data;
};

export const changeRole = async (id, data) => {
  const res = await ApiHandler.put(`${USER}/role?id=${id}`, data, {
    headers: { "Content-Type": "application/json" },
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

  const res = await ApiHandler.put(`${USER}/password?id=${id}`, {
    new_password: newPassword, 
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



export const generateUsers = async (prefix, start, end, academic_year) => {

  const params = new URLSearchParams();
  params.append('prefix', String(prefix));
  params.append('start', String(start));
  params.append('end', String(end));


  const requestBody = {
    academic_year: String(academic_year),
  };


  const res = await ApiHandler.post(
    `${USER}/generate?${params.toString()}`, 
    requestBody                           

  );
  return res.data;
};

export const logoutUser = async () => {
  try {
    await ApiHandler.post(`${USER}/logout`);
  } catch (error) {
    console.warn("Logout backend gagal (mungkin token sudah expired), lanjut clear lokal.", error);
  }
};