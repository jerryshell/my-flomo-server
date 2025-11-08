import api from "./api";

const authApi = {
  login: (data: { email: string; password: string }) => {
    return api.post("/auth/login", data);
  },
};

export default authApi;
