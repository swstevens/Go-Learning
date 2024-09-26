just a place for me to house some small files while I learn Golang

Currently I'm working on building a web server using Golang. It runs on a Raspberry Pi and can be accessed at https://swstevens.duckdns.com. My current goal is to set up websockets for online web game multiplayer.

API's include:
```
https://swstevens.duckdns/api
```
Makes a Call to the database and returns a row of data from the NYC squirrel census. In the future will be built out to include more options on calls to the database.

```
https://swstevens.duckdns/counter
```
Returns how many times all endpoints have been accessed.
```
https://swstevens.duckdns/lissajous?cycles=?
```
Creates a lissajous gid that can be changed based on what value is input for cycles. Cycles field is optional.

