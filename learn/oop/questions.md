# Bộ câu hỏi OOP — tự trả lời, không có đáp án sẵn

Giáo án: [`../../docs/planning/oop-curriculum.md`](../../docs/planning/oop-curriculum.md)

**Làm sau khi xong [`runner/`](./runner/README.md).** Làm trước thì chỉ chép được định nghĩa.

## Cách dùng

Viết câu trả lời **của bạn** vào dưới mỗi câu. Viết xong một tầng thì nhờ review —
tôi chỉ ra thiếu gì, sai gì, chỗ nào nghe như học thuộc.

Không có đáp án trong file này, và sẽ không có. Đọc đáp án hay thì tưởng mình biết;
viết ra rồi bị chỉ chỗ hổng mới thật sự biết.

## Tiêu chí chấm mọi câu tầng 1

Mỗi câu phải có **hai vế**:

> Vế 1 — định nghĩa, ≤ 2 câu, đủ chính xác để không bị bắt bẻ.
> Vế 2 — **"nhưng trong code thật…"**: dẫn chứng từ Go/Python hoặc từ chính bài `runner`.

Thiếu vế 2 = **chưa đạt**, dù vế 1 đúng hoàn toàn.
Trả lời dài quá 45 giây nói ra miệng = cũng chưa đạt.

Stack đích là Go + Python. Không có câu nào về Java ở đây, và đó là **cố ý**.

---

# Tầng 1 — Định nghĩa (chiếm ~90% vòng OOP)

### 1. OOP là gì? Khác procedural chỗ nào?
> Gợi ý chấm: Go có phải ngôn ngữ OOP không? Câu trả lời của bạn có xử lý được ca đó không?

**Trả lời:**

### 2. Encapsulation là gì?
> Gợi ý chấm: Python không có `private` thật, Go dùng chữ hoa/thường. Hai cơ chế khác nhau —
> vậy encapsulation là **cơ chế ngôn ngữ** hay **ý đồ thiết kế**?

**Trả lời:**

### 3. Abstraction là gì? Kể vài kỹ thuật abstraction.
> Gợi ý chấm: phân biệt được abstraction với encapsulation. Rất nhiều người gộp làm một.

**Trả lời:**

### 4. Inheritance là gì? Có những loại nào?
> Gợi ý chấm: có nêu được **nhược điểm** không, hay chỉ kể lợi ích?

**Trả lời:**

### 5. Polymorphism là gì? Overriding và overloading khác nhau thế nào?
> Gợi ý chấm: Go **không có** overloading. Vì sao lại thiết kế như vậy, và mất gì được gì?

**Trả lời:**

### 6. Class và object khác nhau thế nào?
**Trả lời:**

### 7. Access specifier hoạt động ra sao?
> Gợi ý chấm: so sánh thẳng `_name` của Python với chữ hoa/thường của Go.
> Cái nào được ngôn ngữ ép, cái nào chỉ là thoả thuận?

**Trả lời:**

### 8. Constructor là gì? Destructor/finalizer là gì?
> Gợi ý chấm: Python có `__init__` và `__del__`; Go **không có** cả hai. Go làm sao sống được?
> (gợi ý: zero value, hàm `NewX`, `defer`)

**Trả lời:**

### 9. Cohesion và coupling là gì?
> Gợi ý chấm: lấy ví dụ thẳng từ `runner` — `Runner` và reporter coupling ở mức nào, nhờ đâu?

**Trả lời:**

### 10. Abstract class và interface khác nhau thế nào? Khi nào dùng cái nào?
> Gợi ý chấm: **câu hay bị hỏi nhất tầng 1.** Go chỉ có interface, không có abstract class.
> Python có ABC lẫn Protocol — khác nhau chỗ nào, bạn chọn gì trong `runner` và vì sao?

**Trả lời:**

### 11. Inheritance, mixin và composition khác nhau thế nào?
> Gợi ý chấm: **câu quan trọng nhất cả bộ.** Bài `runner` được thiết kế riêng cho câu này —
> dùng thẳng câu 1, 2, 3 trong `runner/NOTES.md` làm dẫn chứng.

**Trả lời:**

---

# Tầng 2 — Riêng SDET, chỗ ăn điểm

### 12. Design a test automation framework from scratch.
> Câu mở, không có đáp án đúng. Trả lời trong [`framework-design.md`](./framework-design.md),
> ở đây chỉ cần **bản rút gọn 90 giây**.

**Trả lời (90 giây):**

### 13. Page Object Model là gì? Giải quyết vấn đề gì?
> Gợi ý chấm: nói được **vấn đề trước**, pattern sau. Nói pattern trước là dấu hiệu học thuộc.

**Trả lời:**

### 14. Factory pattern dùng để làm gì trong test automation?
> Gợi ý chấm: ví dụ kinh điển là sinh ChromeDriver/FirefoxDriver/EdgeDriver sau một interface.
> Bạn có gặp dạng bài tương tự trong `runner` không?

**Trả lời:**

### 15. Singleton là gì? Dùng ở đâu trong test automation?
> Gợi ý chấm: **bẫy.** Nêu được công dụng (WebDriver, config, logger dùng chung) là đủ điểm sàn.
> Nêu thêm được **nhược điểm** — trạng thái toàn cục trá hình, khó test, khó chạy song song —
> mới là điểm cộng. Liên hệ: bài race ở Phase 0B nói gì về trạng thái dùng chung?

**Trả lời:**

### 16. Test pyramid là gì? Áp dụng thế nào khi là SDET?
**Trả lời:**

### 17. Xử lý flaky test thế nào?
> Gợi ý chấm: retry là **giảm triệu chứng**, không phải chữa. Bài `runner` yêu cầu 5 làm retry —
> nói được vì sao retry vẫn cần dù nó không chữa gốc không?

**Trả lời:**

---

# Tầng 3 — SOLID

### 18. SOLID gồm 5 chữ gì?
> Gợi ý chấm: kể đủ là điểm sàn. Nói được **chữ nào bạn thật sự dùng hàng ngày, chữ nào hiếm**
> mới là điểm cộng.

**Trả lời:**

### 19. Liskov Substitution Principle là gì?
> Gợi ý chấm: **chữ hay bị hỏi nhất**, vì nó trả lời thẳng "khi nào kế thừa mới an toàn".
> Cho được một ví dụ **vi phạm** LSP thì mới chắc là hiểu.

**Trả lời:**

### 20. Dependency Inversion đã xuất hiện ở đâu trong bài `runner`?
> Gợi ý chấm: câu này bạn phải trả lời được bằng code của chính mình, không bằng định nghĩa.

**Trả lời:**

---

## Tự kiểm trước khi nhờ review

- [ ] Mọi câu tầng 1 đều có vế **"nhưng trong code thật…"**
- [ ] Không câu nào nói ra miệng quá 45 giây
- [ ] Ít nhất 5 câu dẫn chứng bằng chính `runner/` hoặc bài Phase 0 của bạn
- [ ] Câu 4, 10, 11, 15, 19 — nói trôi **không nhìn giấy**
