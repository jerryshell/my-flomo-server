import logo from "../favicon.svg";

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
        {/* 空白区域，可以根据需要添加其他导航元素 */}
      </div>
    </div>
  );
};

export default Header;
