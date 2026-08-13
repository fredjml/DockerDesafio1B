# Full Cycle Rocks — Go + Docker

Desafio de containerização de uma aplicação Go com imagem Docker final inferior a 2 MB.

## Imagem no Docker Hub

https://hub.docker.com/r/fredjml/fullcycle

## Executar

```bash
docker run --rm fredjml/fullcycle
```

Saída esperada:

```text
Full Cycle Rocks!!
```

Para baixar a imagem antes da execução:

```bash
docker pull fredjml/fullcycle:latest
docker run --rm fredjml/fullcycle:latest
```

## Construir localmente

```bash
docker build -t fredjml/fullcycle:latest .
docker run --rm fredjml/fullcycle:latest
```

## Verificar o tamanho

```bash
docker image inspect fredjml/fullcycle:latest --format='{{.Size}} bytes'
docker images fredjml/fullcycle:latest
```

O `Dockerfile` usa multi-stage build. A aplicação é compilada com `CGO_ENABLED=0`, remoção de símbolos de depuração e `trimpath`. O estágio final usa `scratch` e contém somente o binário estático.

## Publicar manualmente no Docker Hub

```bash
docker login
docker build -t fredjml/fullcycle:latest .
docker push fredjml/fullcycle:latest
```

## Estrutura

```text
.
├── .dockerignore
├── .github
│   └── workflows
│       └── docker-publish.yml
├── Dockerfile
├── README.md
├── go.mod
└── main.go
```

## Repositório

https://github.com/fredjml/DockerDesafio1B.git
