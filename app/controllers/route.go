package controllers

import (
	"log"
	"net/http"
	"number-guessing-game/app/models"
	"number-guessing-game/game"
	"strconv"
)

type Data struct {  //clearに送るデータの構造体
	NowUser models.User
	TopUser []models.User
}

func signup(w http.ResponseWriter, r *http.Request) {  //signup画面を出力
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		name := r.PostFormValue("name")
		models.PlayedUser.Name = name //受け取った名前をUserのNameにする
		http.Redirect(w, r, "/game", 302)
	}
}

func game_route(w http.ResponseWriter, r *http.Request) {//予想値を受け取り、条件に合わせて画面を出力
	if r.Method == "GET" {
		generateHTML(w, game.GameSentence, "layout", "game")

	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		guessnum, _ := strconv.Atoi(r.PostFormValue("number"))

		game.CheckGuess(guessnum)
		if game.GameStatus.Clear {
			models.PlayedUser.Score = game.GameStatus.NumberOfLife//Scoreを追加
			models.PlayedUser.CreateUser()//Userを追加
			topusers := models.GetTopUser()//Score上位10名のデータ
			data := Data{NowUser: models.PlayedUser, TopUser: topusers}//まとめたデータ
			generateHTML(w, data, "layout", "clear")
			game.StartGame()//ゲームが終了したので初期化
		} else if game.GameStatus.Gameover {
			generateHTML(w, nil, "layout", "gameover")
			game.StartGame()//ゲームが終了したので初期化
		} else {
			generateHTML(w, game.GameSentence, "layout", "game")
		}

	}
}
