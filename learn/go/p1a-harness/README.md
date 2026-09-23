# Phase 1 — Checkpoint A: Mutation harness khung xương

Giáo án: [`../../../docs/planning/phase-1-checkpoints.md`](../../../docs/planning/phase-1-checkpoints.md)

**Khoá cho tới khi 0B được xác nhận đạt.**

## Đề bài

> Chạy N task giả song song bằng goroutine + `WaitGroup`, gom kết quả về, in tổng kết.
> Chạy thử `-race` để tự xác nhận không có race condition trong code checkpoint.

Đây là **khung xương thật** của mutation harness ở Phase 2 — nhưng task vẫn giả
(`time.Sleep` + kết quả ngẫu nhiên), chưa gọi ra ngoài.

## File bạn tự tạo

`main.go` — `package main`.

## Tiêu chí đạt

- [ ] Có struct `Result` mang ít nhất: id task, trạng thái pass/fail, thời gian chạy
- [ ] N task chạy **song song** (thời gian tổng nhỏ hơn hẳn tổng thời gian từng task)
- [ ] Gom kết quả bằng **channel**, không phải slice dùng chung
- [ ] `WaitGroup` đợi đủ N, đóng channel đúng chỗ (goroutine riêng: `wg.Wait()` rồi `close`)
- [ ] Task có thể lỗi thì trả `error`; harness **không** sập theo, vẫn gom đủ N kết quả
- [ ] In tổng kết: tổng số task, số pass, số fail, tỉ lệ
- [ ] `go run -race` sạch
- [ ] N khai báo ở **một chỗ duy nhất** (flag hoặc hằng), không hardcode rải rác

## Xác minh

```powershell
cd D:\calibrate\learn\go
go vet ./p1a-harness
go run -race ./p1a-harness
```

## Nối về dự án

Cái này ánh xạ thẳng vào `TODO(weekN-checkpoint)` trong
`go-orchestrator/cmd/calibrate/main.go`. **Không đụng file đó bây giờ** — nó là đích đến ở Phase 2.
