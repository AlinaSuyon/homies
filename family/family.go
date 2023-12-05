package family

type FamilyMember struct {
	Name   string
	Age    int
	Role   string
	Weight float64
}

func NewFamilyMember(name string, age int, role string, weight float64) FamilyMember {
	return FamilyMember{Name: name, Age: age, Role: role, Weight: weight}
}

func FamilyMembers() []FamilyMember {
	members := []FamilyMember{
		NewFamilyMember("Никита", 25, "Папа", 98.4),
		NewFamilyMember("Алина", 26, "Мама", 55.5),
		NewFamilyMember("Моня", 0, "Сладкая булочка", 2.5),
	}
	return members
}
