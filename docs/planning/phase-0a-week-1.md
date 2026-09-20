# Phase 0A — Tuần 1: Type system, control flow, functions

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0A — Tuần 1".

**Mục tiêu thật của tuần này:** hiểu `(result, error)` + `if err != nil` như một **triết lý thiết kế**,
không phải cú pháp cần thuộc. Mọi dòng Go viết sau này — kể cả orchestrator gọi gRPC ở Phase 2 —
đều dựng trên nó.

---

## 1. Zero value — Go không có "biến chưa khởi tạo"

```go
var n int      // 0
var s string   // ""
var ok bool    // false
var p *int     // nil
```

Khai báo là có giá trị ngay. Không có `None` / `undefined` như Python / JS.

Đây **không phải chi tiết vụn**. Nó là nền của cả tuần: kiểu `error` có zero value là `nil`,
nên **"không có lỗi" là trạng thái mặc định tự nhiên** của một biến error — không cần khởi tạo gì.

Hai cách khai báo:

```go
var count int = 5   // đầy đủ
count := 5          // suy ra kiểu, CHỈ dùng được bên trong function
```

`:=` là *khai báo mới + gán*. `=` là *gán vào biến đã có*. Nhầm hai cái này là lỗi compile
phổ biến nhất tuần đầu.

## 2. `for` là vòng lặp duy nhất

Go không có `while`, không có `do-while`:

```go
for i := 0; i < 3; i++ { }  // dạng đếm
for x < 10 { }              // đây chính là "while"
for { break }               // vòng lặp vô hạn
for i, v := range xs { }    // duyệt slice/map
```

`if` không cần ngoặc tròn: `if x > 0 {`. Ngoặc nhọn **bắt buộc**, kể cả thân 1 dòng.

## 3. Multiple return values

Hàm trả nhiều giá trị là chuyện bình thường trong Go:

```go
func split(s string) (string, string) {
	return s[:1], s[1:]
}
a, b := split("go")
```

**Quy ước cứng của toàn hệ sinh thái Go:** khi hàm có thể lỗi, `error` là giá trị **cuối cùng** —
`(T, error)`, không bao giờ `(error, T)`.

## 4. Error as value — điểm cốt lõi của tuần

Trong Python, lỗi là **luồng điều khiển ẩn**. Nhìn `x = f()` không biết `f` có ném exception hay
không — phải đọc docs hoặc đọc ruột hàm. Trong Go, lỗi là **giá trị trả về bình thường**,
nên chữ ký hàm nói thẳng điều đó ra.

Góc nhìn QA: mọi điểm có thể hỏng đều hiện rõ ngay tại chỗ gọi, không nấp trong stack.
Đánh đổi: code dài dòng hơn hẳn. Đó là chủ ý thiết kế, không phải thiếu sót.

Tạo error:

```go
import "errors"

e1 := errors.New("connection refused")     // message cố định
e2 := fmt.Errorf("port %d busy", 8080)     // cần chèn biến
```

Convention: message **chữ thường, không dấu chấm cuối**. Lý do: error hay bị bọc thành chuỗi
(`"dial tcp: connection refused"`) — viết hoa giữa chuỗi trông sai. Wrapping học ở T3.

Idiom chuẩn, sẽ lặp lại hàng nghìn lần:

```go
v, err := doSomething()
if err != nil {
	return err
}
// tới đây mới được dùng v
```

**Hợp đồng quan trọng nhất:** khi `err != nil` thì giá trị còn lại coi như **rác** — đừng đọc nó.
Ngược lại khi `err == nil` thì giá trị chắc chắn dùng được.

## 5. `defer`

```go
func f() {
	defer fmt.Println("2")
	fmt.Println("1")
}   // in 1 rồi 2
```

Khác `finally` ở hai điểm:

- `defer` gắn với **function**, không phải block `{}`. Đặt `defer` trong `for` thì nó tích lại
  tới khi cả hàm kết thúc mới chạy — nguồn leak kinh điển.
- Nhiều `defer` chạy ngược thứ tự khai báo (LIFO).

Bẫy: **tham số được đánh giá ngay lúc khai báo `defer`**, không phải lúc chạy:

```go
i := 0
defer fmt.Println(i)  // in 0, KHÔNG phải 1
i++
```

Tuần này chỉ cần *hiểu* `defer`. **Không dùng nó trong bài tập** — bài chia số không có resource
nào cần đóng. `defer` vô cớ là dấu hiệu chép mẫu.

---

# Bài tập tự đứng

Đề gốc — `ONBOARDING.md`:

> Viết hàm chia 2 số, trả về `(kết quả, error)` thay vì throw exception.
> Gọi hàm với input gây lỗi (chia 0), xử lý đúng bằng `if err != nil`.

## Khởi tạo workspace (người học tự chạy)

```powershell
mkdir D:\calibrate\learn\go\w1-divide
cd D:\calibrate\learn\go
go mod init calibrate/learn
```

Rồi tự tạo `D:\calibrate\learn\go\w1-divide\main.go` từ **file trắng**.

## Tiêu chí đạt — tự tick trước khi nhờ review

- [ ] Hàm trả 2 giá trị, giá trị thứ hai kiểu `error`, đứng cuối
- [ ] Chia 0 → error **không nil**; không `panic`, không `fmt.Println` lỗi bên trong hàm
- [ ] Trường hợp hợp lệ → error là `nil`
- [ ] `main` gọi **cả hai** trường hợp và xử lý bằng `if err != nil`
- [ ] Khi có lỗi thì **không dùng** giá trị kết quả
- [ ] Không dùng `defer`

## Xác minh

```powershell
cd D:\calibrate\learn\go
go vet ./w1-divide
go run ./w1-divide
```

Đạt khi: `go vet` sạch, và output thể hiện rõ **cả** nhánh thành công lẫn nhánh lỗi.

---

## Checklist review (trợ lý dùng khi chấm)

Lỗi kinh điển của người mới, xếp theo mức nghiêm trọng:

1. Dùng `panic` thay vì trả error — hiểu nhầm nền tảng nhất, chặn ngay.
2. Nhận `err` nhưng vẫn dùng giá trị kết quả khi `err != nil` — phá hợp đồng ở mục 4.
3. Nuốt lỗi bằng `_` (`result, _ := Divide(...)`).
4. Trả `string` làm thông báo lỗi thay vì kiểu `error`.
5. `fmt.Println` lỗi **bên trong** hàm thay vì trả lên cho caller quyết định.
6. Thứ tự trả về ngược `(error, T)`.
7. Nhầm `:=` với `=`, hoặc shadowing `err` trong scope con.
8. `defer` vô cớ — dấu hiệu chép mẫu.
9. Message lỗi viết hoa hoặc có dấu chấm cuối.

---

**Qua tuần khi:** người học xác nhận bài chạy đúng + review sạch + hiểu rõ *vì sao* Go chọn
error-as-value. → Mở T2: struct, receiver, interface (`Shape`).
