enum Suit_T {
    hearts = 0,
    diamonds = 1,
    clubs = 2,
    spades = 3
};

typedef struct Card {
    int rank;
    int value;
    char* name;
    enum Suit_T suit;
} Card_T;