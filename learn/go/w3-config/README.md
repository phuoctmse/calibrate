# 0A T3 — `ParseConfig` + error wrapping

Giáo án: [`../../../docs/planning/phase-0a-week-3.md`](../../../docs/planning/phase-0a-week-3.md)

## Đề bài

> Viết 1 hàm `ParseConfig` giả lập đọc config, cố tình gây lỗi ở 1 bước con,
> wrap lỗi đó với context (`failed to parse field X: %w`), ở `main` in ra full error chain
> và dùng `errors.Is` kiểm tra loại lỗi gốc.

Mục tiêu thật: hiểu **error chain**. Ở Phase 2, khi orchestrator gọi gRPC hỏng, bạn phải trả lời
được "hỏng ở tầng nào" — chỉ error wrapping mới cho phép điều đó.

## File bạn tự tạo

`main.go` — `package main`.

## Tiêu chí đạt

- [ ] Có **sentinel error** khai báo ở package level: `var ErrMissingField = errors.New(...)`
- [ ] Có ít nhất **2 tầng** hàm: `ParseConfig` gọi một hàm con, hàm con sinh lỗi gốc
- [ ] Tầng trên wrap bằng `fmt.Errorf` với verb `%w`, không phải `%v`
- [ ] `main` in ra full chain: thấy được **cả** context lẫn lỗi gốc trong cùng một chuỗi
- [ ] `errors.Is(err, ErrMissingField)` trả `true` qua ít nhất 2 tầng wrap
- [ ] Config biểu diễn bằng `map[string]string`, đọc key bằng dạng `v, ok := m[key]`

## Mở rộng bắt buộc

Đổi một chỗ `%w` thành `%v` rồi chạy lại: `errors.Is` phải trở thành `false`.
**Tự mắt thấy** khác biệt rồi hãy đổi về `%w`.

## Xác minh

```powershell
cd D:\calibrate\learn\go
go vet ./w3-config
go run ./w3-config
```

## Bẫy hay mắc

Dùng `%v` tưởng nhầm là wrap · so sánh error bằng `==` hoặc so chuỗi thay vì `errors.Is` ·
tạo error mới ở tầng trên làm mất lỗi gốc · không kiểm tra `ok` khi đọc map
(map trả zero value cho key không tồn tại, im lặng, không panic).
