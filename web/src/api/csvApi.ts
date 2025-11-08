import api from "./api";

const csvApi = {
  csvImport: (data: FormData) => {
    return api.post("/csvImport", data);
  },
};

export default csvApi;
