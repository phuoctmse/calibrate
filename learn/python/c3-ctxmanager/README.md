# 0C T3 — Context manager tự chế + type hints

Giáo án: [`../../../docs/planning/phase-0c-week-3.md`](../../../docs/planning/phase-0c-week-3.md)

## Đề bài

> Viết 1 context manager tự chế bằng class (`__enter__` / `__exit__`) mô phỏng việc
> mở kết nối và đóng kết nối, in log ra để thấy rõ thứ tự gọi.

Mục tiêu thật: hiểu cách Python quản lý resource. Ở Phase 2 sẽ dùng khi mở gRPC channel.
Đây cũng là chỗ đối chiếu thẳng với `defer` của Go.

## File bạn tự tạo

`main.py`.

## Tiêu chí đạt

- [ ] Class có `__enter__` và `__exit__`, dùng được bằng `with ... as ...`
- [ ] `__enter__` in log "mở", `__exit__` in log "đóng"; `__enter__` trả về object dùng được
- [ ] **Chứng minh bằng thực nghiệm:** raise exception giữa block `with`,
      cho thấy `__exit__` vẫn chạy — kết nối vẫn được đóng
- [ ] Giải thích được ý nghĩa 3 tham số của `__exit__` và ý nghĩa giá trị trả về
      (`True` nuốt exception, `False` để nó bay tiếp) — chọn đúng cái bạn muốn và nói vì sao
- [ ] **Toàn bộ** function/method trong file có type hints đầy đủ (tham số + kiểu trả về)
- [ ] Hoàn tất mục "Khởi tạo môi trường" trong [`../README.md`](../README.md): venv chạy được

## Xác minh

```powershell
cd D:\calibrate\learn\python
.\.venv\Scripts\Activate.ps1
python c3-ctxmanager\main.py
```

Đạt khi output cho thấy rõ thứ tự: mở → thân block → đóng, **kể cả** khi có exception.

## Bẫy hay mắc

Đóng resource trong thân `with` thay vì trong `__exit__` · `__enter__` không trả gì
(`as x` nhận `None`) · trả `True` từ `__exit__` rồi nuốt mất lỗi thật mà không biết ·
tưởng type hints được Python kiểm tra lúc chạy (không — chúng chỉ là chú thích,
công cụ ngoài như `mypy` mới kiểm).
