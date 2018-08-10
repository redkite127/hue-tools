# Docker

```bash
docker build -t "hue-tools" .
docker run -d --name hue-tools_01 -p 2002:2002 --restart=always hue-tools
```

or using the pre-build image from Docker Hub: https://hub.docker.com/r/redkite/hue-tools/
```bash
docker run -d --name hue-tools_01 -p 2002:2002 --restart=always redkite/hue-tools
```

# Docker Compose

```bash
docker-compose up -d
```
