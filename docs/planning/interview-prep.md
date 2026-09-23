# Interview Prep Track — OOP + DSA (song song Calibrate)

Chốt 2026-09-23. Track này **không gate** Calibrate và Calibrate cũng không gate nó.
Nguồn sự thật cho lộ trình dự án vẫn là [`../../ONBOARDING.md`](../../ONBOARDING.md).

## Ràng buộc đã chốt

| | |
|---|---|
| Mục tiêu DSA | Luyện phỏng vấn SDET |
| Stack đích | Python + Go (**không** thêm lớp Java) |
| Mốc phỏng vấn | ~3-6 tháng kể từ 2026-09-23 → lấy mốc tính toán là **5 tháng** |
| Quỹ thời gian | **~1h/ngày, tổng cho tất cả** |
| Cách chạy | Song song, không chặn Calibrate |
| Ngôn ngữ DSA | Giải Python → viết lại Go (có điều chỉnh, xem mục "Hai ngôn ngữ") |

## Ngân sách — phần khó chịu, đọc kỹ

1h/ngày × 6 ngày/tuần × 22 tuần ≈ **132 giờ khả dụng**.

| Hạng mục | Chi phí thật |
|---|---|
| Calibrate Phase 0 + 1 (onboarding) | ~65h |
| Calibrate Phase 2 (dự án thật) | ~65h |
| DSA tới mức tự tin đi phỏng vấn (~80 bài) | ~55h |
| OOP | ~12h |
| **Tổng** | **~197h** |

Thiếu ~65h. Không phải sai số ước lượng — thiếu một phần ba.
Ba thứ này không cùng vào được 1h/ngày trong 5 tháng. Phải cắt, và cắt có chủ đích.

## Đã cắt gì

**1. Calibrate Phase 2 ra ngoài cửa sổ phỏng vấn.** (cắt lớn nhất, ~65h)

Lý do: vòng phỏng vấn hỏi *về* dự án chứ hiếm khi đọc code dự án. Thứ được đem ra bàn là
`README.md` (vấn đề, kiến trúc, ADR polyglot split, failure mode) và Checkpoint E
(5 kịch bản mutation) — những thứ **xong ngay cuối Phase 1**. Một dự án dừng đúng chỗ,
có ADR rõ và biết nói vì sao dừng, mạnh hơn một dự án làm dở mà không giải thích được gì.

Phase 2 vẫn làm — sau mốc phỏng vấn, hoặc bằng thời gian dôi ra.

**2. DSA từ ~80 bài xuống ~60 bài, bỏ hẳn Dynamic Programming.** (~15h)

Đây không phải cắt liều. Bank câu hỏi SDET thực tế chia đúng 4 nhóm:
arrays/strings, linked lists, trees/graphs, sorting/searching. **DP không xuất hiện.**
DP là đặc sản vòng phỏng vấn SWE product company, không phải SDET. Chi tiết:
[`dsa-curriculum.md`](./dsa-curriculum.md).

**3. "Viết lại Go" chỉ áp dụng cho ~15 bài chọn lọc, không phải mọi bài.** (~12h)

Làm mọi bài hai lần là nhân đôi chi phí để nhận về lợi ích giảm dần rất nhanh —
sau khoảng bài thứ mười thì việc dịch Python sang Go không còn dạy thêm gì.
Giữ lại đúng những bài mà Go **bắt nhìn thấy thứ Python giấu**: con trỏ, cấp phát,
chia sẻ mảng nền. Danh sách cụ thể trong `dsa-curriculum.md`.

## Còn lại: ~122h, vừa đủ 132h với biên an toàn nhỏ

| Hạng mục | Giờ |
|---|---|
| Calibrate Phase 0 + 1 | ~65h |
| DSA (~60 bài, 15 bài viết lại Go) | ~45h |
| OOP | ~12h |

## Phân bổ theo tháng

Cột "mỗi ngày" là cách chia 1h đó ra.

| Tháng | Calibrate | DSA | OOP |
|---|---|---|---|
| **1** (T9-T10/2026) | 40' — Phase 0A T1→T3 | 20' — mảng, two pointer (**Go**) | — |
| **2** (T10-T11) | 40' — Phase 0B concurrency | 20' — hash map, string (**Go**) | — |
| **3** (T11-T12) | 30' — Phase 0C (Python) | 30' — chuyển sang **Python**; linked list viết **cả hai** | xen kẽ, 2h tổng |
| **4** (T12-T1/2027) | 30' — Phase 1 A→C | 30' — stack/queue, binary search, cây | 5h — bài thiết kế `runner` |
| **5** (T1-T2) | 20' — Phase 1 D, E | 30' — cây/đồ thị, sort/heap | 10' — drill bộ câu hỏi |
| **6** (nếu có) | Phase 2 khởi động | ôn + mock | mock vòng OOP |

Vì sao DSA bắt đầu bằng **Go** chứ không phải Python như bạn chọn: tháng 1-2 bạn chưa học
Python (Phase 0C ở tháng 3). Không thể "giải Python trước" khi chưa biết Python.
Đổi lại, mảng/hash map trong Go củng cố thẳng slice/map/pointer của Phase 0A T3 — hai track
gia cố lẫn nhau thay vì tranh chỗ. Từ tháng 3 đảo về đúng thứ tự bạn chọn: Python trước, Go sau.

## Hai ngôn ngữ — quy tắc cuối

- **Tháng 1-2:** chỉ Go (chưa có Python).
- **Từ tháng 3:** giải Python trước (nhanh, tập trung thuật toán), rồi viết lại Go
  **chỉ với bài nằm trong danh sách 15 bài** ở `dsa-curriculum.md`.
- Bài tháng 1-2 đã giải bằng Go: khi vào tháng 3 thì giải lại bằng Python như bài khởi động
  — vừa ôn thuật toán vừa làm quen cú pháp Python. Không tính vào 60 bài.

## Rủi ro đã biết

- **Chậm tiến độ tháng 1-2 là bình thường**, không phải tín hiệu bỏ cuộc. Bài DSA đầu tiên
  có thể mất 90 phút. Tốc độ chỉ lên sau khoảng 20 bài.
- **Cám dỗ bỏ Calibrate để cày DSA** khi gần phỏng vấn. Đừng. Calibrate là thứ duy nhất
  phân biệt bạn với ứng viên khác cũng giải được hai con trỏ.
- **1h/ngày là giả định, không phải sự thật.** Nếu sau 3 tuần thấy thực tế là 30 phút,
  báo ngay để cắt lại lần hai — cắt sớm còn giữ được cả ba, cắt muộn thì mất một.

## Liên quan

- [`oop-curriculum.md`](./oop-curriculum.md) — vòng OOP hỏi gì, dự án thật dùng gì, học ở đâu
- [`dsa-curriculum.md`](./dsa-curriculum.md) — thứ tự chủ đề, số bài, quy tắc hai ngôn ngữ
- [`progress.md`](./progress.md) — tiến độ, gồm cả hai track này
- [`working-agreement.md`](./working-agreement.md) — luật review, có điều khoản riêng cho DSA
