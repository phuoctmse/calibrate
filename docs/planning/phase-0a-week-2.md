# Phase 0A — Tuần 2: Struct, method/receiver, interface, composition

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0A — Tuần 2".
Bài tập: [`learn/go/w2-shape/`](../../learn/go/w2-shape/README.md).

**Mục tiêu thật:** hiểu vì sao Go bỏ hẳn class và kế thừa, mà vẫn làm được mọi thứ OOP làm —
và vì sao cách đó khiến code **dễ test hơn hẳn**. Đây là tuần quan trọng nhất với vai trò SDET.

> Trạng thái: khung + đề bài đã chốt. Phần giảng chi tiết (ví dụ, so sánh, hỏi đáp)
> diễn ra trực tiếp khi mở tuần này, giống cách đã làm ở Tuần 1.

---

## 1. Struct — gom dữ liệu, không gom hành vi

```go
type Circle struct {
	Radius float64
}
c := Circle{Radius: 2}
```

Struct chỉ là **dữ liệu**. Không có `private`, không có constructor, không có kế thừa.
Chữ hoa đầu tên = exported (ra ngoài package thấy được), chữ thường = không. Chỉ có vậy.

Zero value của struct là struct có mọi field ở zero value — không phải `nil`.
Nghĩa là `var c Circle` đã dùng được ngay. Nối thẳng với bài học Tuần 1.

## 2. Method và receiver — chỗ sai phổ biến nhất

Method không nằm trong struct. Nó là hàm có thêm **receiver** đứng trước tên:

```go
func (c Circle) Area() float64 { ... }   // value receiver
func (c *Circle) Scale(k float64) { ... } // pointer receiver
```

Quy tắc quyết định, học thuộc:

| Dùng | Khi |
|---|---|
| value receiver `(c Circle)` | chỉ **đọc** field, struct nhỏ |
| pointer receiver `(c *Circle)` | cần **sửa** field, hoặc struct lớn (tránh copy) |

Value receiver nhận **bản sao**. Sửa field trong đó thì bên ngoài không thấy gì đổi —
lỗi im lặng, không báo compile. Đây đúng là dạng bug mà góc nhìn QA phải bắt:
*hàm chạy xong, không lỗi, nhưng trạng thái không đổi.*

Quy ước: một kiểu thì **nhất quán** một loại receiver, đừng trộn.

## 3. Interface — implicit satisfaction

```go
type Shape interface {
	Area() float64
}
```

Không có `implements`. Kiểu nào có đủ method đúng signature thì **tự động** thoả interface.

Hệ quả cho việc test — lý do thật sự tuần này quan trọng:

- Định nghĩa interface được ở **phía người dùng**, không phải phía tác giả thư viện.
- Muốn mock cái gì thì khai một interface nhỏ đúng chỗ cần, rồi viết struct giả thoả nó.
  Không cần thư viện mock, không cần sửa code gốc.
- Ở Phase 2, Go orchestrator gọi Python scoring service qua interface → test được
  toàn bộ orchestrator mà **không cần** Python server chạy.

Quy ước: interface càng **nhỏ** càng mạnh. 1 method là lý tưởng.

## 4. Cái bẫy receiver ↔ interface

Nếu `Area()` khai bằng **pointer receiver** thì chỉ `*Circle` thoả `Shape`, `Circle` thì không:

```go
var s Shape = Circle{}   // lỗi compile nếu Area dùng pointer receiver
var s Shape = &Circle{}  // ok
```

Thông báo lỗi Go đưa ra ở đây khó đọc với người mới. Gặp thì nhớ lại mục này.

## 5. Composition qua embedding

```go
type Base struct{ ID string }
type Job struct {
	Base            // embedded, không có tên field
	Name string
}
j.ID  // truy cập thẳng, "thăng cấp" từ Base
```

Trông như kế thừa nhưng **không phải**: `Job` không "là" `Base`, nó chỉ *chứa* `Base` và mượn
tên. Không có override, không có gọi `super`, không có đa hình qua struct — đa hình chỉ qua interface.

---

## Checklist review (trợ lý dùng khi chấm)

Xếp theo mức nghiêm trọng:

1. Khai gì đó để "đăng ký" struct với interface — hiểu nhầm nền tảng, chặn ngay.
2. Không giải thích được vì sao chọn value/pointer receiver — chưa nắm mục 2, chặn.
3. Đặt field vào interface, hoặc interface phình to nhiều method không cần thiết.
4. Hàm nhận `Circle` / `Rectangle` cụ thể thay vì nhận `Shape` — mất sạch ý nghĩa bài.
5. Trộn value và pointer receiver trên cùng một kiểu.
6. Bỏ qua phần mở rộng `[]Shape` — đó là chỗ interface mới thật sự chứng minh giá trị.
7. Tự viết `math.Pi` bằng `3.14` — vụn, nhưng nhắc.

## Hỏi lại để xác nhận hiểu (không phải code)

- "Nếu muốn test một hàm gọi service ngoài mà không cần service đó chạy, bạn làm thế nào?"
  → trả lời được bằng interface + struct giả thì tuần này coi như đạt.

---

**Qua tuần khi:** bài chạy đúng + review sạch + trả lời được câu hỏi trên.
→ Mở T3: slice, map, pointer, error wrapping.
