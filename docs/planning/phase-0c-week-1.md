# Phase 0C — Tuần 1: Dynamic typing, mutability, control flow

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0C — Tuần 1".
Bài tập: [`learn/python/c1-mutable/`](../../learn/python/c1-mutable/README.md).

**Mục tiêu thật:** sau 4.5 tuần Go, bạn đã quen "biến có kiểu, truyền là copy, lỗi là giá trị".
Python **ngược lại cả ba**. Tuần này để thấy rõ chỗ ngược đó — vì mọi bug Python của người
vừa học Go đều nằm ở đó.

> Trạng thái: khung + đề bài đã chốt. Phần giảng chi tiết diễn ra khi mở tuần.

---

## 0. Trước khi viết dòng Python nào

Dựng venv theo [`learn/python/README.md`](../../learn/python/README.md).
Chưa cần cài dependency gì — Phase 0C chạy bằng Python thuần.

## 1. Dynamic typing

```python
x = 5
x = "năm"   # hợp lệ
```

Kiểu gắn với **giá trị**, không gắn với biến. Không khai báo kiểu, không compile —
lỗi kiểu nổ **lúc chạy**, đúng dòng gây lỗi, không sớm hơn.

Đánh đổi so với Go: viết nhanh hơn nhiều, nhưng compiler không còn bắt lỗi giúp.
Bù lại bằng test và type hints (T3).

## 2. Mutable vs immutable — trọng tâm tuần

| Immutable | Mutable |
|---|---|
| `int`, `float`, `str`, `tuple`, `bool`, `frozenset` | `list`, `dict`, `set`, hầu hết object tự định nghĩa |

Python truyền **tham chiếu tới object**, luôn luôn. Hàm nhận được chính object đó,
không phải bản sao.

- Object **mutable** → sửa trong hàm là bên ngoài thấy đổi.
- Object **immutable** → không sửa được; `x = x + 1` tạo object **mới**, biến ngoài không đổi.

Cùng một cơ chế truyền, hai hành vi khác nhau — chỉ do bản chất object. Đây là chỗ
người từ Go sang hiểu sai nhiều nhất, vì Go copy struct mặc định.

Công cụ chứng minh: `id(x)` trả địa chỉ object. Cùng `id` = cùng object.
Bài tập bắt in `id` chứ không chỉ in giá trị — nói thì ai cũng gật, `id` mới là bằng chứng.

```python
a = [1, 2]
b = a          # KHÔNG phải copy, cùng một list
b = a[:]       # shallow copy
b = copy.deepcopy(a)   # deep copy, cần khi list lồng list
```

Shallow copy của list chứa list: lớp ngoài tách, lớp trong vẫn chung. Bẫy tiếp theo sau bẫy đầu.

## 3. Bẫy default argument

```python
def add(item, bucket=[]):   # SAI
    bucket.append(item)
    return bucket
```

Default argument được đánh giá **một lần** lúc định nghĩa hàm, không phải mỗi lần gọi.
Gọi 2 lần thì lần sau thấy dữ liệu lần trước.

Cách đúng: `def add(item, bucket=None):` rồi `if bucket is None: bucket = []`.

Bài tập bắt tự tái hiện lỗi này rồi mới sửa — cùng tinh thần bài race ở 0B:
thấy bug thật rồi mới fix.

## 4. Control flow và comprehension

Thụt lề là cú pháp, không phải thẩm mỹ. `if/elif/else`, `for x in xs`, `while`
(Python **có** `while`, khác Go).

```python
squares = [x*x for x in xs if x > 0]
by_name = {u.name: u for u in users}
```

Comprehension là cách viết Python đặc trưng, đọc code người khác sẽ gặp liên tục.
Giới hạn: lồng quá 2 tầng thì viết `for` thường, dễ đọc hơn.

## 5. `try/except/finally`

```python
try:
    ...
except ValueError as e:
    ...
finally:
    ...
```

Ngược hẳn Go: lỗi là **luồng điều khiển ẩn**. Nhìn `x = f()` không biết `f` có ném gì không.

Luật:

- Bắt **đúng loại** exception. `except:` trần hoặc `except Exception:` nuốt cả `KeyboardInterrupt`,
  cả lỗi lập trình của chính bạn.
- Đừng bắt rồi `pass`. Lỗi bị nuốt im lặng là thứ tệ hơn cả crash — góc nhìn QA hiểu ngay vì sao.
- Custom exception: `class ConfigError(Exception): pass`. Kế thừa từ `Exception`, không từ `BaseException`.

Đối chiếu: sentinel error + `errors.Is` của Go ↔ class exception + `except SpecificError` của Python.
Cùng ý đồ "phân loại lỗi", khác cơ chế.

---

## Checklist review (trợ lý dùng khi chấm)

1. Không in `id()` → không chứng minh được gì, bắt làm lại.
2. `b = a` tưởng là copy.
3. Dùng `.copy()` cho list lồng nhau rồi tưởng đã tách hẳn.
4. Không làm phần default argument — đó là phần học thật, bắt làm.
5. `except:` trần, hoặc `except` rồi `pass`.
6. Kết luận chung chung "Python truyền tham chiếu" mà không phân biệt mutable/immutable.
7. Không có comprehension nào trong bài.
8. Thử sửa `tuple`/`str` mà không ghi lại thông báo lỗi nhận được.

## Hỏi lại để xác nhận hiểu

- "Hàm nhận một `int` và một `list`, sửa cả hai bên trong. Sau khi gọi, biến nào ở ngoài đổi?
  Vì sao — cùng một cách truyền mà?" → trả lời đúng theo mục 2 thì tuần này đạt.

---

**Qua tuần khi:** bài chạy đúng + review sạch + trả lời được câu hỏi trên.
→ Mở T2: function, class, OOP.
