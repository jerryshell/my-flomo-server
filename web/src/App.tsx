import "./App.css";
import { useEffect } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import Footer from "./components/Footer";
import LoginPage from "./pages/LoginPage";
import Header from "./components/Header";
import memoApi from "./api/memoApi";
import HomePage from "./pages/HomePage";
import { useRecoilValue, useSetRecoilState } from "recoil";
import { atoms } from "./atoms/atoms";

function App() {
  const token = useRecoilValue(atoms.token);
  const setMemoList = useSetRecoilState(atoms.memoList);

  const navigate = useNavigate();

  const fetchMemoList = () => {
    return memoApi.list().then((response) => {
      const success = response.data.success;
      if (success) {
        const memoList = response.data.data;
        setMemoList(memoList);
      }
    });
  };

  useEffect(() => {
    if (token) {
      fetchMemoList().then(() => navigate("/home"));
    } else {
      navigate("/login");
    }
  }, [token]);

  return (
    <>
      <Header />

      <Routes>
        <Route path="/login" element={<LoginPage />} />

        {token && (
          <Route
            path="/home"
            element={<HomePage fetchMemoList={fetchMemoList} />}
          />
        )}

        <Route path="*" element={<>404</>} />
      </Routes>

      <Footer />
    </>
  );
}

export default App;
