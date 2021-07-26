# ori

Bot para para ajudar a jogar o RPG Pocket Dragon

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/ori)](https://goreportcard.com/report/github.com/thiago-scherrer/ori)
[![Security check](https://github.com/thiago-scherrer/ori/actions/workflows/security_scan.yml/badge.svg?branch=main)](https://github.com/thiago-scherrer/ori/actions/workflows/security_scan.yml)
[![Test](https://github.com/thiago-scherrer/ori/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/thiago-scherrer/ori/actions/workflows/test.yml)

## Requerimentos

- API Key do telegram
- docker > 20.10.7

## Executando com o docker

1 - Para fazer o build, vá até a pasta deste repo então:

```sh
docker build . -t ori
```

2 - Exporte a chave do bot:

```sh
export TG_TOKEN="42:example"
```

3 - Execute o bot:

```sh
docker run -it -e TG_TOKEN=${TG_TOKEN} ori
```

## Créditos

- [bibliotecaelfica](https://www.bibliotecaelfica.org/pocket-dragon/pocket-dragon-manual-de-regras/)
