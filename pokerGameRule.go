package util

type Poker struct {
	originList []int32 //基础牌型
	t          int32   //比大小牌型
	subT       int32   //如果pokerType相等，用这个再对比大小
}

func NewPoker(originList []int32, Type int32, subType int32) *Poker {
	return &Poker{originList: originList, t: Type, subT: subType}
}

func (p *Poker) Fight(poker *Poker) bool {
	if p.t == poker.t {
		return p.subT > poker.subT
	}
	return p.t > poker.t
}

func (p *Poker) GetType() int32 {
	return p.t
}