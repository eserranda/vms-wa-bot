FROM golang:1.24-alpine

# Mengupdate dan mengupgrade paket, serta menambahkan dependensi build
RUN apk update && apk upgrade && apk add --no-cache build-base sqlite

# Mengatur CGO untuk mengaktifkan SQLite
ENV CGO_ENABLED=1

# menggunakan toolchain otomatis jika diperlukan.
ENV GOTOOLCHAIN=auto

# Mengatur direktori kerja
WORKDIR /go/src/app

# Copy file go.mod dan go.sum terlebih dahulu, lalu download dependensi
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh kode ke dalam image
COPY . .

# Menjalankan aplikasi
CMD ["go", "run", "main.go"]
