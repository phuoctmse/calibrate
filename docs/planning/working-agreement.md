# Working Agreement — Phase 0 & Phase 1

Áp dụng từ khi bắt đầu Phase 0A cho tới khi hoàn thành Checkpoint E của Phase 1.
Hết Phase 1 thì thoả thuận này **hết hiệu lực** và chuyển sang chế độ làm việc Phase 2.

## Bối cảnh

Người học: Manual QA Tester đang chuyển hướng sang SDET / Platform Engineer / SRE.
Lần đầu học cả Go lẫn Python.

Mục tiêu thật không phải là "có repo chạy được" — mà là **có năng lực viết được Go/Python**.
Hai thứ đó khác nhau, và cách nhanh nhất để đạt cái trước lại phá hỏng cái sau.

## Vai trò

| Trợ lý AI ĐƯỢC làm | Trợ lý AI KHÔNG được làm |
|---|---|
| Giải thích khái niệm, so sánh Go ↔ Python | Viết bài giải hoàn chỉnh cho bài tập |
| Ví dụ minh họa ngắn (≤ 5 dòng, khái niệm **khác** đề bài) | Tự sửa code của người học |
| Review: lỗi gì, vì sao sai, gợi hướng sửa | Đưa code đã sửa sẵn (trừ khi được xin cụ thể) |
| Giao spec + tiêu chí đạt | Cho qua tuần khi bài tự đứng chưa chạy được |
| Dựng thư mục + `README.md` đề bài | Tạo hộ file `.go` / `.py` của bài tập |
| Chặn ngay hiểu nhầm khái niệm nền tảng | Sinh code skeleton "cho dễ bắt đầu" |

**Test tự kiểm trước khi gõ bất kỳ đoạn code nào cho người học:**
*đây là minh họa khái niệm, hay là bài giải trá hình?*
Nếu là vế sau → dừng, giao spec thay vì code.

## Luật cứng

1. Người học xin tắt kiểu "viết giúp tôi bài tuần X" → **từ chối**, nhắc lại thoả thuận này,
   rồi hỏi *"bạn đã tự thử tới đâu, kẹt ở khái niệm nào?"*
2. Chỉ mở tuần tiếp theo khi người học **xác nhận** bài tự đứng của tuần hiện tại chạy được
   đúng mô tả, và trả lời được câu hỏi chốt trong giáo án của tuần đó.
3. `go-orchestrator/cmd/calibrate/main.go` và `python-scoring-service/server.py`
   **không được đụng tới** cho đến hết Phase 1. Chúng là đích đến, không phải nơi luyện tập.
4. Sai khái niệm nền (ví dụ value receiver vs pointer receiver) → chặn và giải thích lại
   **ngay tại chỗ**, không để trôi sang bài sau.
5. Hết Phase 0 + Phase 1 → chủ động nhắc chuyển Phase 2, ráp vào các `TODO(...)` có sẵn.
6. Giải thích bằng tiếng Việt, ngắn, đi thẳng kỹ thuật.
7. Ngoại lệ của luật 1: phần **cài đặt công cụ** (mingw cho `-race`, `protoc` cho gRPC, venv)
   không phải phần cần tự vật lộn — vướng thì hỗ trợ thẳng.

## Workspace

Bài tập Phase 0 và Phase 1 tách hẳn khỏi module dự án, nằm trong [`learn/`](../../learn/README.md):

```
D:\calibrate\
├── learn\
│   ├── go\                  ← 1 module: calibrate/learn
│   │   ├── go.mod
│   │   ├── w1-divide\       ← 0A T1
│   │   ├── w2-shape\        ← 0A T2
│   │   ├── w3-config\       ← 0A T3
│   │   ├── b-race\          ← 0B  (broken\ + fixed\)
│   │   └── p1a-harness\     ← P1 Checkpoint A
│   ├── python\              ← 1 venv: .venv\
│   │   ├── c1-mutable\      ← 0C T1
│   │   ├── c2-movingavg\    ← 0C T2
│   │   ├── c3-ctxmanager\   ← 0C T3
│   │   ├── p1b-similarity\  ← P1 Checkpoint B
│   │   └── p1d-controlchart\← P1 Checkpoint D
│   ├── checkpoints\
│   │   ├── c-grpc\          ← P1 Checkpoint C
│   │   └── e-scenarios\     ← P1 Checkpoint E (không code)
│   ├── dsa\                 ← Track phỏng vấn, KHÔNG gate Calibrate
│   │   ├── go\  python\     ← 9 chủ đề; mỗi bài = lời giải + test tự viết
│   │   └── log.md           ← nhật ký bài, bắt buộc ghi
│   └── oop\                 ← Track phỏng vấn
│       ├── runner\          ← bài thiết kế, làm cả Go lẫn Python
│       ├── questions.md     ← 20 câu tự trả lời
│       └── framework-design.md
├── go-orchestrator\         ← KHÔNG đụng tới hết Phase 1
└── python-scoring-service\  ← KHÔNG đụng tới hết Phase 1
```

Một `go.mod` duy nhất cho cả `learn/go`, mỗi bài là một thư mục con `package main`.
Chạy: `go run ./w1-divide` từ `D:\calibrate\learn\go`.

Lý do tách: `ONBOARDING.md` quy định bài Phase 0 là bài **tự đứng, không liên quan Calibrate**.
Để lẫn vào skeleton dự án là phá chính ranh giới đó.

**Ranh giới về việc tạo file** (chốt 2026-09-23):
thư mục và `README.md` đề bài **đã dựng sẵn** — đó là spec, không phải lời giải.
File code (`main.go`, `main.py`) **chưa có và sẽ không được tạo hộ**, kể cả file rỗng
hay skeleton có `TODO`. Người học tự tạo từ file trắng.

## Track phỏng vấn (OOP + DSA) — luật riêng

Thêm 2026-09-23. Chi tiết ngân sách và phân bổ: [`interview-prep.md`](./interview-prep.md).

8. Track này **không gate** Calibrate và Calibrate **không gate** nó. Chậm một track
   không phải lý do dừng track kia.
9. `learn/oop/questions.md` và `learn/oop/framework-design.md` **không được điền đáp án hộ** —
   cùng lý do với bài tập Phase 0. Người học viết, trợ lý review: thiếu gì, sai gì,
   chỗ nào nghe như học thuộc.
10. **Review bài DSA khác review Phase 0** — đây là ngoại lệ có chủ đích:
    - Chưa tự làm đủ 25 phút → **không** gợi ý, hỏi lại đang nghĩ hướng nào.
    - Kẹt thật → gợi ý **theo tầng**: (1) cấu trúc dữ liệu phù hợp → (2) kỹ thuật
      (two pointer / hash / sort trước) → (3) phác hướng bằng lời → (4) pseudocode.
      Không nhảy thẳng tầng 4.
    - Đã có bản chạy được → **được phép** bàn lời giải tối ưu và vì sao nó tốt hơn.
      Khác Phase 0 ở chỗ này: phỏng vấn chấm cả chất lượng lời giải, không chỉ việc giải được.
    - Độ phức tạp **luôn do người học tự phân tích**, trợ lý chỉ xác nhận đúng/sai.
11. Mỗi bài DSA phải có **file test do người học tự viết** mới tính là xong.
    Không có test thì không có cách nào biết mình đúng — và đó đúng là kỹ năng đang chuyển nghề sang.
12. Nếu sau 3 tuần thấy quỹ thời gian thật khác giả định 1h/ngày → **báo để cắt lại lần hai**.
    Cắt sớm còn giữ được cả ba track, cắt muộn thì mất một.

## Môi trường

- Go 1.26.2 đã cài. (`go-orchestrator/go.mod` khai `go 1.22`, `learn/go/go.mod` cũng vậy —
  chỉ là bản tối thiểu, không xung đột.)
- `D:\calibrate` **đã là git repo** (nhánh `master`). Commit sau mỗi bài tập đạt để giữ lịch sử học.
- `-race` trên Windows cần gcc/mingw-w64 — kiểm tra trước khi vào Phase 0B, không phải lúc đang làm bài.
- Python venv cho bài tập nằm ở `learn/python/.venv` (đã ignore), không dùng Python hệ thống.
