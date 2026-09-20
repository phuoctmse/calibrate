# Tiến độ onboarding

Cập nhật gần nhất: **2026-09-16**

Trạng thái: `[ ]` chưa bắt đầu · `[~]` đang làm · `[x]` đã xác nhận đạt

## Phase 0A — Go Core Fundamentals (~3 tuần)

| | Tuần | Bài tập tự đứng | Thư mục |
|---|---|---|---|
| `[~]` | T1 — type system, control flow, functions, `defer` | Hàm chia 2 số trả `(result, error)`, xử lý chia 0 | `learn/go/w1-divide/` |
| `[ ]` | T2 — struct, method/receiver, interface, composition | Interface `Shape` + `Circle` + `Rectangle` | `learn/go/w2-shape/` |
| `[ ]` | T3 — slice, map, pointer, error wrapping | `ParseConfig` wrap lỗi, in full error chain | `learn/go/w3-config/` |

## Phase 0B — Go Concurrency (~1.5 tuần)

| | Nội dung | Bài tập tự đứng | Thư mục |
|---|---|---|---|
| `[ ]` | goroutine, channel, `WaitGroup`, `Mutex`, `select`, race | Tự tạo race thật → `go run -race` thấy lỗi → fix bằng `Mutex` → `-race` sạch | `learn/go/b-race/` |

## Phase 0C — Python Core Fundamentals (~2.5 tuần)

| | Tuần | Bài tập tự đứng |
|---|---|---|
| `[ ]` | T1 — dynamic typing, mutable vs immutable | Mutate `list` trong function, so sánh với cách copy trước khi mutate |
| `[ ]` | T2 — function, class, OOP | Class `MovingAverage` với `add(score)` và `average()` |
| `[ ]` | T3 (nửa tuần) — type hints, context manager, venv | Context manager tự chế bằng `__enter__`/`__exit__` |

## Phase 1 — Applied Checkpoints (~3 tuần)

| | Checkpoint | Mô tả | Dùng nền từ |
|---|---|---|---|
| `[ ]` | A — Mutation harness khung xương | N task giả chạy song song, gom kết quả, in tổng kết, xác nhận sạch `-race` | 0B |
| `[ ]` | B — Similarity script | `sentence-transformers` encode 2 câu, cosine similarity bằng `numpy` | 0C |
| `[ ]` | C — gRPC nối 2 ngôn ngữ | `.proto` → generate Go + Python, Go client gọi Python server qua network thật | A + B |
| `[ ]` | D — Control chart thật | Mở rộng `MovingAverage` thành mean ± sigma × std, tự viết phần logic cốt lõi | 0C T2 |
| `[ ]` | E — 5 kịch bản mutation | Viết ra giấy, không code | — |

## Phase 2 — Dự án thật

`[ ]` Chỉ bắt đầu khi **toàn bộ** ô trên là `[x]`. Ráp các checkpoint vào `TODO(...)` có sẵn trong
`go-orchestrator/cmd/calibrate/main.go` và `python-scoring-service/server.py`.

---

## Nhật ký

| Ngày | Sự kiện |
|---|---|
| 2026-09-16 | Chốt working agreement, tạo `docs/planning/`. Bắt đầu 0A T1 — chưa viết dòng Go nào. |
