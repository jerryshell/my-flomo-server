import api from "./api";

const uploadApi = {
  upload: (data: FormData) => {
    return api.post("/upload", data);
  },
};

export default uploadApi;
