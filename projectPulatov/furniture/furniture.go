package furniture

type Furniture struct {
	Name  string
	Count int
	Size  float64
}

func NewObject() []Furniture {
	return []Furniture{
		{Name: "Ковер", Count: 1, Size: 3.59},
		{Name: "Шкафы", Count: 2, Size: 2.99},
		{Name: "Столы", Count: 3, Size: 1.45},
		{Name: "Кровати", Count: 5, Size: 20.4},
		{Name: "Тумба", Count: 1, Size: 1},
	}
}
