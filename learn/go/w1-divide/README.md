# 0A T1 — Chia 2 số, trả `(result, error)`

Giáo án: [`../../../docs/planning/phase-0a-week-1.md`](../../../docs/planning/phase-0a-week-1.md)

## Đề bài

> Viết hàm chia 2 số, trả về `(kết quả, error)` thay vì throw exception.
> Gọi hàm với input gây lỗi (chia 0), xử lý đúng bằng `if err != nil`.

Mục tiêu thật: hiểu **error as value** như một triết lý thiết kế, không phải cú pháp.

## File bạn tự tạo

`main.go` — `package main`, từ file trắng. Không copy mẫu.

## Tiêu chí đạt

- [ ] Hàm trả 2 giá trị, giá trị thứ hai kiểu `error`, đứng **cuối**
- [ ] Chia 0 → error **không nil**; không `panic`, không `fmt.Println` lỗi bên trong hàm
- [ ] Trường hợp hợp lệ → error là `nil`
- [ ] `main` gọi **cả hai** trường hợp và xử lý bằng `if err != nil`
- [ ] Khi có lỗi thì **không dùng** giá trị kết quả
- [ ] Không dùng `defer` (bài này không có resource nào cần đóng)

## Xác minh

```powershell
cd D:\calibrate\learn\go
go vet ./w1-divide
go run ./w1-divide
```

Đạt khi `go vet` sạch và output thể hiện rõ **cả** nhánh thành công lẫn nhánh lỗi.

## Bẫy hay mắc

`panic` thay vì trả error · dùng kết quả khi `err != nil` · nuốt lỗi bằng `_` ·
trả `string` thay cho `error` · thứ tự `(error, T)` · message lỗi viết hoa hoặc có dấu chấm cuối.
