package relatives

type Relatives struct {
	Member string
	Age    int
}

func NewRelative() []Relatives {
	return []Relatives{
		{Member: "Бабушка", Age: 81},
		{Member: "Дедушка", Age: 90},
		{Member: "Тетя", Age: 45},
		{Member: "Кузина", Age: 5},
		{Member: "Кузен", Age: 5},
	}
}
