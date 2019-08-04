package runtime

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"
)

const pprofDumpLink = "/tmp/pprof-dump.log"

// EnableGoroutineDump enables dumping stack trace of all goroutines with OS
// signal SIGUSR1. To trigger the dump, run command "kill -SIGUSR1 <pid>".
func EnableGoroutineDump() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			ts := time.Now()
			fmt.Println("=== BEGIN pprof dump ===")
			fmt.Println("Timestamp: ", ts.Format(time.RFC3339Nano))

			runtime.GC() // get up-to-date statistics
			ps := pprof.Profiles()

			go dumpToFile(ts, ps)
			go dumpHeapProfile(ts)

			for _, p := range ps {
				fmt.Printf("--- BEGIN %s dump ---\n", p.Name())
				p.WriteTo(os.Stdout, 2)
				fmt.Printf("--- END %s dump ---\n", p.Name())
			}
			fmt.Println("=== END pprof dump ===")
		}
	}()
}

func dumpToFile(ts time.Time, ps []*pprof.Profile) {
	fn := fmt.Sprintf("/tmp/pprof-dump-%s.log", ts.Format("20060102-150405"))
	f, err := os.Create(fn)
	if err != nil {
		fmt.Printf("Error to create pprof dump log file: %v.\n", err)
	}
	defer func() {
		f.Close()
		os.Remove(pprofDumpLink)
		os.Symlink(fn, pprofDumpLink)
	}()

	fmt.Fprintf(f, "Timestamp: %s\n", ts.Format(time.RFC3339Nano))
	for _, p := range ps {
		fmt.Fprintf(f, "--- BEGIN %s dump ---\n", p.Name())
		p.WriteTo(f, 2)
		fmt.Fprintf(f, "--- END %s dump ---\n", p.Name())
	}
}

func dumpHeapProfile(ts time.Time) {
	fn := fmt.Sprintf("/tmp/dump-heap-%s.mprof", ts.Format("20060102-150405"))
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
