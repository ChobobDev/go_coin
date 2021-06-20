package person

type Person struct {
	name        string
	age         int
	nationality string
}

func (p *Person) SetDetails(name string, age int, nationality string) {
	p.name = name
	p.age = age
	p.nationality = nationality
}
