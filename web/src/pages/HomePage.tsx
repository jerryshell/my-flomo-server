import MemoCreate from "../components/MemoCreate";
import FlomoImport from "../components/FlomoImport";
import CsvExport from "../components/CsvExport";
import CsvImport from "../components/CsvImport";
import DangerousArea from "../components/DangerousArea";
import PluginToken from "../components/PluginToken";
import MemoList from "../components/MemoList";
import { useSetRecoilState } from "recoil";
import { atoms } from "../atoms/atoms";
import UserPasswordUpdate from "../components/UserPasswordUpdate";
import UserSettings from "../components/UserSettings";

const HomePage = (props: { fetchMemoList(): void }) => {
  const setEmail = useSetRecoilState(atoms.email);
  const setToken = useSetRecoilState(atoms.token);

  const logout = () => {
    setEmail("");
    setToken("");
    localStorage.removeItem("email");
    localStorage.removeItem("token");
    localStorage.removeItem("expiresAt");
  };

  return (
    <>
      <MemoCreate fetchMemoList={props.fetchMemoList} />

      <button onClick={logout} style={{ color: "#9E3B37" }}>
        登出
      </button>

      <FlomoImport fetchMemoList={props.fetchMemoList} />

      <CsvExport />

      <CsvImport fetchMemoList={props.fetchMemoList} />

      <DangerousArea logout={logout} />

      <PluginToken />

      <UserPasswordUpdate />

      <UserSettings />

      <MemoList />
    </>
  );
};

export default HomePage;
