# 0B — Race condition thật → `-race` → fix bằng `Mutex`

Giáo án: [`../../../docs/planning/phase-0b-concurrency.md`](../../../docs/planning/phase-0b-concurrency.md)

## Đề bài

> Viết chương trình có race condition thật (2 goroutine cùng tăng 1 biến đếm không có lock),
> chạy `go run -race` xem Go tự báo lỗi. Sau đó fix bằng `sync.Mutex`, chạy lại `-race`
> để xác nhận hết lỗi.

Đây là bài **thấy bug thật rồi tự fix**. Không được bỏ qua bước chạy bản hỏng.

## File bạn tự tạo

Giữ **cả hai** để so sánh:

- `broken/main.go` — bản có race, `package main`
- `fixed/main.go` — bản đã fix, `package main`

Hai thư mục đã dựng sẵn, chỉ thiếu file code.

## Tiêu chí đạt — bản `broken`

- [ ] Từ 2 goroutine trở lên cùng ghi vào 1 biến đếm chung, **không** lock
- [ ] Dùng `sync.WaitGroup` để `main` đợi xong — không dùng `time.Sleep` để "đợi"
- [ ] Số vòng lặp đủ lớn (từ 1000 mỗi goroutine) để sai lệch thấy rõ
- [ ] Chạy `-race` thấy Go báo `DATA RACE`. Dán lại output đó khi nhờ review.
- [ ] Chạy thường vài lần: tổng in ra **khác nhau giữa các lần**

## Tiêu chí đạt — bản `fixed`

- [ ] Dùng `sync.Mutex`, `Lock` / `Unlock` bao đúng vùng ghi chung
- [ ] `Unlock` bằng `defer` trong đúng scope — đây là chỗ `defer` **đáng dùng**
- [ ] `-race` sạch
- [ ] Tổng in ra **luôn đúng** qua nhiều lần chạy

## Xác minh

```powershell
cd D:\calibrate\learn\go
go run -race ./b-race/broken
go run -race ./b-race/fixed
```

Bản đầu PHẢI báo `WARNING: DATA RACE`. Bản sau PHẢI sạch.

Ghi chú: `-race` trên Windows cần gcc/mingw. Nếu báo thiếu compiler thì nói ngay,
xử lý trước khi bắt đầu bài.

## Bẫy hay mắc

Dùng `time.Sleep` thay `WaitGroup` · quên `wg.Add(1)` trước khi `go` ·
truyền `WaitGroup` bằng value thay vì pointer (bị copy thì treo vĩnh viễn) ·
`Lock` bao cả phần không cần khiến mất hết lợi ích song song ·
tưởng chạy thường không lỗi nghĩa là không có race.
