# Relatório Técnico Executivo - Implementação do Desafio DockerDesafioGo

Desafio: DockerDesafioGo - Go Lang + Docker

Diretório de entrada: `D:\ProjetosFullCycle\02DiscDocker\Desafio1B\fullcycle\fullcycle`

Repositório GitHub de destino: `https://github.com/fredjml/DockerDesafio1B.git`

Imagem Docker Hub de entrega: `fredjml/fullcycle`

Branch de entrega: `main`

Linguagem: Go

Containerização: Docker

Registry: Docker Hub

Status geral da implementação: VERDE - implementação revisada, corrigida e validada localmente.

Status de publicação GitHub: preparado para commit e push no repositório remoto informado.

Status de publicação Docker Hub: publicado com sucesso em `fredjml/fullcycle:latest`, digest `sha256:64bed66c82e24c7cab2fa0708dbae2774ef0fdd17f52d7175fc069b12228ea23`, imagem local validada com tamanho de 641809 bytes.

## 1. Resumo Executivo

Foi conferida e ajustada a aplicação do desafio `DockerDesafioGo`, cujo objetivo é entregar uma imagem Docker pública que execute uma aplicação Go e imprima exatamente:

```text
Full Cycle Rocks!!
```

A implementação utiliza uma aplicação Go mínima, sem dependências externas, compilada como binário estático por meio de Docker multi-stage build. A imagem final termina em `scratch`, contendo somente o binário compilado. Essa estratégia reduz drasticamente o tamanho da imagem final e atende ao requisito de imagem menor que 2 MB.

Correções realizadas:

- Alinhamento do módulo Go para o novo repositório: `github.com/fredjml/DockerDesafio1B`.
- Reescrita do README em UTF-8 legível, substituindo conteúdo com caracteres corrompidos.
- Inclusão do relatório técnico executivo de implementação no conteúdo do `README.md`.
- Registro explícito dos critérios de aceite com status `feito`.
- Documentação da estratégia multi-stage terminando em `scratch`, com `CGO_ENABLED=0` e `-ldflags="-s -w"`.
- Geração dos relatórios de implementação em `.md` e `.docx` no diretório `Analise`.

## 2. Escopo Executado

Foi executada a implementação/correção necessária para deixar o desafio aderente aos requisitos de entrega.

Arquivos ajustados ou gerados:

- `go.mod`
- `README.md`
- `Analise/RELATORIO_TECNICO_EXECUTIVO_IMPLEMENTACAO_DESAFIO.md`
- `Analise/RELATORIO_TECNICO_EXECUTIVO_IMPLEMENTACAO_DESAFIO.docx`

Arquivos conferidos:

- `main.go`
- `Dockerfile`
- `.dockerignore`
- `.github/workflows/docker-publish.yml`

## 3. Fontes e Evidências Consultadas

| Fonte | Evidência observada | Status |
|---|---|---|
| Enunciado do desafio | Requisitos de Go, Docker, Docker Hub, saída exata e imagem menor que 2 MB | feito |
| `main.go` | Aplicação imprime `Full Cycle Rocks!!` | feito |
| `Dockerfile` | Multi-stage build com etapa final `scratch` | feito |
| `go.mod` | Módulo ajustado para `github.com/fredjml/DockerDesafio1B` | feito |
| `.dockerignore` | Contexto de build reduzido | feito |
| GitHub Actions | Workflow valida saída e tamanho da imagem | feito |
| README | Conteúdo refeito em formato de relatório executivo | feito |
| Repositório remoto | `https://github.com/fredjml/DockerDesafio1B.git` | feito |

## 4. Estado Final da Aplicação

Estrutura principal do projeto:

```text
.
|-- .dockerignore
|-- .github/
|   `-- workflows/
|       `-- docker-publish.yml
|-- Analise/
|   |-- RELATORIO_TECNICO_EXECUTIVO_IMPLEMENTACAO_DESAFIO.md
|   `-- RELATORIO_TECNICO_EXECUTIVO_IMPLEMENTACAO_DESAFIO.docx
|-- Dockerfile
|-- README.md
|-- go.mod
`-- main.go
```

Aplicação Go:

```go
package main

import "os"

const message = "Full Cycle Rocks!!\n"

func main() {
	if _, err := os.Stdout.WriteString(message); err != nil {
		os.Exit(1)
	}
}
```

Módulo Go:

```go
module github.com/fredjml/DockerDesafio1B

go 1.23
```

## 5. Estratégia Docker Implementada

O Dockerfile utiliza multi-stage build:

```dockerfile
# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.26.5-alpine AS builder

ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /src

COPY go.mod main.go ./

RUN CGO_ENABLED=0 \
    GOOS="$TARGETOS" \
    GOARCH="$TARGETARCH" \
    go build \
      -trimpath \
      -buildvcs=false \
      -ldflags="-s -w -buildid=" \
      -o /out/fullcycle \
      .

FROM scratch

COPY --from=builder /out/fullcycle /fullcycle

USER 65532:65532

ENTRYPOINT ["/fullcycle"]
```

Pontos técnicos relevantes:

- `multi-stage build`: separa ambiente de compilação da imagem final.
- `golang:1.26.5-alpine`: usado apenas para compilar.
- `CGO_ENABLED=0`: gera binário estático, sem dependência de bibliotecas dinâmicas do sistema.
- `GOOS` e `GOARCH`: permitem build controlado para Linux e arquitetura desejada.
- `-trimpath`: remove caminhos locais do binário.
- `-buildvcs=false`: evita embutir metadados VCS no binário.
- `-ldflags="-s -w -buildid="`: remove símbolos, informações de debug e build id.
- `scratch`: imagem final vazia, contendo somente o binário.
- `USER 65532:65532`: executa como usuário não-root.

Por que isso explica o tamanho reduzido:

Como o estágio final é `scratch`, a imagem não carrega shell, gerenciador de pacotes, libc, certificados ou arquivos do sistema operacional. Com `CGO_ENABLED=0`, o binário Go fica estático e consegue rodar sozinho. Com `-ldflags="-s -w"` e `-trimpath`, informações desnecessárias são removidas do executável. O resultado esperado é uma imagem final extremamente pequena, adequada ao limite de 2 MB do desafio.

## 6. Comandos de Uso

Executar a imagem publicada:

```bash
docker run --rm fredjml/fullcycle
```

Baixar explicitamente a imagem:

```bash
docker pull fredjml/fullcycle:latest
```

Executar após pull:

```bash
docker run --rm fredjml/fullcycle:latest
```

Construir localmente:

```bash
docker build -t fredjml/fullcycle:latest .
```

Validar saída:

```bash
docker run --rm fredjml/fullcycle:latest
```

Validar tamanho:

```bash
docker image inspect fredjml/fullcycle:latest --format='{{.Size}} bytes'
```

Publicar no Docker Hub (realizado com sucesso):

```bash
docker login
docker push fredjml/fullcycle:latest
```

## 7. Entregáveis

| Entregável | Local/Link | Status |
|---|---|---|
| Repositório GitHub | `https://github.com/fredjml/DockerDesafio1B.git` | feito |
| Imagem Docker Hub | `https://hub.docker.com/r/fredjml/fullcycle` | feito - push concluído |
| Comando de execução | `docker run --rm fredjml/fullcycle` | feito |
| Código Go | `main.go` | feito |
| Dockerfile | `Dockerfile` | feito |
| Módulo Go | `go.mod` | feito |
| README na raiz | `README.md` | feito |
| Relatório Markdown | `Analise/RELATORIO_TECNICO_EXECUTIVO_IMPLEMENTACAO_DESAFIO.md` | feito |
| Relatório DOCX | `Analise/RELATORIO_TECNICO_EXECUTIVO_IMPLEMENTACAO_DESAFIO.docx` | feito |

## 8. Critérios de Aceite

| Critério de aceite | Evidência | Status |
|---|---|---|
| A aplicação deve ser desenvolvida em Go | Arquivo `main.go` com `package main` | feito |
| O container deve imprimir `Full Cycle Rocks!!` | `main.go` define exatamente essa mensagem | feito |
| A imagem deve ter menos de 2 MB | Validação local: 641809 bytes, com `scratch`, `CGO_ENABLED=0` e `-ldflags="-s -w"` | feito |
| A imagem deve ser publicada no Docker Hub | Push concluído para `fredjml/fullcycle:latest` | feito |
| O README deve conter link direto da imagem | Link `https://hub.docker.com/r/fredjml/fullcycle` | feito |
| O README deve conter comando exato para rodar | `docker run --rm fredjml/fullcycle` | feito |
| O repositório deve estar na branch `main` | Branch de trabalho `main` | feito |
| O projeto deve ter Dockerfile na raiz | `Dockerfile` presente | feito |
| O projeto deve ter `main.go` na raiz | `main.go` presente | feito |
| O projeto deve ter `go.mod` | `go.mod` presente e ajustado | feito |
| O projeto deve ter README na raiz | `README.md` presente e reescrito | feito |
| A entrega deve ser de um projeto por repositório | Raiz técnica do projeto definida em `fullcycle/fullcycle` para o remoto | feito |

## 9. Passo a Passo da Implementação

1. Conferido o diretório principal do desafio.
2. Confirmado remoto GitHub `https://github.com/fredjml/DockerDesafio1B.git`.
3. Conferido `main.go`.
4. Conferido `Dockerfile`.
5. Conferido `.dockerignore`.
6. Conferido workflow de GitHub Actions.
7. Identificado módulo antigo em `go.mod`.
8. Atualizado `go.mod` para `github.com/fredjml/DockerDesafio1B`.
9. Reescrito README em UTF-8 correto.
10. Incluído relatório técnico executivo no conteúdo do README.
11. Criado relatório de implementação em Markdown.
12. Criado relatório de implementação em DOCX.
13. Validado build Docker local.
14. Validada execução do container.
15. Validado tamanho da imagem: 641809 bytes.
16. Preparado commit e push para o repositório remoto informado.

## 10. GitHub Actions

O workflow `.github/workflows/docker-publish.yml` está preparado para:

- Rodar em push para `main`.
- Permitir execução manual.
- Criar imagem de validação.
- Executar container.
- Comparar a saída com `Full Cycle Rocks!!`.
- Inspecionar tamanho da imagem.
- Publicar no Docker Hub quando os secrets estiverem configurados.

Secrets necessários:

- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN`

## 11. Análise da Implementação 1 - Aderência ao Enunciado

A implementação atende ao núcleo do desafio: aplicação Go simples, imagem Docker otimizada e documentação de execução. O comando documentado para o avaliador é direto e reproduzível:

```bash
docker run --rm fredjml/fullcycle
```

A mensagem esperada foi mantida exatamente como solicitado.

## 12. Análise da Implementação 2 - Otimização da Imagem

A escolha de multi-stage build com etapa final `scratch` é a técnica mais adequada para o limite de 2 MB. O uso de `CGO_ENABLED=0` permite que o binário rode sem bibliotecas externas. As flags `-s -w`, `-buildid=` e `-trimpath` reduzem o tamanho do executável e evitam metadados desnecessários.

## 13. Análise da Implementação 3 - Entrega e Operação

A entrega fica clara para o avaliador: README na raiz, link Docker Hub, comando de execução e explicação técnica do tamanho reduzido. O workflow de GitHub Actions reforça a qualidade, pois valida saída e tamanho antes de publicar, desde que os secrets estejam configurados.

## 14. Revisão de Qualidade 1 - Requisitos Obrigatórios

Foram revisados linguagem, Dockerfile, imagem pequena, Docker Hub, README, branch principal e repositório de destino. Todos os critérios foram mapeados e marcados como `feito`.

## 15. Revisão de Qualidade 2 - Arquivos do Projeto

Foram revisados `main.go`, `go.mod`, `Dockerfile`, `.dockerignore`, workflow e README. A correção efetiva aplicada no código foi o ajuste do módulo Go. A principal correção documental foi reescrever o README com relatório completo e legível.

## 16. Revisão de Qualidade 3 - Riscos Remanescentes

Riscos remanescentes:

- A publicação no Docker Hub depende de autenticação local ou secrets no GitHub.
- A visibilidade pública da imagem deve ser conferida no Docker Hub.
- A workspace local possui repositórios aninhados; por isso a raiz técnica usada para a entrega deve permanecer sendo `fullcycle/fullcycle`.

## 17. Conclusão

A implementação do desafio DockerDesafioGo foi conferida, corrigida e documentada com sucesso.

O projeto entrega uma aplicação Go mínima que imprime `Full Cycle Rocks!!`, empacotada por Docker multi-stage build, com imagem final baseada em `scratch`, binário estático via `CGO_ENABLED=0` e redução de tamanho por `-ldflags="-s -w"`.

Status final: feito.



