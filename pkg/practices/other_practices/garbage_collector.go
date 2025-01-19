package other_practices

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

/*print all info about garbage collector property*/
func GCinfo() {
	var stats runtime.MemStats
	err := os.Setenv("GOGC", "50")
	if err != nil {
		log.Println(err)
	}

	runtime.ReadMemStats(&stats)
	fmt.Printf("garbage collector is set to %s\n", os.Getenv("GOGC"))
	fmt.Println("HeapAlloc:", stats.HeapAlloc)         // Текущий размер кучи
	fmt.Println("HeapSys:", stats.HeapSys)             // Общий размер кучи
	fmt.Println("GCCPUFraction:", stats.GCCPUFraction) // Процент времени, затрачиваемого на GC
}

// one of the best practice in topic 'garbage collector' is using the pool
func PoolInGC() {
	pool := sync.Pool{
		New: func() any {
			buf := make([]byte, 1024)
			return buf
		},
	}

	for i := 0; i < 15; i++ {
		data := pool.Get().([]byte)
		someAction(data)

		pool.Put(data)
	}
}

func someAction(data []byte) {
	fmt.Println(string(data))
	fmt.Println(len(data))
	fmt.Println()
}
