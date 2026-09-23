# Phase 1 — Checkpoint E: 5 kịch bản mutation

Giáo án: [`../../../docs/planning/phase-1-checkpoints.md`](../../../docs/planning/phase-1-checkpoints.md)

**Không code.** Đây là checkpoint thiết kế — và là checkpoint gần với nghề QA của bạn nhất,
nên đừng coi nhẹ nó vì "chỉ viết chữ".

## Đề bài

> Thiết kế 5 kịch bản mutation (viết ra, không code).

Mutation ở đây nghĩa là: **cố tình làm hệ thống AI kém đi theo cách kiểm soát được**,
rồi xem oracle có phát hiện ra không. Oracle không phát hiện được = oracle không đáng tin.
Tỉ lệ phát hiện chính là **Mutation Kill Rate** — deliverable trung tâm của Calibrate.

## Việc phải làm

Điền [`scenarios.md`](./scenarios.md) trong thư mục này.

## Tiêu chí đạt

- [ ] Đủ **5** kịch bản, mỗi cái nhắm vào một **cơ chế suy giảm khác nhau**
      (không phải 5 biến thể của cùng một ý)
- [ ] Mỗi kịch bản có: cách mutate, điều gì *phải* thay đổi ở output, tín hiệu oracle *phải* bắt được
- [ ] Ít nhất 1 kịch bản mà bạn **dự đoán oracle sẽ KHÔNG bắt được** — và nói rõ vì sao.
      Một bộ kịch bản mà cái nào cũng bắt được là bộ kịch bản quá dễ, vô giá trị
- [ ] Mỗi kịch bản phải **kiểm soát được và lặp lại được** — không phải "làm model tệ đi" chung chung
- [ ] Nói rõ kịch bản nào cần embedding similarity, cái nào chỉ cần rule-based oracle

## Gợi ý hướng nghĩ (không phải đáp án)

Nghĩ theo tầng bị tác động, không theo triệu chứng: tầng model, tầng input đưa vào model,
tầng ngữ cảnh/dữ liệu, tầng tham số sinh, tầng hậu xử lý.
Với mỗi tầng hỏi: *làm hỏng ở đây thì output đổi kiểu gì, và tín hiệu nào lộ ra ngoài?*

Câu hỏi tự kiểm: **nếu oracle bỏ sót kịch bản này, người dùng thật có nhận ra không?**
Nếu người dùng cũng không nhận ra, thì kịch bản đó có đáng đưa vào không?
