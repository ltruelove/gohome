#include <WiFi.h>
#include <WiFiClient.h>
#include <WebServer.h>
#include <ESPmDNS.h>
#include <DHT.h>
#include <ArduinoJson.h>
#include <Preferences.h>
#include <HTTPClient.h>

#define DHTPIN 4
#define DHTTYPE DHT11

Preferences prefs;
WebServer server(80);
DHT dht(DHTPIN, DHTTYPE);

const int led = 2;
const int doorSensorPin = 16;
const char *prefsName = "sensorPrefs";

unsigned long previousMillis = 0;
unsigned long interval = 30000;
int doorState = 0;
bool isConnected = false;
bool hasPreferences = false;
String ssid;
String pass = "";
String apiHost = "";
uint16_t apiPort = 8080;
String name = "";
String sensorId = "";
String isGarage = "";
String st;
String content;
int statusCode;

void setup() {
  Serial.begin(115200);
  
  WiFi.disconnect();
  //clearPreferences();

  prefs.begin(prefsName, true);
  ssid = prefs.getString("ssid", "");
  pass = prefs.getString("pass", "");
  apiHost = prefs.getString("apiHost", "");
  apiPort = prefs.getUInt("apiPort", 80);
  sensorId = prefs.getString("sensorId");
  name = prefs.getString("name");
  isGarage = prefs.getString("isGarage");
  prefs.end();

  if(!ssid.isEmpty()){

    connectToWifi(ssid.c_str(), pass.c_str());
    hasPreferences = true;

    Serial.println("Preferences found");
    Serial.println(ssid);
    Serial.println(pass);
    Serial.println(apiHost);
    Serial.println(apiPort);
    Serial.println(sensorId);

    if(WiFi.status() == WL_CONNECTED){
      isConnected = true;

      if(sensorId.isEmpty()){
        registerSensorWithAPI(apiHost, apiPort);
      }else{
        //this should be where the device operates most of the time
        updateIPWithAPI();
        dht.begin();
        pinMode(led, OUTPUT);
        digitalWrite(led, 1);
        launchSensorWeb();
      }
    } else {
      setupAccessPoint();
    }
  }else{
    setupAccessPoint();
  }
}

void loop() {
  server.handleClient();

  if(!hasPreferences){
    delay(1);
  }else{
    delay(100);
    doorState = digitalRead(doorSensorPin);

    // increment our time counter
    unsigned long currentMillis = millis();
    // if WiFi is down, try reconnecting
    if ((WiFi.status() != WL_CONNECTED) && (currentMillis - previousMillis >=interval)) {
      Serial.println(millis());
      Serial.println("Reconnecting to WiFi...");
      WiFi.disconnect();
      if(WiFi.reconnect()){
        if(!sensorId.isEmpty()){
          updateIPWithAPI();
        }else{
          registerSensorWithAPI(apiHost, apiPort);
        }
      }
      previousMillis = currentMillis;
    }
  }
}

bool registerSensorWithAPI(String host, int port){
  bool registerSuccess = false;
  String registerEndpoint = "http://" + host;
  if(port != 80){
    registerEndpoint = registerEndpoint + ":" + (String)port;
  }
  registerEndpoint = registerEndpoint + "/temps/registersensor";

  if(WiFi.status() != WL_CONNECTED){
    Serial.println("Not connected to WiFi. Cannot register.");
    return registerSuccess;
  }

  WiFiClient client;
  HTTPClient http;

  Serial.println(WiFi.localIP());
  http.begin(client, registerEndpoint);
  http.addHeader("Content-Type", "application/json");

  DynamicJsonDocument doc(1024);

  doc["name"] = name;
  doc["isGarage"] = isGarage == "yes" ? 1:0;
  doc["ipAddress"] = WiFi.localIP().toString();

  String serializedString;
  serializeJson(doc, serializedString);

  Serial.println("Payload:");
  Serial.println(serializedString);
  int httpResponseCode = http.POST(serializedString);

  if(httpResponseCode != 200){
    http.end();
    Serial.print("Register request failed with response code: ");
    Serial.println(httpResponseCode);
    clearPreferences();
    return registerSuccess;
  }

  registerSuccess = true;
  String sensor = http.getString();
  Serial.println(sensor);
  DynamicJsonDocument responseDoc(1024);
  deserializeJson(responseDoc, sensor);
  String sensorId = responseDoc["sensorId"];
  Serial.println(sensorId);

  http.end();

  prefs.begin(prefsName, false);
  prefs.putString("sensorId", sensorId);
  prefs.end();
  ESP.restart();

  return registerSuccess;
}

void updateIPWithAPI(){
  String updateEndpoint = "http://" + apiHost;
  if(apiPort != 80){
    updateEndpoint = updateEndpoint + ":" + (String)apiPort;
  }
  updateEndpoint = updateEndpoint + "/temps/updatesensor";

  WiFiClient client;
  HTTPClient http;

  http.begin(client, updateEndpoint);
  http.addHeader("Content-Type", "application/json");

  DynamicJsonDocument doc(1024);

  doc["name"] = name;
  doc["isGarage"] = isGarage == "yes" ? 1:0;
  doc["ipAddress"] = WiFi.localIP().toString();
  doc["sensorId"] = sensorId;

  String serializedString;
  serializeJson(doc, serializedString);
  Serial.println("Payload:");
  Serial.println(serializedString);
  int httpResponseCode = http.PUT(serializedString);
  http.end();
}

void clearPreferences()
{
  prefs.begin(prefsName, false);
  prefs.clear();
  prefs.end();
  ESP.restart();
}

void connectToWifi(const char *ssid, const char *key){
  WiFi.mode(WIFI_STA);
  Serial.println("connecting");
  WiFi.begin(ssid, key);
  int c = 0;

  while(c < 20){
    if(WiFi.status() == WL_CONNECTED){
      Serial.println("connected");
      IPAddress ip = WiFi.localIP();
      Serial.println(ip.toString());
      return;
    }
    delay(500);
    Serial.print(".");
    c++;
  }
}

void setupAccessPoint(void)
{
  WiFi.disconnect();
  WiFi.mode(WIFI_AP);
  delay(100);
  int n = WiFi.scanNetworks();
  Serial.println("scan done");
  if (n == 0){
    Serial.println("no networks found");
  } else {
    Serial.print(n);
    Serial.println(" networks found");
    for (int i = 0; i < n; ++i)
    {
      // Print SSID and RSSI for each network found
      Serial.print(i + 1);
      Serial.print(": ");
      Serial.print(WiFi.SSID(i));
      Serial.print(" (");
      Serial.print(WiFi.RSSI(i));
      Serial.print(")");
      delay(10);
    }
  }

  Serial.println("");
  st = "<select name='ssid'>";

  for (int i = 0; i < n; ++i) {
    // Print SSID and RSSI for each network found
    st += "<option value='";
    st += WiFi.SSID(i);
    st += "'>";
    st += WiFi.SSID(i);
    st += " (";
    st += WiFi.RSSI(i);

    st += ")";
    st += "</option>";
  }

  st += "</select>";
  delay(100);

  WiFi.softAP("GoHome Temp Sensor", "");
  launchWeb();
}

void launchWeb(){
  Serial.println("");

  Serial.print("SoftAP IP: ");
  Serial.println(WiFi.softAPIP());

  server.on("/", homePage);
  server.on("/setting", setParameters);
  server.onNotFound(handleNotFound);
  /*
  server.on("/scan", scanPage);
  */

  server.begin();
  Serial.println("Web Server started");
}

void launchSensorWeb(){
  server.on("/", handleRoot);
  server.onNotFound(handleNotFound);
  server.begin();
}

void handleNotFound() {
  String message = "File Not Found\n\n";
  message += "URI: ";
  message += server.uri();
  message += "\nMethod: ";
  message += (server.method() == HTTP_GET) ? "GET" : "POST";
  message += "\nArguments: ";
  message += server.args();
  message += "\n";

  for (uint8_t i = 0; i < server.args(); i++) {
    message += " " + server.argName(i) + ": " + server.arg(i) + "\n";
  }

  server.send(404, "text/plain", message);
}

// returns the state of the temp/humidity sensor as well as
// the garage door's open/closed status if it's the garage sensor
void handleRoot() {
  digitalWrite(led, 1);

  StaticJsonDocument<500> doc;
  doc["humidity"] = dht.readHumidity();
  doc["celcius"] = dht.readTemperature();
  doc["fahrenheit"] = dht.readTemperature(true);
  if(isGarage == "yes"){
    doc["doorClosed"] = doorState;
  }
  
  String output;
  serializeJson(doc, output);

  server.sendHeader("Access-Control-Allow-Origin", "*");
  server.send(200, "application/json", output);
  digitalWrite(led, 0);
}

void homePage(){
    IPAddress ip = WiFi.softAPIP();

    prefs.begin(prefsName, true);
    ssid = prefs.getString("ssid", "");
    pass = prefs.getString("pass", "");
    apiHost = prefs.getString("apiHost", "");
    apiPort = prefs.getUInt("apiPort", 80);
    prefs.end();

    char port[8];
    String sPort;
    Serial.println("convert prefs func");
    itoa(apiPort, port, 10);

    sPort = port;

    Serial.println("build page");
    content = "<!DOCTYPE HTML>\r\n<html>Welcome to Wifi Credentials Update page";
    content += "<form action=\"/scan\" method=\"POST\"><input type=\"submit\" value=\"scan\"></form>";
    content += "<p>";
    content += ip.toString();
    content += "</p>";
    content += "<form method='get' action='setting'>";
    content += "<table><tr><td width='100'>";
    content += "<label>SSID: </label></td><td>" + st +"</td></tr>";
    content += "<tr><td><label>Key: </label></td><td><input type='password' name='pass' value='" + pass + "' length=64></td></tr>";
    content += "<tr><td><label>Sensor Name: </label></td><td><input type='text' name='name'></td></tr>";
    content += "<tr><td><label>API Host: </label></td><td><p>http:// <input type='text' name='apiHost' value='" + apiHost + "' length=32></p></td></tr>";
    content += "<tr><td><label>API PORT: </label></td><td><input type='number' name='apiPort' value='" + sPort + "'  length=5></td></tr>";
    content += "<tr><td><label>Is Garage: </label></td><td><input type='checkbox' name='isGarage' value='yes'></td></tr>";
    content += "<tr><td>&nbsp;</td><td><input type='submit'></td></tr></form>";
    content += "</html>";

    Serial.println("send page");
    server.send(200, "text/html", content);
}

void setParameters(){
  ssid = server.arg("ssid");
  pass = server.arg("pass");
  apiHost = server.arg("apiHost");
  name = server.arg("name");
  isGarage = server.arg("isGarage");

  Serial.println(apiHost);
  apiPort = atoi(server.arg("apiPort").c_str());

  if (ssid.length() > 0 && pass.length() > 0) {
      prefs.begin(prefsName, false);
      prefs.putString("ssid", ssid);
      prefs.putString("pass", pass);
      prefs.putString("apiHost", apiHost);
      prefs.putUInt("apiPort", apiPort);
      prefs.putString("isGarage", isGarage);
      prefs.putString("name", name);
      prefs.end();

    content = "{\"Success\":\"saved to eeprom... reset to boot into new wifi\"}";
    statusCode = 200;
    ESP.restart();
  } else {
    content = "{\"Error\":\"404 not found\"}";
    statusCode = 404;
    Serial.println("Sending 404");
  }
  server.sendHeader("Access-Control-Allow-Origin", "*");
  server.send(statusCode, "application/json", content);
}
