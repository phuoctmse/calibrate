# Calibrate — Onboarding Roadmap (Core-First)

Bản này thay bản trước — không chỉ học "vừa đủ để chạy checkpoint", mà học đúng nền tảng ngôn ngữ trước, rồi mới áp dụng vào project. Đánh đổi thật: **dài hơn bản trước đáng kể** (xem mục Ước tính thời gian ở cuối) — chấp nhận vì mục tiêu là hiểu sâu, không phải chỉ để chạy được 1 lần.

Cấu trúc 3 lớp, không được nhảy cóc:
- **Phase 0 — Core Language Fundamentals:** học ngôn ngữ như một môn học riêng, bài tập tự đứng độc lập, KHÔNG liên quan trực tiếp đến Calibrate. Mục tiêu: hiểu bản chất, không phải thuộc cú pháp để chép.
- **Phase 1 — Applied Checkpoints:** dùng nền tảng Phase 0 để làm đúng 4 checkpoint đã map với kiến trúc dự án (giữ nguyên từ bản trước).
- **Phase 2 — Dự án thật.**

---

# PHASE 0A — Go Core Fundamentals (~3 tuần)

### Tuần 1: Type system, control flow, functions
- Biến, kiểu dữ liệu cơ bản (int, string, bool, float), zero value là gì (khác Python — Go không có `None` mặc định, biến luôn có giá trị khởi tạo).
- `if/for/switch` — Go chỉ có 1 loại vòng lặp (`for`), không có `while`.
- Function: multiple return values (đặc trưng Go, dùng để trả `(result, error)` — nền tảng cho error handling triết lý Go).
- `defer` — chạy code khi function kết thúc, dùng để đóng resource (giống `finally` nhưng khác cách dùng thực tế).

**Bài tập tự đứng (không liên quan Calibrate):** viết hàm chia 2 số, trả về `(kết quả, error)` thay vì throw exception. Gọi hàm với input gây lỗi (chia 0), xử lý đúng bằng `if err != nil`. Đây là bài tập bắt buộc hiểu triết lý "error as value" — khác hẳn Python (`try/except`), sẽ dùng xuyên suốt code Go sau này.

### Tuần 2: Struct, interface, composition
- Struct — cách Go định nghĩa "object" (không có class).
- Method gắn vào struct (receiver — phân biệt value receiver vs pointer receiver, đây là lỗi phổ biến nhất người mới học Go hay mắc).
- Interface — Go dùng **implicit satisfaction** (không cần khai báo "implements" như Java/TS) — hiểu tại sao điều này quan trọng cho việc test (mock interface dễ dàng).
- Composition qua embedding (Go không có kế thừa như OOP truyền thống).

**Bài tập tự đứng:** định nghĩa interface `Shape` với method `Area()`. Viết 2 struct `Circle` và `Rectangle` cùng implement interface đó (không khai báo gì thêm, chỉ cần đúng method signature). Viết hàm nhận `Shape` làm tham số, gọi được với cả 2 struct — chứng minh hiểu implicit interface.

### Tuần 3: Slice, map, pointer, error wrapping
- Slice vs array (Go dùng slice là chính, array hiếm dùng) — cách slice chia sẻ underlying array (nguồn bug phổ biến nếu không hiểu).
- Map — cú pháp, kiểm tra key tồn tại (`v, ok := m[key]`).
- Pointer — khi nào dùng `*T` thay vì `T` (mutate dữ liệu qua function, tránh copy struct lớn).
- Error wrapping: `fmt.Errorf("...: %w", err)`, `errors.Is`, `errors.As` — cách Go xây dựng error chain có ngữ cảnh (sẽ dùng khi orchestrator gọi gRPC lỗi, cần biết lỗi gốc từ đâu).

**Bài tập tự đứng:** viết 1 hàm `ParseConfig` giả lập đọc config, cố tình gây lỗi ở 1 bước con, wrap lỗi đó với context ("failed to parse field X: %w"), ở `main` in ra full error chain và dùng `errors.Is` kiểm tra loại lỗi gốc.

---

# PHASE 0B — Go Concurrency Fundamentals (~1.5 tuần)

Tách riêng khỏi fundamentals cơ bản vì đây là phần khó nhất và quan trọng nhất cho vai trò orchestrator.

- Goroutine — chạy hàm bất đồng bộ, khác gì thread hệ điều hành (nhẹ hơn nhiều, do Go runtime quản lý).
- Channel — cách goroutine giao tiếp an toàn (thay vì shared memory + lock).
- `sync.WaitGroup` — đợi N goroutine hoàn thành.
- `sync.Mutex` — khi nào cần lock thật sự (khi nhiều goroutine cùng ghi vào 1 biến chung, channel không giải quyết hết mọi trường hợp).
- `select` — chờ nhiều channel cùng lúc, dùng cho timeout.
- **Race condition là gì và cách phát hiện:** chạy `go run -race` để tự thấy 1 race condition xảy ra khi cố tình viết sai (2 goroutine cùng ghi vào 1 map không có lock) — đây là bài học quan trọng nhất, phải tự mắt thấy lỗi race để hiểu tại sao cần channel/mutex.

**Bài tập tự đứng:** viết chương trình có race condition thật (2 goroutine cùng tăng 1 biến đếm không có lock), chạy `go run -race` xem Go tự báo lỗi. Sau đó fix bằng `sync.Mutex`, chạy lại `-race` để xác nhận hết lỗi. Đây là dạng bài tập "thấy bug thật rồi tự fix" — không chỉ đọc lý thuyết.

---

# PHASE 0C — Python Core Fundamentals (~2.5 tuần)

### Tuần 1: Type system, mutability, control flow
- Dynamic typing — khác Go: biến không cần khai báo kiểu.
- Mutable vs immutable: `list`/`dict`/`set` (mutable) vs `tuple`/`str`/`int` (immutable) — hiểu vì sao điều này gây bug phổ biến (truyền `list` vào function rồi bị mutate ngoài ý muốn — khác hẳn cách Go truyền value theo mặc định).
- List/dict/set comprehension.
- `try/except/finally`, custom exception class.

**Bài tập tự đứng:** viết hàm nhận vào 1 `list`, cố tình mutate nó bên trong function, gọi hàm rồi in `list` gốc ra ngoài để tự thấy nó đã đổi — sau đó viết lại đúng cách (copy trước khi mutate) để so sánh.

### Tuần 2: Function, class, OOP
- `*args`, `**kwargs`, default argument (và bẫy kinh điển: default argument là mutable object).
- Closure — hàm trả về hàm, giữ lại biến từ scope ngoài.
- Class, `__init__`, dunder methods cơ bản (`__repr__`, `__eq__`).
- Inheritance vs composition — Python cho phép đa kế thừa, nhưng ưu tiên composition khi có thể (giống triết lý Go dù cơ chế khác).

**Bài tập tự đứng:** viết class `MovingAverage` giữ 1 danh sách điểm số, có method `add(score)` và `average()`. Đây chính là building block thật sẽ dùng cho control chart baseline ở Phase 1 — nhưng ở đây làm như bài tập độc lập trước, chưa cần đúng logic SPC.

### Tuần 3 (nửa tuần): Type hints, context manager, virtual env
- Type hints (`def f(x: int) -> str:`) — quan trọng vì code mẫu trong skeleton project đã dùng type hints, cần đọc hiểu được.
- Context manager (`with open(...) as f:`) — cách Python quản lý resource, sẽ dùng khi mở gRPC server/connection.
- `venv`, `pip install -r requirements.txt` — thực hành tạo virtual env, cài đúng dependency của `python-scoring-service/requirements.txt`.

**Bài tập tự đứng:** viết 1 context manager tự chế bằng class (`__enter__`/`__exit__`) mô phỏng việc "mở kết nối" và "đóng kết nối", in log ra để thấy rõ thứ tự gọi.

---

# PHASE 1 — Applied Checkpoints (map trực tiếp vào kiến trúc Calibrate)

*(Giữ nguyên tinh thần bản roadmap trước, giờ dựa trên nền Phase 0 vững hơn)*

### Checkpoint A — Mutation harness khung xương (dùng Phase 0B)
Chạy N "task giả" song song bằng goroutine + `WaitGroup`, gom kết quả về, in tổng kết. Yêu cầu thêm so với bản cũ: chạy thử `-race` để tự xác nhận không có race condition trong code checkpoint.

### Checkpoint B — Similarity script (dùng Phase 0C)
Nhận 2 câu văn bản, dùng `sentence-transformers` encode thành vector, tính cosine similarity bằng `numpy`, in điểm 0-1.

### Checkpoint C — gRPC nối 2 ngôn ngữ
Định nghĩa `.proto`, generate code Go + Python, Go client gọi Python server thật qua network (không phải hàm gọi trực tiếp).

### Checkpoint D — Control chart thật
Dùng lại class `MovingAverage` từ bài tập Phase 0C tuần 2, mở rộng thành control chart đúng công thức (mean ± sigma × std), tự viết không copy thư viện có sẵn cho phần logic cốt lõi.

### Checkpoint E — Thiết kế 5 kịch bản mutation (viết ra giấy, không code)

---

# PHASE 2 — Dự án thật
*(giữ nguyên như bản roadmap trước — ráp nối các checkpoint vào skeleton `main.go` / `server.py` đã có sẵn TODO)*

---

## Ước tính thời gian (thành thật, không né)

| Phase | Thời lượng (1h/ngày) |
|---|---|
| 0A — Go fundamentals | ~3 tuần |
| 0B — Go concurrency | ~1.5 tuần |
| 0C — Python fundamentals | ~2.5 tuần |
| 1 — Applied checkpoints | ~3 tuần |
| **Tổng onboarding** | **~10 tuần** |
| 2 — Dự án thật (đã ước tính trước) | ~9-10 tuần |
| **Tổng Calibrate end-to-end** | **~19-20 tuần** |

So với bản roadmap trước (7 tuần onboarding), bản core-first này dài hơn ~3 tuần vì thêm hẳn 2 tuần Go fundamentals thuần (type system, struct/interface) và tách riêng phần concurrency ra làm kỹ hơn, cộng nửa tuần Python OOP/type hints — đây là chi phí thật của việc học "cốt lõi" thay vì "vừa đủ dùng". Đổi lại, kiến thức có được sẽ dùng lại được cho cả ResilienceKit (Go) và bất kỳ project Go/Python nào sau này, không chỉ dùng 1 lần cho Calibrate.