# Phase 0C — Tuần 3 (nửa tuần): Type hints, context manager, virtual env

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0C — Tuần 3".
Bài tập: [`learn/python/c3-ctxmanager/`](../../learn/python/c3-ctxmanager/README.md).

**Mục tiêu thật:** ba thứ thực dụng để đọc và chạy được code Python thật —
`python-scoring-service/server.py` đã dùng type hints, gRPC sẽ cần context manager,
và không có venv thì không cài nổi dependency của Checkpoint B.

Đây là tuần **cuối** của Phase 0. Xong tuần này là hết phần học ngôn ngữ thuần.

> Trạng thái: khung + đề bài đã chốt. Phần giảng chi tiết diễn ra khi mở tuần.

---

## 1. Type hints

```python
def compute(ref: str, cand: str) -> float: ...
scores: list[float] = []
def find(k: str) -> float | None: ...
```

Điều quan trọng nhất: Python **không kiểm tra** chúng lúc chạy. Sai kiểu vẫn chạy.
Chúng là tài liệu cho người đọc và cho công cụ (`mypy`, IDE) — không phải cho interpreter.

Khác Go: Go compiler ép kiểu và từ chối build. Python hint chỉ là lời hứa.
Hiểu đúng chỗ này để không trông cậy nhầm vào nó.

Cú pháp hiện đại (Python 3.10+): `list[str]` thay `List[str]`, `X | None` thay `Optional[X]`.
`typing` chỉ cần cho thứ phức tạp hơn.

Nối lại: `X | None` chính là chỗ Python biểu diễn "có thể không có giá trị" — vai trò mà Go
làm bằng `(T, error)` và zero value.

## 2. Context manager

```python
class Connection:
    def __enter__(self) -> "Connection":
        print("mở")
        return self
    def __exit__(self, exc_type, exc_value, traceback) -> bool:
        print("đóng")
        return False
```

`with obj as x:` → gọi `__enter__`, gán giá trị **trả về** cho `x`, chạy thân block,
rồi **luôn luôn** gọi `__exit__` — kể cả khi có exception, kể cả khi `return` giữa chừng.

Ba tham số của `__exit__` là thông tin exception (hoặc ba `None` nếu thoát bình thường).
Giá trị trả về:

| Trả | Nghĩa |
|---|---|
| `False` (hoặc `None`) | exception tiếp tục bay lên — **mặc định nên dùng** |
| `True` | **nuốt** exception, coi như không có gì xảy ra |

Trả `True` mà không cố ý là cách tự tạo bug im lặng. Bài tập bắt chọn và giải thích.

Đối chiếu với Go:

| | Go `defer` | Python `with` |
|---|---|---|
| Phạm vi | cả function | đúng block `with` |
| Khai ở đâu | chỗ mở resource | chỗ *dùng* resource |
| Ai viết logic dọn | người gọi | tác giả class (một lần, dùng lại mãi) |

`with` gói logic dọn dẹp vào chính object, nên người dùng không thể quên đóng.
`defer` thì người gọi vẫn phải nhớ viết. Đây là khác biệt thiết kế thật, không chỉ cú pháp.

Bản ngắn gọn `@contextlib.contextmanager` tồn tại — **biết là có**, nhưng bài tập yêu cầu
viết bằng class để nhìn rõ cơ chế.

## 3. Virtual env

Không có venv: `pip install` đổ vào Python hệ thống, hai dự án đòi hai version khác nhau
thì hỏng cả hai. Venv là một thư mục chứa Python + thư viện riêng cho một dự án.

```powershell
cd D:\calibrate\learn\python
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt   # chỉ khi tới Checkpoint B
```

Dấu hiệu đã vào: prompt có `(.venv)`. Mỗi terminal mới phải activate lại.
`.venv/` **không commit** (đã có trong `learn/.gitignore`) — nó dựng lại được từ `requirements.txt`.

Đó là lý do phải pin version trong `requirements.txt`: để máy khác, CI, và bạn 6 tháng sau
dựng ra **đúng** môi trường đó. Reproducibility — cùng lý do `proto/scoring.proto` có field
`model_version`.

Đối chiếu Go: `go.mod` làm việc tương tự nhưng gắn liền vào ngôn ngữ, không cần "activate".

---

## Checklist review (trợ lý dùng khi chấm)

1. Không chứng minh `__exit__` vẫn chạy khi có exception — phần học thật, bắt làm.
2. Trả `True` từ `__exit__` mà không biết mình đang nuốt exception.
3. Đóng resource trong thân `with` thay vì trong `__exit__`.
4. `__enter__` không `return` gì → `as x` nhận `None`.
5. Type hints thiếu ở một số hàm, hoặc thiếu kiểu trả về.
6. Tưởng type hints được kiểm tra lúc chạy — hiểu nhầm nền tảng, chặn ngay.
7. Chưa dựng được venv, hoặc quên activate rồi tưởng lỗi ở code.

## Hỏi lại để xác nhận hiểu

- "Trong thân `with`, code raise exception. `__exit__` có chạy không? Exception có bay ra ngoài không?"
  → phải trả lời: có chạy; bay hay không tuỳ giá trị `__exit__` trả về.

---

**Qua tuần khi:** bài chạy đúng + review sạch + venv hoạt động + trả lời được câu hỏi trên.

**Hết Phase 0.** → Mở [Phase 1](./phase-1-checkpoints.md).
Từ đây bài tập bắt đầu **có liên quan tới Calibrate**.
