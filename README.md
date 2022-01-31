# 数字当てゲーム
 1〜100までのランダムで与えられた数字を当てるゲームです。

 ライフが10与えられて、それがなくなるまでに数字を当てましょう。

 ライフは数字を予想するごとに1減っていき、当てた時の残りライフが得点になります。

 またラッキーナンバーとアンラッキーナンバーがあり、ラッキーナンバーが当たるとライフが減らず、アンラッキーナンバーを選ぶとライフがさらに1減ります。

# DEMO
以下のリンクからこのゲームが遊べます。 

https://gentle-sands-90094.herokuapp.com/game

# Features
golangの勉強を最近していたので、golangで記述しました。参考(まだ途中です。）は以下に載せます。

フロントエンドはbootstrapを用いました。

ラッキーナンバーとアンラッキーナンバーを用いることで、得点のバラつきを大きくしました。

名前と得点がデータベースに保存され、クリア画面で得点上位10名が表示されるようにしました。ランキングにのれるように何度も遊んでもらいたいです。

クリア画面で名前と点数が出るようにしました。

# Requirement

* go 1.17
* github.com/lib/pq 1.10.4
* gopkg.in/go-ini/ini.v1 1.66.3
 
# Installation
 
```bash
//postgresql
go get github.com/lib/pq

//ini
go get gopkg.in/go-ini/ini.v1
```
 
# 参考
【Go入門】Golang基礎入門 + 各種ライブラリ + 簡単なTodoWebアプリケーション開発(Go言語)

作成者: M.A EduTech

https://www.udemy.com/course/golang-webgosql/
