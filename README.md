# ori

Bot para para ajudar a jogar o RPG Pocket Dragon

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
