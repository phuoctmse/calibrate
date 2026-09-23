# Phase 1 — Checkpoint B: Similarity script

Giáo án: [`../../../docs/planning/phase-1-checkpoints.md`](../../../docs/planning/phase-1-checkpoints.md)

**Khoá cho tới khi 0C T3 được xác nhận đạt.**

## Đề bài

> Nhận 2 câu văn bản, dùng `sentence-transformers` encode thành vector,
> tính cosine similarity bằng `numpy`, in điểm 0-1.

## Chuẩn bị

```powershell
cd D:\calibrate\learn\python
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
```

Lần chạy đầu sẽ tải model về (vài trăm MB). Chuyện bình thường, không phải treo.

## File bạn tự tạo

`main.py`.

## Tiêu chí đạt

- [ ] Load model một lần, tái sử dụng — **không** load lại mỗi lần tính
- [ ] Tự viết công thức cosine similarity bằng `numpy` (dot product / tích 2 norm),
      **không** gọi hàm `util.cos_sim` có sẵn — đây là phần phải hiểu
- [ ] Kết quả nằm trong `[-1, 1]`; giải thích được vì sao với embedding thường thấy `[0, 1]`
- [ ] Thử **3 cặp câu**: gần như giống hệt, cùng chủ đề khác chữ, hoàn toàn khác chủ đề.
      Ghi lại 3 điểm số và nhận xét xem thứ tự có đúng trực giác không
- [ ] Có type hints; hàm tính similarity **tách riêng**, nhận 2 vector, không dính tới model
- [ ] Xử lý được input rỗng (raise rõ ràng, không trả `nan` im lặng)

## Vì sao tách hàm tính ra riêng

Vì Checkpoint C sẽ gọi hàm đó qua gRPC. Hàm dính chặt vào model thì không bọc được.
Đây là bài học thiết kế, không phải yêu cầu hình thức.

## Xác minh

```powershell
python p1b-similarity\main.py
```

## Nối về dự án

Ánh xạ vào `TODO` trong `python-scoring-service/server.py`. **Không đụng file đó bây giờ.**
