package gl

type User struct {
	Margin    int64
	Lposition uint64
	Lprice    uint64
	Sposition uint64
	Sprice    uint64
	Block     uint64
}

var Users map[string]map[string]User
var LatestBlocks map[string]int64

func init() {
	Users = make(map[string]map[string]User)
	LatestBlocks = make(map[string]int64)
}
