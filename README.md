go-analytics
============

A badly written (probably) and ridiculously simple Go HTTP script to take an incoming request and persist it in Redis for deriving useful information from later.

### Running:

1. `git clone git@github.com:joshmcarthur/go-analytics.git`
2. `cd go-analytics`
3. `go get menteslibres.net/gosexy/redis`
4. (Ensure Redis is installed and running)
5. `go build analytics.go`
6. `./analytics`
7. `open http://localhost:8080/analytics.js` 
8. Run `redis-cli`, and then `LLEN analytics` to see that request data is being pushed onto the list

### To derive useful information

I'm still working that out myself at the moment - all this does is serializes the request and stores it as JSON. What you might do though, is run a worker that does something along the following lines:

1. Use an `RPOP` Redis command to get the oldest item added to the list (FIFO style)
2. De-serialize the JSON
3. Geolocate the remote IP to find out where the visitor came from - maybe try and work out their ISP
4. Take a look at the `User-Agent` and map that to previous requests
5. Have a look at the URL they requested, along with any params
6. Inspect any custom headers you might have added to the request - maybe you're interested in rendering time, the current user, or how many times they've visited before.



### Tests:

* Same instructions as above up to 4]
* `go test -v`
* All tests should be passing, but require that a Redis server is runnning at `localhost:6379`.

### License

See `LICENSE.txt`
