FROM golang:1.17-alpine

# Set the working directory
WORKDIR /app

COPY . .
RUN go mod tidy

# Copy the go.mod and go.sum files
# COPY go.mod go.sum ./
# COPY ./app/*.go ./


# # COPY . .  
# RUN go mod tidy
# RUN go build -o app .
# RUN chmod -R a+x ./app
# CMD ["./app"]