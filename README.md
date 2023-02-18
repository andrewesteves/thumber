# Thumber (Thumbnails)

## Link do YouTube
https://youtu.be/e4797vRzrnM

## Dependências
https://imagemagick.org

## Comandos
```
// sequencial
go run .

// simultâneo
WORKERS=true go run .
```

Para verificar o tempo de execução de uma abordagem para a outra é importante remover os thumbnails criados anteriormente para manter o mesmo número de imagens a serem processadas.

```
find ./photos -type f -name "thumb*" -delete
```