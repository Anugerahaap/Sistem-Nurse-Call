import React, { useState, useEffect, useEffectEvent } from "react";
import { TbAlertCircle } from "react-icons/tb";
import { TfiMore, TfiRssAlt } from "react-icons/tfi";
import { updateDevice } from "../api/api";

export default function Card({
  rn = "",
  deviceID = "",
  bg = {},
  resetFn,
  onChangePTN,
  onChangeDr,
  btnSave,
  patientName,
  displayPtn = localStorage.getItem(deviceID + "ptn")
    ? localStorage.getItem(deviceID + "ptn")
    : "",
  doctorName,
  displayDr = localStorage.getItem(deviceID + "dr")
    ? localStorage.getItem(deviceID + "dr")
    : "",
}) {
  const [openEdit, setOpenEdit] = useState(false);
  const [roomName, setRoomName] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault;

    const params = new URLSearchParams({
      action: "update-room",
    });

    const body = {
      "device-id": deviceID,
      "room-name": roomName,
    };
    try {
      await updateDevice(params, body);
    } catch (error) {
      console.error("Gagal update room name:", error);
    }
  };

  useEffect(() => {
    if (patientName === undefined) {
      onChangePTN({
        target: {
          value: displayPtn,
        },
      });
    }
    if (doctorName === undefined) {
      onChangeDr({
        target: {
          value: displayDr,
        },
      });
    }
  }, [patientName, displayPtn, doctorName, displayDr]);

  return (
    <>
      {openEdit && (
        <div className="fixed inset-0 z-[100] flex items-center justify-center ">
          {/* background */}
          <div
            className="absolute inset-0 bg-black/40 backdrop-blur-sm"
            onClick={() => setOpenEdit(false)}
          ></div>
          {/* form */}
          <div className="relative w-[450px] bg-white rounded-xl shadow-2xl p-6">
            <h2 className="text-3xl font-semibold mb-5 justify-between flex">
              Edit Informasi Ruangan
              <span className="text-xs"> {deviceID}</span>
            </h2>{" "}
            {/* Nama Ruangan */}
            <div className="mb-4">
              <label className="block mb-1 font-medium" id={deviceID}>
                Nama Ruangan
              </label>

              <input
                type="text"
                value={roomName ?? ""}
                onChange={(e) => setRoomName(e.target.value)}
                className="w-full border rounded-md px-3 py-2 outline-none focus:ring-2 focus:ring-sky-400"
              />
            </div>
            {/* Nama Pasien */}
            <div className="mb-4">
              <label className="block mb-1 font-medium">Nama Pasien</label>

              <input
                type="text"
                onChange={onChangePTN}
                value={patientName ?? ""}
                className="w-full border rounded-md px-3 py-2 outline-none focus:ring-2 focus:ring-sky-400"
              />
            </div>
            {/* Dokter */}
            <div className="mb-6">
              <label className="block mb-1 font-medium">Dokter</label>

              <input
                type="text"
                onChange={onChangeDr}
                value={doctorName ?? ""}
                className="w-full border rounded-md px-3 py-2 outline-none focus:ring-2 focus:ring-sky-400"
              />
            </div>{" "}
            {/* Button */}
            <div className="flex justify-end gap-3">
              <button
                onClick={() => {
                  setOpenEdit(false);
                }}
                className="px-4 py-2 rounded-md bg-slate-200 hover:bg-slate-300"
              >
                Batal
              </button>

              <button
                onClick={async (e) => {
                  btnSave();
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
              <TfiMore size={20} className="mx-3" /> {/* tombol edit */}
            </button>
            <p className="text-2xl mx-2  flex-1 text-center">{rn}</p>

            <TbAlertCircle
              size={20}
              className={`mx-1 items-end ${(bg && bg[deviceID] === "DARURAT") || (bg[deviceID] === "INFUS" ? "visible" : "invisible")} `}
            />
          </div>
          <div
            onClick={resetFn}
            className={` ${bg && bg[deviceID] === "INFUS" ? "bg-yellow-200/70 cursor-pointer animate-pulse" : bg[deviceID] === "DARURAT" ? "bg-red-500/70 cursor-pointer animate-pulse" : "bg-slate-700/10"} rounded-b-md  text-center overflow-hidden `}
          >
            <p className="text-2xl mt-2">{displayPtn}/</p>
            <p className="text-xl">{displayDr} </p>
          </div>
        </div>
      </div>
      {/* end here */}
    </>
  );
}
