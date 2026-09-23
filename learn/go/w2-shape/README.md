# 0A T2 — Interface `Shape` + `Circle` + `Rectangle`

Giáo án: [`../../../docs/planning/phase-0a-week-2.md`](../../../docs/planning/phase-0a-week-2.md)

## Đề bài

> Định nghĩa interface `Shape` với method `Area()`. Viết 2 struct `Circle` và `Rectangle`
> cùng implement interface đó (không khai báo gì thêm, chỉ cần đúng method signature).
> Viết hàm nhận `Shape` làm tham số, gọi được với cả 2 struct.

Mục tiêu thật: chứng minh hiểu **implicit interface satisfaction** — Go không có từ khoá
`implements`. Đây là thứ làm việc mock trong test trở nên dễ, nền của Phase 2.

## File bạn tự tạo

`main.go` — `package main`.

## Tiêu chí đạt

- [ ] `type Shape interface { Area() float64 }` — interface chỉ chứa method, không chứa field
- [ ] `Circle` và `Rectangle` là struct riêng, **không** khai báo gì để nối với `Shape`
- [ ] Mỗi struct có method `Area()` đúng signature (tên, tham số, kiểu trả về khớp tuyệt đối)
- [ ] Có hàm nhận tham số kiểu `Shape`, không nhận struct cụ thể
- [ ] `main` gọi hàm đó với **cả hai** struct
- [ ] Giải thích được bằng lời: vì sao chọn value receiver hay pointer receiver cho `Area()`

## Mở rộng bắt buộc — để tự thấy cái bẫy

Thêm một `[]Shape` chứa cả hai loại, duyệt bằng `for range` và in tổng diện tích.
Đây chính là lý do interface tồn tại: **một slice chứa nhiều kiểu khác nhau**.

## Xác minh

```powershell
cd D:\calibrate\learn\go
go vet ./w2-shape
go run ./w2-shape
```

## Bẫy hay mắc

Dùng pointer receiver rồi gán value vào `Shape` (không thoả interface, lỗi compile khó hiểu) ·
đặt field vào interface · method viết thường `area()` thay vì `Area()` ·
tưởng phải khai báo gì đó để "đăng ký" struct với interface.
