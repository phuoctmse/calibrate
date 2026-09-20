"""
Calibrate — Python scoring service (SKELETON)

Vai trò: 1 service nhỏ, chỉ làm 2 việc:
  1. Tính cosine similarity giữa 2 đoạn văn bản (embedding).
  2. Kiểm tra 1 điểm số có vượt control limit (SPC) so với baseline không.

Không biết gì về mutation harness, CI gate, hay Go orchestrator —
chỉ nhận request qua gRPC, trả kết quả, đúng tinh thần tách theo
"ngôn ngữ nào có hệ sinh thái phù hợp hơn cho phần việc này".

TODO(week4-checkpoint): thay hàm compute_similarity giả lập bên dưới
bằng bản dùng sentence-transformers thật (đã luyện ở checkpoint tuần 4).

TODO(week6-checkpoint): thay hàm check_control_limit giả lập bằng
bản dùng control chart thật (đã luyện ở checkpoint tuần 6).

TODO(week5-checkpoint): sau khi generate code từ proto/scoring.proto,
implement class Servicer kế thừa đúng ScoringServiceServicer generated,
thay vì hàm thuần bên dưới.
"""

import numpy as np


def compute_similarity(reference_text: str, candidate_text: str) -> float:
    """
    GIẢ LẬP — placeholder. Bản thật (Phase 2) sẽ:
      1. Load model sentence-transformers 1 lần khi service khởi động
         (không load lại mỗi request — quan trọng cho performance).
      2. Encode cả 2 đoạn văn bản thành vector.
      3. Tính cosine similarity giữa 2 vector.
      4. Trả về model_version để đảm bảo reproducibility khi so sánh
         kết quả giữa các lần chạy mutation testing.
    """
    # Placeholder: trả về giá trị giả để pipeline chạy end-to-end được
    # trong lúc chưa hoàn thành onboarding tuần 4.
    return 0.0


def check_control_limit(
    baseline_scores: list[float],
    new_score: float,
    sigma_multiplier: float = 3.0,
) -> dict:
    """
    GIẢ LẬP — placeholder. Bản thật (Phase 2) sẽ dùng công thức X-bar chart:
      mean = trung bình baseline_scores
      std  = độ lệch chuẩn baseline_scores
      upper_limit = mean + sigma_multiplier * std
      lower_limit = mean - sigma_multiplier * std
      out_of_control = new_score nằm ngoài [lower_limit, upper_limit]

    Đây chính là bài tập checkpoint tuần 6 — đừng copy code mẫu nào khác,
    tự viết lại công thức này để chắc chắn hiểu bản chất trước khi dùng
    trong dự án thật (vì đây là con số quyết định pass/fail của CI gate).
    """
    if not baseline_scores:
        return {
            "out_of_control": False,
            "upper_limit": 0.0,
            "lower_limit": 0.0,
            "z_score": 0.0,
        }

    mean = float(np.mean(baseline_scores))
    std = float(np.std(baseline_scores))
    upper = mean + sigma_multiplier * std
    lower = mean - sigma_multiplier * std
    z = (new_score - mean) / std if std > 0 else 0.0

    return {
        "out_of_control": new_score > upper or new_score < lower,
        "upper_limit": upper,
        "lower_limit": lower,
        "z_score": z,
    }


if __name__ == "__main__":
    print("Calibrate scoring service — SKELETON")
    print("Chưa khởi động gRPC server thật. Hoàn thành ONBOARDING.md tuần 4-5 trước.")
    print(check_control_limit([0.8, 0.82, 0.79, 0.81], 0.55))
