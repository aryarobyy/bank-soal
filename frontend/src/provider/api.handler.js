import axios from "axios";
import { API_BASE_URL } from "../core/constant";

const ApiHandler = axios.create({
  baseURL: API_BASE_URL,
  timeout: 30000,
  withCredentials: true,
});

ApiHandler.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token) config.headers["Authorization"] = `Bearer ${token}`;
    return config;
  },
  (error) => Promise.reject(error),
);

ApiHandler.interceptors.response.use(
  (response) => {
    const newToken = response.headers["x-new-access-token"];
    if (newToken) {
      localStorage.setItem("token", newToken);
    }
    return response;
  },
  (error) => {
    if (!error.response) return Promise.reject(error);

    if (error.response.status === 401) {
      localStorage.removeItem("token");
    }

    return Promise.reject(error.response?.data?.error || error);
  },
);

export default ApiHandler;
