#include <api.h>

String Api::sendEvent(const char *deviceID, const char *event,
                      const char *url) {
  String urlEvent = String(url) + "nurse-events";

  HTTPClient http;

  http.begin(urlEvent.c_str());
  http.addHeader("Content-Type", "application/json");
  String payload = "{\"device-id\":\"" + String(deviceID) + "\",\"event\":\"" +
                   String(event) + "\"}";

  int httpcode = http.POST(payload);
  if (httpcode < 0) {

    String error = http.errorToString(httpcode);

    http.end();

    return error;
  }

  String response = http.getString();

  http.end();

  return response;
}

String Api::getEvent(const char *deviceID, const char *url) {
  String urlEvent = String(url) + "nurse-events/" + String(deviceID);

  HTTPClient http;

  http.begin(urlEvent.c_str());

  int httpcode = http.GET();
  if (httpcode < 0) {

    String error = http.errorToString(httpcode);

    http.end();

    return error;
  }

  String response = http.getString();

  http.end();

  return response;
}

void Api::parseResponse(ApiResponse &response, String r) {

  JsonDocument doc;

  deserializeJson(doc, r);

  response.code = doc["code"];
  response.status = doc["status"].as<String>();
  response.message = doc["message"].as<String>();

  if (response.code == 200) {
    /* code */
    response.data = doc["data"];
  }
}
