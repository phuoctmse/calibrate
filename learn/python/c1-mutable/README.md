# 0C T1 — Mutable vs immutable

Giáo án: [`../../../docs/planning/phase-0c-week-1.md`](../../../docs/planning/phase-0c-week-1.md)

## Đề bài

> Viết hàm nhận vào 1 `list`, cố tình mutate nó bên trong function, gọi hàm rồi in `list`
> gốc ra ngoài để tự thấy nó đã đổi. Sau đó viết lại đúng cách (copy trước khi mutate)
> để so sánh.

Mục tiêu thật: hiểu Python truyền **tham chiếu tới object**, không phải copy — khác hẳn
mặc định của Go. Đây là nguồn bug im lặng phổ biến nhất khi viết Python sau khi quen Go.

## File bạn tự tạo

`main.py`.

## Tiêu chí đạt

- [ ] Có hàm mutate list tại chỗ (`append` / gán theo index), in list gốc **sau khi gọi** để thấy đã đổi
- [ ] Có hàm bản sạch: copy trước rồi mới sửa, trả về list mới; list gốc **không đổi**
- [ ] In `id()` của list trong và ngoài hàm để chứng minh bản mutate là **cùng một object**
- [ ] Làm thêm với `tuple` hoặc `str`: thử sửa tại chỗ và **ghi lại lỗi nhận được**
- [ ] Bẫy kinh điển: viết một hàm có default argument là `[]`, gọi 2 lần liên tiếp,
      in kết quả để tự thấy list bị giữ lại giữa 2 lần gọi
- [ ] Có ít nhất một comprehension (list hoặc dict) trong bài

## Xác minh

```powershell
cd D:\calibrate\learn\python
.\.venv\Scripts\Activate.ps1
python c1-mutable\main.py
```

Đạt khi output tự nó chứng minh được sự khác nhau — người đọc không cần đọc code cũng hiểu.

## Bẫy hay mắc

Dùng `list2 = list1` tưởng là copy · dùng `.copy()` cho list lồng nhau rồi tưởng đã an toàn
(shallow vs deep copy) · in giá trị mà không in `id()` nên không chứng minh được gì ·
kết luận "Python truyền tham chiếu" chung chung thay vì "truyền tham chiếu tới object,
object mutable thì sửa được, immutable thì không".
