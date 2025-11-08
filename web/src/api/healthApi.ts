import api from "./api";

const healthApi = {
  health: () => {
    return api.get("/health");
  },
};

export default healthApi;
