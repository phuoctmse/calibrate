# Planning — Onboarding Calibrate

Thư mục này ghi lại **cách học và cách làm việc** trong giai đoạn onboarding của Calibrate.
Không chứa logic dự án, không chứa bài giải.

Nguồn sự thật về lộ trình vẫn là [`../../ONBOARDING.md`](../../ONBOARDING.md).
Thư mục này chỉ *thi hành* lộ trình đó, không thay thế nó.

## Điều hành

| File | Nội dung |
|---|---|
| [`working-agreement.md`](./working-agreement.md) | Quy tắc bất biến giữa người học và trợ lý AI trong Phase 0 + Phase 1 |
| [`progress.md`](./progress.md) | Trạng thái hiện tại của cả ba track: đang ở đâu, đã xong bài nào |
| [`interview-prep.md`](./interview-prep.md) | Track phỏng vấn: ngân sách giờ, đã cắt gì và vì sao, phân bổ theo tháng |

## Giáo án

Mỗi file = mục tiêu thật của mốc đó, khái niệm cốt lõi, checklist review, điều kiện qua mốc.
Đề bài + tiêu chí đạt nằm trong `README.md` của thư mục bài tập tương ứng dưới `learn/`.

| Mốc | Giáo án | Bài tập |
|---|---|---|
| 0A T1 | [`phase-0a-week-1.md`](./phase-0a-week-1.md) | [`learn/go/w1-divide/`](../../learn/go/w1-divide/README.md) |
| 0A T2 | [`phase-0a-week-2.md`](./phase-0a-week-2.md) | [`learn/go/w2-shape/`](../../learn/go/w2-shape/README.md) |
| 0A T3 | [`phase-0a-week-3.md`](./phase-0a-week-3.md) | [`learn/go/w3-config/`](../../learn/go/w3-config/README.md) |
| 0B | [`phase-0b-concurrency.md`](./phase-0b-concurrency.md) | [`learn/go/b-race/`](../../learn/go/b-race/README.md) |
| 0C T1 | [`phase-0c-week-1.md`](./phase-0c-week-1.md) | [`learn/python/c1-mutable/`](../../learn/python/c1-mutable/README.md) |
| 0C T2 | [`phase-0c-week-2.md`](./phase-0c-week-2.md) | [`learn/python/c2-movingavg/`](../../learn/python/c2-movingavg/README.md) |
| 0C T3 | [`phase-0c-week-3.md`](./phase-0c-week-3.md) | [`learn/python/c3-ctxmanager/`](../../learn/python/c3-ctxmanager/README.md) |
| Phase 1 A–E | [`phase-1-checkpoints.md`](./phase-1-checkpoints.md) | xem bảng trong file đó |

## Track phỏng vấn (thêm 2026-09-23)

Chạy song song, **không gate** Calibrate. Stack đích Python + Go, không có lớp Java.

| Track | Giáo án | Bài tập |
|---|---|---|
| OOP | [`oop-curriculum.md`](./oop-curriculum.md) | [`learn/oop/`](../../learn/oop/README.md) |
| DSA | [`dsa-curriculum.md`](./dsa-curriculum.md) | [`learn/dsa/`](../../learn/dsa/README.md) |

`oop-curriculum.md` chứa kết quả tra thực tế (2026-09-23) về **vòng OOP hỏi gì** so với
**dự án thật dùng gì**, kèm nguồn. Đọc phần 2 của file đó trước khi trả lời bất kỳ câu OOP nào.

Giáo án Tuần 1 được viết ở mức chi tiết đầy đủ (đã giảng xong). Các mốc sau có **khung,
khái niệm cốt lõi, đề bài và tiêu chí đạt** chốt sẵn; phần giảng chi tiết — ví dụ, so sánh,
hỏi đáp — diễn ra trực tiếp khi mở mốc đó, và được bổ sung ngược lại vào file.

## Vì sao cần thư mục này

`ONBOARDING.md` nói **học gì**. Nó không nói **đang học tới đâu** và **ai được viết dòng code nào**.
Thiếu hai thứ đó thì mỗi phiên làm việc mới phải dựng lại ngữ cảnh từ đầu — và đã từng dẫn tới
việc bắt đầu sai điểm trong lộ trình (nhầm bài goroutine của Phase 1 thành bài tuần 2 của Phase 0A,
do đọc phải bản `ONBOARDING.md` cũ).
