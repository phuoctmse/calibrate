# 0C T2 — Class `MovingAverage`

Giáo án: [`../../../docs/planning/phase-0c-week-2.md`](../../../docs/planning/phase-0c-week-2.md)

## Đề bài

> Viết class `MovingAverage` giữ 1 danh sách điểm số, có method `add(score)` và `average()`.

Đây là building block **thật** sẽ dùng lại cho control chart baseline ở Checkpoint D.
Nhưng ở đây làm như bài độc lập — **chưa** cần đúng logic SPC.

## File bạn tự tạo

`main.py`.

## Tiêu chí đạt

- [ ] Class có `__init__` khởi tạo danh sách rỗng **bên trong** `__init__`,
      không phải ở class body (bẫy shared mutable state giữa các instance)
- [ ] `add(score)` thêm 1 điểm; `average()` trả trung bình
- [ ] `average()` khi chưa có điểm nào: **quyết định rõ** trả gì (raise? trả `None`? trả 0)
      và giải thích được vì sao chọn cách đó
- [ ] Có `__repr__` in ra dạng đọc được, hữu ích khi debug
- [ ] Có ít nhất một custom exception class, dùng khi input không hợp lệ
- [ ] Tự chạy thử: tạo **2 instance** riêng, thêm điểm vào instance A,
      in instance B để chứng minh chúng không dùng chung danh sách

## Mở rộng bắt buộc

Thêm tham số `window` (chỉ tính trung bình N điểm gần nhất). Làm xong rồi tự hỏi:
nên lưu bằng `list` hay `collections.deque(maxlen=N)`, và vì sao.

## Xác minh

```powershell
cd D:\calibrate\learn\python
.\.venv\Scripts\Activate.ps1
python c2-movingavg\main.py
```

## Bẫy hay mắc

Khai báo `scores = []` ở class body (mọi instance dùng chung) · quên `self` ·
`average()` chia cho 0 khi rỗng · dùng `print` thay vì raise khi input sai
(che lỗi thay vì báo lỗi) · viết `__str__` mà bỏ `__repr__` (debug khó hơn hẳn).
