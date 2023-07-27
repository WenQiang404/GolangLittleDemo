package channel

import (
	"net/http"
	"time"
)

// 使用channel的select实现接口超时控制
func Tools() string {
	time.Sleep(400 * time.Millisecond)
	return "success"
}

func OverTimeControl(w http.ResponseWriter, req *http.Request) {
	var resp string

	data := make(chan struct{}, 1)
	go func() {
		resp = Tools()
		data <- struct{}{}

	}()
	select {
	case <-data:
	case <-time.After(300 * time.Millisecond):
		resp = "timeout"
	}
	w.Write([]byte(resp))
}
