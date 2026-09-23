# Tiến độ onboarding

Cập nhật gần nhất: **2026-09-23**

Ba track: **Calibrate** (dưới đây) · **DSA** · **OOP** — hai track sau không gate Calibrate,
xem [`interview-prep.md`](./interview-prep.md).

Trạng thái: `[ ]` chưa bắt đầu · `[~]` đang làm · `[x]` đã xác nhận đạt

Quy tắc tick: chỉ chuyển sang `[x]` khi bài chạy đúng **và** review sạch **và**
trả lời được câu hỏi chốt trong giáo án của mốc đó.

## Phase 0A — Go Core Fundamentals (~3 tuần)

| | Tuần | Bài tập tự đứng | Thư mục | Giáo án |
|---|---|---|---|---|
| `[~]` | T1 — type system, control flow, functions, `defer` | Hàm chia 2 số trả `(result, error)`, xử lý chia 0 | [`learn/go/w1-divide/`](../../learn/go/w1-divide/README.md) | [T1](./phase-0a-week-1.md) |
| `[ ]` | T2 — struct, method/receiver, interface, composition | Interface `Shape` + `Circle` + `Rectangle` | [`learn/go/w2-shape/`](../../learn/go/w2-shape/README.md) | [T2](./phase-0a-week-2.md) |
| `[ ]` | T3 — slice, map, pointer, error wrapping | `ParseConfig` wrap lỗi, in full error chain | [`learn/go/w3-config/`](../../learn/go/w3-config/README.md) | [T3](./phase-0a-week-3.md) |

## Phase 0B — Go Concurrency (~1.5 tuần)

| | Nội dung | Bài tập tự đứng | Thư mục | Giáo án |
|---|---|---|---|---|
| `[ ]` | goroutine, channel, `WaitGroup`, `Mutex`, `select`, race | Tự tạo race thật → `-race` thấy lỗi → fix bằng `Mutex` → `-race` sạch | [`learn/go/b-race/`](../../learn/go/b-race/README.md) | [0B](./phase-0b-concurrency.md) |

Điều kiện tiên quyết: `-race` chạy được trên máy (cần gcc/mingw). Kiểm tra trước khi mở phase.

## Phase 0C — Python Core Fundamentals (~2.5 tuần)

| | Tuần | Bài tập tự đứng | Thư mục | Giáo án |
|---|---|---|---|---|
| `[ ]` | T1 — dynamic typing, mutable vs immutable | Mutate `list` trong function, so sánh với cách copy trước khi mutate | [`learn/python/c1-mutable/`](../../learn/python/c1-mutable/README.md) | [T1](./phase-0c-week-1.md) |
| `[ ]` | T2 — function, class, OOP | Class `MovingAverage` với `add(score)` và `average()` | [`learn/python/c2-movingavg/`](../../learn/python/c2-movingavg/README.md) | [T2](./phase-0c-week-2.md) |
| `[ ]` | T3 (nửa tuần) — type hints, context manager, venv | Context manager tự chế bằng `__enter__`/`__exit__` | [`learn/python/c3-ctxmanager/`](../../learn/python/c3-ctxmanager/README.md) | [T3](./phase-0c-week-3.md) |

## Phase 1 — Applied Checkpoints (~3 tuần)

Giáo án chung: [`phase-1-checkpoints.md`](./phase-1-checkpoints.md)

| | Checkpoint | Mô tả | Thư mục | Nền từ |
|---|---|---|---|---|
| `[ ]` | A — Mutation harness khung xương | N task giả chạy song song, gom kết quả, in tổng kết, `-race` sạch | [`learn/go/p1a-harness/`](../../learn/go/p1a-harness/README.md) | 0B |
| `[ ]` | B — Similarity script | `sentence-transformers` encode 2 câu, cosine similarity tự viết bằng `numpy` | [`learn/python/p1b-similarity/`](../../learn/python/p1b-similarity/README.md) | 0C |
| `[ ]` | C — gRPC nối 2 ngôn ngữ | `.proto` rút gọn tự viết → generate Go + Python, Go client gọi Python server qua network thật | [`learn/checkpoints/c-grpc/`](../../learn/checkpoints/c-grpc/README.md) | A + B |
| `[ ]` | D — Control chart thật | Mở rộng `MovingAverage` thành mean ± sigma × std, tự viết phần logic cốt lõi | [`learn/python/p1d-controlchart/`](../../learn/python/p1d-controlchart/README.md) | 0C T2 + C |
| `[ ]` | E — 5 kịch bản mutation | Điền `scenarios.md`, không code | [`learn/checkpoints/e-scenarios/`](../../learn/checkpoints/e-scenarios/README.md) | A–D |

## Phase 2 — Dự án thật

`[ ]` Chỉ bắt đầu khi **toàn bộ** ô trên là `[x]`. Khi đó
[`working-agreement.md`](./working-agreement.md) hết hiệu lực, mở khoá
`go-orchestrator/cmd/calibrate/main.go` và `python-scoring-service/server.py`,
ráp các checkpoint vào `TODO(...)` có sẵn.

---

# Track phỏng vấn — chạy song song, không gate Calibrate

Kế hoạch + ngân sách: [`interview-prep.md`](./interview-prep.md).
Mốc phỏng vấn ~3-6 tháng kể từ 2026-09-23. Quỹ thời gian ~1h/ngày **cho tất cả các track**.

## DSA (~60 bài / 22 tuần)

Giáo án: [`dsa-curriculum.md`](./dsa-curriculum.md) · Nhật ký bài: [`learn/dsa/log.md`](../../learn/dsa/log.md)

| | Tuần | Chủ đề | Bài | Ngôn ngữ |
|---|---|---|---|---|
| `[ ]` | 1-3 | Mảng, two pointer, sliding window | 9 | Go |
| `[ ]` | 4-5 | Hash map, set, đếm tần suất | 6 | Go |
| `[ ]` | 6-7 | Chuỗi | 6 | Go → Python |
| `[ ]` | 8-9 | Linked list | 7 | cả hai |
| `[ ]` | 10-11 | Stack, queue | 5 | Python |
| `[ ]` | 12-13 | Binary search | 5 | Python (2 bài Go) |
| `[ ]` | 14-17 | Cây nhị phân, BST, BFS/DFS | 12 | Python (3 bài Go) |
| `[ ]` | 18-19 | Sắp xếp, heap, top-K | 6 | Python |
| `[ ]` | 20-21 | Đồ thị cơ bản | 4 | Python |
| `[ ]` | 22 | Ôn + mock (2 bài lạ / 45 phút, nói ra trong lúc làm) | — | — |

Đã bỏ có chủ đích: **Dynamic Programming**, backtracking nâng cao, đồ thị có trọng số,
trie, segment tree. Lý do trong [`dsa-curriculum.md`](./dsa-curriculum.md#phạm-vi--và-thứ-đã-cố-tình-bỏ).

Mốc kiểm: tuần 5 → 15 bài · tuần 9 → 28 · tuần 13 → 39 · tuần 17 → 51 · tuần 21 → ~60.

## OOP (~12h, dồn tháng 4-5)

Giáo án: [`oop-curriculum.md`](./oop-curriculum.md)

| | Phần | Dạng | Thư mục | Mở khi |
|---|---|---|---|---|
| `[ ]` | `runner` — cùng bài toán, Go và Python | code | [`learn/oop/runner/`](../../learn/oop/runner/README.md) | xong 0C T2 |
| `[ ]` | Framework design | viết, không code | [`learn/oop/framework-design.md`](../../learn/oop/framework-design.md) | sau `runner` |
| `[ ]` | 20 câu drill (tầng 1/2/3) | tự trả lời + review | [`learn/oop/questions.md`](../../learn/oop/questions.md) | sau `runner` |

Điều kiện xong track: nói trôi 4 trụ + abstract class vs interface + LSP trong 30 giây mỗi câu,
không nhìn giấy, và mọi câu tầng 1 đều có vế **"nhưng trong code thật…"**.

---

## Nhật ký

| Ngày | Sự kiện |
|---|---|
| 2026-09-16 | Chốt working agreement, tạo `docs/planning/`. Bắt đầu 0A T1 — chưa viết dòng Go nào. |
| 2026-09-23 | Dựng đủ workspace `learn/` (12 thư mục bài tập + `go.mod` + `.gitignore`) và 7 giáo án còn lại. Chốt ranh giới: README đề bài dựng sẵn, file code người học tự tạo. |
| 2026-09-23 | Thêm track phỏng vấn (OOP + DSA) chạy song song. Tra thực tế vòng OOP/SDET hỏi gì. Chốt ngân sách: thiếu ~65h so với nhu cầu → cắt Phase 2 ra ngoài cửa sổ phỏng vấn, bỏ DP, giới hạn "viết lại Go" còn ~15 bài. |
