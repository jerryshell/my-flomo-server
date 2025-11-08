import api from "./api";

const memoApi = {
  list: () => {
    return api.get("/memo/list");
  },
  create: (data: { content: string }) => {
    return api.post("/memo/create", data);
  },
  update: (data: { id: string; content: string }) => {
    return api.post("/memo/update", data);
  },
  deleteById: (id: string) => {
    return api.post(`/memo/delete/id/${id}`);
  },
};

export default memoApi;
