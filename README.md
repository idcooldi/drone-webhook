# drone-webhook
[![Go Report Card](https://goreportcard.com/badge/github.com/idcooldi/drone-webhook)](https://goreportcard.com/report/github.com/idcooldi/drone-webhook)
[![Build Status](https://cloud.drone.io/api/badges/idcooldi/drone-webhook/status.svg)](https://cloud.drone.io/idcooldi/drone-webhook)
[![GoDoc](https://godoc.org/github.com/idcooldi/drone-webhook?status.svg)](https://godoc.org/github.com/idcooldi/drone-webhook)
[![Docker Pulls](https://img.shields.io/docker/pulls/idcooldi/drone-webhook.svg)](https://hub.docker.com/r/idcooldi/drone-webhook)
[![](https://images.microbadger.com/badges/image/idcooldi/drone-webhook.svg)](https://microbadger.com/images/idcooldi/drone-webhook "Get your own image badge on microbadger.com")



[Drone](https://github.com/drone/drone) plugin for sending webhook with custom authorization.


```steps:
   - name: send
     image: idcooldi/drone-webhook
     settings:
      bearer:
        from_secret: bearer_token
       urls: https://your.webhook/...
       debug: true
```

## Parameter Reference

**urls**

Payload gets sent to this list of URLs

**username**

Username for basic auth

**password**

Password for basic auth

**method**

HTTP submission method, defaults to POST

**content_type**

Content type, defaults to application/json

**skip_verify**

Skip SSL verification

**debug**

Enable debug information

**headers**

Custom headers 

Simple usage headers:
```
   - name: send
     image: idcooldi/drone-webhook
     settings:
       urls:
         - https://your.webhook/...
       headers:
         - "Content-Length=123"
```
