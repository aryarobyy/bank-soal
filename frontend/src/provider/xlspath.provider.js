// src/provider/xlspath.provider.js

import { XLSPATH } from "../core/constant"; // Anda perlu menambahkan XLSPATH = 'xlspath' ke file constant.js Anda
import ApiHandler from "./api.handler";


export const getAllXlsPaths = async (limit = 10, offset = 0) => {
  const params = new URLSearchParams();
  params.append('limit', limit);
  params.append('offset', offset);
  
  const res = await ApiHandler.get(`/${XLSPATH}/?${params.toString()}`);

  return res.data.data; 
};


export const getXlsPathById = async (id) => {
  const res = await ApiHandler.get(`/${XLSPATH}/id?id=${id}`);
  return res.data.data; 
};


export const deleteXlsPath = async (id) => {
  const res = await ApiHandler.delete(`/${XLSPATH}/?id=${id}`);
  return res.data;
};

export const downloadXlsFile = async (id) => {
  const res = await ApiHandler.get(
    `/${XLSPATH}/download?id=${id}`,
    {
      responseType: 'blob',
    }
  );
  return res; 
};