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

- Truy cập **Swagger UI** để kiểm thử các API: [http://DOMAIN:PORT/swagger/index.html]
  vd: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Thiết kế & Quyết định kỹ thuật

### Database

- Chọn **MySQL** vì:
  - Hỗ trợ **unique index** tốt, tránh trùng code.
  - Dễ tích hợp với Gorm.
  - Phổ biến, dễ deploy production.

### Thiết kế API

- RESTful API là chuẩn thiết kế API dựa trên HTTP, dễ triển khai, với URL rõ ràng, dễ đọc và dễ debug. Nó hỗ trợ khả năng mở rộng, caching hiệu quả, tách client-server, và tương thích với nhiều ngôn ngữ cũng như nền tảng.

### Thuật toán tạo mã ngắn

- Sử dụng **Base62** `[a-zA-Z0-9]` để mã ngắn dễ đọc và gọn.
- Sinh ngẫu nhiên số từ **0 → 61** (tương ứng 62 ký tự) với **length = 8** bằng **crypto/rand**.
- Dựa vào số ngẫu nhiên chọn ký tự tương ứng từ Base62 → tạo chuỗi ngắn, bảo đảm **khó đoán**.
- Dễ kết hợp với cơ chế **xử lý trùng lặp bằng retry**, đảm bảo mỗi mã **duy nhất**.

### Xử lý conflict / duplicate

- Trường `code` trong database được đặt **unique** để tránh trùng lặp, đặc biệt khi có nhiều request cùng lúc.
- Trước khi lưu, hệ thống **kiểm tra xem code đã tồn tại trong database**.
- Nếu code vừa sinh đã tồn tại:
  - Sinh lại một mã mới, tối đa 5 lần để tránh tình trạng timeout.
- Cơ chế này đảm bảo **mỗi short code là duy nhất** và vẫn giữ tính **ngẫu nhiên**.

### Xử lý kiểm tra URL hợp lệ

- Kiểm tra url không phải là url của server
- Kiểm tra phải bắt đầu bằng protocol hợp lệ ( http/https )
- Kiểm tra phải có domain hợp lệ
- Kiểm tra url không có 1 số ký tự đặt biệt

### Xử lý nếu database lớn số short link nhiều

- Index và khóa chính: Trường id là khóa chính, còn code được đánh dấu unique, MySQL tự động tạo index, giúp truy vấn code nhanh, ngay cả với số lượng lớn bản ghi.
- Phân trang: Khi lấy danh sách short link hỗ trợ phân trang, tránh load toàn bộ dữ liệu, tăng tốc độ truy vấn và giảm áp lực lên database.

### Xủ lý bảo mật

- Short link khó đoán: Mã rút gọn được tạo ngẫu nhiên (Base62 + crypto/rand), đảm bảo không thể giải mã trực tiếp ra URL gốc.
- Quyền truy cập theo tài khoản: Khi tích hợp chức năng tài khoản, chỉ người dùng đã xác thực mới có thể xem chi tiết, quản lý hoặc chỉnh sửa short link của mình.

### Xủ lý khả năng mở rộng

- RESTful API phù hợp cho việc dễ dàng mở rộng và phát triển các tính năng
- Nếu traffic tăng cao có thể dùng redis để cache dữ liệu giúp tăng tốc độ truy vấn cũng như giảm tải cho database

## Trade-offs

- **Chọn MySQL**: Phổ biến, dễ triển khai, hỗ trợ unique index tốt để tránh trùng code.
- Nhược điểm là không có một số tính năng của NoSQL nhưng vẫn phù hợp với ứng dụng này.

- **Lưu URL cùng với code duy nhất**: Đảm bảo mỗi short code là duy nhất và tránh duplicate.
- Nhược điểm: khi traffic cao, việc kiểm tra database liên tục có thể gây quá tải. Có thể giảm tải bằng cách **sử dụng cache** và **tối ưu truy vấn**.

- **Chọn Gorm thay vì các thư viện khác** vì nó hỗ trợ ORM đầy đủ, ánh xạ dữ liệu từ struct sang table trong database. Gorm tự động sinh SQL, hạn chế việc viết SQL thủ công, giúp phát triển nhanh chóng, code dễ đọc và dễ bảo trì.
- Nhược điểm: với một số query phức tạp, hiệu năng có thể hơi thấp và SQL sinh ra không tối ưu bằng việc viết tay.

- **Sử dụng random code ngẫu nhiên**: mỗi URL đều được sinh một short code ngẫu nhiên, ngay cả khi cùng một URL được nhiều người tạo. Cách này giúp **tránh xung đột khi nhiều user quản lý và truy cập cùng một short URL**.
- Nhược điểm: cùng một URL gốc sẽ có nhiều short code khác nhau, làm **tăng dung lượng lưu trữ trong database**.

  - **Ý tưởng giải pháp**: phân thành 2 loại short URL
    - **Short URL do User tạo**: short code được sinh **ngẫu nhiên, không trùng**, phù hợp khi nhiều user quản lý và truy cập cùng một URL gốc.
    - **Short URL do Customer (người dùng không có tài khoản)**: nếu trùng URL thì **trùng short code**, giúp tiết kiệm dung lượng lưu trữ trong database.
  - **Kết quả mong muốn**: vừa đáp ứng nhu cầu của User, vừa **hạn chế lưu trữ dư thừa trong database**.

- **Sử dụng `crypto/rand` thay vì `math/rand`**: Mặc dù `crypto/rand` chậm hơn `math/rand`, nhưng đảm bảo **bảo mật và khó đoán**, phù hợp cho short URL.

- **Sử dụng Base62**: Tạo URL ngắn, dễ đọc, dễ copy, dễ nhớ.
- Nhược điểm là không dùng được ký tự đặc biệt, nhưng vẫn đủ ký tự cho nhu cầu tạo short code.

## Challenges

- Lựa chọn thuật toán **generate short code** và thiết kế logic phù hợp cho việc tạo và lưu trữ short URL.
  - Mỗi thuật toán và logic tạo code đều có **ưu nhược điểm riêng**, ảnh hưởng đến toàn bộ luồng xử lý của hệ thống.
  - Nghiên cứu kỹ các ưu nhược điểm của thuật toán, tham khảo **logic xử lý từ các nền tảng khác**.
  - Trong quá trình giải quyết, học được cách **tránh trùng dữ liệu, tiết kiệm dung lượng database** và hiểu sâu hơn về **trade-offs của từng giải pháp**.

## Limitations & Improvements

### Limitations

- Chưa có giao diện để người dùng tương tác
- Chưa có tích hợp thêm tài khoản cũng như quản lý và tạo short link theo tài khoản
- Chưa có các chức năng như phân tích thống kê về truy cập của short link
- Chưa có format kiểu response trả về thống nhất
- Chưa có caching → mỗi request đều query trực tiếp database.
- Lượng lớn link cũ tạo ra làm tốn dung lượng database

### Improvements

- Viết code format kiểu response trả về thống nhất
- Làm thêm phần giao diện để người dùng tương tác
- Làm thêm các chức năng của tài khoản cũng như phần xác thực và phân quyền ( có thể làm phân cấp loại tài khoản )
- Làm thêm các chức năng phân tích thống kê về truy cập của short link
- Tích hợp redis vào để cache dữ liệu giúp tối ưu hệ thống, tối ưu các câu lệnh SQL
- Làm thêm chức năng hạn sử dụng cho short link để hạn chế dữ liệu rác trong database
- Nghiên cứu và phát triển thêm các tính năng như đăng nhập google,tích hợp thanh toán online để nâng cấp tài khoản...
