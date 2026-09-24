#pragma once
#include <Arduino.h>

constexpr uint8_t BUTTON_DARURAT = 25;
constexpr uint8_t BUTTON_INFUS = 26;
constexpr uint8_t BUTTON_RESET = 27;

// pin output
constexpr uint8_t LED_RED = 33;    // darurat led
constexpr uint8_t LED_YELLOW = 32; // infus led
constexpr uint8_t BUZZER =
    14; // turn on whenever darurat or infus being pressed