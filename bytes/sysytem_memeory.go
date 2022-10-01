package main

import (
        "fmt"
        "runtime"
        "runtime/debug"
)

// Allocate 1 MB
var big = make([]byte, 1<<20)

func main() {
        var ms1, ms2 runtime.MemStats

        // Drop the link to 1 MB
        big = nil
        // Force GC
        runtime.GC()
        runtime.ReadMemStats(&ms1)
        // Force memory release
        debug.FreeOSMemory()
        runtime.ReadMemStats(&ms2)
        fmt.Println("1MB in bytes: ", 1<<20)
        fmt.Println("Idle memory before: ", ms1.HeapIdle)
        fmt.Println("Idle memory after: ", ms2.HeapIdle)
        fmt.Println("Idle memory delta: ", int64(ms2.HeapIdle)-int64(ms1.HeapIdle))
        fmt.Println("Released memory before: ", ms1.HeapReleased)
        fmt.Println("Released memory after: ", ms2.HeapReleased)
        fmt.Println("Released memory delta: ", ms2.HeapReleased - ms1.HeapReleased)
}
