package main

type GoodsInfoType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Num  int    `json:"num"`
}

func (g *GoodsInfoType) Buy() bool {
	if g.Num < 1 {
		return false
	}
	g.Num -= 1
	return true
}
