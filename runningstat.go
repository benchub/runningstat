package runningstat

import (
	"math"
)

type RunningStat struct {
	m_n    int64
	m_oldM float64
	m_newM float64
	m_oldS float64
	m_newS float64
}


// https://www.johndcook.com/blog/standard_deviation/
func (oldRS *RunningStat) Push(x float64) {
	rs := oldRS
	rs.m_n += 1
	if rs.m_n == 1 {
		rs.m_oldM = x
		rs.m_newM = x
		rs.m_oldS = 0
	} else {
		rs.m_newM = rs.m_oldM + (x-rs.m_oldM)/float64(rs.m_n)
		rs.m_newS = rs.m_oldS + (x-rs.m_oldM)*(x-rs.m_newM)

		rs.m_oldM = rs.m_newM
		rs.m_oldS = rs.m_newS
	}
}

func (rs *RunningStat) Reset() {
	rs.m_n = 0
	rs.m_oldM = 0
	rs.m_newM = 0
	rs.m_oldS = 0
	rs.m_newS = 0
}

func (rs *RunningStat) Init(count int64,mean float64,deviation float64) {
	rs.m_n = count
	rs.m_oldM = mean
	rs.m_newM = mean
	rs.m_oldS = 0
	rs.m_newS = deviation

	if count > 1 {
		rs.m_oldS = deviation*deviation*float64(count) - deviation*deviation
		rs.m_newS = rs.m_oldS		
	}
}

func (rs RunningStat) RunningStatCount() int64 {
	return rs.m_n
}

func (rs RunningStat) RunningStatMean() float64 {
	if rs.m_n > 0 {
		return rs.m_newM
	}

	return 0
}

func (rs RunningStat) RunningStatVariance() float64 {
	if rs.m_n > 1 {
		return rs.m_newS / float64(rs.m_n-1)
	}

	return 0
}

func (rs RunningStat) RunningStatDeviation() float64 {

	return math.Sqrt(rs.RunningStatVariance())
}
