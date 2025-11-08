import api from "./api";

const authApi = {
  deleteMyAccount: () => {
    return api.post("/deleteMyAccount");
  },
};

export default authApi;
