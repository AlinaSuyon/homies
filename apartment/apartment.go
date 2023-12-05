package apartment

import "fmt"

type Kitchen struct {
	Area         float64
	Appliances   []string
	CabinetStyle string
}

type Bedroom struct {
	Area   float64
	Bed    string
	Desks  int
	Chairs int
}

type Bathroom struct {
	Toilet  string
	Bathtub string
	Sink    string
}

type Apartment struct {
	Kitchen  Kitchen
	Bedroom  Bedroom
	Corridor float64
	Bathroom Bathroom
	Area     float64
}

func NewApartment() Apartment {
	return Apartment{
		Kitchen: Kitchen{
			Area:         9.0,
			Appliances:   []string{"Холодильник", "Плита"},
			CabinetStyle: "Кухонный гарнитур белого цвета, серая столешница и мраморный фартук",
		},
		Bedroom: Bedroom{
			Area:   10.0,
			Bed:    "Кровать 160x200",
			Desks:  2,
			Chairs: 2,
		},
		Corridor: 4.0,
		Bathroom: Bathroom{
			Toilet:  "1",
			Bathtub: "1",
			Sink:    "1",
		},
		Area: 60.0, // Общая площадь квартиры
	}
}

func (a *Apartment) PrintDescription() {
	fmt.Println("\nОписание квартиры:")
	fmt.Printf("Кухня: Площадь %.2f кв.м, Техника: %v, Стиль кухонного гарнитура: %s\n", a.Kitchen.Area, a.Kitchen.Appliances, a.Kitchen.CabinetStyle)
	fmt.Printf("Спальня: Площадь %.2f кв.м, Кровать: %s, Столы: %d, Кресла: %d\n", a.Bedroom.Area, a.Bedroom.Bed, a.Bedroom.Desks, a.Bedroom.Chairs)
	fmt.Printf("Коридор: %.2f кв.м\n", a.Corridor)
	fmt.Printf("Ванная: Площадь %.2f кв.м, Унитаз: %s, Ванна: %s, Раковина: %s\n", a.Bathroom, a.Bathroom.Toilet, a.Bathroom.Bathtub, a.Bathroom.Sink)
	fmt.Printf("Общая площадь квартиры: %.2f кв.м\n", a.Area)
}
