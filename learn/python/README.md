# learn/python — Sân tập Phase 0C & Phase 1 (B, D)

Một venv duy nhất cho toàn bộ bài Python. Mỗi bài là một thư mục con.

## Khởi tạo môi trường (làm 1 lần, đầu Phase 0C)

```powershell
cd D:\calibrate\learn\python
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install --upgrade pip
```

Dấu hiệu đã vào venv: prompt có tiền tố `(.venv)`.
Mỗi phiên terminal mới phải `Activate.ps1` lại.

Nếu PowerShell chặn script:

```powershell
Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
```

## Dependency

`requirements.txt` ở đây **chỉ cần cài khi tới Checkpoint B** (Phase 1) — nó kéo về
`sentence-transformers` + `torch`, vài GB. Phase 0C T1–T3 chạy bằng Python thuần, không cần gì.

```powershell
# chỉ chạy khi bắt đầu Checkpoint B
pip install -r requirements.txt
```

Đây cũng là bài học `venv` của 0C T3: hiểu vì sao dependency phải bị nhốt trong 1 thư mục,
thay vì cài thẳng vào Python hệ thống.

## Chạy bài

```powershell
cd D:\calibrate\learn\python
python c1-mutable\main.py
```

## Thư mục

| Thư mục | Bài |
|---|---|
| `c1-mutable/` | 0C T1 — mutable vs immutable |
| `c2-movingavg/` | 0C T2 — class `MovingAverage` |
| `c3-ctxmanager/` | 0C T3 — context manager tự chế |
| `p1b-similarity/` | Phase 1 Checkpoint B — cosine similarity |
| `p1d-controlchart/` | Phase 1 Checkpoint D — control chart |
