## Golang-Labs-Auction
Abertura e fechamento do Leilão - Go Routines

Como usar
Clone o repositório:

```bash

git clone git@github.com:felipecardosodeoliveira/Golang-Labs-Auction.git
cd Golang-Labs-Auction

Dentro da pasta raiz, execute o comando para iniciar os serviços com Docker Compose:
docker compose up

Para rodar o teste da nova implementação:

go test ./...

Para criar uma auction:

POST http://localhost:8080/auction

{
"product_name": "iphone",
"category": "eletronics",
"description": "auction 1",
"condition": 0
}

HTTP/1.1 201 Created
Date: Sat, 16 Nov 2024 19:49:20 GMT
Content-Length: 0


Para buscar as auctions:

GET http://localhost:8080/auction?status=0

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 16 Nov 2024 19:49:36 GMT
Content-Length: 383

[
  {
    "id": "77f3d8d2-e418-46ee-b986-11284a91f234",
    "product_name": "notebook",
    "category": "eletronics",
    "description": "a normal notebook",
    "condition": 0,
    "status": 1,
    "timestamp": "2024-11-16T19:27:28Z"
  },
  {
    "id": "76ad444c-4032-47d2-b7c6-998acc574831",
    "product_name": "iphone",
    "category": "eletronics",
    "description": "auction 1 description",
    "condition": 0,
    "status": 0,
    "timestamp": "2024-11-16T19:49:20Z"
  }
]

Para buscar auction por id:

GET http://localhost:8080/auction/winner/77f3d8d2-e418-46ee-b986-11284a91f234