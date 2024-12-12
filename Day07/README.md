# Day 7 - Advent of Docker

There were a few issues with folllowing along with this stage to cement my learnings.

The examples listed the following to mount and persist the volume.

```sh
$ docker run -p 8080:8080 -v mydata:/data hello-world-go
```

This did not work, and after going deep into trying to find this, whilst pairing with a colleague - the problem was the container mount `/data` the containers data location was actually `go/data`.

To persist the data in the Docker Volume I needed to run:

```sh
$ docker run -p 8080:8080 -v mydata:/go/data hello-world-go
```