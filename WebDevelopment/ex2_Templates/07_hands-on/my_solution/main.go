package main

import (
	"log"
	"os"
	"text/template"
)

type Resaurants []Restaurant

type Restaurant struct {
	Name string
	Menu Menu
}

type Menu struct {
	Breakfast Dishes
	Lunch     Dishes
	Dinner    Dishes
}

type Dishes []Dish

type Dish struct {
	Name        string
	Price       float64
	Ingridients []Ingridient
}

type Ingridient string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	fried_egg := Dish{
		Name:        "Fried egg with bacon",
		Price:       4.99,
		Ingridients: []Ingridient{"two eggs", "bacon", "pepper"},
	}

	pizza := Dish{
		Name:        "Pizza",
		Price:       15.99,
		Ingridients: []Ingridient{"dough", "tomato sauce", "mozzarella", "bacon", "olives", "champignons"},
	}

	spaghetti_carbonara := Dish{
		Name:        "Spaghetti Carbonara",
		Price:       12.99,
		Ingridients: []Ingridient{"spaghetti", "bacon", "eggs", "herbs", "parmezan"},
	}

	toast_chesse_cham := Dish{
		Name:        "Tosts with chesse and cham",
		Price:       4.99,
		Ingridients: []Ingridient{"toast", "mozzarella", "cham"},
	}

	chicken_salad := Dish{
		Name:        "Chicken salad",
		Price:       6.99,
		Ingridients: []Ingridient{"lettuce", "chopped roasted chicken", "coctail tomatos", "olive", "herbs"},
	}

	fruits_salad := Dish{
		Name:        "Fruit salad",
		Price:       6.99,
		Ingridients: []Ingridient{"lettuce", "chopped roasted chicken", "coctail tomatos", "olive", "herbs"},
	}

	waffless_jam := Dish{
		Name:        "waffless with jam",
		Price:       3.99,
		Ingridients: []Ingridient{"waffless", "jam"},
	}

	Menu1 := Menu{
		Breakfast: Dishes{fried_egg, chicken_salad},
		Lunch:     Dishes{spaghetti_carbonara},
		Dinner:    Dishes{toast_chesse_cham, waffless_jam},
	}

	Menu2 := Menu{
		Breakfast: Dishes{toast_chesse_cham, waffless_jam},
		Lunch:     Dishes{pizza},
		Dinner:    Dishes{fruits_salad, toast_chesse_cham},
	}

	SlowFood := Restaurant{
		Name: "SlowFood",
		Menu: Menu1,
	}

	EatNgo := Restaurant{
		Name: "EatNgo",
		Menu: Menu2,
	}

	Resaurants := []Restaurant{
		SlowFood,
		EatNgo,
	}

	err := tpl.Execute(os.Stdout, Resaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
