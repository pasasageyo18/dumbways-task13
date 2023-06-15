package main

import (
	"context"
	"fmt"
	"go-latihan1/connection"

	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Title        string
	Content      string
	StartDate    string
	EndDate      string
	Duration     string
	Technologies []string
}

var ProjectsData = []Project{
	// {
	// 	Title:        "Auouo",
	// 	Content:      "kamu nanya? kamu bertanya-tanya?",
	// 	StartDate:    "01-01-2023",
	// 	EndDate:      "20-01-2023",
	// 	Duration:     "2 hours",
	// 	Technologies: []string{"StackIO", "NodeJS", "javascript"},
	// },
	// {
	// 	Title:        "Mamamia Lezatoz",
	// 	Content:      "Chocolatoz mamamia lezatoz",
	// 	StartDate:    "05-01-2023",
	// 	EndDate:      "30-01-2023",
	// 	Duration:     "2 days",
	// 	Technologies: []string{"javascript"},
	// },
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()
	// e.GET("/hehe", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.POST("/edit-project/:id", editProject)
	e.POST("/delete-project/:id", deleteProject)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/form-project", formAddProject)
	e.POST("/add-project", addProject)

	e.Logger.Fatal(e.Start("localhost:7000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")

	fmt.Println(data)

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"projects": ProjectsData,
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contactnew.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var DetailProject = Project{}

	for index, item := range ProjectsData {
		if id == index {
			DetailProject = Project{
				Title:        item.Title,
				Content:      item.Content,
				StartDate:    item.StartDate,
				EndDate:      item.EndDate,
				Duration:     item.Duration,
				Technologies: item.Technologies,
			}
		}
	}

	item := map[string]interface{}{
		"Project": DetailProject,
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), item)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/addpro.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	title := c.FormValue("inputTitle")
	startDate := c.FormValue("inputStartDate")
	endDate := c.FormValue("inputEndDate")
	content := c.FormValue("inputProContent")
	technologies := c.Request().Form["technologies"]

	startDateReal, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return err
	}

	endDateReal, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return err
	}

	duration := endDateReal.Sub(startDateReal)
	durationDays := int(duration.Hours() / 24)
	durationWeeks := int(duration.Hours() / 24 / 7)
	durationMonths := int(duration.Hours() / 24 / 30)
	durationYears := int(duration.Hours() / 24 / 365)

	println("Title : " + title)
	println("startDate : " + startDate)
	println("endDate : " + endDate)
	println("Content : " + content)
	fmt.Println("Technologies: ", strings.Join(technologies, ", "))
	println(durationDays, "days")
	println(durationWeeks, "weeks")
	println(durationMonths, "months")
	println(durationYears, "years")

	var newProject = Project{
		Title:        title,
		Content:      content,
		StartDate:    startDate,
		EndDate:      endDate,
		Duration:     fmt.Sprintf("%d days, %d weeks, %d months, %d years", durationDays, durationWeeks, durationMonths, durationYears),
		Technologies: technologies,
	}

	ProjectsData = append(ProjectsData, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(delete echo.Context) error {
	i, _ := strconv.Atoi(delete.Param("id"))

	fmt.Println("index : ", i)

	ProjectsData = append(ProjectsData[:i], ProjectsData[i+1:]...)

	return delete.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(edit echo.Context) error {
	id, _ := strconv.Atoi(edit.Param("id"))
	fmt.Println("index : ", id)

	ProjectsData = append(ProjectsData[:id], ProjectsData[id+1:]...)
	return edit.Redirect(http.StatusMovedPermanently, "/addProject")
}
