# Phase 1 — Checkpoint C: gRPC nối 2 ngôn ngữ

Giáo án: [`../../../docs/planning/phase-1-checkpoints.md`](../../../docs/planning/phase-1-checkpoints.md)

**Khoá cho tới khi Checkpoint A và B đều được xác nhận đạt.**

## Đề bài

> Định nghĩa `.proto`, generate code Go + Python, Go client gọi Python server **thật qua network**
> (không phải hàm gọi trực tiếp).

Đây là checkpoint khó nhất Phase 1: lần đầu hai ngôn ngữ nói chuyện với nhau qua một hợp đồng.

## Dùng proto nào

Dùng **bản rút gọn của bạn tự viết** trong thư mục này, **không** dùng thẳng
[`../../../proto/scoring.proto`](../../../proto/scoring.proto).

Lý do: proto dự án có 2 RPC và 4 message — quá nhiều thứ hỏng cùng lúc cho lần đầu.
Bắt đầu bằng **1 RPC, 1 request, 1 response** của riêng bạn. Đọc proto dự án để lấy *hình mẫu*
(cách đặt `package`, `option go_package`, đánh số field), rồi rút gọn.

Khi bản rút gọn chạy được mới mở rộng dần cho tới lúc khớp `ComputeSimilarity` thật.

## Thư mục

- `proto/` — file `.proto` bạn tự viết (chưa dựng sẵn, tự tạo)
- `server-py/` — Python gRPC server, tái dùng hàm similarity từ Checkpoint B
- `client-go/` — Go client gọi sang

## Chuẩn bị

```powershell
# Python
cd D:\calibrate\learn\python
.\.venv\Scripts\Activate.ps1
pip install grpcio grpcio-tools

# Go
cd D:\calibrate\learn\go
go get google.golang.org/grpc google.golang.org/protobuf
```

Cần thêm `protoc` (hoặc dùng `python -m grpc_tools.protoc` cho phía Python và
`protoc-gen-go` / `protoc-gen-go-grpc` cho phía Go). Vướng ở bước cài thì hỏi ngay —
đây là chỗ tốn thời gian nhất và **không** phải phần kiến thức cần tự vật lộn.

## Tiêu chí đạt

- [ ] File `.proto` tự viết, có `syntax`, `package`, `option go_package`, 1 service 1 RPC
- [ ] Generate được code **cả hai** phía; code sinh ra **không** commit (đã có trong `.gitignore`)
- [ ] Python server chạy, lắng nghe trên 1 port cụ thể
- [ ] Go client kết nối qua `localhost:<port>`, gửi request, in response ra
- [ ] Server và client chạy ở **2 terminal khác nhau** — chứng minh đúng là qua network
- [ ] Logic similarity **tái dùng** từ Checkpoint B, không viết lại
- [ ] **Thử nghiệm bắt buộc:** tắt server rồi chạy client → client báo lỗi rõ ràng,
      không panic, không treo vô hạn. Có timeout bằng `context.WithTimeout`
- [ ] Giải thích được: vì sao dùng `.proto` thay vì gửi JSON qua HTTP

## Nối về dự án

Thử nghiệm "tắt server" ở trên chính là **failure mode đã chốt** trong `README.md`:
Python scoring service chết thì Go orchestrator degrade gracefully, không sập theo.
Ở đây bạn mới chỉ cần báo lỗi đúng cách. Phần degrade thật thuộc Phase 2.
