#include <stdio.h>
#include <stdlib.h>

typedef struct first
{
    int firstInt;
    char firstChar;
} first;

typedef struct second
{
    char secondChar;
    char thirdChar;
    int secondInt;
} second;

typedef struct third
{
    char thirdCharF;
    char thirdCharS;
    int thirdInt;
} third;

typedef struct forth
{
    short forthShort;
    int forthInt;
} forth;

void testSecond(void *test)
{
    second *secondExamplePtr = (second *)test;

    printf("Second int: %d second character: %c Hexint: %x HexSChar: %x\n", secondExamplePtr->secondInt, secondExamplePtr->secondChar, secondExamplePtr->secondInt, secondExamplePtr->secondChar);
}

void testThird(void *test)
{
    third *thirdExamplePtr = (third *)test;

    printf("Third int: %d third characterF: %c third characterS: %c Hexint: %x HexFChar: %x HexSChar: %x\n", thirdExamplePtr->thirdInt, thirdExamplePtr->thirdCharF, thirdExamplePtr->thirdCharS, thirdExamplePtr->thirdInt, thirdExamplePtr->thirdCharF, thirdExamplePtr->thirdCharS);
}

void testForth(void *test)
{
    forth *forthExamplePtr = (forth *)test;
    printf("Forth int: %d forth short: %d Hexint: %x Hexshort: %x\n", forthExamplePtr->forthInt, forthExamplePtr->forthShort, forthExamplePtr->forthInt, forthExamplePtr->forthShort);
}

int main()
{
    // printf("int: %lu char: %lu short: %lu long: %lu long long: %lu, float: %lu, double: %lu, long double: %lu",
    //        sizeof(int), sizeof(char), sizeof(short), sizeof(long), sizeof(long long), sizeof(float), sizeof(double), sizeof(long double));

    first firstExample;
    first *firstExamplePtr = calloc(1, sizeof(first));

    firstExamplePtr->firstInt = 0x45444342;
    firstExamplePtr->firstChar = 0x41;

    printf("First int: %d and the first character: %c Hexint: %x HexChar: %x\n", firstExamplePtr->firstInt, firstExamplePtr->firstChar, firstExamplePtr->firstInt, firstExamplePtr->firstChar);
    testSecond((void *)firstExamplePtr);
    testThird((void *)firstExamplePtr);
    testForth((void *)firstExamplePtr);
}