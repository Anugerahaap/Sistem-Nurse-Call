export const API_URL = "http://localhost:3000/api";

export async function callApi({
  url = "",
  method = "GET",
  body = null,
  headers = {},
}) {
  const response = await fetch(`${API_URL}/${url}`, {
    method,
    headers: { "Content-Type": "application/json", ...headers },
    body: body ? JSON.stringify(body) : null,
  });
  if (!response.ok) {
    throw new Error(`HTTP Error:${response.status}`);
  }
  return response.json();
}

// devices func

export async function getDevices() {
  return callApi({ url: "/devices", method: "GET" });
}

export async function getDevicesID(deviceID = "") {
  return callApi({ url: `/devices/${deviceID}`, method: "GET" });
}

export async function registerNewDevice(data) {
  return callApi({ url: "/devices", method: "POST", body: data });
}
export async function updateDevice(params, data) {
  return callApi({
    url: `/devices?${params.toString()}`,
    method: "PATCH",
    body: data,
  });
}

// nurse-event func

export async function getNurseEvents() {
  return callApi({ url: "/nurse-events", method: "GET" });
}

export async function getNurseEvent(deviceID) {
  return callApi({ url: `/nurse-events/${deviceID}`, method: "GET" });
}

export async function registerNewEvents(data) {
  return callApi({ url: "/nurse-events", method: "POST", body: data });
}

export async function getLatestEvents() {
  return callApi({ url: "/alerts", method: "GET" });
}
