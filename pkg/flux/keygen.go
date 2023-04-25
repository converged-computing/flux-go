package flux

/*
#include <unistd.h>
#include <flux/core.h>
#include <czmq.h>
#include <sodium.h>
#include <stddef.h>
#include <stdlib.h>
#include "../flux/cgo_helpers.h"

void flux_zcert_set_meta (zcert_t *cert, const char *field, const char *name) {
    zcert_set_meta (cert, field, "%s", name);
}
*/
import "C"
import (

    "fmt"
	"unsafe"
)

// Keygen uses zeromq, and was originally part of the flux/cmd directory
// so we put it under flux implying that, since "cmd" has special meaning
// in Go!

// KeyGen generates a curve certificate
func KeyGen(name string, hostname string, path string) {

    // Create the new certificate (likely want to check for error here)
    cert := C.zcert_new()

    // Use wrapper to set the cert metadata
    // Name (typically the hostname but doesn't need to be)
    // used in overlay logging
    nameValue := C.CString(name)
    nameField := C.CString("name")
    defer C.free(unsafe.Pointer(nameValue))
    defer C.free(unsafe.Pointer(nameField))
    C.flux_zcert_set_meta(cert, nameField, nameValue)

    // Hostname
    hostnameValue := C.CString(hostname)
    hostnameField := C.CString("keygen.hostname")
    defer C.free(unsafe.Pointer(hostnameValue))
    defer C.free(unsafe.Pointer(hostnameField))
    C.flux_zcert_set_meta(cert, hostnameField, hostnameValue)

    // Note that we can also generate keygen.time, keygen.userid,
    // And other version metadata. See
    cpath := C.CString(path)
    defer C.free(unsafe.Pointer(cpath))
    fmt.Printf("Saving to %s\n", path)    
    C.zcert_save_secret (cert, cpath)
    C.zcert_destroy (&cert)
}