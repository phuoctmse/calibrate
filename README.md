# Calibrate

AI Feature Regression CI Gate với Self-Validating Test Oracle — verified bằng mutation testing trên chính hệ thống AI.

> **Trạng thái:** Onboarding phase. Xem [`ONBOARDING.md`](./ONBOARDING.md) trước khi đọc code — mọi file trong `go-orchestrator/` và `python-scoring-service/` hiện là **skeleton có chủ đích**, đánh dấu rõ `TODO(weekN-checkpoint)` map với từng mốc học trong lộ trình.

## Vì sao dự án này tồn tại

Tool có sẵn (promptfoo, DeepEval, Ragas) chấm điểm output AI 1 lần, tĩnh — không ai kiểm định lại **chính bộ chấm điểm (oracle) đó có đáng tin hay không**. Calibrate thêm 1 tầng: cố tình làm hệ thống AI suy giảm chất lượng theo cách kiểm soát được (downgrade model, cắt context, prompt injection), rồi đo xem oracle có phát hiện được không → **Mutation Kill Rate**.

Xem đầy đủ vấn đề, kiến trúc, và phân tích trade-off tại [artifact báo cáo dự án](https://claude.ai/artifact/TDzoY9H22mtkfJydmLHn5c).

## Kiến trúc

```
[Playwright E2E] → click UI thật, lấy AI output thật
        │
        ▼
┌───────────────────────┐         gRPC          ┌──────────────────────────┐
│   Go orchestrator      │ ─────────────────────▶│  Python scoring service  │
│  (control plane)        │                       │                          │
│  - mutation harness      │◀─────────────────────│  - embedding similarity  │
│  - rule-based oracle      │                      │  - SPC control chart     │
│  - CI gate / kill-rate    │                      │                          │
└───────────────────────┘                         └──────────────────────────┘
```

## ADR: Polyglot service split

```
Title: Polyglot service split for statistical scoring layer in Calibrate
Context: Core deliverable (mutation kill rate) requires statistically
         correct control-chart and embedding-similarity computation.
         Go's numerical ecosystem (gonum) is immature relative to Python's
         (scipy/statsmodels), risking silent correctness bugs in the
         project's central metric.
Decision: Split into Go control-plane (orchestration, mutation harness,
          rule-based oracle, CI gate) and Python scoring service
          (embedding similarity, SPC), connected via gRPC with a
          versioned protobuf contract (see proto/scoring.proto).
Consequences:
  + Correctness of core metric backed by mature statistical libraries
  + Independent testability/scalability of scoring layer
  + Sidecar pattern is directly transferable to platform engineering roles
  - Added operational complexity (2 deployables, network boundary)
  - Requires explicit failure-mode design (degrade-gracefully policy)
    to avoid coupling availability of whole pipeline to scoring service
```

## Cấu trúc thư mục

```
calibrate/
├── ONBOARDING.md               # Lộ trình học (nguồn sự thật) — đọc trước khi đọc code
├── docs/planning/              # Thi hành lộ trình: working agreement, tiến độ, giáo án, kế hoạch phỏng vấn
├── learn/                      # Sân tập — KHÔNG phải code dự án
│   ├── go/                     #   Calibrate Phase 0+1: 1 module calibrate/learn
│   ├── python/                 #   Calibrate Phase 0+1: 1 venv
│   ├── checkpoints/            #   Phase 1: gRPC (C) và kịch bản mutation (E)
│   ├── dsa/                    #   Track phỏng vấn: ~60 bài, mỗi bài kèm test tự viết
│   └── oop/                    #   Track phỏng vấn: bài thiết kế + bộ câu hỏi
├── proto/
│   └── scoring.proto           # Contract gRPC giữa 2 service
├── go-orchestrator/
│   ├── go.mod
│   └── cmd/calibrate/main.go   # Skeleton mutation harness — khoá tới hết Phase 1
└── python-scoring-service/
    ├── requirements.txt
    └── server.py               # Skeleton similarity + control chart — khoá tới hết Phase 1
```

Bắt đầu onboarding: [`docs/planning/progress.md`](./docs/planning/progress.md) cho biết đang ở đâu,
[`learn/README.md`](./learn/README.md) cho biết làm bài nào tiếp theo.

Song song với onboarding có một track luyện phỏng vấn SDET (OOP + DSA) —
ngân sách giờ và các khoản đã cắt nằm ở
[`docs/planning/interview-prep.md`](./docs/planning/interview-prep.md).

## Failure mode đã chốt (áp dụng khi implement thật)

Khi Python scoring service down/timeout: Go orchestrator **degrade gracefully** — chỉ dùng rule-based oracle, đánh dấu kết quả report là "partial confidence". Không để cả pipeline crash theo 1 service phụ.

## Roadmap

- [ ] Phase 1 — Onboarding (xem `ONBOARDING.md`, ~10 tuần, ~1h/ngày; tiến độ ở `docs/planning/progress.md`)
- [ ] Phase 2 — Implement thật: nối gRPC, chọn SUT (OSS demo app + AI feature), viết 5 kịch bản mutation thật
- [ ] Phase 3 — Playwright E2E lấy output thật từ UI
- [ ] Phase 4 — Chạy full experiment, đo Mutation Kill Rate thật, viết báo cáo kết quả
- [ ] Phase 5 — Đóng gói CI gate (GitHub Actions), viết README hoàn chỉnh cho portfolio
