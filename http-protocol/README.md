```
$ go run http-protocol/main.go
$ curl http://localhost:1729/foo
$ netcat localhost 1729
GET /foo HTTP/1.1
Host: localhost
$ netcat localhost 1729
POST /login HTTP/1.1
Host: localhost
Content-Length: 28

username=arpit&password=pass
```
