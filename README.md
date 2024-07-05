# Example API service in Go

This is an example of an API service application written in Go.
The application uses [Gin](https://github.com/gin-gonic/gin/) for the REST API framework.

## Docker image

The application is available as a Docker image on [Docker Hub](https://hub.docker.com/r/bluitel/api/tags).
To build one locally, clone the repository and run make command:

```bash
make image
```

## Deployment

This application us deployable to Kubernetes cluster using [Pulumi](https://www.pulumi.com/).
To deploy the application using Pulumi, you need to have [Pulumi installed](https://www.pulumi.com/docs/install/)
on your machine.

After installing Pulumi, run the following commands to deploy the application:

```bash
pulumi up --stack dev --config "api:kube-context=kubernetes-admin@kubernetes"
```

Use the `--config` flag and specify the Kubernetes context to use for deployment.

The `kubernetes-admin@kubernetes` default k8s context which is used in the above command is deploying the
application to the Kubernetes cluster running in raspberry pi configured using ansible script
@[binodluitel/rpi-ansible](https://github.com/binodluitel/rpi-ansible).

### Run using Docker

If you wish to run using Docker, you can do that too. Simply `docker run` the image.

```bash
docker run bluitel/api:latest
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)
[GIN-debug] Listening and serving HTTP on :8080
{"level":"info","ts":1719970919.3356018,"caller":"api/main.go:32","msg":" ----- Welcome to the API service example ----- "}
{"level":"debug","ts":1719970919.3363056,"caller":"api/main.go:39","msg":"Application build information","name":"api-service","version":"5ce1db0","build_time":"2024-07-02T23:16:14Z","ref_name":"main","ref_sha":"5ce1db0d5d6ab557ed35756f53edba06ebe137fd"}
[GIN-debug] Listening and serving HTTP on :9090
```

## Metrics

The application exposes metrics on `/metrics` endpoint. The metrics are exposed in Prometheus format.

```bash
$ curl -s http://127.0.0.1:9090/metrics
```

## 418 (I'm a teapot)

Currently, all REST verbs for / (Base) URL returns 418 (I'm a teapot) status code with no content because
base route does not have any implementation and is not intended to be used.

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

## API Endpoints

The application has the following REST API endpoints:

- `GET v1/stream/logs` - Stream logs from the application (currently streams giberish)
