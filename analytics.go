package main

import (
  "fmt"
  "encoding/json"
  "net/http"
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

func jsHandler(w http.ResponseWriter, r *http.Request) {
  serializedRequest, serializeError := json.Marshal(r)

  if serializeError != nil {
    fmt.Printf("Error: %s", serializeError.Error())
    http.Error(w, serializeError.Error(), http.StatusInternalServerError)
  }

  client.LPush(redisKey, serializedRequest)
  return
}

func main() {
  redisConnect("localhost", 6379)
  http.HandleFunc("/analytics.js", jsHandler)
  http.ListenAndServe(":8080", nil)
  client.Quit()
}