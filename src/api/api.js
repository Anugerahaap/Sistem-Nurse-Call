export const BASE_URL = "http://localhost:3000/api";

/**
 *
 * boilerplate atau template func
 * @param {string} endpoint-endpoint server
 * @param {string} method-ya method dari web yg mau di kirim kwkwk
 * @param {object} body-sebuah body json yg berupa objek yg akan dikirim ke server
 * @param {{}} headers- untuk menambah headers
 * @param {int} timeout -untuk mengatur batas waktu jika server terlalu lama mengrespon
 *
 */
export async function callApi({
  endpoint = "",
  method = "GET",
  body = null,
  headers = {},
  timeout,
}) {
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), timeout);
  try {
    const response = await fetch(`${BASE_URL}/${endpoint}`, {
      method,
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        ...headers,
      },
      body: body ? JSON.stringify(body) : null,
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(
        errorData.message || `HTTP Error! Status: ${response.status}`,
      );
    }

    return await response.json();
  } catch (error) {
    // console.log(error);
    clearTimeout(timer);

    if (error.name === "AbortError") {
      throw new Error("Timeout: Server terlalu lama merespon.");
    }

    throw error;
  }
}
/**
 * mendapatkan semua device
 */
export async function getDevices() {
  return await callApi({ endpoint: "/devices", timeout: 3 });
}

export async function updateRoomNames(deviceID, roomname) {
  const params = new URLSearchParams({ "room-name": roomname });
  return await callApi({
    endpoint: `/devices/${deviceID}?${params.toString()}`,
    method: "PATCH",
    timeout: 5,
  });
}

/**
 *mendapatkan last event terbaru,jika ada selain RESET maka alarm berbunyi
 */
export async function getLatestEvents() {
  return await callApi({ endpoint: "/nurse-events", timeout: 3 });
}
/**
 *mendapatkan last event terbaru dari device ID
 */
export async function getLastEvent(deviceID) {
  return await callApi({ endpoint: `/nurse-events/${deviceID}` });
}

/**
 *buat logs or history,mendapatkan semua events
 */
export async function getLogsEvents() {
  return await callApi({ endpoint: "/nurse-events?logs=true", timeout: 3 });
}

export async function registerNewEvent(data) {
  return await callApi({
    endpoint: "/nurse-events",
    method: "POST",
    body: data,
    timeout: 5,
  });
}
