import { useEffect, useState } from "react";
import { TfiMore, TfiRssAlt } from "react-icons/tfi";
import { TbAlertCircle } from "react-icons/tb";
import { updateRoomNames, registerNewEvent } from "../api/api";
export default function BodyCard({
  deviceID = "default-id",
  roomname = "default-name",
  event = "",
  refreshfn,
}) {
  const [openEdit, setOpenEdit] = useState(false);
  const [updateRoomName, setUpdateName] = useState(roomname);
  const [updateDoctorName, setUpdateDoctorName] = useState(
    getItems("doctor-" + deviceID) || "",
  );
  const [updatePatientName, setUpdatePatientName] = useState(
    getItems("patient-" + deviceID) || "",
  );

  //
  function storeItems(key = "", value = "") {
    localStorage.setItem(key, value);
  }
  function getItems(key = "") {
    return localStorage.getItem(key);
  }
  //

  const handleSubmit = async (e) => {
    e.preventDefault;
    try {
      const res = await updateRoomNames(deviceID, updateRoomName);
      await refreshfn();
    } catch (error) {
      console.error("gagal mengupdate room-name");
    }
  };

  const sendReset = async () => {
    try {
      const body = {
        "device-id": deviceID,
        event: "RESET",
      };
      await registerNewEvent(body);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      {/* Body  of content*/}

      {/* start here */}
      {openEdit && (
        <div className="fixed inset-0 z-[100] flex items-center justify-center ">
          {/* background */}
          <div
            className="absolute inset-0 bg-black/40 backdrop-blur-sm"
            onClick={() => {
              setOpenEdit(false);
              setUpdateName(roomname);
            }}
          ></div>
          {/* form */}
          <div className="relative w-[450px] bg-white rounded-xl shadow-2xl p-6">
            <h2 className="text-3xl font-semibold mb-5 justify-between flex">
              Edit Informasi Ruangan
              <span className="text-xs"> {deviceID}</span>
              {/* <span className="text-xs">test nama ruangan</span> */}
            </h2>{" "}
            {/* Nama Ruangan */}
            <div className="mb-4">
              <label className="block mb-1 font-medium" id={deviceID}>
                Nama Ruangan
                <input
                  id="room"
                  type="text"
                  value={updateRoomName}
                  onChange={(e) => setUpdateName(e.target.value)}
                  className="w-full border rounded-md px-3 py-2 outline-none focus:ring-2 focus:ring-sky-400"
                />
              </label>
            </div>
            {/* Nama Pasien */}
            <div className="mb-4">
              <label className="block mb-1 font-medium">
                Nama Pasien
                <input
                  id="patient"
                  type="text"
                  onChange={(e) => setUpdatePatientName(e.target.value)}
                  value={updatePatientName}
                  className="w-full border rounded-md px-3 py-2 outline-none focus:ring-2 focus:ring-sky-400"
                />
              </label>
            </div>
            {/* Dokter */}
            <div className="mb-6">
              <label className="block mb-1 font-medium">
                Dokter
                <input
                  id="doctor"
                  type="text"
                  onChange={(e) => setUpdateDoctorName(e.target.value)}
                  value={updateDoctorName}
                  className="w-full border rounded-md px-3 py-2 outline-none focus:ring-2 focus:ring-sky-400"
                />
              </label>
            </div>{" "}
            {/* Button */}
            <div className="flex justify-end gap-3">
              <button
                onClick={() => {
                  setOpenEdit(false);
                  setUpdateName(roomname);
                  setUpdateDoctorName(getItems("doctor-" + deviceID) || "");
                  setUpdatePatientName(getItems("patient-" + deviceID) || "");
                }}
                className="px-4 py-2 rounded-md bg-slate-200 hover:bg-slate-300"
              >
                Batal
              </button>

              <button
                onClick={async (e) => {
                  //   btnSave();
                  storeItems("doctor-" + deviceID, updateDoctorName);
                  storeItems("patient-" + deviceID, updatePatientName);
                  setOpenEdit(false);
                  handleSubmit(e);
                }}
                className="px-4 py-2 rounded-md bg-sky-500 text-white hover:bg-sky-600"
              >
                Simpan
              </button>
            </div>
          </div>
        </div>
      )}
      {/* start here */}
      <div className="bg-slate-200 m-2 w-[382px] h-[172px]  rounded-md ">
        <div className="h-full grid grid-rows-[1fr_3fr] ">
          <div className="rounded-md p-3 flex shadow-md">
            <button className="" onClick={() => setOpenEdit(true)}>
              <TfiMore size={20} className="mx-3" />
            </button>
            <p className="text-2xl mx-2  flex-1 text-center">{roomname}</p>

            <TbAlertCircle
              size={20}
              className={`mx-1 items-end ${(event && event === "DARURAT") || (event === "INFUS" ? "visible" : "invisible")} `}
            />
          </div>
          <div
            onClick={sendReset}
            className={` ${event && event === "INFUS" ? "bg-yellow-200/70 cursor-pointer animate-pulse" : event === "DARURAT" ? "bg-red-500/70 cursor-pointer animate-pulse" : "bg-slate-700/10"} rounded-b-md  text-center overflow-hidden `}
          >
            <p className="text-2xl mt-2">
              {getItems("patient-" + deviceID) || ""}/
            </p>
            <p className="text-xl">{getItems("doctor-" + deviceID) || ""} </p>
          </div>
        </div>
      </div>
      {/* end here */}
    </>
  );
}
