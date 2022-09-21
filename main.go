package main

import (
	"context"
	"fmt"
	"idnmedia/configs"
	"idnmedia/infrastructure"
	"idnmedia/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	missionRepository "idnmedia/repositories/mission"
	playerRepository "idnmedia/repositories/player"
	playerMissionRepository "idnmedia/repositories/player_mission"

	authUsecase "idnmedia/usecases/auth"
	missionUsecase "idnmedia/usecases/mission"
	playerMissionUsecase "idnmedia/usecases/player_mission"

	authController "idnmedia/controllers/auth"
	"idnmedia/controllers/middlewares"
	missionController "idnmedia/controllers/mission"
	playerMissionController "idnmedia/controllers/player_mission"
)

func main() {
	// setup config
	conf := configs.New()
	ctx := context.Background()

	// setup db
	postgresDB := infrastructure.SetupPostgresDB(conf.Postgres)

	// setup middleware
	jwt := middlewares.SetupJWT(conf.JWT)

	// setup repositories
	playerRepo := playerRepository.NewPostgreRepository(postgresDB)
	missionRepo := missionRepository.NewPostgreRepository(postgresDB)
	playerMissionRepo := playerMissionRepository.NewPostgreRepository(postgresDB)
	dbTx := repositories.NewDBTx(postgresDB)

	// setup usecases
	authUC := authUsecase.NewUsecase(jwt, playerRepo)
	missionUC := missionUsecase.NewUsecase(missionRepo)
	playerMissionUC := playerMissionUsecase.NewUsecase(missionRepo, playerMissionRepo, playerRepo, dbTx)

	// setup controllers
	authCtrl := authController.NewHttpController(authUC)
	missionCtrl := missionController.NewHttpController(missionUC)
	playerMissionCtrl := playerMissionController.NewHttpController(playerMissionUC)

	// Setup http server
	router := mux.NewRouter()
	router.Use(mux.MiddlewareFunc(middlewares.MethodLogger()))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	v1 := router.PathPrefix("/v1").Subrouter()
	// auth controller
	apiAuth := v1.PathPrefix("/auth").Subrouter()
	apiAuth.HandleFunc("/login", authCtrl.Login).Methods(http.MethodPost)
	apiAuth.HandleFunc("/profile", middlewares.Adapt(http.
		HandlerFunc(authCtrl.Profile),
		jwt.MuxMiddleware,
	).ServeHTTP)
	// profile := apiAuth.PathPrefix("/profile").Subrouter()
	// profile.Use(jwt.MuxMiddleware)
	// profile.HandleFunc("", authCtrl.Profile).Methods(http.MethodGet)

	// mission controller
	apiMission := v1.PathPrefix("/mission").Subrouter()
	apiMission.Use(jwt.MuxMiddleware)
	apiMission.HandleFunc("", missionCtrl.Create).Methods(http.MethodPost)
	apiMission.HandleFunc("", missionCtrl.FindAllPagination).Methods(http.MethodGet)

	// player mission controller
	apiPlayerMission := v1.PathPrefix("/playerMission").Subrouter()
	apiPlayerMission.Use(jwt.MuxMiddleware)
	apiPlayerMission.HandleFunc("", playerMissionCtrl.FindAll).Methods(http.MethodGet)
	apiPlayerMission.HandleFunc("/missionId/{missionId}", playerMissionCtrl.Assign).Methods(http.MethodPost)
	apiPlayerMission.HandleFunc("/missionId/{missionId}/progess", playerMissionCtrl.Progess).Methods(http.MethodPost)
	apiPlayerMission.HandleFunc("/missionId/{missionId}/complete", playerMissionCtrl.Complete).Methods(http.MethodPost)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.App.Port),
		Handler: router,
	}
	log.Printf("Starting server on port %d \n", conf.App.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(ctx, "%v", err)
	}
}
