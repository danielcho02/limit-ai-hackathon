package presenter

import (
	"fmt"
	"time"
)

func formatTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	// 7일(168시간) 이상 차이날 경우: 날짜 포맷 (YYYY-MM-DD)
	if diff >= 7*24*time.Hour {
		return t.Format("2006-01-02")
	}

	// 1일 이상 차이날 경우: n일 전
	if diff >= 24*time.Hour {
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%d일 전", days)
	}

	// 1시간 이상 차이날 경우: n시간 전
	if diff >= time.Hour {
		hours := int(diff.Hours())
		return fmt.Sprintf("%d시간 전", hours)
	}

	// 1시간 미만일 경우 : n분 전 또는 방금 전
	minutes := int(diff.Minutes())
	if minutes > 0 {
		return fmt.Sprintf("%d분 전", minutes)
	}
	return "방금 전"
}