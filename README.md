basic-gin-service

Mali REST API (Go + Gorilla/Mux) za klijente i naloge.
Docker Compose podiže sve: Postgres, migracije i aplikaciju.

🚀 Quick start (Docker)
# pokreni bazu, migracije i app
docker compose up --build

Baza (za DBeaver / psql):

Host: localhost
Port: 5432
DB: basic_gin
User: postgres
Pass: 1234


🔌 API rute
Clients
GET    /clients
GET    /clients/{id}
POST   /clients
PUT    /clients
DELETE /clients/{id}

Accounts
GET    /clients/{id}/accounts
GET    /accounts/{id}
POST   /accounts
PUT    /accounts


Lista klijenata
curl http://localhost:8080/clients

Klijent po ID
curl http://localhost:8080/clients/1

Kreiraj klijenta
curl -X POST http://localhost:8080/clients \
  -H "Content-Type: application/json" \
  -d '{
        "first_name":"Pera",
        "last_name":"Perić",
        "email":"pera@example.com",
        "residence_address":"Beograd",
        "birth_date":"1990-01-01"
      }'

Izmeni klijenta
curl -X PUT http://localhost:8080/clients \
  -H "Content-Type: application/json" \
  -d '{
        "id": 1,
        "first_name":"Pera",
        "last_name":"Perić",
        "email":"pera+new@example.com",
        "residence_address":"Novi Beograd",
        "birth_date":"1990-01-01"
      }'

Obriši klijenta
curl -X DELETE http://localhost:8080/clients/2


Nalozi za klijenta
curl http://localhost:8080/clients/1/accounts

Račun po ID
curl http://localhost:8080/accounts/1

Kreiraj račun (za postojeći client_id)
curl -X POST http://localhost:8080/accounts \
  -H "Content-Type: application/json" \
  -d '{"client_id": 1}'

Ažuriraj račun (uplata/isplata)

deposit=true → uplata (povećava balans)

deposit=false → isplata (smanjuje balans)

curl -X PUT http://localhost:8080/accounts \
  -H "Content-Type: application/json" \
  -d '{
        "id": 1,
        "amount": 2500.00,
        "deposit": true
      }'