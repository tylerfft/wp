package tool

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var GenId func() string

func init() {
	GenId = func() func() string {
		var cnt int = 1000000
		var GenIdLock sync.RWMutex
		return func() string {
			GenIdLock.Lock()
			defer GenIdLock.Unlock()
			out1 := fmt.Sprint(time.Now().Unix(), cnt)
			out1 = strings.Replace(out1, " ", "", -1)
			out2 := fmt.Sprint(cnt)
			out2 = strings.Replace(out2, " ", "", -1)
			out := out1 + "_" + out2
			cnt++
			if cnt == 10000000 {
				cnt = 1000000
			}
			return out
		}
	}()
}
