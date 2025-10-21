// src/provider/api.handler.js
import axios from "axios";
import { API_BASE_URL } from "../core/constant";

const ApiHandler = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  withCredentials: true,
});

ApiHandler.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    const skipAuth = config.url.includes('/login') || config.url.includes('/refresh');
    if (token && !skipAuth) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);


ApiHandler.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (!error.response) return Promise.reject(error);
    const originalRequest = error.config;
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        const refreshResponse = await axios.post(
          `${API_BASE_URL}/user/refresh`,
          {},
          { withCredentials: true }
        );
        const newAccessToken = refreshResponse.data.data;
        localStorage.setItem("token", newAccessToken);
        originalRequest.headers["Authorization"] = `Bearer ${newAccessToken}`;
        return ApiHandler(originalRequest);
      } catch (refreshError) {
        console.error("Refresh token failed:", refreshError);
        localStorage.removeItem("token");
      }
    }
    return Promise.reject(error.response?.data?.error || error);
  }
);

export default ApiHandler;
