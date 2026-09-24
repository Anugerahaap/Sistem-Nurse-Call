#pragma once
#include <Arduino.h>
#include <HTTPClient.h>

class Api {
private:
  char *device_id;

public:
  // Api(char *device_id) { device_id = device_id; }

  String sendEvent(const char *deviceID, const char *event, const char *url);
  String getEvent(const char *deviceID, const char *url);
};
