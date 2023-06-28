FROM golang:1.19 as build

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy




# COPY . .  
# RUN go mod download
# RUN go build -o app .

# FROM scratch  
# COPY --from=build /app/app .
# CMD ["./app"]