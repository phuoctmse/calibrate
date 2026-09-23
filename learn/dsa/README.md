# learn/dsa — Luyện DSA cho phỏng vấn SDET

Giáo án: [`../../docs/planning/dsa-curriculum.md`](../../docs/planning/dsa-curriculum.md)
Nhật ký: [`log.md`](./log.md) ← **ghi sau mỗi bài, không bỏ**

Track này **không gate** Calibrate. Mục tiêu ~60 bài trong 22 tuần, 20-30 phút/ngày.

## Điểm khác biệt quan trọng: bài DSA được **kiểm bằng test bạn tự viết**

Không có lời giải mẫu để đối chiếu. Cách biết mình đúng là **tự viết test** —
đúng kỹ năng bạn đang chuyển nghề sang làm, và cũng là thứ phỏng vấn SDET hỏi:
*"bạn kiểm tra hàm này thế nào?"*

Mỗi bài gồm **2 file**: file lời giải + file test. Thiếu file test thì bài chưa xong.

Test tối thiểu phải có: case bình thường · case rỗng · case 1 phần tử ·
case biên (đầu/cuối mảng) · case không tìm thấy.

## Bố cục

```
dsa/
├── log.md                  ← nhật ký bài, cột "sai ở đâu" là cột quan trọng nhất
├── go/                     ← module calibrate/dsa
│   ├── arrays/  hashmap/  strings/  linkedlist/  stackqueue/
│   └── binarysearch/  trees/  sorting/  graphs/
├── python/                 ← cùng 9 chủ đề
└── requirements.txt        ← pytest
```

Mỗi chủ đề là **một package/thư mục chứa nhiều bài**, không phải một bài một thư mục.

## Quy ước đặt tên

| | Go | Python |
|---|---|---|
| Lời giải | `arrays/two_sum.go` | `arrays/two_sum.py` |
| Test | `arrays/two_sum_test.go` | `arrays/test_two_sum.py` |

Go: mỗi thư mục là một package cùng tên thư mục (`package arrays`), **không** phải `package main`.

## Chạy

```powershell
# Go — chạy hết, hoặc một chủ đề
cd D:\calibrate\learn\dsa\go
go test ./...
go test ./arrays -v

# Python — dùng chung venv với learn/python
cd D:\calibrate\learn\python
.\.venv\Scripts\Activate.ps1
pip install -r ..\dsa\requirements.txt
cd D:\calibrate\learn\dsa\python
pytest -v
pytest arrays -v
```

## Luật làm bài

1. Tự làm **25 phút** trước khi xin gợi ý. Xin thì phải nói được kẹt ở đâu.
2. **Viết độ phức tạp trước khi chạy code.** Đoán rồi kiểm, không phải chạy xong mới tra.
3. Ghi `log.md` **ngay** sau mỗi bài. Cột "sai ở đâu" quan trọng hơn cột "xong".
4. Bài sai → hẹn làm lại sau **7 ngày**, đánh dấu trong log.
5. Không đọc lời giải trước khi tự có một bản chạy được, **kể cả bản chậm**.
6. Ít nhất **1 bài/tuần nói ra trong lúc làm**. Phỏng vấn là bài thi nói.

## Ngôn ngữ nào cho bài nào

- **Tuần 1-5:** chỉ Go (chưa học Python).
- **Từ tuần 6:** giải Python trước. Viết lại Go **chỉ với** nhóm bài trong bảng
  ở [`dsa-curriculum.md`](../../docs/planning/dsa-curriculum.md#quy-tắc-hai-ngôn-ngữ):
  toàn bộ linked list, 2 bài binary search, 3 bài cây.
- Ngoài danh sách đó, **một ngôn ngữ là đủ**.
