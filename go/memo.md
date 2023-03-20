# go

- rune
- pointer
- struct
- interface

## 文法
### 命名規則
- アンダースコアは使われない
  - example: `update_user` -> `updateUser`
- 大文字からはじまるものは公開シンボルとして扱われる

### 条件分岐
セミコロンを用いて以下のように書くこともできます。
```go
 if user, err := userName(); err == nil {
     fmt.Println(user.Name)
 } else {
     log.Println(err)
}
```

### ループ
GoにはLabeled Breakという記法があり、多重ループの内側から一気に外側のループまで 抜けることができるようになっています。
```go
 finished:
     for i := 0; i < 100; i++ {
         for j := 0; j < 100; j++ {
             if check(i, j) {
                 break finished
             }
} }
```

これらの for ループとは別に、channel(後述)の読み取りを for で書くことができます。
```go
for v := range ch {
     // do something
}
```
この場合 、`channel` が `close(ch)` により閉じられるとforループが終了します 。

### 構造体
 関数の引数などに `struct` のデータを渡すと都度、コピーが行われます。コピーのオーバーヘッドをなくすのであれば、ポインタを使うと良いでしょう。

 ```go
func showName(user *User) {
              fmt.Println(user.wName)
}
func main() {
    user := User {
    Name: "Bob",
    Age: 18,
    }
    showName(&user)
}
```
ここでは `&` を使っています。これは参照と呼び、ポインタを受け取る関数に実体の変数を渡す際 に利用します。

### ポインタ
`&` を使ってポインタを得て `*` による実体の参照ができます。

### コンストラクタ
```go
func NewUser(name string, age int) *User {
    return &User {
        Name: name,
        Age: age,
    }
}
```

### goroutine

Go のランタイムで管理された軽量なスレッド

コンパイル時に-raceを付けて実行することでランタイムが`race condition`を検出してくれます。

- CPIをつかわない処理
    - 通信
    - ファイルのI/O

`race condition`が発生した場合にはどのような結果になるかが保証されていません。 データを保護するには `sync.Mutex` を使って保護する必要があります。

#### channel

`channel` は `goroutine` に対してデータを送受信できるしくみです。




