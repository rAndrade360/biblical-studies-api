package http

import (
	"database/sql"

	altcontroller "github.com/rAndrade360/biblical-studies-api/api/handlers/http/alternative"
	qcontroller "github.com/rAndrade360/biblical-studies-api/api/handlers/http/question"
	qgcontroller "github.com/rAndrade360/biblical-studies-api/api/handlers/http/questiongroup"
	alrepository "github.com/rAndrade360/biblical-studies-api/api/internal/repositories/alternative"
	qrepository "github.com/rAndrade360/biblical-studies-api/api/internal/repositories/question"
	qgrepository "github.com/rAndrade360/biblical-studies-api/api/internal/repositories/questiongroup"
	altservice "github.com/rAndrade360/biblical-studies-api/api/services/alternative"
	qservice "github.com/rAndrade360/biblical-studies-api/api/services/question"
	qgservice "github.com/rAndrade360/biblical-studies-api/api/services/questiongroup"
)

type HttpControllers struct {
	QuestionGroupController qgcontroller.QuestionGroupController
	QuestionController      qcontroller.QuestionController
	AlternativeController   altcontroller.AlternativeController
}

func Load(db *sql.DB) HttpControllers {
	qgrepo := qgrepository.NewQuestionGroupRepository(db)
	qrepo := qrepository.NewQuestionRepository(db)
	altrepo := alrepository.NewAlternativeRepository(db)

	qgsvc := qgservice.NewQuestionGroupService(qgrepo)
	qsvc := qservice.NewQuestionService(qrepo, qgsvc)
	altsvc := altservice.NewAlternativeService(altrepo, qsvc)

	qgctrl := qgcontroller.NewQuestionGroupController(qgsvc)
	qctrl := qcontroller.NewQuestionController(qsvc)
	altctrl := altcontroller.NewAlternativeController(altsvc)

	return HttpControllers{
		QuestionGroupController: qgctrl,
		QuestionController:      qctrl,
		AlternativeController:   altctrl,
	}
}
