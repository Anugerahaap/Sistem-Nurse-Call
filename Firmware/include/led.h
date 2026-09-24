#pragma once
#include <pin.h>

class Led {
public:
  Led() {
    pinMode(LED_RED, OUTPUT);
    pinMode(LED_YELLOW, OUTPUT);
    pinMode(BUZZER, OUTPUT);
    pinMode(LED_BUILTIN, OUTPUT);
  }

  void daruratOn() { digitalWrite(LED_RED, HIGH); }
  void infusOn() { digitalWrite(LED_YELLOW, HIGH); }
  void buzzerOn() { digitalWrite(BUZZER, HIGH); }

  void daruratOff() { digitalWrite(LED_RED, LOW); }
  void infusOff() { digitalWrite(LED_YELLOW, LOW); }
  void buzzerOff() { digitalWrite(BUZZER, LOW); }
};

class BlinkLed : public Led {
private:
  int pin;

  bool ledState = LOW;
  unsigned long previousMillis = 0;

public:
  BlinkLed(int pinLed) { pin = pinLed; }

  bool led = false;

  void ledSwitch(bool cons) { led = cons; }

  void blink(unsigned long interval) {
    unsigned long currentMillis = millis();

    if (currentMillis - previousMillis >= interval) {
      previousMillis = currentMillis;

      ledState = !ledState;

      digitalWrite(pin, ledState);
    }
  }
};