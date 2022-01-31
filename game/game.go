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

type Sentence struct {
	LastResult string
	LowOrHi    string
	Candid     string
	Life       string
}

type Status struct {
	RandomNumber int
	NumberOfLife int
	LuckyNumber []int
	UnluckyNumber []int
	CandidSlice []string
	Clear      bool
	Gameover   bool
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
	if userGuess == GameStatus.RandomNumber{
		GameStatus.Clear = true
		GameStatus.NumberOfLife++
	} else  if userGuess < GameStatus.RandomNumber {
		GameSentence.LastResult = "不正解！！"
		GameSentence.LowOrHi = "小さすぎます。もっと大きい値です"
	} else {
		GameSentence.LastResult = "不正解！！"
		GameSentence.LowOrHi = "大きすぎます。もっと小さい値です"
	}
	if StringInSlice(userGuess, GameStatus.LuckyNumber) {
		GameSentence.LastResult = "ラッキーナンバーです"
		GameStatus.NumberOfLife++
	}
	if StringInSlice(userGuess, GameStatus.UnluckyNumber) {
		GameSentence.LastResult = "アンラッキーナンバーです"
		GameStatus.NumberOfLife--
	}
	if GameStatus.NumberOfLife<=0{
		GameStatus.Gameover = true
	}
	a := strconv.Itoa(userGuess)
	GameStatus.CandidSlice = append(GameStatus.CandidSlice, a)
	GameSentence.Candid = strings.Join(GameStatus.CandidSlice, " ")
	num := math.Max(0, float64(GameStatus.NumberOfLife))//数値が0以下にならないようにする
	GameSentence.Life = strings.Repeat("♡", int(num))
}
