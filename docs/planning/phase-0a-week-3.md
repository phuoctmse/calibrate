# Phase 0A — Tuần 3: Slice, map, pointer, error wrapping

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0A — Tuần 3".
Bài tập: [`learn/go/w3-config/`](../../learn/go/w3-config/README.md).

**Mục tiêu thật:** hai thứ. (1) Hiểu slice/map đủ sâu để không dính bug chia sẻ bộ nhớ.
(2) Biết dựng **error chain có ngữ cảnh** — thứ quyết định ở Phase 2 bạn có debug nổi
một lỗi gRPC hay không.

> Trạng thái: khung + đề bài đã chốt. Phần giảng chi tiết diễn ra khi mở tuần.

---

## 1. Array vs slice

```go
var a [3]int          // array — độ dài là MỘT PHẦN của kiểu, hiếm dùng
s := []int{1, 2, 3}   // slice — thứ dùng hàng ngày
```

Slice là một cấu trúc 3 phần: **con trỏ tới mảng nền, length, capacity**.
Nó **không** chứa dữ liệu, nó *trỏ* tới dữ liệu.

Hệ quả phải tự tay thấy:

```go
a := []int{1, 2, 3}
b := a[:2]
b[0] = 99     // a[0] cũng thành 99 — chung mảng nền
```

`append` khi còn capacity thì ghi đè tại chỗ; khi hết capacity thì **cấp mảng mới** và
từ đó hai slice không còn liên quan nhau nữa. Cùng một dòng `append`, hai hành vi khác nhau
tuỳ trạng thái runtime. Đây là nguồn bug slice kinh điển của Go.

Cách an toàn khi muốn tách hẳn: `copy(dst, src)` hoặc `slices.Clone`.

## 2. Map

```go
m := map[string]string{"host": "localhost"}
v, ok := m["port"]   // v = "", ok = false
```

Điểm chết người: đọc key không tồn tại **không lỗi, không panic** — trả zero value.
`""` là "không có key" hay "có key với giá trị rỗng"? Chỉ `ok` mới phân biệt được.
Bỏ `ok` là tự tạo bug im lặng. Bài tập tuần này ép dùng dạng `v, ok`.

Thêm: map **phải** `make` trước khi ghi (`var m map[string]string` là nil map — đọc được,
ghi là panic). Và thứ tự duyệt map là **ngẫu nhiên có chủ đích** — đừng phụ thuộc vào nó.

## 3. Pointer

```go
x := 5
p := &x    // lấy địa chỉ
*p = 10    // sửa qua con trỏ → x == 10
```

Go không có số học con trỏ. Dùng pointer khi: cần sửa dữ liệu của caller, hoặc tránh copy
struct lớn, hoặc cần biểu diễn "không có giá trị" (`nil`).

Nối lại Tuần 2: pointer receiver chính là chuyện này. Nối tới đây thì hai tuần khớp nhau.

Lưu ý: slice và map **đã** chứa con trỏ bên trong — truyền chúng vào hàm là hàm sửa được
dữ liệu gốc, dù trông như truyền value. Khác hẳn struct.

## 4. Error wrapping — phần quan trọng nhất tuần

Tuần 1 học *trả* lỗi. Tuần này học *dựng ngữ cảnh* cho lỗi.

```go
var ErrMissingField = errors.New("missing field")

// tầng dưới
return fmt.Errorf("field %q: %w", name, ErrMissingField)
// tầng trên
return fmt.Errorf("parse config: %w", err)
```

`%w` = wrap, giữ được lỗi gốc bên trong. `%v` = chỉ nhúng chữ, **mất** lỗi gốc.
Nhìn output in ra thì hai cái giống hệt nhau — khác biệt chỉ lộ ra khi gọi `errors.Is`.
Bài tập bắt bạn tự đổi `%w` thành `%v` để thấy tận mắt.

Kiểm tra lỗi:

| Cách | Dùng khi |
|---|---|
| `errors.Is(err, ErrX)` | hỏi "có phải loại lỗi này không", xuyên qua mọi tầng wrap |
| `errors.As(err, &target)` | cần lấy **dữ liệu** trong error (mã lỗi, field nào hỏng) |
| `err == ErrX` | gần như luôn sai — vỡ ngay khi ai đó wrap thêm một tầng |
| so sánh `err.Error()` | luôn sai. Message là cho người đọc, không phải hợp đồng |

Quy ước viết message wrap: **không** lặp lại chữ "error"/"failed to" ở mọi tầng,
vì chúng nối chuỗi lại với nhau. Mỗi tầng chỉ thêm đúng phần ngữ cảnh của mình.

Góc nhìn QA: error chain chính là **traceability**. `"parse config: field port: missing field"`
cho biết hỏng ở đâu và hỏng cái gì. `"error"` thì không cho biết gì.

---

## Checklist review (trợ lý dùng khi chấm)

1. Dùng `%v` mà tưởng là wrap — trọng tâm cả tuần, chặn ngay.
2. So sánh error bằng `==` hoặc so chuỗi thay vì `errors.Is`.
3. Tạo error mới ở tầng trên, nuốt mất lỗi gốc.
4. Đọc map không kiểm tra `ok`.
5. Chỉ có 1 tầng hàm → không có chain nào để nhìn, làm lại cho đủ 2 tầng.
6. Sentinel error khai bên trong hàm thay vì package level (không ai `errors.Is` được).
7. Message wrap lặp "failed to" ở mọi tầng.
8. Không làm phần mở rộng `%w` → `%v` — bắt làm, đó là chỗ học thật.

## Hỏi lại để xác nhận hiểu

- "Hai slice cùng trỏ vào một mảng nền, sửa cái này cái kia đổi theo. Khi nào thì **hết** đổi theo?"
  → trả lời được bằng `append` vượt capacity thì đạt.

---

**Qua tuần khi:** bài chạy đúng + review sạch + trả lời được câu hỏi trên.
→ Mở Phase 0B: concurrency. Hết Phase 0A.
