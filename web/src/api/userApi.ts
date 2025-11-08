import api from "./api";

const userApi = {
  updatePassword: (data: { password: string }) => {
    return api.post("/user/updatePassword", data);
  },
  updateSettings: (data: { dailyReviewEnabled: boolean }) => {
    return api.post("/user/updateSettings", data);
  },
};

export default userApi;
