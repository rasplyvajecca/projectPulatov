package plants

type Plants struct {
	Plant string
	Count int
}

func NewPlant() []Plants {
	return []Plants{
		{Plant: "Кактусы", Count: 3},
		{Plant: "Пион", Count: 1},
		{Plant: "Фиалки", Count: 2},
		{Plant: "Ромашки", Count: 18},
		{Plant: "Одуванчик", Count: 1},
	}
}
