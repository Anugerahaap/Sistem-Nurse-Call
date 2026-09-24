import React, { useState, useEffect } from "react";

//  192.168.1.8
const NurseCallSimulator = () => {
  // State untuk status tombol dan LED
  const [btn1Pressed, setBtn1Pressed] = useState(false); // tombol merah
  const [btn2Pressed, setBtn2Pressed] = useState(false); // tombol kuning
  const [led1On, setLed1On] = useState(false); // led merah
  const [led2On, setLed2On] = useState(false); // led kuning

  // logika simulasi untuk tombol merah
  useEffect(() => {
    let interval;
    if (btn1Pressed) {
      setBtn2Pressed(false);
      setLed2On(false);
      interval = setInterval(() => {
        setLed1On((prev) => !prev);
      }, 500);
    } else {
      setLed1On(false);
    }

    return () => clearInterval(interval);
  }, [btn1Pressed]);

  // logika simulasi untuk tombol kuning
  useEffect(() => {
    let interval;
    if (btn2Pressed) {
      setBtn1Pressed(false);
      setLed1On(false);
      interval = setInterval(() => {
        setLed2On((prev) => !prev);
      }, 500);
    } else {
      setLed2On(false);
    }

    return () => clearInterval(interval);
  }, [btn2Pressed]);

  return (
    <div className="p-6 bg-gray-100 rounded-xl shadow-inner max-w-md mx-auto my-8 border border-gray-300">
      <h3 className="text-xl font-bold text-gray-800 mb-4 text-center">
        Simulasi Nurse Call ESP32
      </h3>
      {/* Area LED */}
      <div
        className={`flex justify-center gap-8 mb-8 p-4 rounded-lg shadow-sm ${btn1Pressed || btn2Pressed ? "cursor-pointer" : "cursor-default"}${
          led2On
            ? " shadow-[0_0_20px_rgba(234,179,8,0.8)] bg-yellow-200 "
            : "bg-white "
        } ${
          led1On
            ? "shadow-[0_0_20px_rgba(239,68,68,0.8)] bg-red-200 "
            : "bg-white "
        }`}
        onClick={() => {
          setBtn1Pressed(false);
          setBtn2Pressed(false);
          setLed1On(false);
          setLed2On(false);
        }}
      >
        <div className="flex flex-col items-center  ">
          <div
            className={`w-16 h-16 rounded-full border-4 transition-colors duration-300 ${
              led1On
                ? "bg-red-500 border-red-700 shadow-[0_0_20px_rgba(239,68,68,0.8)]"
                : "bg-red-100 border-red-300"
            }`}
          ></div>
          <span className="mt-2 text-sm font-medium text-gray-600">
            LED Darurat
          </span>
        </div>
        <div className="flex flex-col items-center">
          <div
            className={`w-16 h-16 rounded-full border-4 transition-colors duration-300 ${
              led2On
                ? "bg-yellow-500 border-yellow-700 shadow-[0_0_20px_rgba(234,179,8,0.8)]"
                : "bg-yellow-100 border-yellow-300"
            }`}
          ></div>
          <span className="mt-2 text-sm font-medium text-gray-600">
            LED Infus
          </span>
        </div>
      </div>

      {/* Area Tombol */}
      <div className="flex justify-center gap-6">
        <button
          onClick={() => setBtn1Pressed(true)}
          className={`px-6 py-4 rounded-lg font-bold text-white shadow-md transform transition active:scale-95 ${
            btn1Pressed
              ? "bg-gray-800 translate-y-1"
              : "bg-gray-600 hover:bg-gray-700"
          }`}
        >
          Tombol Darurat
        </button>

        <button
          onClick={() => setBtn2Pressed(true)}
          className={`px-6 py-4 rounded-lg font-bold text-white shadow-md transform transition active:scale-95 ${
            btn2Pressed
              ? "bg-gray-800 translate-y-1"
              : "bg-gray-600 hover:bg-gray-700"
          }`}
        >
          Tombol Infus
        </button>
      </div>

      <div className="mt-6 text-xs text-gray-500 text-center bg-gray-200 p-2 rounded">
        <p className="">
          💡 <strong>Cara Pakai:</strong> Klik & Tahan tombol untuk
          mensimulasikan tekanan saklar fisik.
        </p>
        <p className="">
          💡 <strong>Cara Matikan:</strong> Klik pada bagian kotak led.
        </p>
      </div>
    </div>
  );
};

export default NurseCallSimulator;
