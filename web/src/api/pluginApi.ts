import api from "./api";

const pluginApi = {
  getToken: () => {
    return api.get("/plugin/getToken");
  },
  createToken: () => {
    return api.post("/plugin/createToken");
  },
};

export default pluginApi;
