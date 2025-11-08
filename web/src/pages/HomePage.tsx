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
import ServerStatusCheck from "../components/ServerStatusCheck";

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
    <div className="container mx-auto px-4 py-8">
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* 左侧功能区域 */}
        <div className="lg:col-span-1 space-y-6">
          {/* 服务器状态卡片 */}
          <div className="card bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">服务器状态</h2>
              <ServerStatusCheck />
            </div>
          </div>

          {/* 创建备忘录卡片 */}
          <div className="card bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">创建备忘录</h2>
              <MemoCreate fetchMemoList={props.fetchMemoList} />
            </div>
          </div>

          {/* 导入导出卡片 */}
          <div className="card bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">数据管理</h2>
              <div className="space-y-4">
                <FlomoImport fetchMemoList={props.fetchMemoList} />
                <CsvExport />
                <CsvImport fetchMemoList={props.fetchMemoList} />
              </div>
            </div>
          </div>

          {/* 用户设置卡片 */}
          <div className="card bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">用户设置</h2>
              <div className="space-y-4">
                <UserPasswordUpdate />
                <UserSettings />
                <PluginToken />

                <div className="divider"></div>

                <button
                  className="btn btn-outline btn-error w-full"
                  onClick={logout}
                >
                  登出
                </button>
              </div>
            </div>
          </div>

          {/* 危险区域卡片 */}
          <div className="card bg-base-100 shadow-xl border-2 border-error">
            <div className="card-body">
              <h2 className="card-title text-error">危险区域</h2>
              <DangerousArea logout={logout} />
            </div>
          </div>
        </div>

        {/* 右侧备忘录列表 */}
        <div className="lg:col-span-2">
          <div className="card bg-base-100 shadow-xl">
            <div className="card-body">
              <h2 className="card-title">我的备忘录</h2>
              <MemoList />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;
