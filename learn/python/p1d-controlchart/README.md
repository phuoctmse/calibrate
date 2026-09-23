# Phase 1 — Checkpoint D: Control chart thật

Giáo án: [`../../../docs/planning/phase-1-checkpoints.md`](../../../docs/planning/phase-1-checkpoints.md)

**Khoá cho tới khi Checkpoint C được xác nhận đạt.**

## Đề bài

> Dùng lại class `MovingAverage` từ bài 0C T2, mở rộng thành control chart đúng công thức
> (mean ± sigma × std), tự viết phần logic cốt lõi, không copy thư viện có sẵn.

Đây là **metric trung tâm** của Calibrate. Sai chỗ này thì cả dự án sai mà không ai biết —
đúng cái rủi ro mà ADR "Polyglot service split" trong `README.md` nói tới.

## File bạn tự tạo

`main.py` — bắt đầu bằng cách copy class `MovingAverage` của chính bạn từ `c2-movingavg/`.

## Tiêu chí đạt

- [ ] Tính được `mean` và `std` của baseline; **tự viết công thức**, không gọi `numpy.std`
      cho phần cốt lõi (được dùng `numpy` để đối chiếu kết quả)
- [ ] Phân biệt rõ và **chọn có lý do**: std của population (chia N) hay sample (chia N-1)
- [ ] Tính `UCL = mean + k*std` và `LCL = mean - k*std`, `k` là tham số (mặc định 3)
- [ ] Có method kiểu `is_out_of_control(value)` trả về bool
- [ ] Yêu cầu tối thiểu số điểm baseline trước khi cho phép đánh giá; chưa đủ thì raise rõ ràng
      (baseline 2 điểm mà đòi kết luận là vô nghĩa về mặt thống kê)
- [ ] **Đối chiếu:** in kết quả của bạn cạnh `numpy.mean` / `numpy.std` để chứng minh khớp
- [ ] Test tay ít nhất 3 case: điểm trong ngưỡng, điểm vượt trên, điểm vượt dưới
- [ ] Có type hints đầy đủ

## Bẫy hay mắc

Nhầm population std với sample std rồi lệch ngưỡng · `std = 0` khi mọi điểm giống nhau
(mọi giá trị khác đều thành out-of-control — xử lý hay chấp nhận? quyết định và ghi lại) ·
đưa chính điểm đang kiểm tra vào baseline rồi tự pha loãng ngưỡng ·
hardcode `k = 3` rải rác khắp file.

## Nối về dự án

Ánh xạ vào `TODO` control chart trong `python-scoring-service/server.py`.
**Không đụng file đó bây giờ.**
