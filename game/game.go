package game

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var GameStatus Status
var GameSentence Sentence

type Sentence struct { //gameで表示される分がまとめられている
	LastResult string //状態を表す
	LowOrHi    string //大きいか小さいかを表す
	Candid     string //これまでの予想値を表す
	Life       string //ライフを表す
}

type Status struct { //ゲームに必要なデータがまとめられている
	RandomNumber  int //答え
	NumberOfLife  int //ライフの数
	LuckyNumber   []int //ラッキーナンバーのスライス
	UnluckyNumber []int //アンラッキーナンバーのスライス
	CandidSlice   []string //これまでの予想値のスライス
	Clear         bool //クリアしたかどうかを示す
	Gameover      bool //ゲームオーバーかどうか示す
}

func init() {
	StartGame()
}

func StartGame() {                             //各種設定の初期化を行う
	//GameStatusの初期化
	rand.Seed(time.Now().UnixNano())
	GameStatus.RandomNumber = rand.Intn(100)+1
	fmt.Println(GameStatus.RandomNumber)
	GameStatus.NumberOfLife = 10
	GameStatus.CandidSlice = nil
	GameStatus.LuckyNumber, GameStatus.UnluckyNumber = Lucky_unlucky_number(GameStatus.RandomNumber)
	GameStatus.Clear = false
	GameStatus.Gameover = false
	//GameSentenceの初期化
	GameSentence.LastResult = ""
	GameSentence.LowOrHi = ""
	GameSentence.Candid = ""
	GameSentence.Life = strings.Repeat("♡", GameStatus.NumberOfLife)
}

func StringInSlice(a int, list []int) bool {   //要素がスライスの中にあるか判定
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Lucky_unlucky_number(RandomNumber int) (luckyNumber []int, unluckyNumber []int) { //ラッキーナンバー、アンラッキーナンバーを重複なしで30個ずつ選択
	var a []int
	a = append(a, RandomNumber)
	i := 0
	for i < 60 {
		num := rand.Intn(100)
		if !StringInSlice(num, a) {
			a = append(a, num)
			i++
		}
	}
	luckyNumber = a[1:31]
	unluckyNumber = a[31:61]
	return
}

func CheckGuess(userGuess int) { //ユーザーが選択数字から状態を更新する
	GameStatus.NumberOfLife--
	if userGuess == GameStatus.RandomNumber{//ゲームがクリアされた時
		GameStatus.Clear = true
		GameStatus.NumberOfLife++
	} else  if userGuess < GameStatus.RandomNumber {//予想値が答えより小さかった時
		GameSentence.LastResult = "不正解！！"
		GameSentence.LowOrHi = "小さすぎます。もっと大きい値です"
	} else {                                        //予想値が答えより大きかった時
		GameSentence.LastResult = "不正解！！"
		GameSentence.LowOrHi = "大きすぎます。もっと小さい値です"
	}
	if StringInSlice(userGuess, GameStatus.LuckyNumber) {//予想値がラッキーナンバーだった時
		GameSentence.LastResult = "ラッキーナンバーです"
		GameStatus.NumberOfLife++
	}
	if StringInSlice(userGuess, GameStatus.UnluckyNumber) {//予想値がアンラッキーナンバーだった時
		GameSentence.LastResult = "アンラッキーナンバーです"
		GameStatus.NumberOfLife--
	}
	if GameStatus.NumberOfLife<=0{//ゲームオーバーになった時
		GameStatus.Gameover = true
	}
	a := strconv.Itoa(userGuess)
	GameStatus.CandidSlice = append(GameStatus.CandidSlice, a) 
	GameSentence.Candid = strings.Join(GameStatus.CandidSlice, " ") //これまでの予想値をスライスから文字列に変換している
	num := math.Max(0, float64(GameStatus.NumberOfLife))//数値が0以下にならないようにする
	GameSentence.Life = strings.Repeat("♡", int(num)) //ライフの数をからハートの文字列に変換している
}
