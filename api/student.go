package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/osmanonurcan/go-test/db"
	"github.com/osmanonurcan/go-test/model"

	"encoding/json"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()
	students := []model.Student{}

	//GET STUDENTS FROM DB TO STUDENTS VARIABLE
	db.Find(&students)

	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, students)
}

func GetStudent(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()
	student := model.Student{}

	//READ ID PARAMETER
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//GET A STUDENT FROM DB TO STUDENT VARIABLE
	db.Find(&student, id)

	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, student)
}

func AddStudent(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()

	student := model.Student{}

	defer c.Request().Body.Close()

	//READ REQUEST BODY
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "Failed reading the request body")
	}

	//DECODE JSON BODY TO STUDENT VARIABLE
	err = json.Unmarshal(b, &student)
	if err != nil {
		log.Printf("Failed unmarshaling in plan: %s", err)
		return c.String(http.StatusInternalServerError, "Failed unmarshaling in plan")
	}

	//log.Print(plan)

	//SAVE TO DB
	db.Create(&student)
	return c.String(http.StatusOK, "we got your plan")

}

func UpdateStudent(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()

	student := model.Student{}

	defer c.Request().Body.Close()

	//READ REQUEST BODY
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//DECODE JSON BODY TO PLAN VARIABLE
	err = json.Unmarshal(b, &student)
	if err != nil {
		log.Printf("Failed unmarshaling in user: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//READ ID PARAMETER
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	student_db := model.Student{}

	//GET THE FIRST PLAN FROM DB WHERE ID=ID TO PLAN_DB VARIABLE
	db.First(&student_db, id)

	//CHANGE ATTIRIBUTES IF IT IS NOT EMPTY
	if student.Name != "" {
		student_db.Name = student.Name
	}
	if student.Surname != "" {
		student_db.Surname = student.Surname
	}
	if student.Mail != "" {
		student_db.Mail = student.Mail
	}

	// log.Print(id)
	log.Print(student)
	log.Print("-----------------")

	//SAVE TO DB
	db.Save(&student_db)
	return c.String(http.StatusOK, "")
}

func DeleteStudent(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()
	student_db := model.Student{}

	//READ ID PARAMETER
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//GET THE FIRST PLAN FROM DB WHERE ID=ID TO PLAN_DB VARIABLE
	db.First(&student_db, id)

	//DELETE THE PLAN WHERE ID=ID
	db.Delete(&student_db)
	return c.String(http.StatusOK, "")
}
