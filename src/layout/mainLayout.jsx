import Navbar from "../components/navbar";
import Footbar from "../components/footbar";
import { useEffect, useState } from "react";
import { getLatestEvents } from "../api/api";
export default function MainLayout({ children, highligh }) {
  const [data, setData] = useState([]);
  const [event, setEvent] = useState({});
  const [alert, setAlert] = useState(0);

  async function loadEvents() {
    try {
      const event = await getLatestEvents();
      setData(event.data);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    const initialLoad = async () => {
      await Promise.all([loadEvents()]);
    };
    initialLoad();

    const interval = setInterval(async () => {
      await loadEvents();
    }, 1000);

    return () => {
      clearInterval(interval);
    };
  }, []);

  useEffect(() => {
    if (data.length > 0) {
      for (const item of data) {
        setEvent((prev) => ({
          ...prev,
          [item["device-id"]]: item["event"],
        }));
      }
    }
  }, [data]);

  useEffect(() => {
    if (!event) return;
    const count = data.filter((item) => {
      const currentEvent = event[item["device-id"]];
      return currentEvent !== undefined && currentEvent !== "RESET";
    }).length;

    setAlert(count);
  }, [event]);
  return (
    <>
      <Navbar alerts={alert} />
      <main>{children}</main>
      <Footbar highligh={highligh} />
    </>
  );
}
