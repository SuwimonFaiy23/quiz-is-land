package main

import (
	"github.com/SuwimonFaiy23/quiz-is-land/config"
	"github.com/SuwimonFaiy23/quiz-is-land/src/question"
	"github.com/SuwimonFaiy23/quiz-is-land/src/session"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()            // โหลด config ก่อน
	db := config.ConnectDatabase() // เชื่อม DB

	// session
	sessionRepo := session.NewRepository(db)
	sessionService := session.NewService(sessionRepo)
	sessionHandler := session.NewHandler(sessionService)

	//question
	questionRepo := question.NewRepository(db)
	questionService := question.NewService(questionRepo, sessionRepo)
	questionHandler := question.NewHandler(questionService)

	e := echo.New()
	e.POST("/api/v1/quiz/session", sessionHandler.CreateSession)

	// question
	e.POST("/api/v1/quiz/:session_id", questionHandler.GetQuestion)
	e.POST("/api/v1/quiz/summary/:session_id", questionHandler.GetSummary)

	e.Logger.Fatal(e.Start(":8080"))
}
