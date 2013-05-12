package analytics

import (
  "testing"
  "net/http"
  "io/ioutil"
  "net/http/httptest"
)

var host = "localhost"
var port = uint(6379)


func TestRedisConnect(t *testing.T) {
  var err error
  err = redisConnect(host, port)
  // Real simple - does it connect?
  if err != nil {
    t.Fatalf("Failed")
  }
}

func TestRedisConnectFailed(t *testing.T) {
  var err error
  err = redisConnect(host, 6378)

  // Simple again - does it return an error
  if err == nil {
    t.Fatalf("Failed")
  }
}

func TestRedisStore(t *testing.T) {
  redisConnect(host, port)
  redisStore("testing")

  // Ensure it's added to our Redis list
  value, _ := client.LPop("analytics")

  if value != "testing" {
    t.Fatalf("Failed: %s", value)
  }
}

func TestJsHandlerDoesNotError(t *testing.T) {
  dummy := httptest.NewServer(http.HandlerFunc(jsHandler))
  _, err := http.Get(dummy.URL)

  if err != nil {
    t.Fatalf("Failed")
  }
}

func TestJsHandlerHasContentType(t *testing.T) {
  dummy := httptest.NewServer(http.HandlerFunc(jsHandler))
  res, _ := http.Get(dummy.URL)

  if res.Header.Get("Content-Type") != "application/javascript" {
    t.Fatalf("Failed")
  }
}


func TestJsHandlerHasCorrectStatus(t *testing.T) {
  dummy := httptest.NewServer(http.HandlerFunc(jsHandler))
  res, _ := http.Get(dummy.URL)

  if res.StatusCode != http.StatusCreated {
    t.Fatalf("Failed")
  }
}

func TestJsHandlerHasBody(t *testing.T) {
  dummy := httptest.NewServer(http.HandlerFunc(jsHandler))
  res, _ := http.Get(dummy.URL)
  body, _ := ioutil.ReadAll(res.Body)

  if string(body) == "" {
    t.Fatalf("Failed")
  }
}