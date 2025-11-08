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
      <div className="grid grid-cols-1 xl:grid-cols-4 gap-6">
        {/* 左侧功能区域 */}
        <div className="xl:col-span-1 space-y-6">
          {/* 服务器状态卡片 */}
          <div className="card bg-base-100 shadow-lg border border-base-300/50">
            <div className="card-body p-6">
              <div className="flex items-center space-x-3 mb-4">
                <div className="w-3 h-3 bg-info rounded-full"></div>
                <h2 className="card-title text-lg font-semibold">服务器状态</h2>
              </div>
              <ServerStatusCheck />
            </div>
          </div>

          {/* 记录想法卡片 */}
          <div className="card bg-base-100 shadow-lg border border-base-300/50">
            <div className="card-body p-6">
              <div className="flex items-center space-x-3 mb-4">
                <div className="w-3 h-3 bg-success rounded-full"></div>
                <h2 className="card-title text-lg font-semibold">记录想法</h2>
              </div>
              <MemoCreate fetchMemoList={props.fetchMemoList} />
            </div>
          </div>

          {/* 导入导出卡片 */}
          <div className="card bg-base-100 shadow-lg border border-base-300/50">
            <div className="card-body p-6">
              <div className="flex items-center space-x-3 mb-4">
                <div className="w-3 h-3 bg-warning rounded-full"></div>
                <h2 className="card-title text-lg font-semibold">数据管理</h2>
              </div>
              <div className="space-y-4">
                <FlomoImport fetchMemoList={props.fetchMemoList} />
                <CsvExport />
                <CsvImport fetchMemoList={props.fetchMemoList} />
              </div>
            </div>
          </div>

          {/* 用户设置卡片 */}
          <div className="card bg-base-100 shadow-lg border border-base-300/50">
            <div className="card-body p-6">
              <div className="flex items-center space-x-3 mb-4">
                <div className="w-3 h-3 bg-primary rounded-full"></div>
                <h2 className="card-title text-lg font-semibold">用户设置</h2>
              </div>
              <div className="space-y-4">
                <UserPasswordUpdate />
                <UserSettings />
                <PluginToken />

                <div className="divider"></div>

                <button
                  className="btn btn-outline btn-error w-full gap-2"
                  onClick={logout}
                >
                  <svg
                    className="w-4 h-4"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                    />
                  </svg>
                  登出
                </button>
              </div>
            </div>
          </div>

          {/* 危险区域卡片 */}
          <div className="card bg-base-100 shadow-lg border-2 border-error">
            <div className="card-body p-6">
              <div className="flex items-center space-x-3 mb-4">
                <div className="w-3 h-3 bg-error rounded-full"></div>
                <h2 className="card-title text-lg font-semibold text-error">
                  危险区域
                </h2>
              </div>
              <DangerousArea logout={logout} />
            </div>
          </div>
        </div>

        {/* 右侧想法列表 */}
        <div className="xl:col-span-3">
          <div className="card bg-base-100 shadow-lg border border-base-300/50">
            <div className="card-body p-6">
              <div className="flex items-center space-x-3 mb-6">
                <div className="w-3 h-3 bg-secondary rounded-full"></div>
                <h2 className="card-title text-xl font-semibold">我的想法</h2>
              </div>
              <MemoList />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;
