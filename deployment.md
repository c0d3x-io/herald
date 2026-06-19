# docker
docker image builder 
```docker
docker build -t herald:v0.1.0 .
```

docker run 
```
docker run -p 8080:8080 --env-file .env herald:v0.1.0
```

# .env
HERALD_UPSTREAM_URL=https://your-upstream-api.com
HERALD_LISTEN_ADDR=:8080
HERALD_LOG_LEVEL=debug

# Self sign certificated 
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/CN=localhost"
