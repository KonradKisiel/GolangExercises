package main

import (
	"log"
	"os"
	"text/template"
)

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

	Menu := Menu{
		Breakfast: Dishes{fried_egg, chicken_salad, toast_chesse_cham, waffless_jam},
		Lunch:     Dishes{pizza, spaghetti_carbonara},
		Dinner:    Dishes{fruits_salad, toast_chesse_cham, waffless_jam},
	}

	err := tpl.Execute(os.Stdout, Menu)
	if err != nil {
		log.Fatalln(err)
	}
}
