// Calibrate — Go control plane (orchestrator)
//
// Vai trò: điều phối mutation harness, gọi Python scoring service qua gRPC,
// tổng hợp kết quả thành Mutation Kill Rate, chạy như CI gate.
//
// File này là SKELETON — chỉ dựng khung để bắt đầu Phase 2 sau khi hoàn
// thành ONBOARDING.md. Từng phần TODO tương ứng đúng 1 checkpoint đã luyện.
package main

import (
	"fmt"
	"sync"
)

// MutationScenario đại diện cho 1 kịch bản làm suy giảm chất lượng AI
// có kiểm soát (downgrade model, cắt context, prompt injection...).
// TODO(week7-checkpoint): điền 5 kịch bản đã viết ra giấy ở tuần 7.
type MutationScenario struct {
	Name        string
	Description string
	// Apply sẽ gọi vào SUT thật (AI feature) theo cách đã mutate.
	// TODO(phase2): implement gọi SUT thật thay vì hàm giả lập bên dưới.
	Apply func() (output string, err error)
}

// MutationResult là kết quả sau khi oracle chấm 1 kịch bản mutation.
type MutationResult struct {
	Scenario MutationScenario
	Detected bool // oracle có "giết" (phát hiện) được mutation này không
	Err      error
}

// RunMutationHarness chạy tất cả kịch bản song song bằng goroutine,
// đây chính là khung xương từ checkpoint tuần 2 (chạy N task giả song song).
//
// TODO(phase2): thay hàm giả lập scoreScenario bằng gọi gRPC thật sang
// Python scoring service (ComputeSimilarity + CheckControlLimit).
func RunMutationHarness(scenarios []MutationScenario) []MutationResult {
	var wg sync.WaitGroup
	results := make([]MutationResult, len(scenarios))

	for i, sc := range scenarios {
		wg.Add(1)
		go func(idx int, scenario MutationScenario) {
			defer wg.Done()
			results[idx] = scoreScenario(scenario)
		}(i, sc)
	}

	wg.Wait()
	return results
}

// scoreScenario là hàm GIẢ LẬP — placeholder cho việc gọi gRPC scoring
// service thật. Giữ tách riêng để dễ thay thế ở Phase 2 mà không đụng
// vào logic điều phối song song ở trên.
func scoreScenario(sc MutationScenario) MutationResult {
	output, err := sc.Apply()
	if err != nil {
		return MutationResult{Scenario: sc, Err: err}
	}

	// TODO(phase2): thay bằng gọi gRPC ScoringService.ComputeSimilarity
	// rồi ScoringService.CheckControlLimit để quyết định Detected.
	_ = output
	return MutationResult{Scenario: sc, Detected: false}
}

// ComputeKillRate tính Mutation Kill Rate — con số trung tâm của dự án.
func ComputeKillRate(results []MutationResult) float64 {
	if len(results) == 0 {
		return 0
	}
	killed := 0
	for _, r := range results {
		if r.Detected {
			killed++
		}
	}
	return float64(killed) / float64(len(results))
}

func main() {
	// TODO(week7-checkpoint): thay bằng 5 kịch bản mutation thật đã thiết kế.
	placeholderScenarios := []MutationScenario{
		{Name: "model-downgrade", Description: "TODO", Apply: func() (string, error) { return "todo", nil }},
		{Name: "context-truncation", Description: "TODO", Apply: func() (string, error) { return "todo", nil }},
	}

	results := RunMutationHarness(placeholderScenarios)
	killRate := ComputeKillRate(results)

	fmt.Printf("Calibrate — Mutation Kill Rate: %.2f%% (%d scenarios)\n", killRate*100, len(results))
	fmt.Println("⚠️  Đây là skeleton — cần hoàn thành ONBOARDING.md và nối gRPC thật trước khi số liệu này có ý nghĩa.")
}
