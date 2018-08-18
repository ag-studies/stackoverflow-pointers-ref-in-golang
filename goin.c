#include <stdio.h>
#include "libtest.h"

int main() {
    //pointer
    GoUintptr p;
    double* pp;
    //call function
    p = ExportedFunction(3.0, 3.1, 3.2, 2);
    //convert result in double pointer
    pp = (double*) p;
    //print result
    printf("P=%lu R0=%f R1=%f\n", p, pp[0], pp[1]);
    //exit
    return 0;
}