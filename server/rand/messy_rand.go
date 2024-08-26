package rand

type IRandom interface {
	Get() int64
}

type MessyRandom struct {
	IRandom
	curVal int64
	seed   int64
	count  int64
}

func CreateMessyRandom(seed int64) *MessyRandom {
	return &MessyRandom{
		curVal: seed,
		count:  1,
	}
}

func (m *MessyRandom) Get() int64 {
	defer m.getRandom()

	return m.curVal
}

func (m *MessyRandom) getRandom() {
	//m.curVal = (m.curVal ^ (m.curVal % 153)) - m.curVal + (m.curVal%31)>>m.count%5&m.curVal%m.count
	m.curVal = (m.seed ^ (m.curVal % (m.count + 1)) + 1)
	m.curVal *= (m.count%255)<<(m.count%7) + 47
	m.curVal <<= 2
	m.count++
}
