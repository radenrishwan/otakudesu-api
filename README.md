# Otakudesu API
an Unofficial API for [otakudesu](https://otakudesu.video)

you can access demo here [https://scraping-2wepigvexa-et.a.run.app/](https://scraping-2wepigvexa-et.a.run.app/)
## Usage
### NOTE
before running server, you need to export otakudesu endpoint. you can see at .env file for command or type command below
```bash
export ENDPOINT=PUT-OTAKUDESU-ENDPOINT
```
### Run from Source Code
1. Clone the repository
```bash
git clone https://github.com/radenrishwan/otakudesu-api
```
2. Run Server
```bash
go run main.go
```
3. Build 

    you can also build from source code and run

    ```bash
    go build -o main
    ```

    and run using
    
    ```bash
    ./main
    ```

### Build using Docker

1. Clone the repository
```bash
git clone https://github.com/radenrishwan/otakudesu-api
```

2. Build docker image
```bash
docker build -t otakudesu-api:1.0.0 . 
```
or you can also push from dockerhub
```
docker push radenrishwan/otakudesu-api:1.0.0
```

3. Create container and run
```
docker run -it -p 8080:8080 [IMAGE-NAME]
```

### Using docker compose
you can also running using docker compose 
1. Change endpoint variable on docker-compose.yml file
2. Run command below
```bash
docker-copomse up -d
```

## TODO
1. add more endpoint
2. create a demo app
3. create api docs