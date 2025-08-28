# grt

Purposes:
- write basic raytracing/rasterization algorithms from scratch
- cpu based
- familiarize myself with golang low-level optimizations (builtin asm e.g.)
- have some fun :)

# decisions

- no asserts inside basic primitives
- color represenation - float vs byte, see comments in color.go, first attempt with bytes

# todo's

- add perf test to compare matrix inversion alogrithms

## done

+ add support for ppm file format - a primitive way to see results
+ add fast inversion path for SRT matrices (see org doc - matrix inversion for details)
+ add testify package to simplify asserts
