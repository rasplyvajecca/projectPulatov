package pets

type Pets struct {
	Pet   string
	Count int
}

func NewPet() []Pets {
	return []Pets{
		{Pet: "Кошки", Count: 4},
		{Pet: "Собаки", Count: 2},
		{Pet: "Рыбки", Count: 20},
		{Pet: "Хомяк", Count: 1},
		{Pet: "Мыши", Count: 3},
	}
}
