import { useEffect, useRef, useState } from "react";
import { TfiMore, TfiRssAlt } from "react-icons/tfi";
import alarmSound from "../assets/Hidup-jokowi.mp3";

export default function Navbar({
  tittle = "Nurse Call",
  company = "RS Permata Hati",
  description = "Monitoring station",
  alerts = 0,
}) {
  const [status, setStatus] = useState(true);
  const [sound, setSound] = useState(false);

  const alarmRef = useRef(new Audio(alarmSound));
  const playAlarm = () => {
    const alarm = alarmRef.current;

    alarm.loop = true;
    alarm.currentTime = 0;

    alarm.play();
  };

  const stopAlarm = () => {
    const alarm = alarmRef.current;

    alarm.pause();
    alarm.currentTime = 0;
  };
  useEffect(() => {
    if (alerts > 0) {
      setSound(true);
    } else {
      setSound(false);
    }
  }, [alerts]);

  useEffect(() => {
    if (sound) {
      playAlarm();
    } else {
      stopAlarm();
    }
  }, [sound]);

  useEffect(() => {
    function handleOnline() {
      setStatus(true);
    }

    function handleOffline() {
      setStatus(false);
    }
    window.addEventListener("online", handleOnline);
    window.addEventListener("offline", handleOffline);

    return () => {
      window.removeEventListener("online", handleOnline);
      window.removeEventListener("offline", handleOffline);
    };
  }, []);
  return (
    <>
      {/* Navbar */}
      <div
        className=" w-screen h-16 m-2 px-4 font-serif grid grid-cols-3 text-2xl shadow-md text-slate-800 fixed "
        onClick={() => {
          alerts++;
        }}
      >
        <div className="bg-sky-100 grid  rounded-l-md ">
          <span className="mx-2">{tittle}</span>
          <span className="mx-2">{description}</span>
        </div>
        <div className="bg-sky-100 flex items-center justify-center ">
          <span className="p-1">{company}</span>
        </div>
        <div className="bg-sky-100 items-center rounded-r-md grid grid-cols-2 ">
          <p
            className={`mx-2 flex text-center ${status ? "text-green-600" : "text-red-600"} font-semibold text-shadow-sm`}
          >
            {status ? "ONLINE" : "OFFLINE"}
            <TfiRssAlt className="mx-3" />
          </p>
          <p className="mx-2  text-center ">
            {alerts < 1
              ? "No Active Call"
              : alerts == 1
                ? "1 Active Call"
                : `${alerts} Active Calls`}
          </p>
        </div>
      </div>
    </>
  );
}
