#include "card.h"
#include <stdio.h>

int main(int argc, char **argv) {
    Card_T card = create_card(1, spades);
    print_card(&card);
}