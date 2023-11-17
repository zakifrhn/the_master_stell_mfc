#! buat container
FROM golang:1.20.4-alpine

#! buat folder untuk menyimpan code
WORKDIR /gotestdev

#! Copy semua file
#? titik yg kedua lokasi penyimpanan
COPY . .

#! instal depedency
#? run kedua berfungsi untuk build app
RUN go mod download
RUN go build -v -o /gotestdev/goback ./cmd/main.go

#!open port
EXPOSE 8590

#! run app
ENTRYPOINT [ "/gotestdev/goback" ]

#! docker run --name <nama container> --net <info di networks> -e DB_HOST=<nama image postgres di docker> -p <port luar>:<port dalam> <nama image>

#! 1.
#! docker build -t itsfarhanz/gotests .  

#! 2.
#! docker run --name db-test-hiring -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Fazztrak2023 -e POSTGRES_DB=go_test_hiring -p 5470:5432 postgres:alpine3.18

#! 3.
#! docker run --name gotest --net bridge-db-app -e DB_HOST=db-test-hiring -e DB_NAME=go_test_hiring -e DB_USER=postgres -e DB_PASS=Fazztrak2023 -e DB_PORT=5432 -p 9092:8590 itsfarhanz/gotests