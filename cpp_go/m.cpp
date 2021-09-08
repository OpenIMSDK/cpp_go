#include <stdio.h>
#include <unistd.h>


#include "libcallback.h"




void gocallback(void* s,int len) {

    printf("%s\n", (char*)s);

}




int main() {
        const char* a = "cstring input";
	doSomethingCallback(gocallback, (char*)"cstring hello", (char*)a);
	pause();
}


