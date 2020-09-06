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

### Config staticfy in `config.yml` file

```yaml
staticfy:
  port: 9000
  prefix: "/assets"
  directory: "/statics"
```

### Using upload and delete file

- Environment

```.env
TZ=Asia/Bangkok
HOST=http://localhost:4000
X-API-KEY=api-secret
JWT_SECRET=jwt-secret

# PostgresQL
DATABASE_DRIVER_NAME=postgres
DATABASE_CONNECTION_STRING=postgres://user:password@localhost/dbName?sslmode=disable

# MySQL
DATABASE_DRIVER_NAME=mysql
DATABASE_CONNECTION_STRING=user:password@(localhost:3306)/dbName?charset=utf8&parseTime=true
```

- Config staticfy in `config.yml` file

```yaml
staticfy:
  port: 9000
  host: "http://localhost:9000"
  prefix: "/assets"
  directory: "/statics"
  routes:
    upload:
      path: "/v1/images"
      method: "POST"
      support: ['.png', '.jpg', '.jpeg', '.gif']
      directory: "/images"
    delete:
      path: "/v1/images/delete"
      method: "POST"
      directory: "/images"
```

- Upload File API

```shell script
POST http://localhost:9000/v1/images
Header:
  x-api-key: df2ec6bd-eaf4-47cd-a397-1fbf1b1d60c7
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.FT288R0fAPyqDnBCVChX9MB_8Fjj_-d6y2N402Ws5eM
Request:
  file: @/Users/staticfy/images/innotech-development.png
```

- Delete File API

```shell script
POST http://localhost:9000/v1/images/delete
Header:
  x-api-key: df2ec6bd-eaf4-47cd-a397-1fbf1b1d60c7
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.FT288R0fAPyqDnBCVChX9MB_8Fjj_-d6y2N402Ws5eM
Request:
  id: 1
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

-> staticfy serving /assets on :9000
```

- Listening [http://localhost:9000/assets/hello.txt](http://localhost:9000/assets/hello.txt)


### Example Project

[https://github.com/prongbang/staticfy-example](https://github.com/prongbang/staticfy-example)
