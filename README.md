# mini-redis-server

## How to run application 
1) go get github.com/christinakim/mini-redis-server
2) cd into application
3) to start the application `go run application`

## How to use application 

Example of `String SET`: 
 `curl -H "Content-Type: application/json" -X POST -d '{"key":"one","value":"1"}' http://localhost:8000/api/v1/strings/set`
 
Example of `String GET`: 
 `curl -X GET -d '{"key":"one"}' http://localhost:8000/api/v1/strings/get`

Example of `String DELETE`: 
`curl -X DELETE -d '{"key":"one"}' http://localhost:8000/api/v1/strings/get`


Example of `list GET`:
 `curl -X GET -d '{"key":"one"}' http://localhost:8000/api/v1/lists/get`

 Example of `list SET`:
 `curl -H "Content-Type: application/json" -X POST -d '{"key":"one","value":["1"]}' http://localhost:8000/api/v1/lists/set`

Example of `list APPEND`:
`curl -H "Content-Type: application/json" -X POST -d '{"key":"one","append":"2"}' http://localhost:8000/api/v1/lists/append`

Example of `list DELETE`:
`curl -X DELETE -d '{"key":"one"}' http://localhost:8000/api/v1/lists/pop`

Example of `map GET`:
 `curl -X GET -d '{"key":"one"}' http://localhost:8000/api/v1/maps/get`

 Example of `map SET`:
 `curl -H "Content-Type: application/json" -X POST -d '{"key":"one","value":{"two":"two"}}' http://localhost:8000/api/v1/maps/set`

Example of `map DELETE`:
`curl -X DELETE -d '{"key":"one"}' http://localhost:8000/api/v1/lists/pop`

Example of `map MAPGET`:
 `curl -X GET -d '{"key":"one","mapkey":"two"}' http://localhost:8000/api/v1/maps/mapget`

 Example of `map MAPSET`:
 `curl -H "Content-Type: application/json" -X POST -d '{"key":"one","mapkey":"two", "mapvalue":"three"}' http://localhost:8000/api/v1/maps/mapset`
 
 Example of `map MAPDELETE`:
 `curl -H "Content-Type: application/json" -X POST -d '{"key":"one","mapkey":"two"}' http://localhost:8000/api/v1/maps/mapset`


Example of search:
 `curl -X GET http://localhost:8000/api/v1/search/one`
