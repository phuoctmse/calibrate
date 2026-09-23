# Phase 0C — Tuần 2: Function, class, OOP

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0C — Tuần 2".
Bài tập: [`learn/python/c2-movingavg/`](../../learn/python/c2-movingavg/README.md).

**Mục tiêu thật:** viết được class sạch — và làm quen với `MovingAverage`, class sẽ được
mở rộng thành control chart ở Checkpoint D. Đây là tuần duy nhất của Phase 0 mà sản phẩm
được dùng lại nguyên vẹn về sau.

> Trạng thái: khung + đề bài đã chốt. Phần giảng chi tiết diễn ra khi mở tuần.

---

## 1. `*args`, `**kwargs`, default argument

```python
def f(a, b=1, *args, **kwargs): ...
```

`*args` gom tham số vị trí thừa thành tuple, `**kwargs` gom tham số tên thành dict.

Bẫy mutable default đã học ở T1 — tuần này **không được tái phạm**.

Keyword-only argument (`def f(*, k=3)`) buộc gọi bằng tên. Dùng cho tham số cấu hình
để chỗ gọi tự giải thích: `check(score, sigma=3)` dễ đọc hơn `check(score, 3)`.

## 2. Closure

```python
def multiplier(n):
    def inner(x):
        return x * n     # giữ n từ scope ngoài
    return inner
```

Hàm là object hạng nhất: gán được, truyền được, trả về được.

Bẫy classic: closure bắt **biến**, không bắt **giá trị**. Tạo closure trong vòng lặp thì
tất cả cùng thấy giá trị cuối. Cùng bản chất với bẫy loop-variable capture ở goroutine (0B).

## 3. Class

```python
class MovingAverage:
    def __init__(self, window: int = 0) -> None:
        self._scores = []     # đúng: instance attribute
```

**Bẫy số một của tuần** — attribute khai ở class body dùng chung cho **mọi instance**:

```python
class Bad:
    scores = []    # SAI: mọi instance dùng chung một list
```

Cùng gia đình với bẫy mutable default ở T1. Bài tập bắt tạo 2 instance để tự chứng minh.

`self` là tham số đầu tiên, luôn phải viết ra (khác Java/JS có `this` ẩn).

Quy ước quyền truy cập: Python **không có** `private` thật. `_name` = "nội bộ, đừng đụng"
theo thoả thuận. Khác Go dùng chữ hoa/thường và được compiler ép.

## 4. Dunder methods

| Method | Vai trò |
|---|---|
| `__init__` | khởi tạo (không phải constructor — object đã tồn tại rồi) |
| `__repr__` | chuỗi cho **lập trình viên**, hiện khi debug/log/in trong list |
| `__str__` | chuỗi cho **người dùng cuối** |
| `__eq__` | định nghĩa `==`; mặc định là so sánh **identity**, không so nội dung |
| `__len__` | cho `len(obj)` |

Nếu chỉ viết một cái, viết `__repr__`. Không có nó thì debug thấy
`<MovingAverage object at 0x...>` — vô dụng.

Viết `__eq__` mà không viết `__hash__` thì object mất khả năng làm key của dict/set.

## 5. Inheritance vs composition

Python cho đa kế thừa, nhưng ưu tiên composition — cùng kết luận với Go dù cơ chế khác hẳn.
Kế thừa chỉ dùng khi quan hệ thật sự là "là một", và cây kế thừa nông.

Đối chiếu: Go **buộc** bạn composition (không có kế thừa). Python **cho phép** kế thừa
nhưng khuyên đừng lạm dụng. Cùng bài học, hai cách thực thi.

## 6. Thiết kế `MovingAverage` — câu hỏi phải tự trả lời

Đây là phần khó nhất tuần, và nó là **thiết kế**, không phải cú pháp:

- `average()` khi chưa có điểm nào thì trả gì? Raise / `None` / `0`?
  Ba lựa chọn, ba hệ quả khác nhau ở chỗ gọi. Chọn và **bảo vệ được lựa chọn**.
- `add()` nhận giá trị không phải số thì sao? Raise ngay, hay để nó nổ ở `average()` sau?
  (Nguyên tắc: hỏng sớm, hỏng gần nguồn.)
- `window` lưu bằng `list` hay `deque(maxlen=N)`? Cái sau tự bỏ phần tử cũ, O(1) hai đầu.

Không có đáp án duy nhất. Có đáp án **giải thích được** hoặc không.

---

## Checklist review (trợ lý dùng khi chấm)

1. Attribute mutable khai ở class body — bẫy trọng tâm, chặn ngay.
2. Không trả lời được "`average()` khi rỗng trả gì và vì sao".
3. Không có `__repr__`.
4. `print` thay vì raise khi input sai — che lỗi thay vì báo lỗi.
5. Tái phạm mutable default argument sau khi đã học T1.
6. Không tạo 2 instance để chứng minh không dùng chung state.
7. Không có custom exception.
8. Chia cho 0 khi danh sách rỗng.
9. Dùng kế thừa ở chỗ composition đủ dùng.

## Hỏi lại để xác nhận hiểu

- "Hai instance `MovingAverage` khác nhau. Thêm điểm vào cái A, cái B có đổi không?
  Viết thế nào thì nó **sẽ** đổi?" → trả lời được cả hai vế thì đạt.

---

**Qua tuần khi:** bài chạy đúng + review sạch + trả lời được câu hỏi trên.
→ Mở T3: type hints, context manager, venv. Giữ file bài này lại, Checkpoint D sẽ dùng.
