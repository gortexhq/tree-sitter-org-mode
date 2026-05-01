package tree_sitter_orgmode

// #cgo CFLAGS: -I${SRCDIR}/../../src -std=c11 -fPIC
// #include "../../src/parser.c"
// #include "../../src/scanner.c"
import "C"

import "unsafe"

// Language returns the tree-sitter Language for Org-mode.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_orgmode())
}
