package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"number-guessing-game/config"
	"os"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {  //htmlを出力
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {  //サーバーを起動
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", signup)
	http.HandleFunc("/game", game_route)

	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
}
