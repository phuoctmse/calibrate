# Phase 0B — Go Concurrency (~1.5 tuần)

Nguồn: [`ONBOARDING.md`](../../ONBOARDING.md) mục "PHASE 0B".
Bài tập: [`learn/go/b-race/`](../../learn/go/b-race/README.md).

**Mục tiêu thật:** phần khó nhất và quan trọng nhất cho vai trò orchestrator.
Không phải học cú pháp `go func()` — mà là **tự tay tạo ra một race condition, nhìn nó,
rồi sửa nó**. Đọc lý thuyết về race không bao giờ đủ.

> Trạng thái: khung + đề bài đã chốt. Phần giảng chi tiết diễn ra khi mở phase.
> Đây là phase dài nhất tính theo độ khó — không ép tiến độ.

---

## 0. Kiểm tra trước khi bắt đầu

`-race` trên Windows cần gcc (mingw-w64). Chạy thử ngay hôm mở phase:

```powershell
cd D:\calibrate\learn\go
go run -race ./w1-divide
```

Báo thiếu compiler thì xử lý **trước**, đừng để vướng giữa bài. Không có `-race`
thì cả phase này mất hết ý nghĩa.

## 1. Goroutine

```go
go doWork()   // chạy song song, trả về ngay
```

Nhẹ hơn OS thread rất nhiều (khởi tạo vài KB stack, Go runtime tự ghép lên thread thật).
Chạy hàng nghìn goroutine là chuyện bình thường.

Hai điều phải nhớ ngay:

- `main` kết thúc thì **mọi goroutine bị giết**, không cần biết đang làm gì.
- Không có cách nào "đợi một goroutine" từ bên ngoài. Muốn đợi thì phải tự dựng cơ chế.

Đó là lý do tồn tại của `WaitGroup` và channel.

## 2. `sync.WaitGroup`

```go
var wg sync.WaitGroup
for i := 0; i < n; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		...
	}()
}
wg.Wait()
```

Luật: `Add` **trước** khi `go`, `Done` bằng `defer` ngay dòng đầu goroutine.
Truyền `WaitGroup` vào hàm phải bằng **pointer** — copy nó thì `Wait` treo vĩnh viễn.

Dùng `time.Sleep` để "đợi cho xong" là sai, luôn luôn. Nó chỉ là race chưa nổ.

## 3. Race condition là gì

Hai goroutine cùng chạm một ô nhớ, ít nhất một bên **ghi**, không có gì đồng bộ hoá.

`counter++` trông như một thao tác, thực ra là ba: đọc, cộng, ghi. Hai goroutine xen kẽ
vào giữa ba bước đó thì mất bản cập nhật.

Điều làm race nguy hiểm hơn bug thường: **chạy 100 lần đúng cả 100 vẫn không chứng minh
được gì**. Nó phụ thuộc lịch chạy, đổi máy/đổi tải là đổi kết quả. Đây chính xác là loại
bug mà tư duy "test thủ công vài lần thấy ổn" bó tay — và là lý do `-race` tồn tại.

```powershell
go run -race ./b-race/broken
```

`-race` không đoán. Nó ghi lại thứ tự truy cập bộ nhớ lúc chạy và báo khi thấy hai truy cập
xung đột — kể cả khi lần chạy đó kết quả vẫn đúng. Đổi lại: chậm hơn ~10 lần, tốn RAM.
Nên bật trong CI/test, không bật ở production.

Giới hạn phải biết: `-race` chỉ thấy đường code **thực sự chạy** trong lần đó.
Sạch `-race` nghĩa là "không phát hiện race ở những gì vừa chạy", không phải "không có race".

## 4. `sync.Mutex`

```go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()   // thực tế: defer mu.Unlock() ngay sau Lock
```

Bao đúng vùng dữ liệu chung, không bao rộng hơn cần thiết (bao cả hàm thì hết song song,
coi như viết code tuần tự cho phức tạp).

Quy ước Go: đặt mutex **ngay cạnh** dữ liệu nó bảo vệ, thường là field trong cùng struct.
Nhìn vào là biết nó khoá cái gì.

## 5. Channel

```go
ch := make(chan Result)      // unbuffered
ch := make(chan Result, 10)  // buffered
ch <- v      // gửi
v := <-ch    // nhận
close(ch)
for v := range ch { }   // nhận tới khi channel đóng
```

Triết lý Go: *"Đừng giao tiếp bằng cách chia sẻ bộ nhớ; hãy chia sẻ bộ nhớ bằng cách giao tiếp."*

Unbuffered: gửi **chặn** cho tới khi có bên nhận — đó vừa là truyền dữ liệu vừa là đồng bộ hoá.
Buffered: gửi không chặn khi buffer còn chỗ.

Luật đóng channel: **bên gửi đóng, bên nhận không bao giờ đóng.** Gửi vào channel đã đóng
là panic. Nhận từ channel đã đóng thì trả zero value ngay (dùng `v, ok := <-ch` để phân biệt).

Mẫu chuẩn khi có N goroutine cùng gửi:

```go
go func() { wg.Wait(); close(ch) }()
for v := range ch { ... }
```

Deadlock hay gặp: gửi vào unbuffered channel mà chưa ai nhận, hoặc `range` một channel
không bao giờ được đóng. Go phát hiện được trường hợp toàn bộ chương trình kẹt và panic
`all goroutines are asleep - deadlock!` — nhưng chỉ khi *toàn bộ* kẹt.

## 6. `select`

```go
select {
case v := <-ch:
	...
case <-time.After(2 * time.Second):
	// timeout
}
```

Chờ nhiều channel, cái nào sẵn trước chạy trước; nhiều cái cùng sẵn thì chọn ngẫu nhiên.
Có `default` thì thành non-blocking.

Đây là cách dựng timeout — và ở Phase 2 chính là cách orchestrator không treo vĩnh viễn
khi Python scoring service chết. Nối thẳng vào "degrade gracefully" trong `README.md`.

## 7. Channel hay Mutex?

| Dùng | Khi |
|---|---|
| channel | **truyền quyền sở hữu** dữ liệu từ goroutine này sang goroutine khác |
| mutex | nhiều goroutine cùng đọc/ghi một trạng thái **dùng chung tại chỗ** (counter, cache) |

Bài tập phase này cố tình là loại thứ hai, để thấy channel không thay thế được mutex
trong mọi trường hợp.

---

## Checklist review (trợ lý dùng khi chấm)

1. Không chạy bản `broken` với `-race`, hoặc không dán được output `DATA RACE` — bắt làm lại.
   Cả phase xoay quanh việc **tự mắt thấy**.
2. Dùng `time.Sleep` thay `WaitGroup`.
3. Quên `wg.Add(1)` trước `go`, hoặc `Add` bên trong goroutine.
4. Truyền `WaitGroup` bằng value.
5. `Lock` bao cả hàm → mất hết song song.
6. Quên `Unlock` ở một nhánh return (đây là lý do phải dùng `defer`).
7. Bản `broken` chạy ra kết quả đúng vì số vòng lặp quá nhỏ → tăng lên.
8. Kết luận "chạy thường không lỗi nên không có race" — hiểu nhầm nền tảng, chặn ngay.
9. Đóng channel ở phía nhận (nếu bài có channel).

## Hỏi lại để xác nhận hiểu

- "`-race` chạy sạch. Có thể kết luận code không có race không?" → phải trả lời **không**,
  kèm lý do ở mục 3.
- "Khi nào dùng channel, khi nào dùng mutex?" → theo bảng mục 7.

---

**Qua phase khi:** cả hai bản chạy đúng như mô tả + review sạch + trả lời được 2 câu trên.
→ Mở Phase 0C: Python fundamentals. Đổi hẳn ngôn ngữ, đổi hẳn tư duy.
