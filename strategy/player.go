package strategy

type Player struct {
	Name      string
	Strategy  IStrategy
	WinCount  int
	LoseCount int
	GameCount int
}

func (p *Player) nextHand() *Hand {
	return p.Strategy.nextHand()
}

func (p *Player) win() {
	p.Strategy.study(true)
	p.WinCount++
	p.GameCount++
}

func (p *Player) lose() {
	p.Strategy.study(false)
	p.LoseCount++
	p.GameCount++
}

func (p *Player) even() {
	p.GameCount++
}
