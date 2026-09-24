// import
import { MdMenu } from "react-icons/md";

export default function Navbar() {
  const BodyNavbar = [
    { label: "ALL" },
    { label: "NURSECALL" },
    { label: "FIREALARM" },
  ];

  const Footer = [{ label: "CONTACT" }];

  return (
    <nav className="fixed w-screen h-12  mt-3 flex items-center px-4 font-serif">
      {/* top left */}
      <div className="flex-1 items-center">
        <div className="flex">
          <div className="bg-gray-700 rounded-md px-4 py-3 mx-3 flex">
            <span className="cursor-pointer mx-1 text-gray-100 ">
              PORTOFOLIO
            </span>
            <p
              className="mx-1 rounded-sm bg-slate-300 cursor-pointer px-2 hover:bg-slate-100 hover:text-blue-600
              hover:animate-[hop_300ms_ease-out] transition-all duration-300 hover:text-blue-600"
            >
              ABOUT
            </p>
            <p className="mx-1 rounded-sm bg-slate-300 cursor-pointer px-2 hover:bg-slate-100 hover:text-blue-600 hover:animate-[hop_300ms_ease-out] transition-all duration-300 hover:text-blue-600">
              WORKS
            </p>
          </div>
        </div>
      </div>
      {/* top mid */}
      <div className="flex-1 items-center ">
        <div className=" flex justify-center rounded-sm  ">
          <div className="rounded-sm bg-gray-700  px-4 py-3 flex">
            {BodyNavbar.map((item, index) => {
              return (
                <p
                  key={index}
                  className=" mx-1 rounded-sm bg-slate-300 cursor-pointer px-2 hover:bg-slate-100 hover:text-blue-600
              hover:animate-[hop_300ms_ease-out] transition-all duration-300 hover:text-blue-600"
                >
                  {item.label}
                </p>
              );
            })}
          </div>
        </div>
      </div>
      {/* adjustable */}
      <div className="flex-1 "></div>
    </nav>
  );
}
