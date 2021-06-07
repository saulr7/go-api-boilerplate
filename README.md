# BoilerPlate Golang API

## Run

- cd certs && openssl genrsa -out app.rsa 1024
- openssl rsa -in app.rsa -pubout > app.rsa.pub
- go run main.go
