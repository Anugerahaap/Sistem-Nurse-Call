#include <response.h>

Event ApiResponse::toString(Event &e) {

  e.deviceID = data["device-id"].as<String>();
  e.event = data["event"].as<String>();

  return e;
};