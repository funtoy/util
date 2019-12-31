package util

type Poker struct {
	originList []int32 //基础牌型
	t          int32   //比大小牌型
	subT       int32   //如果pokerType相等，用这个再对比大小
}

func NewPoker() *Poker {
	return new(Poker)
}

func (p *Poker) Fight(poker *Poker) bool {
	if p.t == poker.t {
		return p.subT > poker.subT
	}
	return p.t > poker.t
}

func (p *Poker) Reset() {
	p.originList = nil
	p.t = 0
	p.subT = 0
}

func (p *Poker) GetType() int32 {
	return p.t
}

func (p *Poker) SetType(v int32) {
	p.t = v
}

func (p *Poker) GetSubType() int32 {
	return p.subT
}

func (p *Poker) SetSubType(v int32) {
	p.subT = v
}

func (p *Poker) GetList() []int32 {
	return p.originList
}

func (p *Poker) SetList(v []int32) {
	p.originList = v
}

func (p *Poker) AddToList(v []int32) {
	p.originList = append(p.originList, v...)
}
