#include <WiFi.h>
#include <WiFiClient.h>
#include <WebServer.h>
#include <ESPmDNS.h>
#include <DHT.h>
#include <ArduinoJson.h>

const char *ssid = "C3PO";
const char *password = "dexter is cute";
const char* message = "{\"soilReading\" : ";
const int led = 2;
const int SoilSensorPin = 34;

IPAddress local_IP(192, 168, 1, 16);
IPAddress gateway(192, 168, 1, 1);
IPAddress subnet(255, 255, 255, 0);

WebServer server(80);

void setup() {
  // put your setup code here, to run once:
  WiFi.config(local_IP, gateway, subnet);
  WiFi.begin(ssid, password);

  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
  }
  
  String ip = WiFi.localIP().toString();

  server.on("/status", handleStatus);
  server.onNotFound(handleNotFound);

  server.begin();
}

void loop() {
  // put your main code here, to run repeatedly:
  server.handleClient();

}

void handleStatus() {
  String response = String(message);
  response.concat(analogRead(SoilSensorPin));
  response.concat("}");
  
  server.send(200, "application/json", response);
}

void handleNotFound(){
  //digitalWrite(led, 1);
  String message = "File Not Found\n\n";
  message += "URI: ";
  message += server.uri();
  message += "\nMethod: ";
  message += (server.method() == HTTP_GET)?"GET":"POST";
  message += "\nArguments: ";
  message += server.args();
  message += "\n";
  for (uint8_t i=0; i<server.args(); i++){
    message += " " + server.argName(i) + ": " + server.arg(i) + "\n";
  }
  server.send(404, "text/plain", message);
}
