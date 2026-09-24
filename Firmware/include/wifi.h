#pragma once
#include <Arduino.h>
#include <WiFi.h>
#include <secret.h>
// wifi
class Wifi {
private:
  /* data */

public:
  void connectWifi();
  bool isConnected() { return WiFi.status() == WL_CONNECTED; }
};
