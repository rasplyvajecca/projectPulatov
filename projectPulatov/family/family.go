package family

type Family struct {
	Member string
	Age    int
}

func NewMember() []Family {
	return []Family{
		{Member: "Мама", Age: 35},
		{Member: "Папа", Age: 30},
		{Member: "Дочка", Age: 13},
		{Member: "Сын", Age: 4},
		{Member: "Сын", Age: 1},
	}
}
