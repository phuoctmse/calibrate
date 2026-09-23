# learn/oop — OOP cho vòng phỏng vấn SDET

Giáo án: [`../../docs/planning/oop-curriculum.md`](../../docs/planning/oop-curriculum.md)

Ngân sách ~12h, ba phần. Làm theo thứ tự — phần 1 tạo ra trải nghiệm để phần 3 có gì mà nói.

| # | Phần | Dạng | Giờ | Khi nào |
|---|---|---|---|---|
| 1 | [`runner/`](./runner/README.md) | code, cả Go lẫn Python | ~5h | tháng 4 |
| 2 | [`framework-design.md`](./framework-design.md) | viết, không code | ~4h | tháng 4 |
| 3 | [`questions.md`](./questions.md) | tự trả lời rồi được review | ~3h | rải đều tháng 5 |

## Vì sao thứ tự này

Phần lớn người luyện OOP làm ngược: học thuộc định nghĩa trước, rồi cố nhớ lúc phỏng vấn.
Kết quả là câu trả lời đúng mà nghe như đọc sách, và sập ngay ở câu hỏi đuổi
*"vậy trong dự án anh làm thì sao?"*.

Ở đây làm ngược lại: `runner` cho bạn một lần **tự tay** mô hình hoá cùng một bài toán
bằng Python (có kế thừa) và Go (không có gì cả). Sau đó `questions.md` chỉ là việc
gọi tên lại thứ bạn đã làm.

## Nguyên tắc trả lời mọi câu OOP

> "Định nghĩa là X. Nhưng trong code thật tôi dùng Y, vì Z."

Vế sau là thứ phân biệt người có nghề với người thuộc bài.
Không có vế sau thì câu trả lời chưa xong, dù vế đầu đúng hoàn toàn.

## Luật vẫn như Phase 0

`questions.md` và `framework-design.md` **không chứa đáp án** — chỉ có câu hỏi và tiêu chí đạt.
Bạn viết, tôi review: thiếu gì, sai gì, chỗ nào nghe như học thuộc.

Đọc đáp án hay thì tưởng mình biết. Viết ra rồi bị chỉ chỗ hổng mới thật sự biết.
