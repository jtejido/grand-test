package main

/*
#cgo CFLAGS: -I${SRCDIR}/TestU01/include
#cgo LDFLAGS: ${SRCDIR}/TestU01/testu01/.libs/libtestu01.a ${SRCDIR}/TestU01/probdist/.libs/libprobdist.a ${SRCDIR}/TestU01/mylib/.libs/libmylib.a -lm

#include "unif01.h"
#include "bbattery.h"
#define BUFSZ (1<<20)
extern void goRandInt32Batch(void *);

static unsigned int goRandInt32(void) {
	static unsigned int buffer[BUFSZ];
	static int n = 0;
	if (n == 0) {
		goRandInt32Batch(buffer);
		n = BUFSZ;
	}
	return buffer[--n];
}

static int test(void) {
   unif01_Gen *gen;
   gen = unif01_CreateExternGenBits ("rand.Uint32", goRandInt32);
   bbattery_BigCrush (gen);
   unif01_DeleteExternGenBits (gen);
   return 0;
}
*/
import "C"
import "github.com/jtejido/grand"
import "github.com/jtejido/grand/source32"
import "unsafe"

var (
	Test      = source32.NewPcgMcgXshRr32(12345) // change this
	Generator = grand.New(Test)
)

const BUFSZ = C.BUFSZ

//export goRandInt32Batch
func goRandInt32Batch(p unsafe.Pointer) {
	s := (*[BUFSZ]uint32)(p)[:]
	for i := range s {
		s[i] = Generator.Uint32()
	}
}

func main() {
	C.test()
}
