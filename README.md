# Staticfy [![Docker Pulls](https://img.shields.io/docker/pulls/prongbang/staticfy.svg)](https://hub.docker.com/r/prongbang/staticfy/) [![Image Size](https://img.shields.io/docker/image-size/prongbang/staticfy.svg)](https://hub.docker.com/r/prongbang/staticfy/)

> [Staticfy](https://hub.docker.com/r/prongbang/staticfy) links to assets in your project.

## Setup project

```bash
your-project
└── staticfy
    ├── config.yml
    └── statics
        └── hello.txt
```

### Config Staticfy in `config.yml` file

```yaml
staticfy:
  directory: "/statics"
  prefix: "/publisher"
  port: 9000
```

## How to run

### Run with Docker

```yaml
version: '3.7'
services:
  app_staticfy:
    image: prongbang/staticfy:latest
    ports:
      - "9000:9000"
    volumes:
      - "./staticfy:/staticfy"
```

```
$ docker-compose up -d
```

### Run with Golang

```shell script
$ go get -u github.com/prongbang/staticfy
$ cd project
```

#### Default port `9000` from `config.yml`

```bash
$ staticfy
```

- Running

```shell script
   ______       __  _     ___    
  / __/ /____ _/ /_(_)___/ _/_ __
 _\ \/ __/ _ `/ __/ / __/ _/ // /
/___/\__/\_,_/\__/_/\__/_/ \_, / 
                          /___/

-> staticfy serving /publisher on :9000
```

- Listening [http://localhost:9000/publisher/hello.txt](http://localhost:9000/publisher/hello.txt)


### Example Project

[https://github.com/prongbang/staticfy-example](https://github.com/prongbang/staticfy-example)
