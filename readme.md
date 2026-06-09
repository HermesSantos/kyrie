# Kyrie
API HTTP em Go com frases e jaculatórias católicas.

![Kyrie](assets/kyrie.png)


## Como executar

```bash
go run ./cmd
```

Por padrão, a aplicação sobe em:

```text
http://localhost:1323
```

## Rotas

Todas as rotas da API usam o prefixo `/api`.

### Saúde

| Método | Rota | Descrição |
| --- | --- | --- |
| GET | `/api/ping` | Retorna `pong`. |

### Quotes

Objeto retornado:

```json
{
  "id": 1,
  "author": "Santa Teresa de Jesus",
  "quote": "..."
}
```

| Método | Rota | Descrição |
| --- | --- | --- |
| GET | `/api/quotes` | Lista todas as frases. |
| GET | `/api/quotes/random` | Retorna uma frase aleatória. |
| GET | `/api/quotes/:id` | Retorna uma frase pelo `id`. |

Exemplos:

```bash
curl http://localhost:1323/api/quotes
curl http://localhost:1323/api/quotes/random
curl http://localhost:1323/api/quotes/1
```

### Ejaculatories

Objeto retornado:

```json
{
  "id": 1,
  "category": "Deus Onipotente",
  "quote": "..."
}
```

| Método | Rota | Descrição |
| --- | --- | --- |
| GET | `/api/ejaculatories` | Lista todas as jaculatórias. |
| GET | `/api/ejaculatories/random` | Retorna uma jaculatória aleatória. |
| GET | `/api/ejaculatories/categories` | Lista as categorias disponíveis. |
| GET | `/api/ejaculatories/category/:category` | Lista jaculatórias por categoria. |
| GET | `/api/ejaculatories/:id` | Retorna uma jaculatória pelo `id`. |

Exemplos:

```bash
curl http://localhost:1323/api/ejaculatories
curl http://localhost:1323/api/ejaculatories/random
curl http://localhost:1323/api/ejaculatories/categories
curl http://localhost:1323/api/ejaculatories/category/Jesus%20Cristo
curl http://localhost:1323/api/ejaculatories/1
```

## Dados

Os dados usados pela API ficam em:

| Arquivo | Conteúdo |
| --- | --- |
| `data/quotes.json` | Frases com autor. |
| `data/ejaculatories.json` | Jaculatórias organizadas por categoria. |
