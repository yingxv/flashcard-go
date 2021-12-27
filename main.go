package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
	"github.com/yingxv/flashcard-go/src/app"
	"github.com/yingxv/flashcard-go/src/db"
	"github.com/yingxv/flashcard-go/src/middleware"
	"github.com/yingxv/flashcard-go/src/util"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		addr   = flag.String("l", ":8030", "绑定Host地址")
		dbinit = flag.Bool("i", false, "init database flag")
		mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
		mdb    = flag.String("db", "to-do-list", "database name")
		ucHost = flag.String("uc", "http://locahost:8020", "user center host")
		r      = flag.String("r", "localhost:6379", "rdb addr")
	)
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mongoClient := db.NewMongoClient()
	err := mongoClient.Open(*mongo, *mdb, *dbinit)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     *r,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	validate := util.NewValidator()
	trans := util.NewValidatorTranslator(validate)

	app := app.New(validate, trans, ucHost, mongoClient, rdb)
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	//task ctrl
	router.POST("/record/create", app.RecordCreate)
	router.DELETE("/record/remove/:id", app.RecordRemove)
	router.PATCH("/record/update", app.RecordUpdate)
	router.GET("/record/list", app.RecordList)
	router.PATCH("/record/review", app.RecordReview)
	router.GET("/record/review-all", app.RecordReviewAll)
	router.PATCH("/record/random-review", app.RecordRandomReview)
	router.PATCH("/record/set-review-result", app.RecordSetReviewResult)

	srv := &http.Server{Handler: app.IsLogin(middleware.CORS(router)), ErrorLog: nil}
	srv.Addr = *addr

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("server on http port", srv.Addr)

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	cleanup := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range signalChan {
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			go func() {
				_ = srv.Shutdown(ctx)
				cleanup <- true
			}()
			<-cleanup
			mongoClient.Close()
			rdb.Close()
			fmt.Println("safe exit")
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}
