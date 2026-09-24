#pragma once
#include <pin.h>
class Button {
private:
  /* data */
  int pin;

  unsigned long lastDebounceTime = 0;
  const unsigned long debounceDelay = 50;

  bool lastButtonState = true;
  bool buttonState = true;

public:
  Button(int pinButton) {
    pin = pinButton;
    pinMode(pinButton, INPUT_PULLUP);
  }

  int getPin() { return pin; }

  bool isPressed() {
    bool reading = digitalRead(pin);

    if (reading != lastButtonState) {
      lastDebounceTime = millis();
    }

    if ((millis() - lastDebounceTime) > debounceDelay) {
      if (reading != buttonState) {
        buttonState = reading;

        if (buttonState == LOW) {
          lastButtonState = reading;
          return true;
        }
      }
    }

    lastButtonState = reading;
    return false;
  }
};
