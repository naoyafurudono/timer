package timer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/tabwriter"
	"time"
)

// Lap は各計測地点の情報を保持する構造体です。
type Lap struct {
	File     string        // ファイル名
	Line     int           // 行番号
	Message  string        // ループ情報などのメッセージ
	Duration time.Duration // 計測開始からの経過時間
}

// Timer は処理時間計測の本体となる構造体です。
type Timer struct {
	startTime time.Time
	laps      []Lap
}

// New は新しいタイマーを作成し、計測を開始します。
func New() *Timer {
	return &Timer{
		startTime: time.Now(),
		laps:      []Lap{},
	}
}

// Lap は計測地点を記録します。
// messageFormat には fmt.Sprintf と同じ形式で、ループ変数などの情報を渡すことができます。
func (t *Timer) Lap(messageFormat string, args ...any) {
	// runtime.Caller(1) を使うことで、この Lap メソッドを呼び出した元の情報を取得します。
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	lap := Lap{
		File:     filepath.Base(file), // フルパスではなくファイル名のみを記録
		Line:     line,
		Message:  fmt.Sprintf(messageFormat, args...),
		Duration: time.Since(t.startTime),
	}
	t.laps = append(t.laps, lap)
}

// Print は計測結果を整形して標準出力に表示します。
// text/tabwriter を利用して、結果を見やすく整形します。
func (t *Timer) Print() {
	fmt.Println("--- Performance Measurement Results ---")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Elapsed\tLocation\tMessage")
	fmt.Fprintln(w, "-------\t--------\t-------")
	for _, lap := range t.laps {
		// 結果をフォーマットして出力
		fmt.Fprintf(w, "%-15v\t%s:%d\t%s\n", lap.Duration, lap.File, lap.Line, lap.Message)
	}
	w.Flush() // バッファの内容をフラッシュして出力
	fmt.Println("---------------------------------------")
}

// PrintJSON は計測結果を1行のJSON形式で標準出力に表示します。
func (t *Timer) PrintJSON() {
	type lapJSON struct {
		File     string `json:"file"`
		Line     int    `json:"line"`
		Message  string `json:"message"`
		Duration string `json:"duration"`
	}

	lapsJSON := make([]lapJSON, len(t.laps))
	for i, lap := range t.laps {
		lapsJSON[i] = lapJSON{
			File:     lap.File,
			Line:     lap.Line,
			Message:  lap.Message,
			Duration: lap.Duration.String(),
		}
	}

	data, err := json.Marshal(map[string]any{
		"laps": lapsJSON,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal JSON: %v\n", err)
		return
	}
	fmt.Println(string(data))
}
