import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
});

api.interceptors.request.use(async (config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.set("token", token);
  }
  return config;
});

api.interceptors.response.use(async (response) => {
  const token = localStorage.getItem("token");
  if (response.data.code === 401 && token) {
    localStorage.removeItem("email");
    localStorage.removeItem("token");
    localStorage.removeItem("expiresAt");
    window.location.href = "/";
    return Promise.reject(response);
  }
  return response;
});

export default api;
