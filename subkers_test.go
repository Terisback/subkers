package subkers

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	for i := 0; i <= 10; i++ {
		got := timeToString(time.Second*time.Duration(i) + time.Millisecond*time.Duration(99*i))
		want := strings.ReplaceAll(fmt.Sprintf("0:%2.v.%3.v", i, i*99), " ", "0")
		if got != want {
			t.Errorf("Want %s, Got %s", want, got)
		}
	}
}
