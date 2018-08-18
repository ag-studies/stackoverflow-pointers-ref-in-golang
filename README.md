# stackoverflow-pointers-ref-in-golang

This project was created to answer a question about exchange pointer between python and golang.

## How do

To build shared-library
```
$ export GOPATH=$(pwd) 
$ go build -o libtest.so -buildmode=c-shared pyrefs
```

## Tests

To Golang-based test:
```
$ export GOPATH=$(pwd) 
$ go run src/pyrefs/function.go
```

To C-based test:
```
clang -o goinc.test goin.c -L. -ltest
```

In MacOS-10.12, to C-based test:
```
clang -o goinc.test goin.c -L. -L/usr/lib -ltest
```

To Python-based test:
```
python goin.py
```

## About problem:

The developer needs call a golang function from python passing float and int type params and returning an slice, but occour an unkwnow error.

Source: https://stackoverflow.com/questions/51845092/how-to-return-an-array-from-golang-to-python-using-ctypes/51912417

## The solution:

The developer used the code follow to set ```restype``` to lib:


    lib.Function.restype = ndpointer(dtype = c_double, shape = (N,))

In this case ```restype``` is a pointer type (by Numpy doc):

> **def ndpointer(dtype=None, ndim=None, shape=None, flags=None)** 
> 
> *[others texts]*
>
> Returns
> 
> klass : ndpointer type object
> 
> A type object, which is an `_ndtpr` instance containing dtype, ndim, shape and flags information.


The appropriated type in Golang to treat it is ```unsafe.Pointer```.

However yet has a problem: slice pointer like result. This violate rule 2 of ![**Rules for passing pointers between Go and C**](https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md):

> C code may not keep a copy of a Go pointer after the call returns.

The solution is to convert ```unsafe.Pointer``` to ```uintptr```. 

More details about solution: https://stackoverflow.com/questions/51845092/how-to-return-an-array-from-golang-to-python-using-ctypes/51912417#51912417


<!-- About C-interop, unsafe
For the purposes of C-interop, unsafe.Pointer(&bytes) will create a pointer to the first byte of the slice, which is not the first byte of the data (which is usually what C expects)--for this reason, you should use unsafe.Pointer(&bytes[0]) -->
