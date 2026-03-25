# Short URL Service

## Mô tả bài toán

Dự án này là một **URL shortening service**.

- Cho phép người dùng nhập một URL dài và nhận về **short url** (ví dụ `abc123XY`).
- Khi truy cập short url, sẽ **redirect** đến URL gốc.
- Cho phép quản lý các short url đã tạo,
- Có thể thống kê lượt truy cập vào short url, bao gồm thông tin như: ip, thiết bị, ngôn ngữ, thời gian truy cập, v.v.
- Lưu ý của các tính năng:
  - Tạo short URL tự động hoặc tự truyền alias.
  - Tránh trùng short url giữa các URL gốc.
  - Không thể giải mã short url ra URL gốc.

Mục tiêu:

- Xây dựng một backend **URL shortening service** để người dùng có thể tạo ra short link từ url gốc
- Tối ưu **hiệu năng** và **bảo mật** khi sinh mã ngắn.
- API rõ ràng, dễ dùng cho client.

## Cách chạy project

### Yêu cầu

- Go >= 1.25
- MySQL
- Git

### Cấu hình

- Tạo file `.env`:

```
# Database configuration
DB_USER=root               # Tên user kết nối MySQL
DB_PASSWORD=               # Mật khẩu MySQL
DB_HOST=127.0.0.1          # Địa chỉ host MySQL
DB_PORT=3306               # Cổng MySQL
DB_NAME=short_url_api      # Tên database

# Server configuration
PORT=8080                  # Cổng server sẽ chạy
DOMAIN_SHORT=localhost     # Domain gốc cho short URL
PROTOCOL=http              # Giao thức (http hoặc https)
```

### Run project

```bash
# Clone repo
git clone https://github.com/PhuocTai1704/short-url-api.git
cd short-url-api

# Cài dependencies
go mod tidy

# Chạy server
go run cmd/api/main.go
```

### Test API
### Demo: https://short-url-api-production.up.railway.app/swagger/index.html#
- Truy cập **Swagger UI** để kiểm thử các API: [http://DOMAIN:PORT/swagger/index.html]
  vd: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

