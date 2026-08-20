#ifndef QCHELPER_H
#define QCHELPER_H

// Zero arguments, no return value
#define QC0V(name)                          \
struct name ## _args {                      \
};                                          \
typedef struct name ## _args name ## _args; \
void name ## _call(void *uargs) {           \
    (void)uargs;                            \
    name();                                 \
}

// Zero arguments, return value
#define QC0R(name, ret_type)                \
struct name ## _args {                      \
    ret_type ret;                           \
};                                          \
typedef struct name ## _args name ## _args; \
void name ## _call(void *uargs) {           \
    struct name ## _args *args;             \
    args = (struct name ## _args*)uargs;    \
    args->ret = name();                     \
}

// One argument, no return value
#define QC1V(name, arg0_type)               \
struct name ## _args {                      \
    arg0_type arg0;                         \
};                                          \
typedef struct name ## _args name ## _args; \
void name ## _call(void *uargs) {           \
    struct name ## _args *args;             \
    args = (struct name ## _args*)uargs;    \
    name(args->arg0);                       \
}

// One argument, return value
#define QC1R(name, arg0_type, ret_type)     \
struct name ## _args {                      \
    arg0_type arg0;                         \
    ret_type ret;                           \
};                                          \
typedef struct name ## _args name ## _args; \
void name ## _call(void *uargs) {           \
    struct name ## _args *args;             \
    args = (struct name ## _args*)uargs;    \
    args->ret = name(args->arg0);           \
}

// Two arguments, no return value
#define QC2V(name, arg0_type, arg1_type)    \
struct name ## _args {                      \
    arg0_type arg0;                         \
    arg1_type arg1;                         \
};                                          \
typedef struct name ## _args name ## _args; \
void name ## _call(void *uargs) {           \
    struct name ## _args *args;             \
    args = (struct name ## _args*)uargs;    \
    name(args->arg0, args->arg1);           \
}

// Two arguments, return value
#define QC2R(name, arg0_type, arg1_type, ret_type) \
struct name ## _args {                             \
    arg0_type arg0;                                \
    arg1_type arg1;                                \
    ret_type ret;                                  \
};                                                 \
typedef struct name ## _args name ## _args;        \
void name ## _call(void *uargs) {                  \
    struct name ## _args *args;                    \
    args = (struct name ## _args*)uargs;           \
    args->ret = name(args->arg0, args->arg1);      \
}

// Three arguments, no return value
#define QC3V(name, arg0_type, arg1_type, arg2_type) \
struct name ## _args {                              \
    arg0_type arg0;                                 \
    arg1_type arg1;                                 \
    arg2_type arg2;                                 \
};                                                  \
typedef struct name ## _args name ## _args;         \
void name ## _call(void *uargs) {                   \
    struct name ## _args *args;                     \
    args = (struct name ## _args*)uargs;            \
    name(args->arg0, args->arg1, args->arg2);       \
}

// Three arguments, return value
#define QC3R(name, arg0_type, arg1_type, arg2_type, ret_type) \
struct name ## _args {                                        \
    arg0_type arg0;                                           \
    arg1_type arg1;                                           \
    arg2_type arg2;                                           \
    ret_type ret;                                             \
};                                                            \
typedef struct name ## _args name ## _args;                   \
void name ## _call(void *uargs) {                             \
    struct name ## _args *args;                               \
    args = (struct name ## _args*)uargs;                      \
    args->ret = name(args->arg0, args->arg1, args->arg2);     \
}

// Four arguments, no return value
#define QC4V(name, arg0_type, arg1_type, arg2_type, arg3_type) \
struct name ## _args {                                         \
    arg0_type arg0;                                            \
    arg1_type arg1;                                            \
    arg2_type arg2;                                            \
    arg3_type arg3;                                            \
};                                                             \
typedef struct name ## _args name ## _args;                    \
void name ## _call(void *uargs) {                              \
    struct name ## _args *args;                                \
    args = (struct name ## _args*)uargs;                       \
    name(args->arg0, args->arg1, args->arg2, args->arg3);      \
}

// Four arguments, return value
#define QC4R(name, arg0_type, arg1_type, arg2_type, arg3_type, ret_type) \
struct name ## _args {                                                   \
    arg0_type arg0;                                                      \
    arg1_type arg1;                                                      \
    arg2_type arg2;                                                      \
    arg3_type arg3;                                                      \
    ret_type ret;                                                        \
};                                                                       \
typedef struct name ## _args name ## _args;                              \
void name ## _call(void *uargs) {                                        \
    struct name ## _args *args;                                          \
    args = (struct name ## _args*)uargs;                                 \
    args->ret = name(args->arg0, args->arg1, args->arg2, args->arg3);    \
}

#endif
