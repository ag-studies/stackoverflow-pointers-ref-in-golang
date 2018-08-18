from ctypes import c_double, cdll, c_int, c_byte
from numpy.ctypeslib import ndpointer

# load library
lib = cdll.LoadLibrary("libtest.so")
# define args types
lib.ExportedFunction.argtypes = [c_double]*3 + [c_int]

# define a lovely function
def lovely_python_function(s0, s1, s2, N):
    lib.ExportedFunction.restype = ndpointer(dtype = c_double, shape = (N,))
    return lib.ExportedFunction(s0, s1, s2, N)

# test
a = lovely_python_function(2.0, 2.1, 2.2, 10)

# expected (2.0 + 2.1 + 2.2 = 6.3) (N=10 times)
print a 
# print [ 6.3  6.3  6.3  6.3  6.3  6.3  6.3  6.3  6.3  6.3]