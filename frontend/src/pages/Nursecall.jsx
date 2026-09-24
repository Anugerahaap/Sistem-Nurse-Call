import React, { useState, useEffect, use } from "react";
import { TbAlertCircle } from "react-icons/tb";
import { TfiMore, TfiRssAlt } from "react-icons/tfi";
import Card from "../components/Card";
import Sidebar from "../components/Sidebar";
import { LuMail, LuMailWarning, LuMailMinus } from "react-icons/lu";
import {
  getDevices,
  getDevicesID,
  getLatestEvents,
  getNurseEvent,
  getNurseEvents,
  registerNewEvents,
} from "../api/api";
import { data } from "react-router-dom";
import { useRef } from "react";

import alarmSound from "../assets/Hidup-jokowi.mp3";

export default function NurseCallPage() {
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
  const [sound, setSound] = useState(false);
  const [alerts, setAlerts] = useState(0);
  const [devices, setDevices] = useState([]);
  const [online, setOnline] = useState(navigator.onLine);
  const [openMenu, setOpenMenu] = useState(false);
  // pass patient name to Card
  const [patientName, setPtnName] = useState({});
  // pass doctor name to card
  const [doctorName, setDrName] = useState({});

  // latest events
  const [events, setEvents] = useState([]);
  const [eventsBg, setEventsBG] = useState({});

  async function loadEvents() {
    try {
      const event = await getLatestEvents();
      setEvents(event.data);
    } catch (error) {
      console.error("Gagal mengambil events:", error);
    }
  }
  async function loadDevices() {
    try {
      const data = await getDevices();
      setDevices(data.data);
    } catch (error) {
      console.error("Gagal mengambil devices:", error);
    }
  }

  useEffect(() => {
    const initialLoad = async () => {
      await Promise.all([loadEvents(), loadDevices()]);
    };

    initialLoad();
    const interval = setInterval(async () => {
      await loadDevices();
      await loadEvents();
    }, 800);
    function handleOnline() {
      setOnline(true);
    }

    function handleOffline() {
      setOnline(false);
    }

    window.addEventListener("online", handleOnline);
    window.addEventListener("offline", handleOffline);

    return () => {
      window.removeEventListener("online", handleOnline);
      window.removeEventListener("offline", handleOffline);
      clearInterval(interval);
    };
  }, []);

  useEffect(() => {
    if (!events) return;

    const alertCount = events.filter(
      (item) => item["event"] !== "RESET",
    ).length;
    setAlerts(alertCount);

    events.map((item) =>
      setEventsBG((prev) => ({
        ...prev,
        [item["device-id"]]: item["event"],
      })),
    );
  }, [events]);

  useEffect(() => {
    if (sound) {
      playAlarm();
    } else {
      stopAlarm();
    }
  }, [sound]);

  useEffect(() => {
    if (alerts > 0) {
      setSound(true);
    } else {
      setSound(false);
    }
  }, [alerts]);
  return (
    <>
      {/* Navbar */}
      <div className=" w-screen h-16 m-2 px-4 font-serif grid grid-cols-3 text-2xl shadow-md text-slate-800 fixed ">
        <div className="bg-sky-100 grid  rounded-l-md ">
          <span className="mx-2">Nurse Call</span>
          <span className="mx-2">Monitoring station</span>
        </div>
        <div className="bg-sky-100 flex items-center justify-center ">
          <span className="p-1">RS Permata Hati</span>
        </div>
        <div className="bg-sky-100 items-center rounded-r-md grid grid-cols-2 ">
          <p
            className={`mx-2 flex text-center  ${online ? "text-green-600" : "text-red-600"} font-semibold text-shadow-sm`}
          >
            {online ? "ONLINE" : "OFFLINE"}
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
      {/* Body  of content*/}

      <div className="p-20 h-screen w-screen grid grid-rows  overflow-y-auto">
        <div className="flex justify-center items-center flex-wrap">
          {/* start here */}
          {devices.map((item) => {
            const deviceID = item["device-id"];

            return (
              <Card
                rn={item["room-name"]}
                deviceID={deviceID}
                key={deviceID}
                onChangePTN={(e) => {
                  setPtnName((prev) => ({
                    ...prev,
                    [deviceID + "ptn"]: e.target.value,
                  }));
                }}
                patientName={patientName[deviceID + "ptn"]}
                onChangeDr={(e) => {
                  setDrName((prev) => ({
                    ...prev,
                    [deviceID + "dr"]: e.target.value,
                  }));
                }}
                doctorName={doctorName[deviceID + "dr"]}
                btnSave={async () => {
                  await loadEvents();
                  await loadDevices();
                  localStorage.setItem(
                    deviceID + "dr",
                    doctorName[deviceID + "dr"] ?? "",
                  );
                  localStorage.setItem(
                    deviceID + "ptn",
                    patientName[deviceID + "ptn"] ?? "",
                  );
                }}
                bg={eventsBg}
                resetFn={async () => {
                  try {
                    const body = {
                      "device-id": deviceID,
                      event: "RESET",
                    };
                    await registerNewEvents(body);
                    await loadEvents();
                  } catch (error) {
                    console.error(error);
                  }
                }}
              />
            );
          })}
          {/* end here */}
        </div>
      </div>

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

                <button className="block w-full text-left p-2 hover:bg-slate-100">
                  Riwayat Panggilan
                </button>

                <button className="block w-full text-left p-2 hover:bg-slate-100">
                  Pengaturan
                </button>
                <button className="block w-full text-left p-2 hover:bg-slate-100">
                  Logs
                </button>
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
