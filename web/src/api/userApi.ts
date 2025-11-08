import api from "./api";

const userApi = {
  updatePassword: (data: { password: string }) => {
    return api.post("/user/updatePassword", data);
  },
};

export default userApi;
