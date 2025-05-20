# 1. Go o'rnatilgan image'dan foydalanamiz
FROM golang:1.23-alpine AS builder

# 2. Ishchi katalogni yaratamiz
WORKDIR /app

# 3. Go mod va go sum fayllarini copy qilamiz
COPY go.mod go.sum ./

# 4. Go modlarni o'rnatamiz
RUN go mod tidy

# 5. Manba kodini konteynerga copy qilamiz
COPY . .

# 6. Go ilovasini quramiz
RUN go build -o main .

# 7. Ishga tushirish uchun minimal image
FROM alpine:latest  

# 8. Postgresql uchun .env faylni yuklaymiz
ENV DB_HOST=postgres_to_do
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_NAME=to_do
ENV DB_PASSWORD=hamidjon4424
ENV TO_DO=:8087

WORKDIR /root/

# 9. Qurilgan ilovani ko'chiramiz
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# 10. Portni ochamiz
EXPOSE 8087

# 11. Ilovani ishga tushiramiz
CMD ["./main"]
