 - projekat pokretati iz root-a

*go run ./cmd/app*

 - run migracije

*migrate -path ./migrations -database "postgres://postgres:1234@localhost:5432/basic_gin?sslmode=disable" up*
