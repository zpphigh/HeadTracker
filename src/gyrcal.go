package main

const GYRO_CAL_THRESHOLD = 1_000_000

type GyrCal struct {
	offset [3]int32
	sum    [3]int32
	count  [3]int32
}

func (g *GyrCal) apply(x int32, y int32, z int32) (int32, int32, int32) {
	return g.applyAxis(0, x), g.applyAxis(1, y), g.applyAxis(2, z)
}

func (g *GyrCal) applyAxis(i, v int32) int32 {
	tmp := v - g.offset[i]
	if tmp > 0 && tmp > GYRO_CAL_THRESHOLD {
		tmp = 1
	}
	if tmp < 0 && tmp < -GYRO_CAL_THRESHOLD {
		tmp = -1
	}
	g.sum[i] += tmp
	g.count[i]++
	if g.count[i] > 100 {
		g.offset[i] += g.sum[i] / g.count[i] / 2
		g.count[i] = 0
		g.sum[i] = 0
	}
	return v - g.offset[i]
}