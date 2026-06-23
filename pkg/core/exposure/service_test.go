package exposure

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateExposureA8(t *testing.T) {
	tests := []struct {
		vibrationMagnitude float64
		triggerTime        int
		expectedA8         float64
	}{
		{0.0, 0, 0.0},
		{100000.0, 0, 0.0},
		{0.0, 10000000, 0.0},
		{5.0, 60, 0.0},       // 60s = 1min, (1/8)=0 in integer division
		{5.0, 479, 0.0},      // 479s = 7min, (7/8)=0 in integer division
		{5.0, 480, 5.0},      // 480s = 8min, (8/8)=1, sqrt(1)=1
		{10.0, 480, 10.0},    // 8min, sqrt(1)=1
		{5.0, 960, 5.0 * math.Sqrt(2.0)}, // 960s = 16min, (16/8)=2, sqrt(2)
	}

	for _, tc := range tests {
		a8 := generateExposureA8(tc.vibrationMagnitude, tc.triggerTime)
		assert.InDelta(t, tc.expectedA8, a8, 1e-9, "vibrationMagnitude=%v, triggerTime=%v", tc.vibrationMagnitude, tc.triggerTime)
	}
}

func TestGenerateExposurePoints(t *testing.T) {
	tests := []struct {
		vibrationMagnitude float64
		triggerTime        int
		expectedPoints     float64
	}{
		{0.0, 0, 0.0},
		{100000.0, 0, 0.0},
		{0.0, 10000000, 0.0},
		{2.5, 480, 100.0},  // (2.5/2.5)^2 * ((8/8)*100) = 1 * 100 = 100
		{5.0, 480, 400.0},  // (5/2.5)^2 * 100 = 4 * 100 = 400
		{5.0, 240, 200.0},  // 4 * ((4/8)*100) = 4 * 50 = 200
		{5.0, 120, 100.0},  // 4 * ((2/8)*100) = 4 * 25 = 100
		{5.0, 60, 50.0},    // 4 * ((1/8)*100) = 4 * 12.5 = 50
	}

	for _, tc := range tests {
		points := generateExposurePoints(tc.vibrationMagnitude, tc.triggerTime)
		assert.Equal(t, tc.expectedPoints, points, "vibrationMagnitude=%v, triggerTime=%v", tc.vibrationMagnitude, tc.triggerTime)
	}
}
