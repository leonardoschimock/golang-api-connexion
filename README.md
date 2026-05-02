# 📚 Google Books API - Go (Golang)

Este projeto é uma API desenvolvida em **Golang** que realiza consultas na **Google Books API**, retornando informações sobre livros com base em um termo de busca.

O objetivo deste repositório é estudo prático de:
- Estruturação de APIs em Go
- Uso de rotas (mux/router)
- Consumo de APIs externas
- Tratamento de erros e respostas HTTP

---

## 🚀 Tecnologias utilizadas

- Go (Golang)
- net/http
- Encoding/JSON
- API pública do Google Books

---

## 📁 Estrutura do projeto

```
.
├── main.go
├── routes/
│   └── routes.go
├── handlers/
│   └── books.go
├── services/
│   └── google_books.go
└── README.md
```

---

## ⚙️ Como funciona

A API expõe uma rota que recebe um parâmetro de busca (`q`) e faz uma requisição para a Google Books API:

```
GET /books?q=harry+potter
```

A aplicação:
1. Recebe a query
2. Encoda a URL corretamente
3. Faz a requisição externa
4. Retorna os dados em JSON

---

## 🔌 Exemplo de requisição

### Request

```
GET http://localhost:8080/books?q=clean+code
```

### Response (resumido)

```json
{
  "items": [
    {
      "volumeInfo": {
        "title": "Clean Code",
        "authors": ["Robert C. Martin"]
      }
    }
  ]
}
```

---

## 🧪 Testes

Você pode testar usando:

- Postman
- Insomnia
- Curl

---

## ⚠️ Possível erro comum

### Erro 429 - Too Many Requests

```json
{
  "error": {
    "code": 429,
    "message": "Quota exceeded..."
  }
}
```

Isso significa que o limite diário da API foi atingido.

### 💡 Soluções:
- Aguardar o reset da quota diária
- Utilizar uma API Key
- Criar um novo projeto no Google Cloud

---

## 🔑 Uso de API Key (Recomendado)

Adicione sua chave na URL:

```go
googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + query + "&key=SUA_API_KEY"
```

> ⚠️ **Nunca exponha sua API Key em repositórios públicos**

---

## 📌 Objetivo

Este projeto faz parte da disciplina de extensão "backend com Go", focando em boas práticas e integração com APIs externas.
