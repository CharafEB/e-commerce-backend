package main

import (
	"context"
	"database/sql"
	"fmt"
	controller "github/think.com/Controller"
	"github/think.com/Controller/admin"
	middlewares "github/think.com/Middlewares"
	router "github/think.com/Router"
	"github/think.com/model"
	"log"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dbinfo := os.Getenv("POSTGERS_API_LINE")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("There is an err in the conaction ")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("err type: %T\n", err)
		log.Printf("err is: %v\n", err)
		log.Fatal("There is an err in Ping")
	}
	log.Println("Database connection successful")

	add := ":" + os.Getenv("PORT")
	st_ore := model.NewStore(db)
	log.Println("Store has been opened")

	if _, err := os.Stat("Searchindex.bleve"); err == nil {
		fmt.Printf("We have an index file\n")
	} else {
		fmt.Printf("Creating index\n")
		if err := st_ore.Indexer.IndexArticles(context.Background()); err != nil {
			log.Fatalf("Failed to index articles: %v", err)
		}
	}

	bleveIndex, err := bleve.Open(os.Getenv("INDEX_PATH"))
	if err != nil {
		log.Fatalf("failed to open index: %v", err)
	}
	log.Println("Index has been opened")

	app := middlewares.Application{
		Address:          add,
		Storge:           st_ore,
		BleveSearchIndex: bleveIndex,
	}

	ctrlApp := &controller.Application{
		Application: app,
	}

	ctrlAdmin := &admin.Application{
		Application: app,
	}

	cntrolObj := &router.Control{
		Controller:      ctrlApp,
		AdminController: ctrlAdmin,
	}
	appRouter := &router.Application{
		Application: app,
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           86400,
	})

	appRouter.CORSMiddleware = c

	mux := cntrolObj.Moul()

	log.Printf("Server starting with CORS enabled on %s", add)

	if err := appRouter.Run(mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
