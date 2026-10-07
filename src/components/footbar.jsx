import { useState } from "react";
import { LuMail, LuMailWarning, LuMailMinus } from "react-icons/lu";
import { Link } from "react-router-dom";

export default function Footbar({
  list = ["Home", "Setting", "Logs"],
  highligh = "",
}) {
  const [openMenu, setOpenMenu] = useState(false);

  return (
    <>
      {/* footbar */}
      <footer className="text-xl grid grid-cols-3 shadow-md  fixed bottom-0 left-0 w-full z-50  px-3 m-2 h-16">
        <div className="bg-sky-100 rounded-l-md"></div>
        <div className="bg-sky-100 flex items-center justify-center">
          <p>RS Permata Hati</p>
        </div>
        <div className="bg-sky-100 rounded-r-md grid grid-cols-3">
          <div className="relative flex items-center col-start-3">
            {/* popup */}
            {openMenu && (
              <div
                className={`absolute bottom-full right-5 h-125 mb-2 w-64 bg-slate-200 rounded-md shadow-lg p-3
                            transition-all duration-300
                    ${
                      openMenu
                        ? "opacity-100 translate-y-0"
                        : "opacity-0 translate-y-2 pointer-events-none"
                    }`}
              >
                <p
                  onClick={() => setOpenMenu(!openMenu)}
                  className="font-semibold text-center shadow-sm cursor-pointer"
                >
                  Menu
                </p>
                {list.map((item) => (
                  <Link
                    key={item}
                    to={`/${item}`}
                    className={`block  w-full text-left p-2 hover:bg-slate-100 ${highligh === item && "rounded-md bg-slate-100"}`}
                  >
                    {item}
                  </Link>
                ))}
              </div>
            )}

            {/* Tombol */}
            <button
              onClick={() => setOpenMenu(!openMenu)}
              className="cursor-pointer transition-transform duration-200 hover:scale-110 active:scale-90"
            >
              {openMenu ? <LuMailMinus size={25} /> : <LuMail size={25} />}
            </button>
          </div>
        </div>
      </footer>
    </>
  );
}
