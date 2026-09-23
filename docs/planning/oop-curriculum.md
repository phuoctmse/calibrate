# OOP Curriculum — vòng phỏng vấn ↔ dự án thật

Chốt 2026-09-23. Ngân sách: **~12h**, chia làm 3 lớp. Xem [`interview-prep.md`](./interview-prep.md).

Stack đích là **Python + Go**, không có Java. Khái niệm OOP chuyển được; trivia Java
(JDK/JRE/JVM, Comparator vs Comparable, HashMap internals, Executor) **bỏ hẳn**.

---

## Phần 1 — Vòng OOP thực tế hỏi gì

Kết quả tra 2026-09-23 (nguồn ở cuối file). Ba tầng, gần như cố định giữa mọi bank câu hỏi.

### Tầng 1 — Định nghĩa (~90% vòng OOP)

Hỏi gần như nguyên văn, trả lời được trong 30 giây là đạt:

1. OOP là gì; khác procedural chỗ nào
2. **Encapsulation** là gì
3. **Abstraction** là gì; các kỹ thuật abstraction
4. **Inheritance** là gì; có những loại nào
5. **Polymorphism** là gì; **overriding vs overloading**
6. Class vs object
7. Access specifier hoạt động ra sao
8. Constructor; destructor/finalizer
9. **Cohesion vs coupling**
10. **Abstract class vs interface** — khác nhau thế nào, khi nào dùng cái nào
11. **Inheritance vs mixin vs composition**

### Tầng 2 — Riêng SDET, đây là chỗ ăn điểm

12. **"Design a test automation framework from scratch"** — câu mở, hỏi rất thường xuyên
13. **Page Object Model** — pattern phổ biến nhất trong test automation
14. **Factory** — sinh ChromeDriver / FirefoxDriver / EdgeDriver sau một interface chung
15. **Singleton** — quản lý WebDriver, config, logger dùng chung
16. Test pyramid; xử lý flaky test

### Tầng 3 — SOLID

17. 5 chữ là gì
18. **LSP** — chữ hay bị hỏi nhất, vì nó trả lời thẳng "khi nào kế thừa mới an toàn"

---

## Phần 2 — Dự án thật dùng gì

Đây là phần làm câu trả lời của bạn khác câu trả lời học thuộc.

| Vòng phỏng vấn nói | Code thật làm |
|---|---|
| Kế thừa là trụ cột của OOP | Cây kế thừa sâu bị coi là **nợ kỹ thuật**. Thực tế giữ 1-2 tầng, hiếm hơn |
| Kế thừa để tái sử dụng code | **Fragile base class**: subclass dính vào chi tiết nội bộ của base, sửa base là vỡ ngầm. Kế thừa **phá vỡ encapsulation** |
| Interface là "bản hợp đồng" | Đúng — và công dụng lớn nhất là **tiêm phụ thuộc + mock**, thứ vòng phỏng vấn hỏi ít nhất |
| Học đủ GoF pattern | Framework tốt **dùng 2-3 pattern cho chuẩn**, không rải hết catalog. Ưu tiên test dễ đọc hơn kiến trúc đẹp |
| Singleton là pattern cơ bản | Singleton là **trạng thái toàn cục trá hình** — khó test, khó chạy song song. Biết nói ra nhược điểm này là điểm cộng |
| SOLID là 5 nguyên tắc cần thuộc | Dùng hàng ngày chỉ có **S** và **D**. **L** chỉ cần khi thật sự kế thừa. **I** thì Go ép sẵn bằng văn hoá interface nhỏ |

**Bằng chứng mạnh nhất cho composition:** Go không có class, không có kế thừa, vẫn build được
hệ thống lớn. Bạn đang học đúng ngôn ngữ đó — dùng nó làm luận cứ.

**Câu trả lời mẫu về cấu trúc** (không phải chữ, là *hình dạng*):
> "Định nghĩa là X. Nhưng trong code thật tôi dùng Y, vì Z."

Vế sau là thứ phân biệt người có nghề với người thuộc bài. Mọi câu tầng 1 đều nên có vế sau.

---

## Phần 3 — Học ở đâu

Không thêm phase mới. OOP bám vào cái đã có, cộng đúng 2 thứ mới.

| Nội dung | Học ở | Trạng thái |
|---|---|---|
| Encapsulation, class/object, constructor, access specifier | 0C T2 (Python), 0A T2 (Go) | đã có trong lộ trình |
| Abstraction, interface, polymorphism | 0A T2 — implicit interface của Go | đã có |
| Composition vs inheritance | 0A T2 (embedding) + 0C T2 (Python) | đã có |
| Cohesion/coupling, SOLID (S và D), DI, mock | **`learn/oop/runner/`** | **mới** |
| Tầng 2 SDET: framework design, POM, Factory, Singleton | **`learn/oop/framework-design.md`** | **mới** |
| Toàn bộ tầng 1 + 3, dạng trả lời miệng | **`learn/oop/questions.md`** | **mới** |

### Bài `runner` — trọng tâm thực hành (~5h)

Mô hình hoá **cùng một bài toán hai lần**: Python (có class, có kế thừa) và Go
(không có gì cả). Bài toán cố tình chọn loại dụ người ta kế thừa, mà composition mới đúng.

Xong bài này bạn trả lời "inheritance vs composition" bằng **trải nghiệm**, không bằng định nghĩa —
và đó chính là câu hay được hỏi nhất tầng 1.

Đề bài: [`learn/oop/runner/README.md`](../../learn/oop/runner/README.md).

### Bài `framework-design` — tầng 2 (~4h, viết, không code)

Trả lời "design a test automation framework from scratch" thành một tài liệu.
Đây là câu mở, không có đáp án đúng — chấm bằng việc bạn có nói được **đánh đổi** hay không.

Đề bài: [`learn/oop/framework-design.md`](../../learn/oop/framework-design.md).

### Bộ câu hỏi — drill (~3h, rải đều tháng 5)

`learn/oop/questions.md` chứa **câu hỏi + tiêu chí đạt**, không chứa đáp án.
Bạn tự viết câu trả lời vào, tôi review: thiếu gì, sai gì, chỗ nào nghe như học thuộc.

Lý do không phát đáp án sẵn: đọc đáp án hay thì tưởng mình biết; viết ra rồi bị chỉ chỗ hổng
mới thật sự biết. Nguyên tắc này giống hệt working agreement của Phase 0.

---

## Điều kiện coi như xong track OOP

- [ ] `runner` chạy được cả bản Go lẫn bản Python, giải thích được vì sao không dùng kế thừa
- [ ] `framework-design.md` viết xong, nêu được ít nhất 3 đánh đổi
- [ ] `questions.md` trả lời hết, mỗi câu tầng 1 có đủ **vế "nhưng code thật…"**
- [ ] Nói trôi 4 trụ + abstract class vs interface + LSP trong 30 giây mỗi câu, không đọc giấy

---

## Nguồn (tra 2026-09-23)

- [Devinterview-io/oop-interview-questions](https://github.com/Devinterview-io/oop-interview-questions) — bank 52 câu, tầng 1 lấy từ đây
- [Top 45 SDET Interview Questions (testmuai)](https://www.testmuai.com/learning-hub/sdet-interview-questions/) — phân nhóm vòng SDET, gồm cả nhóm DSA
- [Design Patterns in Automation Framework (BrowserStack)](https://www.browserstack.com/guide/design-patterns-in-automation-framework) — POM, Factory, Singleton trong test automation
- [Test Automation Framework Design Patterns (QASkills)](https://qaskills.sh/blog/automation-framework-design-patterns) — "dùng 2-3 pattern cho chuẩn thay vì cả catalog"
- [Composition over inheritance (Wikipedia)](https://en.wikipedia.org/wiki/Composition_over_inheritance) — fragile base class, phá encapsulation
- [Composition over Inheritance in Go (DEV)](https://dev.to/amirsefati/composition-over-inheritance-in-go-the-design-choice-that-makes-microservices-boring-in-the-best-2m0h) — Go không kế thừa vẫn đủ dùng
- [Interview Questions for a Go Developer — OOP (Medium)](https://medium.com/@alsgladkikh/interview-questions-for-a-go-developer-part-3-oop-227acaf3a770) — vòng OOP hỏi gì khi ngôn ngữ là Go
