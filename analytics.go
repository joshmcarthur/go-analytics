package main

import (
  "text/template"
  "encoding/json"
  "net/http"
  "log"
  "menteslibres.net/gosexy/redis"
)

var client *redis.Client
var redisKey = "analytics"

func redisConnect(host string, port uint)(error) {
  var err error
  client = redis.New()
  err = client.Connect(host, port)
  return err
}

func redisStore(value string) {
  client.LPush(redisKey, value)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
  serializedRequest, serializeError := json.Marshal(r)

  if serializeError != nil {
    http.Error(w, serializeError.Error(), http.StatusInternalServerError)
    return
  }

  redisStore(string(serializedRequest))
  w.Header().Set("Content-Type", "application/javascript")
  w.WriteHeader(http.StatusCreated)

  t, _ := template.ParseFiles("templates/analytics.js")
  t.Execute(w, nil)
}

func main() {
  redisConnect("localhost", 6379)
  http.HandleFunc("/analytics.js", jsHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
  client.Quit()
}