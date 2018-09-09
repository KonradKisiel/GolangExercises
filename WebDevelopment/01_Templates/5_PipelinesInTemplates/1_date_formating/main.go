package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

/*
t := time.Now()
	fmt.Println(t)
	tf := t.Format(time.Kitchen)
	fmt.Println(tf)
*/

func monthDayYear(t time.Time) string {
	/*
	   These are predefined layouts for use in Time.Format and time.Parse. The reference time used in the layouts is the specific time:

	   Mon Jan 2 15:04:05 MST 2006

	   which is Unix time 1136239445. Since MST is GMT-0700, the reference time can be thought of as

	   01/02 03:04:05PM '06 -0700
	*/
	// 02-01-2006 -> day month year
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
