# DSA Curriculum — luyện phỏng vấn SDET

Chốt 2026-09-23. Ngân sách: **~45h / ~60 bài / 22 tuần**. Xem [`interview-prep.md`](./interview-prep.md).

Workspace: [`learn/dsa/`](../../learn/dsa/README.md). Nhật ký bài: [`learn/dsa/log.md`](../../learn/dsa/log.md).

---

## Phạm vi — và thứ đã cố tình bỏ

Bank câu hỏi SDET thực tế (tra 2026-09-23) chia đúng **4 nhóm**:
arrays & strings · linked lists · trees & graphs · sorting & searching.

**Dynamic Programming không xuất hiện.** DP là đặc sản vòng phỏng vấn SWE product company,
không phải SDET. Bỏ DP tiết kiệm ~15h — đây là khoản cắt lớn thứ hai của cả track,
và là khoản cắt có căn cứ nhất.

Cũng bỏ: backtracking nâng cao, đồ thị có trọng số (Dijkstra/MST), trie, segment tree,
bit manipulation nâng cao. Gặp trong phỏng vấn SDET là ngoại lệ, không phải thông lệ.

**Nếu JD cụ thể sau này đòi DP** thì báo — thêm 3 tuần vào cuối, không phá phần còn lại.

---

## Thứ tự chủ đề

Thứ tự này bám theo (a) tần suất hỏi và (b) **khái niệm Phase 0 đang học tuần đó**,
để hai track gia cố lẫn nhau thay vì tranh chỗ.

| Tuần | Chủ đề | Bài | Ngôn ngữ | Gia cố cho |
|---|---|---|---|---|
| 1-3 | Mảng, two pointer, sliding window | 9 | **Go** | 0A T3 — slice, mảng nền chung |
| 4-5 | Hash map, set, đếm tần suất | 6 | **Go** | 0A T3 — map, `v, ok :=` |
| 6-7 | Chuỗi | 6 | Go → Python | 0C T1 — immutable `str` |
| 8-9 | Linked list | 7 | **cả hai** | 0A T3 — con trỏ |
| 10-11 | Stack, queue | 5 | Python | — |
| 12-13 | Binary search | 5 | Python (2 bài Go) | biên, off-by-one |
| 14-17 | Cây nhị phân, BST, BFS/DFS | 12 | Python (3 bài Go) | đệ quy |
| 18-19 | Sắp xếp, heap, top-K | 6 | Python | so sánh, ổn định |
| 20-21 | Đồ thị cơ bản (BFS/DFS, thành phần liên thông) | 4 | Python | — |
| 22 | Ôn + mock | — | — | — |
| | **Tổng** | **~60** | | |

Số bài là **mục tiêu, không phải hạn ngạch**. Làm 4 bài mà hiểu hơn 9 bài mà chép.

---

## Quy tắc hai ngôn ngữ

Bạn chọn "giải Python → viết lại Go". Giữ nguyên tinh thần, điều chỉnh hai chỗ vì thực tế:

**1. Tuần 1-5 chỉ có Go** — Phase 0C (Python) mãi tháng 3 mới bắt đầu, chưa biết Python thì
không "giải Python trước" được. Đổi lại, mảng/map trong Go củng cố thẳng 0A T3.

**2. Từ tuần 6: Python trước, Go chỉ với ~15 bài chọn lọc.**

Làm mọi bài hai lần là nhân đôi chi phí để nhận lợi ích giảm dần rất nhanh — qua khoảng
bài thứ mười, dịch Python sang Go không dạy thêm gì. Giữ lại đúng những bài mà **Go bắt nhìn
thấy thứ Python giấu**:

| Nhóm | Vì sao phải viết lại bằng Go |
|---|---|
| Toàn bộ 7 bài **linked list** | Python không có con trỏ tường minh. Go bắt bạn tự quản `*Node`, `nil`, và tự thấy vì sao đảo danh sách cần đúng 3 biến |
| 2 bài **binary search** | `mid := (lo+hi)/2` tràn số trong ngôn ngữ có kiểu cố định — Python `int` vô hạn nên giấu mất bug này |
| 3 bài **cây** (duyệt, độ sâu, cân bằng) | Đệ quy có con trỏ + `nil` receiver, khác hẳn cảm giác đệ quy trong Python |
| 3 bài **mảng** đã làm ở tuần 1-3 | Giải lại bằng Python như bài khởi động tháng 3. **Không** tính vào 60 bài |

Ngoài danh sách này thì **một ngôn ngữ là đủ**. Muốn viết lại thêm thì viết, nhưng đừng lấy
nó làm lý do để chậm chủ đề mới.

---

## Luật làm bài — quan trọng hơn danh sách bài

1. **Tự làm 25 phút trước khi xin gợi ý.** Kẹt thật thì xin, nhưng phải nói được kẹt ở đâu.
2. **Viết độ phức tạp thời gian + bộ nhớ trước khi chạy code.** Đoán rồi kiểm, không phải
   chạy xong mới tra. Vòng phỏng vấn hỏi cái này ngay sau khi bạn code xong.
3. **Ghi vào [`log.md`](../../learn/dsa/log.md) ngay sau mỗi bài**: ngày, tên bài, thời gian,
   độ phức tạp, và **sai ở đâu**. Cột "sai ở đâu" là cột có giá trị nhất — cuối tháng đọc lại
   sẽ thấy bạn sai lặp đúng 3-4 kiểu, và đó mới là thứ cần luyện.
4. **Bài sai thì hẹn làm lại sau 7 ngày**, đánh dấu trong log. Làm lại bài đã sai đáng giá
   hơn làm bài mới.
5. **Không đọc lời giải trước khi tự viết xong một bản chạy được**, kể cả bản chậm.
   Bản brute force chạy được rồi tối ưu là quy trình thật của phỏng vấn.
6. **Nói ra trong lúc làm** ít nhất 1 bài/tuần — phỏng vấn là bài thi nói, không phải bài thi viết.

## Cách tôi review bài DSA

Khác review Phase 0 một chút (xem [`working-agreement.md`](./working-agreement.md)):

- Bạn chưa tự làm đủ 25 phút → tôi **không** gợi ý, hỏi lại bạn đang nghĩ hướng nào.
- Kẹt thật → gợi ý **theo tầng**: (1) nhắc cấu trúc dữ liệu phù hợp → (2) nhắc kỹ thuật
  (two pointer / hash / sort trước) → (3) phác hướng bằng lời → (4) mới tới pseudocode.
  Không nhảy thẳng tầng 4.
- Bạn đã có bản chạy được → tôi chỉ ra bản tối ưu hơn và **vì sao**, nhưng độ phức tạp
  vẫn **bạn tự phân tích**. Tôi chỉ xác nhận đúng/sai.
- Đây là điểm khác Phase 0: sau khi bạn thật sự tự làm xong, bàn lời giải tối ưu là **được phép**
  và cần thiết — phỏng vấn chấm cả chất lượng lời giải, không chỉ việc giải được.

---

## Mốc kiểm tra

| Sau tuần | Phải đạt |
|---|---|
| 5 | 15 bài, tự phân tích được O(n) / O(n²) không cần nhắc |
| 9 | 28 bài, linked list làm được bằng cả Go lẫn Python không nhìn lại bài cũ |
| 13 | 39 bài, binary search viết đúng biên ngay lần đầu |
| 17 | 51 bài, duyệt cây (cả 3 kiểu) viết được từ trí nhớ |
| 21 | ~60 bài, log có ít nhất 10 bài đã làm lại lần hai |
| 22 | Mock: 2 bài lạ trong 45 phút, nói ra trong lúc làm |

---

## Nguồn (tra 2026-09-23)

- [Top 45 SDET Interview Questions (testmuai)](https://www.testmuai.com/learning-hub/sdet-interview-questions/) — nhóm DSA gồm đúng: arrays & strings, linked lists, trees & graphs, sorting & searching; không có DP
