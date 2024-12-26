package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ehsanmsb/Tati/pkg/config"
	"github.com/ehsanmsb/Tati/pkg/render"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) ContactUs(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl")
}

func (m *Repository) Aboute(w http.ResponseWriter, r *http.Request) {
	result := addValue(345, 789)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is my aboute page and sum two number is %d", result))

}

func addValue(x, y int) int {
	sum := x + y
	return sum
}

func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	var num1 float32 = 4545.7
	var num2 float32 = 0
	divideNumber, err := divideValue(num1, num2)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0, please change y value")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("The result of divide %f by %f is %f", num1, num2, divideNumber))

}

func divideValue(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0, Please change the y value")
		return 0, err
	}
	result := x / y
	return result, nil
}
