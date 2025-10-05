# timer

Go言語向けのシンプルなパフォーマンス計測ライブラリです。処理の各ステップにかかった時間を計測し、ファイル名と行番号を自動記録します。

- `New()` → `Lap()` → `Print()` の3ステップで計測できてお手軽です。
- 標準パッケージにのみ依存しています。

## インストール

```bash
go get github.com/naoyafurudono/timer
```

## 使い方

```go
package main

import (
    "time"
    "github.com/naoyafurudono/timer"
)

func main() {
    // タイマーを作成
    t := timer.New("example")

    // 処理の計測
    t.Lap("Start processing")

    time.Sleep(100 * time.Millisecond)
    t.Lap("First step completed")

    // ループ内での計測
    for i := range 3 {
        time.Sleep(50 * time.Millisecond)
        t.Lap("Loop iteration %d", i)
    }

    time.Sleep(200 * time.Millisecond)
    t.Lap("Final step completed")

    // 結果を表示
    t.Print()

    // JSON形式で表示
    t.PrintJSON()
}
```

### 出力例

**テーブル形式（`Print()`）:**

```
--- Performance Measurement Results ---
name: example
Elapsed            Location        Message
-------            --------        -------
123.456µs          main.go:14      Start processing
100.789ms          main.go:17      First step completed
150.234ms          main.go:22      Loop iteration 0
200.567ms          main.go:22      Loop iteration 1
250.891ms          main.go:22      Loop iteration 2
451.234ms          main.go:26      Final step completed
---------------------------------------
```

**JSON形式（`PrintJSON()`）:**

```json
{"laps":[{"file":"main.go","line":14,"message":"Start processing","duration":"123.456µs"},{"file":"main.go","line":17,"message":"First step completed","duration":"100.789ms"},...],"name":"example"}
```

## ライセンス

MIT
