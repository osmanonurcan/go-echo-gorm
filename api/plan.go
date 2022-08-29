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

func GetPlans(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()
	plans := []model.Plan{}
	student := model.Student{}

	//READ ID PARAMETER
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//GET STUDENT FROM DB TO STUDENT VARIABLE
	db.Find(&student, id)

	//log.Print(student)

	//err = db.Model(&student).Association("Plan").Find(&plans).Error
	//log.Print(err)

	db.Where("student_id = ?", id).Find(&plans)

	//log.Print(plans)

	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, plans)
}

func AddPlan(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()

	plans := []model.Plan{}
	plan := model.Plan{}

	defer c.Request().Body.Close()

	//READ REQUEST BODY
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "Failed reading the request body")
	}

	//DECODE JSON BODY TO PLAN VARIABLE
	err = json.Unmarshal(b, &plan)
	if err != nil {
		log.Printf("Failed unmarshaling in plan: %s", err)
		return c.String(http.StatusInternalServerError, "Failed unmarshaling in plan")
	}

	//READ ID PARAMETER
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	plan.StudentID = uint(id)

	//log.Print(plan)

	//CHECK TIME CONFLICT IN PLANS
	db.Where("(start_time BETWEEN ? AND ? ) OR (finish_time BETWEEN ? AND ?) OR (start_time<? AND finish_time>?) AND student_id = ?", plan.StartTime, plan.FinishTime, plan.StartTime, plan.FinishTime, plan.StartTime, plan.FinishTime, id).Find(&plans)
	if len(plans) > 0 {
		log.Printf("There is a time conflict in the plans")
		return c.String(http.StatusInternalServerError, "There is a time conflict in the plans")
	}

	//SAVE TO DB
	db.Create(&plan)
	return c.String(http.StatusOK, "we got your plan")

}

func UpdatePlan(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()

	plan := model.Plan{}
	plans := []model.Plan{}

	defer c.Request().Body.Close()

	//READ REQUEST BODY
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//DECODE JSON BODY TO PLAN VARIABLE
	err = json.Unmarshal(b, &plan)
	if err != nil {
		log.Printf("Failed unmarshaling in user: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//READ ID PARAMETER
	student_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	plan_id, err := strconv.Atoi(c.Param("plan_id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	plan_db := model.Plan{}

	//GET THE FIRST PLAN FROM DB WHERE ID=ID TO PLAN_DB VARIABLE
	db.First(&plan_db, plan_id)

	//CHANGE ATTIRIBUTES IF IT IS NOT EMPTY
	if plan.Name != "" {
		plan_db.Name = plan.Name
	}
	if !plan.StartTime.IsZero() {
		plan_db.StartTime = plan.StartTime
	}
	if !plan.FinishTime.IsZero() {
		plan_db.FinishTime = plan.FinishTime
	}
	if plan.State != "" {
		plan_db.State = plan.State
	}

	// log.Print(id)
	log.Print(plan)
	log.Print("-----------------")

	//CHECK TIME CONFLICT IN PLANS
	db.Where("(start_time BETWEEN ? AND ? ) OR (finish_time BETWEEN ? AND ?) OR (start_time<? AND finish_time>?) AND student_id = ?", plan.StartTime, plan.FinishTime, plan.StartTime, plan.FinishTime, plan.StartTime, plan.FinishTime, student_id).Not("id = ?", plan_id).Find(&plans)
	log.Print(plans)
	if len(plans) > 0 {
		log.Printf("There is a time conflict in the plans")
		return c.String(http.StatusInternalServerError, "There is a time conflict in the plans")
	}

	//SAVE TO DB
	db.Save(&plan_db)
	return c.String(http.StatusOK, "")
}

func DeletePlan(c echo.Context) error {
	//ACCESS TO DB
	db := db.DbManager()
	plan_db := model.Plan{}

	plan_id, err := strconv.Atoi(c.Param("plan_id"))
	if err != nil {
		log.Printf("Failed reading the request param: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	//GET THE FIRST PLAN FROM DB WHERE ID=ID TO PLAN_DB VARIABLE
	db.First(&plan_db, plan_id)

	//DELETE THE PLAN WHERE ID=ID
	db.Delete(&plan_db)
	return c.String(http.StatusOK, "")
}
