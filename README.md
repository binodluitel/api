# API service example

This is an example of a simple API service application written in Go version v1.22.

## All REST verbs for / (Base) URL will return 418 (I'm a teapot) status code.

```bash
$ curl -s -v http://127.0.0.1:8080
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080
> GET / HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/8.6.0
> Accept: */*
>
< HTTP/1.1 418 I'm a teapot
< Date: Mon, 01 Jul 2024 18:12:30 GMT
< Content-Length: 0
<
* Connection #0 to host 127.0.0.1 left intact
```
