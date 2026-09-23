# Bài `runner` — mô hình hoá cùng một bài toán bằng Go và Python

Giáo án: [`../../../docs/planning/oop-curriculum.md`](../../../docs/planning/oop-curriculum.md)

**Khoá cho tới khi Phase 0C T2 xong** (cần biết class Python) — dự kiến tháng 4.

Mục tiêu: sau bài này bạn trả lời "inheritance vs composition" bằng **trải nghiệm**,
không bằng định nghĩa. Đó là câu hay được hỏi nhất vòng OOP.

## Bài toán

Một **test runner tí hon**. Nhận danh sách test case (hàm trả về pass/fail),
chạy hết, rồi báo cáo kết quả.

Cố tình chọn bài toán này vì ba lý do: nó là domain của nghề bạn đang chuyển sang,
nó **dụ người ta dùng kế thừa**, và nó chính là phiên bản thu nhỏ của mutation harness
ở Checkpoint A.

## Yêu cầu chức năng

1. **Reporter** — một hợp đồng tối đa 3 method, ví dụ `Start`, `Record(result)`, `Finish`.
2. Ít nhất **2 bản cài đặt**: in ra console, và ghi ra file JSON.
3. **Runner nhận reporter từ bên ngoài**, không tự tạo bên trong.
4. **Báo cáo ra nhiều nơi cùng lúc** (console *và* JSON) mà **không sửa `Runner`**
   và **không** viết reporter thứ ba biết về hai cái kia.
5. **Retry**: test fail thì chạy lại tối đa N lần. Thêm tính năng này mà **không sửa `Runner`**.
6. Test cho `Runner`: tiêm một reporter giả, kiểm tra nó **nhận đúng** những gì mong đợi.
   Test này **không được** in ra console hay ghi file.

Yêu cầu 4, 5, 6 mới là bài. 1-3 chỉ là dựng sân.

## Làm hai lần

- `python/` — Python có class, có kế thừa, có đa kế thừa.
- `go/` — Go không có gì trong số đó.

Làm Python trước hay Go trước đều được. Nhưng **phải làm cả hai** mới thấy được điều cần thấy.

## Tiêu chí đạt

**Chung:**
- [ ] Hợp đồng reporter tối đa 3 method
- [ ] `Runner` **không** biết reporter cụ thể nào đang chạy
- [ ] Yêu cầu 4 giải bằng cách gói nhiều reporter vào một thứ **cùng thoả hợp đồng đó**
- [ ] Yêu cầu 5 giải bằng **bọc ngoài**, không phải bằng kế thừa `Runner` hay thêm `if` vào `Runner`
- [ ] Có test tiêm reporter giả, chạy được mà không tạo file, không in ra màn hình
- [ ] `Runner` không tự `print` — mọi thứ ra ngoài đều đi qua reporter

**Riêng Python:**
- [ ] Nếu có dùng kế thừa ở đâu đó: **giải thích được vì sao chỗ đó kế thừa mới đúng**
- [ ] Nếu không dùng kế thừa chỗ nào: giải thích được vì sao không cần
- [ ] Type hints đầy đủ; hợp đồng reporter khai bằng `Protocol` hoặc ABC — **chọn một và nói lý do**

**Riêng Go:**
- [ ] Interface khai ở **phía dùng** (package `Runner`), không phải phía reporter
- [ ] `go vet` sạch
- [ ] Reporter giả trong test chỉ là một struct nhỏ trong file test, **không** cần thư viện mock

## Phần viết — bắt buộc, đây mới là thứ đem đi phỏng vấn

Tạo `NOTES.md` trong thư mục này, trả lời:

1. Yêu cầu 4 (báo cáo nhiều nơi) — nếu giải bằng kế thừa thì hỏng ở đâu?
2. Yêu cầu 5 (retry) — nếu làm `RetryRunner` kế thừa `Runner` thì hỏng ở đâu?
3. Bản Go thiếu hẳn kế thừa. Có chỗ nào **thật sự** thấy thiếu không, hay không?
4. Nhờ thiết kế nào mà test ở yêu cầu 6 chạy được không cần file, không cần console?
   Gọi tên nguyên tắc đó ra.
5. Nếu bây giờ cần thêm reporter gửi kết quả lên Slack — phải sửa bao nhiêu file có sẵn?
   Con số đó nói lên điều gì?

Năm câu này gần như là năm câu hỏi đuổi của vòng OOP. Viết được rồi thì
[`../questions.md`](../questions.md) chỉ còn là gọi tên lại.

## Bẫy hay mắc

Tạo reporter **bên trong** `Runner` (hết tiêm, hết test được) ·
thêm `if retry` vào `Runner` (mỗi tính năng mới lại sửa `Runner` một lần) ·
hợp đồng reporter phình ra 6-7 method vì cố gói mọi nhu cầu ·
bản Go khai interface ở package reporter thay vì package dùng ·
test gọi reporter thật rồi kiểm tra bằng cách đọc file vừa ghi.
