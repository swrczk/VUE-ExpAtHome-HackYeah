package main

import (
	"log"
	"context"
	"net/http"

	"github.com/apuigsech/rest-layer-sql"

	"github.com/rs/rest-layer/resource"
	"github.com/rs/rest-layer/rest"
	"github.com/rs/rest-layer/schema"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_DRIVER		= "sqlite3"
	DB_SOURCE		= "file:data.db"
)

var task  = schema.Schema{
    Fields: schema.Fields{
			"id": schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
            "name": { Required: true, Validator: &schema.String{ MaxLen: 150 } },
            "info": { Required: true, Validator: &schema.String{} },
            "exp": { Required: true, Default: 1, Sortable: true, Filterable: true, Validator: &schema.Integer{} },
            "added": { Required: true, Validator: &schema.String{ MaxLen: 20 } },
            "confirmed": { Required: true, Validator: &schema.String{ MaxLen: 20 } },
            "catid": { Required: true, Validator: &schema.String{ MaxLen: 20 } },
            "score": { Required: true, Default: 0, Sortable: true, Filterable: true, Validator: &schema.Integer{} },
    },
}

var category = schema.Schema{
    Fields: schema.Fields{
			"id": schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
            "name": { Required: true, Validator: &schema.String{ MaxLen: 150 } },
            "info": { Required: true, Validator: &schema.String{} },
            "icon": { Required: true, Validator: &schema.String{} },
    },
}

var assign = schema.Schema{
    Fields: schema.Fields{
			"id": schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
            "taskid": { Required: true, Validator: &schema.String{ MaxLen: 20 } },
            "userid": { Required: true, Validator: &schema.String{ MaxLen: 20 } },
    },
}

var user = schema.Schema{
    Fields: schema.Fields{
			"id": schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
            "username": { Required: true, Validator: &schema.String{ MaxLen: 150 }, Filterable: true },
            "name": { Required: true, Validator: &schema.String{} },
            "pass": { Required: true, Validator: &schema.String{}, Hidden: true, Filterable: true },
            "exp": { Required: true, Default: 0, Sortable: true, Filterable: true, Validator: &schema.Integer{} },
            "admin": { Required: true, Default: false, Filterable: true, Validator: &schema.Bool{} },
    },
}

func addBind(name string, sch schema.Schema, index resource.Index) (*resource.Resource, resource.Storer) {
    cfg := sqlStorage.Config{ 2, map[string]string{} }
    s, err := sqlStorage.NewHandler(DB_DRIVER, DB_SOURCE, name, &cfg)
    if err != nil { log.Fatalf("[%s] Error connecting database: %s", name, err) }
    err = s.Create(context.TODO(), &sch)
    if err != nil { log.Fatalf("[%s] Error creating table: %s", name, err) }
    return index.Bind(name, sch, s, resource.DefaultConf), s
}

func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
        next.ServeHTTP(w, r)
    })
}

func main() {
	index := resource.NewIndex()
    _, assignsS := addBind("assigns", assign, index)
    tasks, tasksS  := addBind("tasks", task, index)
    users, _ := addBind("users", user, index)
    cats, _ := addBind("categories", category, index)

    users.Bind("tasks", "userid", assign, assignsS, resource.DefaultConf)
    cats.Bind("tasks", "catid", task, tasksS, resource.DefaultConf)
    tasks.Bind("users", "taskid", assign, assignsS, resource.DefaultConf)


	api, err := rest.NewHandler(index)
	if err != nil { log.Fatalf("Invalid API configuration: %s", err) }
	http.Handle("/", middleware(api))
	log.Print("Serving API on http://localhost:8888")
	if err := http.ListenAndServe(":8888", nil); err != nil { log.Fatal(err) }
}
