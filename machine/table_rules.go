package machine

import "log"

type TableRules interface {
	//检验条件
	Condition( table Table ) bool
	//进行处理
	Action( table Table )
}

//==================================单轮结算====================================
type SettleForRoundRule struct {}
func (this *SettleForRoundRule) Condition( table Table ) bool {
	return false
}
func (this *SettleForRoundRule) Action( table Table ) {

}

//==================================流局====================================
type LiuJuRule struct {}
func (this *LiuJuRule) Condition( table Table ) bool{
	if table.CardsRest.Len() == 0 {
		log.Println("=====rule liuju true=====")
		return true
	} else {
		log.Println("=====rule liuju false=====")
		return false
	}
}
func (this *LiuJuRule) Action( table Table ) {
	log.Println("=====rule liuju action=====")
}