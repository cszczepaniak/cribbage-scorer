use core::fmt::Display;
use std::error::Error;

#[derive(Clone, Copy, Debug, Hash, Eq, PartialEq)]
pub enum Suit {
    Clubs,
    Diamonds,
    Hearts,
    Spades,
    Unknown,
}

impl From<usize> for Suit {
    fn from(n: usize) -> Self {
        match n {
            0 => Suit::Clubs,
            1 => Suit::Diamonds,
            2 => Suit::Hearts,
            3 => Suit::Spades,
            _ => Suit::Unknown,
        }
    }
}

impl From<Suit> for usize {
    fn from(n: Suit) -> Self {
        match n {
            Suit::Clubs => 0,
            Suit::Diamonds => 1,
            Suit::Hearts => 2,
            Suit::Spades => 3,
            Suit::Unknown => usize::MAX,
        }
    }
}

pub fn new_deck() -> [Card; 52] {
    let mut cards = [Card {
        suit: Suit::Unknown,
        value: 0,
        rank: 0,
    }; 52];
    for i in 0..52 {
        cards[i] = Card::from_index(i);
    }
    cards
}

#[derive(Clone, Copy, Debug)]
pub struct Card {
    pub suit: Suit,
    pub value: usize,
    pub rank: usize,
}

impl Card {
    fn from_index(i: usize) -> Card {
        let value = match i % 13 {
            n if n >= 10 => 10,
            n => n + 1,
        };
        Card {
            value,
            suit: (i / 13).into(),
            rank: (i % 13) + 1,
        }
    }
    pub fn from_str(s: &str) -> Result<Card, Box<dyn Error>> {
        let lower_s = &s.to_lowercase();
        let mut c = Card {
            rank: 0,
            value: 0,
            suit: Suit::Unknown,
        };
        match &lower_s[lower_s.len() - 1..] {
            "c" => {
                c.suit = Suit::Clubs;
            }
            "d" => {
                c.suit = Suit::Diamonds;
            }
            "h" => {
                c.suit = Suit::Hearts;
            }
            "s" => {
                c.suit = Suit::Spades;
            }
            _ => Err("invalid suit!")?,
        };
        match &lower_s[..lower_s.len() - 1] {
            "j" => {
                c.rank = 11;
                c.value = 10;
            }
            "q" => {
                c.rank = 12;
                c.value = 10;
            }
            "k" => {
                c.rank = 13;
                c.value = 10;
            }
            "a" => {
                c.rank = 1;
                c.value = 1;
            }
            v => {
                let val = v.parse::<usize>()?;
                c.rank = val;
                c.value = val;
            }
        };
        Ok(c)
    }
}

impl Display for Card {
    fn fmt(
        &self,
        formatter: &mut std::fmt::Formatter<'_>,
    ) -> std::result::Result<(), std::fmt::Error> {
        match self.rank {
            0 => formatter.write_str("A")?,
            10 => formatter.write_str("J")?,
            11 => formatter.write_str("Q")?,
            12 => formatter.write_str("K")?,
            _ => formatter.write_str(&format!("{}", self.value))?,
        }
        match self.suit {
            Suit::Clubs => formatter.write_str("C"),
            Suit::Diamonds => formatter.write_str("D"),
            Suit::Hearts => formatter.write_str("H"),
            Suit::Spades => formatter.write_str("S"),
            Suit::Unknown => formatter.write_str("?"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use claim::*;
    use std::collections::HashSet;

    #[test]
    fn test_new_deck() {
        let d = new_deck();
        assert_eq!(d.len(), 52);
        let mut unique: HashSet<(usize, usize, Suit)> = HashSet::new();

        for c in d.iter() {
            let tup = (c.value, c.rank, c.suit);
            assert!(!unique.contains(&tup));
            unique.insert(tup);
        }
    }

    #[test]
    fn test_from_index_equality() {
        let tests = vec![
            (
                0,
                Card {
                    rank: 0,
                    value: 1,
                    suit: Suit::Clubs,
                },
            ),
            (
                1,
                Card {
                    rank: 1,
                    value: 2,
                    suit: Suit::Clubs,
                },
            ),
            (
                2,
                Card {
                    rank: 2,
                    value: 3,
                    suit: Suit::Clubs,
                },
            ),
            (
                10,
                Card {
                    rank: 10,
                    value: 10,
                    suit: Suit::Clubs,
                },
            ),
            (
                11,
                Card {
                    rank: 11,
                    value: 10,
                    suit: Suit::Clubs,
                },
            ),
            (
                12,
                Card {
                    rank: 12,
                    value: 10,
                    suit: Suit::Clubs,
                },
            ),
            (
                51,
                Card {
                    rank: 12,
                    value: 10,
                    suit: Suit::Spades,
                },
            ),
        ];
        for (index, c) in tests {
            let actual_card = Card::from_index(index);
            assert_eq!(c.rank, actual_card.rank);
            assert_eq!(c.value, actual_card.value);
            assert_eq!(c.suit, actual_card.suit);
        }
    }

    #[test]
    fn test_from_index_full_range() {
        for i in 0..52 {
            let card = Card::from_index(i);
            assert_ge!(card.value, 1);
            assert_le!(card.value, 10);
            match i {
                0..=12 => assert_eq!(card.suit, Suit::Clubs),
                13..=25 => assert_eq!(card.suit, Suit::Diamonds),
                26..=38 => assert_eq!(card.suit, Suit::Hearts),
                _ => assert_eq!(card.suit, Suit::Spades),
            }
        }
    }
}
