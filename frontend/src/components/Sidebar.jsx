import { React, useState } from "react";
import { Link } from "react-router-dom";
import {
  MdMonitor,
  MdHome,
  MdOutlineMenuOpen,
  MdSettings,
  MdOutlineMenu,
} from "react-icons/md";

export default function Sidebar({ preview }) {
  const [isOpen, setIsOpen] = useState(true);

  console.log(preview);

  const menuItems = [
    { label: "Home", icon: <MdHome /> },
    { label: "Dashboard", icon: <MdMonitor /> },
    { label: "Settings", icon: <MdSettings /> },
  ];
  return (
    <>
      <nav
        className={`shadow-md ${preview ? "h-screen" : "h-[500px]"}   bg-slate-400 p-1 flex flex-col duration-700 ${isOpen ? "w-40" : "w-16"} `}
      >
        {/* Header */}
        <div
          className={
            "border px-3 py-2 h-10 flex justify-between items-center rounded-md "
          }
        >
          <MdHome
            className={`
            duration-300
            ${isOpen ? "translate-x-0" : "-translate-x-2 opacity-0"}`}
            size={25}
          />
          <div>
            <MdOutlineMenuOpen
              size={25}
              className={`
            cursor-pointer    
            duration-300
            ${isOpen ? "rotate-0" : "rotate-180"}`}
              onClick={() => setIsOpen(!isOpen)}
            />
          </div>
        </div>
        {/* Body */}
        <ul className="mt-2 flex-1">
          {menuItems.map((item, index) => {
            return (
              <li
                key={index}
                className="flex items-center gap-2 px-3 py-2 rounded-md hover:bg-slate-300 cursor-pointer relative group"
              >
                <div>{item.icon && <span>{item.icon}</span>}</div>
                <p className="duration-700 overflow-hidden">
                  {isOpen && item.label}
                </p>
                <p
                  className={`${isOpen && "hidden"} m-1 absolute left-32 shadow-md rounded-md
                  w-0  
                  p-0
                  overflow-hidden
                  group-hover:w-fit
                  group-hover:p-2
                  group-hover:left-14
                  group-hover:duration-300
                  group-hover:bg-slate-400
                  
                  `}
                >
                  {item.label}
                </p>
              </li>
            );
          })}
        </ul>
      </nav>
    </>
  );
}
