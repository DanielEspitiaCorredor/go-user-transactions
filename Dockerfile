FROM golang:1.23-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify && \
    apk --no-cache add ca-certificates

COPY . ./

# Build
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /go-user-transactions -a -ldflags="-s -w"

# FROM scratch
# COPY --from=build /app/assets/ /assets/
# COPY --from=build /go-user-transactions /go-user-transactions
# COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt


EXPOSE 8080

# Run
CMD [ "/go-user-transactions" ]
