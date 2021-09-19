use crate::cards::Card;
use itertools::Itertools;

pub fn score_hand(hand: &[Card; 4], cut: Card, is_crib: bool) -> usize {
    score_runs_and_pairs_and_nobs(hand, cut)
        + score_fifteens(hand, cut)
        + score_flush(hand, cut, is_crib)
}

fn score_fifteens(hand: &[Card; 4], cut: Card) -> usize {
    let all = [hand[0], hand[1], hand[2], hand[3], cut];
    if (all.iter().map(|c| c.value).reduce(|a, b| a | b).unwrap()) & 1 == 0 {
        // every card is even; no fifteens possible
        return 0;
    }
    let sum: usize = all.iter().map(|c| c.value).sum();
    if sum < 15 || sum > 46 {
        // in these cases, we know we can't have 15s so skip the work
        return 0;
    }
    // functional programming makes this one-liner SO readable!
    (2..=5)
        .map(|n| {
            &all.map(|c| c.value)
                .iter()
                .combinations(n)
                .map(|c| {
                    let s: usize = c.iter().copied().sum();
                    s
                })
                .filter(|&n| n == 15)
                .count()
                * 2
        })
        .sum()
}

fn score_flush(hand: &[Card; 4], cut: Card, is_crib: bool) -> usize {
    let first = hand[0].suit;
    for c in &hand[1..] {
        if c.suit != first {
            return 0;
        }
    }
    match (is_crib, cut.suit == first) {
        (true, false) => 0,
        (false, false) => 4,
        (_, true) => 5,
    }
}

fn score_runs_and_pairs_and_nobs(hand: &[Card; 4], cut: Card) -> usize {
    let mut rank_counts = [0usize; 15];
    let mut ranks = [0usize; 5];
    let all: [Card; 5] = [hand[0], hand[1], hand[2], hand[3], cut];
    for (i, c) in all.iter().enumerate() {
        rank_counts[c.rank] += 1;
        ranks[i] = c.rank
    }
    score_pairs(&rank_counts)
        + score_runs(&ranks, &rank_counts)
        + score_nobs(hand, cut, &rank_counts)
}

fn score_pairs(rank_counts: &[usize; 15]) -> usize {
    let mut res = 0;
    for r in &rank_counts[1..14] {
        match r {
            2 => res += 2,
            3 => res += 6,
            4 => res += 12,
            _ => {}
        }
    }
    res
}

fn score_runs(all: &[usize; 5], rank_counts: &[usize; 15]) -> usize {
    for r in all {
        if rank_counts[r - 1] > 0 {
            // this is not the beginning of a run
            continue;
        }
        let mut next_up = r + 1;
        let mut len = 1;
        let mut mult = rank_counts[*r];
        while rank_counts[next_up] > 0 {
            mult *= rank_counts[next_up];
            len += 1;
            next_up += 1;
        }
        if len >= 3 {
            return len * mult;
        }
    }
    0
}

fn score_nobs(hand: &[Card; 4], cut: Card, rank_counts: &[usize; 15]) -> usize {
    if cut.rank == 11 || rank_counts[11] == 0 {
        // cut is a jack or we don't have a jack; we can't have nobs
        return 0;
    }
    for c in hand {
        if c.rank == 11 && c.suit == cut.suit {
            return 1;
        }
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::convert::TryInto;

    fn make_hand(cut_str: &str, hand_str: &str) -> (Card, [Card; 4]) {
        (
            Card::from_str(cut_str).expect("invalid card"),
            hand_str
                .split(",")
                .map(|s| Card::from_str(s).expect("invalid card"))
                .collect::<Vec<_>>()
                .try_into()
                .expect("wrong size"),
        )
    }

    #[test]
    fn test_a_bunch_of_hands() {
        let cases: Vec<(&str, &str, usize)> = vec![
            ("5H", "5S,5C,5D,JH", 29),
            ("8H", "8S,8C,8D,10H", 12),
            ("KH", "8S,8C,8D,10H", 6),
            ("KH", "8S,8C,2D,10H", 2),
            ("KH", "KS,8C,2D,8H", 4),
            ("8H", "9S,8C,10D,8H", 15),
            ("JH", "8S,8C,9D,10H", 10),
            ("9H", "8S,8C,9D,10H", 16),
            ("KH", "8S,8C,9D,10H", 8),
            ("KH", "8S,2C,9D,10H", 3),
            ("KH", "8S,JC,9D,10H", 4),
            ("QH", "8S,JC,9D,10H", 5),
            ("6H", "JH,KC,10D,8H", 1),
            ("3H", "8D,4D,10D,6D", 6),
            ("4C", "5H,3D,7D,7S", 9),
            ("9C", "1H,5D,7D,7S", 6),
            ("7H", "6D,6S,10H,9C", 6),
            ("7H", "6D,6S,10H,8C", 10),
        ];
        for (cut_str, hand_str, exp) in cases {
            let (cut, hand) = make_hand(cut_str, hand_str);
            assert_eq!(exp, score_hand(&hand, cut, false))
        }
    }

    #[test]
    fn test_score_runs_and_pairs_and_nobs() {
        let cases: Vec<(&str, &str, usize)> = vec![
            ("AS", "10S,5S,3S,4S", 3),
            ("6S", "9S,7S,8S,QS", 4),
            ("2S", "AS,5S,3S,4S", 5),
            ("AS", "AH,5S,3S,4S", 5),
            ("5S", "2H,3S,4S,2S", 10),
            ("3S", "2H,3C,4S,2S", 16),
            ("10S", "9C,9H,8S,9S", 15),
            ("2C", "AH,2H,2S,AS", 8),
            ("5H", "6S,8S,5C,5S", 6),
            ("3C", "3D,KS,3H,3S", 12),
            ("JC", "2D,4S,6H,8S", 0),
            ("8C", "2D,4S,6H,JC", 1),
            ("10C", "2D,4S,6H,8S", 0),
        ];
        for (cut_str, hand_str, exp) in cases {
            let (cut, hand) = make_hand(cut_str, hand_str);
            assert_eq!(exp, score_runs_and_pairs_and_nobs(&hand, cut))
        }
    }
    #[test]
    fn test_score_flush() {
        let cases: Vec<(&str, &str, usize, usize)> = vec![
            ("AH", "AS,2S,3H,4H", 0, 0),
            ("AH", "AS,2H,3H,4H", 0, 0),
            ("AS", "AH,2H,3H,4H", 4, 0),
            ("5H", "AH,2H,3H,4H", 5, 5),
        ];
        for (cut_str, hand_str, exp_non_crib, exp_crib) in cases {
            let (cut, hand) = make_hand(cut_str, hand_str);
            assert_eq!(exp_non_crib, score_flush(&hand, cut, false));
            assert_eq!(exp_crib, score_flush(&hand, cut, true));
        }
    }
    #[test]
    fn test_score_fifteens() {
        let cases: Vec<(&str, &str, usize)> = vec![
            ("10H", "5S,2S,2H,4H", 2),
            ("10H", "5S,5S,5H,5H", 16),
            ("9H", "6S,6S,9H,5H", 8),
            ("5S", "AS,2S,3S,4H", 2),
        ];
        for (cut_str, hand_str, exp) in cases {
            let (cut, hand) = make_hand(cut_str, hand_str);
            assert_eq!(exp, score_fifteens(&hand, cut));
        }
    }
}
