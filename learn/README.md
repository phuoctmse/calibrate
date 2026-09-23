# learn/ — Sân tập Phase 0 & Phase 1

Đây **không phải** code dự án Calibrate. Đây là nơi luyện ngôn ngữ.

`ONBOARDING.md` quy định bài Phase 0 là bài **tự đứng, không liên quan Calibrate**.
Tách hẳn ra đây để không phá ranh giới đó — và để xoá cả thư mục này lúc nào cũng
không ảnh hưởng dự án.

Hai track chạy song song, **không** gate lẫn nhau:

```
learn/
├── go/          ┐
├── python/      ├─ Track Calibrate — Phase 0 + 1 của ONBOARDING.md
├── checkpoints/ ┘
│
├── dsa/         ┐ Track phỏng vấn — chạy nền, xem docs/planning/interview-prep.md
└── oop/         ┘
```

## Luật của thư mục này

1. Mỗi thư mục bài tập có sẵn `README.md` = **đề bài + tiêu chí đạt + lệnh xác minh**.
2. File code (`main.go`, `*.py`) **chưa có, và sẽ không được tạo hộ**. Bạn tự tạo từ file trắng.
3. Tự tick hết checklist "Tiêu chí đạt" **trước khi** nhờ review.
4. Không copy code từ nơi khác vào. Bài này tính bằng năng lực, không tính bằng output.

Xem [`../docs/planning/working-agreement.md`](../docs/planning/working-agreement.md) để biết
trợ lý AI được và không được làm gì.

## Thứ tự làm

| # | Bài | Thư mục | Giáo án |
|---|---|---|---|
| 1 | 0A T1 — error as value | `go/w1-divide/` | [`phase-0a-week-1.md`](../docs/planning/phase-0a-week-1.md) |
| 2 | 0A T2 — struct, interface | `go/w2-shape/` | [`phase-0a-week-2.md`](../docs/planning/phase-0a-week-2.md) |
| 3 | 0A T3 — slice, map, pointer, wrap | `go/w3-config/` | [`phase-0a-week-3.md`](../docs/planning/phase-0a-week-3.md) |
| 4 | 0B — concurrency, race | `go/b-race/` | [`phase-0b-concurrency.md`](../docs/planning/phase-0b-concurrency.md) |
| 5 | 0C T1 — mutability | `python/c1-mutable/` | [`phase-0c-week-1.md`](../docs/planning/phase-0c-week-1.md) |
| 6 | 0C T2 — class, OOP | `python/c2-movingavg/` | [`phase-0c-week-2.md`](../docs/planning/phase-0c-week-2.md) |
| 7 | 0C T3 — type hints, context manager | `python/c3-ctxmanager/` | [`phase-0c-week-3.md`](../docs/planning/phase-0c-week-3.md) |
| 8 | P1 A — mutation harness | `go/p1a-harness/` | [`phase-1-checkpoints.md`](../docs/planning/phase-1-checkpoints.md) |
| 9 | P1 B — similarity | `python/p1b-similarity/` | ↑ |
| 10 | P1 C — gRPC 2 ngôn ngữ | `checkpoints/c-grpc/` | ↑ |
| 11 | P1 D — control chart | `python/p1d-controlchart/` | ↑ |
| 12 | P1 E — 5 kịch bản mutation | `checkpoints/e-scenarios/` | ↑ |

Không nhảy cóc. Mỗi ô chỉ mở khi ô trước được xác nhận đạt trong
[`../docs/planning/progress.md`](../docs/planning/progress.md).

---

## Track phỏng vấn — chạy song song, không gate gì cả

Ngân sách và phân bổ theo tháng: [`../docs/planning/interview-prep.md`](../docs/planning/interview-prep.md).
Mốc phỏng vấn ~3-6 tháng kể từ 2026-09-23, quỹ thời gian ~1h/ngày **cho tất cả**.

| Track | Thư mục | Nhịp | Giáo án |
|---|---|---|---|
| DSA (~60 bài) | [`dsa/`](./dsa/README.md) | 20-30 phút/ngày, liên tục | [`dsa-curriculum.md`](../docs/planning/dsa-curriculum.md) |
| OOP (3 phần) | [`oop/`](./oop/README.md) | dồn vào tháng 4-5 | [`oop-curriculum.md`](../docs/planning/oop-curriculum.md) |

Hai track này **không** phải tick xong mới được đi tiếp Calibrate, và ngược lại.
Nhưng `oop/runner/` cần Phase 0C T2 (class Python) nên mới mở được từ tháng 4.
