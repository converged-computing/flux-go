package main
/*
#include <flux/core.h>
#include <flux/idset.h>
#include <flux/hostlist.h>
#include <stddef.h>
#include <jansson.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"github.com/flux-framework/flux-core/pkg/core"
)

func main() {
	fmt.Println("⭐️ Testing flux-core in Go! ⭐️")

	flux := core.NewFlux()

	// Handle is at flux.Handle
	fmt.Printf("Submitting a Sleep Job: sleep 10\n")

	// Create and submit a jobspec
	jobspec := core.NewJobSpec("sleep 10")
	future := flux.Submit(jobspec)
	fmt.Printf("Flux Future: %s\n", future)

	C.flux_future_wait_all_create()

	// Get the id for it
	// id := new(C.ulong)
	// fluxid := C.flux_job_submit_get_id(future, id)
    // flux_future_destroy (f);
}