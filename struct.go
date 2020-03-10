package rfm

// Stat is struct for keep statistic data
type Stat struct {
	Count int     `json:"count"`
	Mean  float64 `json:"mean"`
	STD   float64 `json:"std"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	P25   float64 `json:"percent_25"`
	P50   float64 `json:"percent_50"`
	P75   float64 `json:"percent_75"`
}

// Point is contain data point in chart
// x and y should contain float64 between 0.00 to 1.00
type Point struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Value float64 `json:"value"`
}

// ClusterStat is struct for keep statistic data
type ClusterStat struct {
	ClusterNum   int
	Stat         *Stat
	Points       []Point `json:"point"`
	PointsCenter *Point  `json:"point_center"`
}
