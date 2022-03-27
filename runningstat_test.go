package runningstat

import (
	"math"
	"testing"
)

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

// Tests for a running stat starting empty, then getting 3 5's pushed to it
// We would expect mean*count == 15, deviation = 0, variance = 0
func InitRunningStatA() RunningStat {
	rs := RunningStat{}
	rs.Push(5)
	rs.Push(5)
	rs.Push(5)

	return rs
}

func TestRunningStatA_Mean(t *testing.T) {
	rs := InitRunningStatA()

	want := float64(15)
	if got := float64(rs.RunningStatCount()) * rs.RunningStatMean(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatA_Deviation(t *testing.T) {
	rs := InitRunningStatA()

	want := float64(0)
	if got := rs.RunningStatDeviation(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatA_Variance(t *testing.T) {
	rs := InitRunningStatA()

	want := float64(0)
	if got := rs.RunningStatVariance(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

// Tests for a running stat starting empty, then getting a 1, 2, and 6 pushed to it
// We would expect mean*count == 9, deviation (rounded) to be 2.646, variance = 7
func InitRunningStatB() RunningStat {
	rs := RunningStat{}
	rs.Push(1)
	rs.Push(2)
	rs.Push(6)

	return rs
}

func TestRunningStatB_Mean(t *testing.T) {
	rs := InitRunningStatB()

	want := float64(9)
	if got := float64(rs.RunningStatCount()) * rs.RunningStatMean(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatB_Deviation(t *testing.T) {
	rs := InitRunningStatB()

	want := float64(2.646)
	if got := roundFloat(rs.RunningStatDeviation(),3); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatB_Variance(t *testing.T) {
	rs := InitRunningStatB()

	want := float64(7)
	if got := rs.RunningStatVariance(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

// Test that we can reset 
func TestRunningStat_ResetCount(t *testing.T) {
	rs := InitRunningStatB()

	rs.Reset()

	want := int64(0)
	if got := rs.RunningStatCount(); got != want {
		t.Errorf("Found %d, wanted %d", got, want)
	}
}

func TestRunningStat_ResetMean(t *testing.T) {
	rs := InitRunningStatB()

	rs.Reset()

	want := float64(0)
	if got := rs.RunningStatMean(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStat_ResetDeviation(t *testing.T) {
	rs := InitRunningStatB()

	rs.Reset()

	want := float64(0)
	if got := rs.RunningStatDeviation(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStat_ResetVariance(t *testing.T) {
	rs := InitRunningStatB()

	rs.Reset()

	want := float64(0)
	if got := rs.RunningStatVariance(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

// Tests for a running stat init of mean=6,count=3,deviation=1,variance=1
func InitRunningStatC() RunningStat {
	rs := RunningStat{}
	rs.Init(3,6,1)

	return rs
}

func TestRunningStatC_Mean(t *testing.T) {
	rs := InitRunningStatC()

	want := float64(18)
	if got := float64(rs.RunningStatCount()) * rs.RunningStatMean(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatC_Deviation(t *testing.T) {
	rs := InitRunningStatC()

	want := float64(1)
	if got := roundFloat(rs.RunningStatDeviation(),3); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatC_Variance(t *testing.T) {
	rs := InitRunningStatC()

	want := float64(1)
	if got := rs.RunningStatVariance(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}


// Tests for a running stat inited with mean=6,count=3,deviation=1,variance=1, then two more 7's added
// expect mean=6.4,count=5,deviation (rounded)=0.894,variance=0.8
func InitRunningStatC_Modified() RunningStat {
	rs := RunningStat{}
	rs.Init(3,6,1)
	rs.Push(7)
	rs.Push(7)

	return rs
}

func TestRunningStatC_ModifiedMean(t *testing.T) {
	rs := InitRunningStatC_Modified()

	want := float64(32)
	if got := float64(rs.RunningStatCount()) * rs.RunningStatMean(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatC_ModifiedDeviation(t *testing.T) {
	rs := InitRunningStatC_Modified()

	want := float64(0.894)
	if got := roundFloat(rs.RunningStatDeviation(),3); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatC_ModifiedVariance(t *testing.T) {
	rs := InitRunningStatC_Modified()

	want := float64(0.8)
	if got := roundFloat(rs.RunningStatVariance(),3); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

// Test merging two different RunningStats
func TestRunningStatAB_Mean(t *testing.T) {
	rs := InitRunningStatA()
	rsB := InitRunningStatB()
	rs.Merge(rsB)

	want := float64(24)
	if got := float64(rs.RunningStatCount()) * rs.RunningStatMean(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatAB_Deviation(t *testing.T) {
	rs := InitRunningStatA()
	rsB := InitRunningStatB()
	rs.Merge(rsB)

	want := float64(2)
	if got := roundFloat(rs.RunningStatDeviation(),3); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}

func TestRunningStatAB_Variance(t *testing.T) {
	rs := InitRunningStatA()
	rsB := InitRunningStatB()
	rs.Merge(rsB)

	want := float64(4)
	if got := rs.RunningStatVariance(); got != want {
		t.Errorf("Found %f, wanted %f", got, want)
	}
}
