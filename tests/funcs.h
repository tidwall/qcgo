#ifndef FUNCS_H
#define FUNCS_H

#include <stdlib.h>
#include <stdint.h>

void noop(void *uargs);

struct add_args {
    int64_t arg0;
    int64_t arg1;
    int64_t ret;
};

void add(void *uargs);

struct bigstack_args {
	size_t size;
    uint64_t ret;
};
void bigstack(void *uargs);


#endif
