#pragma once
#include <ArduinoJson.h>

struct Event {
  String deviceID;
  String event;
};

class ApiResponse {

public:
  int code;
  String status;
  String message;
  JsonDocument data;

  Event toString(Event &e);
};
