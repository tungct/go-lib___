# go-msqueue-worker

Server nhận Request message (POST), đẩy message vào MessageQueue chờ xử lý, WorkerPool thực hiện xử lý message trong queue và ghi ra file

## Yêu cầu
- Go 1.9 hoặc thấp hơn, ubuntu 16.04
- Thư viện cài đặt thêm :

```bash
go get github.com/gorilla/mux
```

## Hướng dẫn

