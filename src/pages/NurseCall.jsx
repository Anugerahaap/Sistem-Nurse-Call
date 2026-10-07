import MainLayout from "../layout/mainLayout";
import BodyCard from "../components/body-card";
import { getDevices, getLatestEvents, getLastEvent } from "../api/api";
import { useEffect, useState } from "react";

/**
 * Page utama / home page
 */
export default function NurseCall() {
  const [devices, setDevices] = useState([]);
  const [events, setEvents] = useState({});

  async function loadEvents(deviceID) {
    try {
      const event = await getLastEvent(deviceID);
      setEvents((prev) => ({
        ...prev,
        [deviceID]: event.data?.event,
      }));
    } catch (error) {
      console.log(error);
    }
  }
  async function loadDevices() {
    try {
      const devices = await getDevices();
      setDevices(devices.data);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    const initialLoad = async () => {
      await Promise.all([loadDevices()]);
    };
    initialLoad();
  }, []);

  useEffect(() => {
    const loadAllEvents = async () => {
      for (const item of devices) {
        await loadEvents(item["device-id"]);
      }
    };
    if (devices.length > 0) {
      loadAllEvents();
    }
  }, [devices]);

  return (
    <MainLayout highligh={"Home"}>
      <div className="p-20 h-screen w-screen flex justify-center items-center flex-wrap overflow-y-auto">
        {devices.map((item, key) => {
          const deviceID = item["device-id"];
          const roomName = item["room-name"];
          return (
            <BodyCard
              key={deviceID}
              deviceID={deviceID}
              roomname={roomName}
              event={events[deviceID]}
              refreshfn={loadDevices}
            />
          );
        })}
        <div className=""></div>
      </div>
    </MainLayout>
  );
}
