import logo from "../favicon.svg";
import ServerStatusCheck from "./ServerStatusCheck";

const Header = () => {
  return (
    <div className="navbar bg-base-100 shadow-lg">
      <div className="navbar-start">
        <div className="flex items-center space-x-2">
          <img src={logo} alt="logo" className="w-8 h-8" />
          <span className="text-xl font-bold">My Flomo</span>
        </div>
      </div>

      <div className="navbar-end">
        <ServerStatusCheck />
      </div>
    </div>
  );
};

export default Header;
