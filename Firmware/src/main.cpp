#include <api.h>
#include <button.h>
#include <iostream>
#include <led.h>
#include <wifi.h>

Wifi wifi;
Api api;

Button redBtn(BUTTON_DARURAT);
Button ylwBtn(BUTTON_INFUS);
Button rstBtn(BUTTON_RESET);

Led led;
BlinkLed blinkWifiError(LED_BUILTIN);
BlinkLed red(LED_RED);
BlinkLed ylw(LED_YELLOW);

const char *DEVICE_ID = "NC-xx1";
const char *base_url = "http://192.168.1.6:3000/api/";

void setup() {
  Serial.begin(115200);

  wifi.connectWifi();
};

void loop() {
  if (!wifi.isConnected()) {
    blinkWifiError.blink(500);
  }

  if (redBtn.isPressed()) {
    ylw.ledSwitch(false);
    led.infusOff();
    red.ledSwitch(true);
    String darurat = api.sendEvent(DEVICE_ID, "DARURAT", base_url);
    Serial.println(darurat);

    Serial.println(api.getEvent(DEVICE_ID, base_url));
  }

  if (red.led) {
    red.blink(300);
    // led.buzzerOn();
  }

  if (ylwBtn.isPressed()) {
    led.daruratOff();
    red.ledSwitch(false);
    ylw.ledSwitch(true);
    String infus = api.sendEvent(DEVICE_ID, "INFUS", base_url);
    Serial.println(infus);
    Serial.println(api.getEvent(DEVICE_ID, base_url));
  }
  if (ylw.led) {
    ylw.blink(300);
    // led.buzzerOn();
  }

  if (rstBtn.isPressed()) {

    red.ledSwitch(false);
    ylw.ledSwitch(false);

    led.buzzerOff();
    led.daruratOff();
    led.infusOff();
    String reset = api.sendEvent(DEVICE_ID, "RESET", base_url);
    Serial.println(reset);
    Serial.println(api.getEvent(DEVICE_ID, base_url));
  };
}