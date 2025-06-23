package foo

import (
	"bytes"
	"context"
	"runtime/pprof"
	"strings"
	"sync"
	"testing"
)

func TestFoo(t *testing.T) {
	start := sync.WaitGroup{}
	start.Add(1)
	end := sync.WaitGroup{}
	end.Add(1)
	label := "whatdoyousay"
	go func() {
		pprof.Do(context.Background(), pprof.Labels(label, "hello"), func(ctx context.Context) {
			//time.Sleep(time.Millisecond) //with this wait test passes
			start.Done()
			end.Wait()
		})
	}()
	start.Wait()
	prof := new(bytes.Buffer)
	if err := pprof.Lookup("goroutine").WriteTo(prof, 1); err != nil {
		t.Fatal(err)
		return
	}
	profileString := prof.String()
	if !strings.Contains(profileString, label) {
		t.Fatalf("no pprof label %s in goroutine profile %s", label, profileString)
		return
	}
	end.Done()
}
