#include <time.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>


bool numguess(int guessNum, int actualNum){
    if(guessNum<actualNum){
        return false;
    }else if (guessNum>actualNum) {
        return false;
    }else{
        return true;
    }
}
int main(){
    int actualNum, guessNum, Attempts;
    int tries = 0;

    srand(time(NULL));
    actualNum = rand()%10000000;
    while(true){
        guessNum = (rand()%10000000)+1;
        if (numguess(guessNum, actualNum)){
            printf("The number was: %d\n", guessNum);
            break;
        }
        tries++;
    }
    printf("It took %d tries\n", tries);
    return 0;
}
