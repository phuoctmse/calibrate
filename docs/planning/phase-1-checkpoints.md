# Phase 1 — Applied Checkpoints (~3 tuần)

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 1".

**Chỉ mở khi toàn bộ Phase 0 đã `[x]` trong [`progress.md`](./progress.md).**

Khác biệt so với Phase 0: bài tập từ đây **có liên quan trực tiếp tới Calibrate**.
Mỗi checkpoint là một mảnh thật của kiến trúc, làm rời ra trước khi ráp lại ở Phase 2.

Điều **không** đổi: working agreement vẫn còn hiệu lực. Vẫn không viết code hộ.
`go-orchestrator/cmd/calibrate/main.go` và `python-scoring-service/server.py`
vẫn **không được đụng** cho tới hết Phase 1.

## Thứ tự và phụ thuộc

```
0B ──────▶ A (harness)  ─┐
                          ├──▶ C (gRPC) ──▶ D (control chart) ──▶ E (kịch bản)
0C ──────▶ B (similarity)─┘
```

A và B độc lập nhau, làm được song song. C cần cả hai. D cần C chạy được
(vì D sẽ là RPC thứ hai). E làm cuối vì cần hiểu toàn bộ pipeline mới thiết kế được kịch bản tốt.

| | Checkpoint | Thư mục | Nền từ |
|---|---|---|---|
| A | Mutation harness khung xương | [`learn/go/p1a-harness/`](../../learn/go/p1a-harness/README.md) | 0B |
| B | Similarity script | [`learn/python/p1b-similarity/`](../../learn/python/p1b-similarity/README.md) | 0C |
| C | gRPC nối 2 ngôn ngữ | [`learn/checkpoints/c-grpc/`](../../learn/checkpoints/c-grpc/README.md) | A + B |
| D | Control chart thật | [`learn/python/p1d-controlchart/`](../../learn/python/p1d-controlchart/README.md) | 0C T2 + C |
| E | 5 kịch bản mutation | [`learn/checkpoints/e-scenarios/`](../../learn/checkpoints/e-scenarios/README.md) | A–D |

Đề bài đầy đủ + tiêu chí đạt nằm trong README của từng thư mục.
File này chỉ ghi **ý đồ** và **chỗ dễ trượt** của từng checkpoint.

---

## A — Mutation harness khung xương

**Ý đồ:** dựng đúng cái khung điều phối sẽ dùng thật, khi task còn giả nên sai cũng rẻ.
Cái phải học ở đây không phải goroutine (đã học 0B) mà là **kiến trúc fan-out/fan-in**:
phát N việc ra, gom N kết quả về, một việc hỏng không kéo cả mẻ hỏng theo.

**Chỗ dễ trượt:** gom kết quả bằng slice dùng chung + mutex thay vì channel (chạy được nhưng
trượt mất bài học); đóng channel sai chỗ gây deadlock; task lỗi làm sập cả harness —
đây chính là mầm của "degrade gracefully" ở Phase 2.

**Câu hỏi chốt:** "Một task panic. Harness nên làm gì?" Trả lời được là hiểu vì sao
orchestrator phải cô lập lỗi từng task.

## B — Similarity script

**Ý đồ:** hiểu embedding similarity là **một phép đo có sai số**, không phải chân lý.
Số 0.87 nghĩa là gì — đó mới là câu hỏi của Calibrate.

**Chỗ dễ trượt:** gọi `util.cos_sim` có sẵn thay vì tự viết (mất hết phần học);
load model trong vòng lặp; không tách hàm tính ra khỏi model nên C không bọc được.

**Câu hỏi chốt:** "Hai câu trái nghĩa nhau nhưng cùng chủ đề. Similarity cao hay thấp?
Điều đó nói gì về việc dùng similarity làm oracle?" Đây là hạt giống của Checkpoint E.

## C — gRPC nối 2 ngôn ngữ

**Ý đồ:** checkpoint khó nhất Phase 1. Lần đầu có **ranh giới mạng** thật giữa hai ngôn ngữ,
và một **hợp đồng** (`.proto`) độc lập với cả hai.

**Nguyên tắc:** bắt đầu bằng proto rút gọn **tự viết** — 1 RPC, 1 request, 1 response.
Không dùng thẳng `proto/scoring.proto` (2 RPC, 4 message) cho lần đầu; quá nhiều thứ
hỏng cùng lúc thì không debug được.

**Chỗ dễ trượt:** phần lớn thời gian sẽ tốn vào cài `protoc` / plugin, **không** vào kiến thức.
Vướng ở đó thì hỏi ngay, đừng tự vật lộn — nó không phải phần cần học.
Ngoài ra: client treo vô hạn khi server chết (thiếu `context.WithTimeout`);
commit code generate vào git.

**Câu hỏi chốt:** "Vì sao `.proto` chứ không phải JSON qua HTTP?" Phải nói được ít nhất:
hợp đồng có kiểu, sinh code cả hai phía, versioning field rõ ràng.

## D — Control chart thật

**Ý đồ:** đây là **metric trung tâm** của cả dự án. Sai ở đây thì mọi con số Calibrate đưa ra
đều vô nghĩa mà không ai phát hiện — đúng cái rủi ro mà ADR "Polyglot service split"
trong `README.md` viện dẫn để tách service Python ra.

**Chỗ dễ trượt:** nhầm population std (chia N) với sample std (chia N-1);
`std = 0` khi mọi điểm giống nhau khiến mọi giá trị khác thành out-of-control;
đưa chính điểm đang kiểm tra vào baseline; baseline quá ít điểm mà vẫn kết luận.

**Bắt buộc:** tự viết công thức, rồi **in cạnh** `numpy.mean` / `numpy.std` để đối chiếu.
Tự viết mà không đối chiếu là tự tin không có căn cứ.

**Câu hỏi chốt:** "Baseline có 3 điểm. Kết luận out-of-control có đáng tin không? Vì sao?"

## E — 5 kịch bản mutation

**Ý đồ:** checkpoint **không code**, và là checkpoint gần nghề QA của bạn nhất.
Đây là lúc bốn checkpoint trước gộp lại thành một câu hỏi: *oracle này có đáng tin không?*

**Chỗ dễ trượt:** 5 biến thể của cùng một cơ chế; kịch bản mơ hồ không lặp lại được;
và nguy hiểm nhất — bộ kịch bản mà **cái nào oracle cũng bắt được**.
Bộ đó vô giá trị: nó chỉ chứng minh bạn đã chọn toàn bài dễ.

**Bắt buộc:** ít nhất 1 kịch bản dự đoán oracle sẽ **trượt**, kèm lý do.

**Câu hỏi chốt:** "Kill rate 100%. Tin được không?"

---

## Điều kiện kết thúc Phase 1

Cả 5 checkpoint `[x]` trong [`progress.md`](./progress.md), mỗi cái trả lời được câu hỏi chốt.

Khi đó:

1. [`working-agreement.md`](./working-agreement.md) **hết hiệu lực**, chuyển chế độ Phase 2.
2. Mở khoá `go-orchestrator/cmd/calibrate/main.go` và `python-scoring-service/server.py`.
3. Ráp các checkpoint vào `TODO(...)` có sẵn trong hai file đó.
