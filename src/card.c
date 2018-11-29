#include "card.h"
#include <stdio.h>

void print_card(Card_T *card) {
    printf(card->name);
}

Card_T create_card(int rank, Suit_T suit) {
    char name[3];
    if(rank >=2 && rank <= 10) {
        sprintf(name, "%d", rank);
    }
    else if(rank == 1) {
        name[0] = 'A';
    }
    else if(rank == 11) {
        name[0] = 'J';
    }
    else if(rank == 12) {
        name[0] = 'Q';
    }
    else if(rank == 13) {
        name[0] = 'K';
    }

    
    switch (suit)
    {
        case hearts:
            name[1] = 'H';
            break;
        case diamonds:
            name[1] = 'D';
            break;
        case clubs:
            name[1] = 'C';
            break;
        case spades:
            name[1] = 'S';
            break;
    
        default:
            break;
    }

    name[2] = '\0';

    Card_T card;
    card.name = name;
    card.rank = rank;
    card.suit = suit;
    card.value = rank > 10 ? 10 : rank;
    return card;
}