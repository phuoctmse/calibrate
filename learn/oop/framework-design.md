# "Design a test automation framework from scratch"

Giáo án: [`../../docs/planning/oop-curriculum.md`](../../docs/planning/oop-curriculum.md)
**Viết, không code.** ~4h. Làm sau [`runner/`](./runner/README.md).

Đây là câu mở hay được hỏi nhất ở vòng SDET. **Không có đáp án đúng.**
Người chấm không tìm kiến trúc đẹp — họ tìm xem bạn có nói được **đánh đổi** hay không.

Viết câu trả lời của bạn vào từng mục dưới. Không có đáp án sẵn trong file này.

---

## Luật khi trả lời câu mở

1. **Hỏi lại trước khi thiết kế.** Test cái gì — web, API, mobile? Ai chạy — dev hay QA?
   Chạy ở đâu — máy cá nhân hay CI? Nhảy thẳng vào "tôi dùng Selenium + POM" là trượt.
2. **Nêu đánh đổi, đừng nêu lựa chọn.** "Tôi chọn A" yếu. "Tôi chọn A, mất B, chấp nhận được
   vì C" mạnh.
3. **Ít pattern, dùng cho chuẩn.** Framework tốt dùng 2-3 pattern, không rải hết catalog.
   Kể tên 8 pattern là dấu hiệu chưa làm thật.
4. **Test dễ đọc thắng kiến trúc đẹp.** Nói được ưu tiên này ra là điểm cộng lớn.

---

## 1. Câu hỏi làm rõ bạn sẽ hỏi trước

Viết ra **5 câu** bạn hỏi người phỏng vấn trước khi thiết kế. Với mỗi câu, ghi thêm:
*câu trả lời khác nhau thì thiết kế đổi ra sao?* Câu hỏi không làm đổi thiết kế là câu hỏi thừa.

**Trả lời:**

## 2. Các tầng của framework

Liệt kê các tầng và trách nhiệm từng tầng. Ràng buộc: mỗi tầng phải nói được
**nó KHÔNG làm gì**. Tầng không có ranh giới là tầng sẽ phình ra.

**Trả lời:**

## 3. Ba quyết định thiết kế và cái giá của chúng

Chọn đúng **3** quyết định, mỗi cái viết: chọn gì · được gì · **mất gì** · vì sao chấp nhận được.

Gợi ý vùng có đánh đổi thật: chạy song song hay tuần tự · test data cố định hay sinh ra ·
chờ tường minh hay chờ ngầm · báo cáo tự viết hay dùng sẵn · bao nhiêu phần trăm E2E so với API.

**Trả lời:**

## 4. Pattern nào bạn dùng, và pattern nào cố tình không dùng

Tối đa 3 pattern. Với mỗi cái: **vấn đề gì** khiến bạn cần nó (vấn đề trước, tên pattern sau).

Rồi kể **1 pattern bạn cố tình không dùng** và vì sao. Mục này quan trọng ngang mục trên —
nó cho thấy bạn chọn chứ không phải gom.

**Trả lời:**

## 5. Cấu hình và bí mật

Chạy trên 3 môi trường (local, staging, CI), mỗi nơi URL và credential khác nhau.
Thiết kế thế nào để không hardcode, không commit secret, và **không** biến config thành
một Singleton toàn cục mà mọi thứ đều đụng vào được?

> Liên hệ: câu 15 trong [`questions.md`](./questions.md) và bài race ở Phase 0B.

**Trả lời:**

## 6. Flaky test

Framework của bạn xử lý flaky test ra sao? Bắt buộc trả lời được: retry là **giảm triệu chứng**
hay **chữa gốc**? Nếu chỉ giảm triệu chứng thì vì sao vẫn làm? Và làm sao để retry
không **che mất** một bug thật?

**Trả lời:**

## 7. Làm sao biết framework này tốt

Nêu 3 chỉ số bạn theo dõi. Ràng buộc: ít nhất một chỉ số **không phải** tỉ lệ pass —
tỉ lệ pass cao có thể chỉ nghĩa là test quá dễ.

> Liên hệ: đây đúng là câu hỏi trung tâm của Calibrate — làm sao biết bộ chấm điểm đáng tin.
> Mutation kill rate là một câu trả lời cho dạng câu hỏi này. Dùng được trong phỏng vấn.

**Trả lời:**

## 8. Bản nói 3 phút

Rút gọn tất cả phần trên thành **3 phút nói**. Đây mới là thứ dùng thật.
Viết ra rồi bấm giờ đọc to. Quá 3 phút thì cắt tiếp.

**Trả lời:**

---

## Tiêu chí đạt

- [ ] 5 câu hỏi làm rõ, mỗi câu nói được nó làm đổi thiết kế thế nào
- [ ] Mỗi tầng có nói **nó không làm gì**
- [ ] 3 quyết định, mỗi cái có nêu **cái mất**
- [ ] Tối đa 3 pattern, **vấn đề trước tên pattern sau**, cộng 1 pattern cố tình không dùng
- [ ] Mục 6 phân biệt được giảm triệu chứng với chữa gốc
- [ ] Mục 7 có ít nhất 1 chỉ số không phải tỉ lệ pass
- [ ] Bản 3 phút bấm giờ đọc to **không quá 3 phút**
