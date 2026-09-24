#include <wifi.h>

void Wifi::connectWifi() {
  WiFi.begin(SSID, PASS);
  int counter = 0;

  Serial.println("Connecting to Wifi");
  while (WiFi.status() != WL_CONNECTED) {
    /* code */
    delay(500);
    Serial.println("Percobaan ke :");
    Serial.print(counter);
    Serial.println(".");
    counter++;

    if (counter == 20) {
      Serial.println("connecting wifi failed");
      return;
    }
  }
  Serial.println();
  Serial.println("WiFi connected!");
  Serial.print("IP Address: ");
  Serial.println(WiFi.localIP());
}
