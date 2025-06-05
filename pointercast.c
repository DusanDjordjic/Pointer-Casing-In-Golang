#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char* Data;
    unsigned int Len;
} StringHeader;

typedef struct {
    char* Data;
    unsigned int Len;
    unsigned int Cap;
} ByteSliceHeader;

ByteSliceHeader createBuffer(unsigned int len, int* error);

int main() {
    int error = 0;
    ByteSliceHeader buffer = createBuffer(10, &error);

    if (error != 0) {
        printf("failed to create buffer, error %d", error);
        return 1;
    }
   
    // Manually add some data to buffer
    buffer.Data[0] = 'A';
    buffer.Data[1] = 'B';
    buffer.Data[2] = 'C';
    buffer.Data[3] = '\0';
    buffer.Len = 4;

    StringHeader string = *((StringHeader*)(&buffer));
   
    printf("\"%s\" %u \n", string.Data, string.Len); // "ABC" 4

    // This is why in golang it's unsafe to use buffer after you have casted it to string as
    // strings in golang are immutable, but if you drop the buffer than it's fine
    buffer.Data[2] = 'D';
    printf("\"%s\" %u \n", string.Data, string.Len); // "ABD" 4

    // Free still works because it's the same allocation
    free(string.Data);
    return 0;
    
}

ByteSliceHeader createBuffer(unsigned int cap, int* error) {
    ByteSliceHeader header = {};

    // calloc sets all bytes to 0
    char* data = calloc(sizeof(char), cap);
    if (data == NULL) {
        *error = 1;
        return header;
    }

    header.Data = data;
    header.Len = 0;
    header.Cap = cap;

    return header;
}
