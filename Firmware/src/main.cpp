#include <api.h>
#include <button.h>
#include <iostream>
#include <led.h>
#include <response.h>
#include <wifi.h>

Wifi wifi;
Api api;
ApiResponse response;
ApiResponse dataEvent;
Event lastEvent;

Button redBtn(BUTTON_DARURAT);
Button ylwBtn(BUTTON_INFUS);
Button rstBtn(BUTTON_RESET);

Led led;
BlinkLed blinkWifiError(LED_BUILTIN);
BlinkLed red(LED_RED);
BlinkLed ylw(LED_YELLOW);

const char *DEVICE_ID = "baruUpdate";
const char *base_url = "http://192.168.1.6:3000/api/";
const unsigned long getInterval = 1000;
unsigned long prevMilis = 0;

void setup() {
  Serial.begin(115200);

  wifi.connectWifi();
};

void loop() {
  if (!wifi.isConnected()) {
    blinkWifiError.blink(500);
  }

  unsigned long cMilis = millis();
  if (cMilis - prevMilis >= getInterval) {
    /* code */
    prevMilis = cMilis;
    String res = api.getEvent(DEVICE_ID, base_url);
    api.parseResponse(dataEvent, res);
    // Serial.println(res);
    if (dataEvent.code == 200) {
      /* code */
      dataEvent.toString(lastEvent);
    }
  };

  if (redBtn.isPressed()) {
    ylw.ledSwitch(false);
    led.infusOff();
    red.ledSwitch(true);
    String darurat = api.sendEvent(DEVICE_ID, "DARURAT", base_url);
    api.parseResponse(response, darurat);
  }

  if (red.led || lastEvent.event == "DARURAT") {
    red.blink(300);
  }

  if (ylwBtn.isPressed()) {
    led.daruratOff();
    red.ledSwitch(false);
    ylw.ledSwitch(true);
    String infus = api.sendEvent(DEVICE_ID, "INFUS", base_url);
    api.parseResponse(response, infus);
  }
  if (ylw.led || lastEvent.event == "INFUS") {
    ylw.blink(300);
  }

  if (rstBtn.isPressed()) {

    red.ledSwitch(false);
    ylw.ledSwitch(false);

    led.buzzerOff();
    led.daruratOff();
    led.infusOff();
    String reset = api.sendEvent(DEVICE_ID, "RESET", base_url);
    api.parseResponse(response, reset);
  }

  if (lastEvent.event == "RESET") {
    /* code */
    red.ledSwitch(false);
    ylw.ledSwitch(false);

    led.buzzerOff();
    led.daruratOff();
    led.infusOff();
  }
}