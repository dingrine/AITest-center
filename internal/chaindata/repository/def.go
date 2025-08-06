package repository

var RegionCertTypeMap = map[int]string{0: "MAT-CN", 1: "MAT-HK", 2: "MAT-US"}
var EvaSystemMap = map[int]string{0: "ProTest", 1: "TradingTest"}
var EnvSystem = []int{0, 1}

type DataRepository interface {
}
