// Copyright 2017 Google Inc. All Rights Reserved.
//
// Distributed under MIT license.
// See file LICENSE for detail or copy at https://opensource.org/licenses/MIT

package cbrotli

// Inform golang build system that it should link brotli libraries.

/*
#cgo CFLAGS: -I${SRCDIR}/libbrotli/include
#include "libbrotli/common/dictionary.c"
#include "libbrotli/dec/bit_reader.c"
#include "libbrotli/dec/decode.c"
#include "libbrotli/dec/huffman.c"
#include "libbrotli/dec/state.c"
*/
import "C"
