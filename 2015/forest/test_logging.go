package forest

import "testing"

func TestThatInfoLoggingIsPrinted(t *testing.T) {
	l := logging{false} // HL
	captureStdout(t, func() {
		l.Logf("%s", "logf")
	}, func(output string) {
		if output != "\tinfo : logf\n" {
			t.Errorf("different output than expected:%s", output)
		}
	})
}
