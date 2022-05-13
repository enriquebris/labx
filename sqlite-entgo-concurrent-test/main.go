package main

import (
	"context"
	"fmt"
	"github.com/enriquebris/labx/sqlite-entgo-concurrent-test/ent"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var (
		totalGRs                = 100
		totalOperationPerWorker = 25
		//done                    = make(chan struct{})
		wg         sync.WaitGroup
		failuresMp sync.Map
	)

	// create GRs
	for i := 0; i < totalGRs; i++ {
		wg.Add(1)
		go worker(i, &wg, client, totalOperationPerWorker, &failuresMp)
	}

	wg.Wait()
	log.Println("\n\nall workers are done !.!\n\n")

	// show failures per worker
	for i := 0; i < totalGRs; i++ {
		failures := 0
		if raw, ok := failuresMp.Load(i); ok {
			failures = raw.(int)
		}
		log.Printf("[%v] failures: %v", i, failures)
	}
}

func worker(id int, wg *sync.WaitGroup, entClient *ent.Client, totalOperations int, failuresMp *sync.Map) {
	for i := 0; i < totalOperations; i++ {
		if _, err := entClient.Dummy.Create().SetName(fmt.Sprintf("%v-dummy", id)).SetAge(id).Save(context.Background()); err != nil {
			log.Printf("[%v] failed adding entry: %v", id, err)

			totalErrors := 1
			raw, loaded := failuresMp.LoadOrStore(id, totalErrors)
			if loaded {
				totalErrors = raw.(int)
				failuresMp.Store(id, totalErrors+1)
			}
		}
	}

	log.Printf("worker %v done", id)

	wg.Done()
	// sending back "done" signal
	//done <- struct{}{}
}
