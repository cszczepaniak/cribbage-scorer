use cribbage_scorer::cards::Card;
use cribbage_scorer::score;
use itertools::Itertools;

fn main() {
    let mut all_scores = [0; 30];
    for hand_vec in cribbage_scorer::cards::new_deck().iter().combinations(5) {
        let all = [
            *hand_vec[0],
            *hand_vec[1],
            *hand_vec[2],
            *hand_vec[3],
            *hand_vec[4],
        ];
        // manually unwrapping here is probably fastest
        all_scores[score::score_hand(&[all[0], all[1], all[2], all[3]], all[4], false)] += 1;
        all_scores[score::score_hand(&[all[4], all[0], all[1], all[2]], all[3], false)] += 1;
        all_scores[score::score_hand(&[all[3], all[4], all[0], all[1]], all[2], false)] += 1;
        all_scores[score::score_hand(&[all[2], all[3], all[4], all[0]], all[1], false)] += 1;
        all_scores[score::score_hand(&[all[1], all[2], all[3], all[4]], all[0], false)] += 1;
    }
    for (i, occ) in all_scores.iter().enumerate() {
        println!("{}: {}", i, occ);
    }
}

fn hands_from_cards<'a>(h: &'a [Card]) -> Vec<(Card, Vec<&'a Card>)> {
    let mut res = Vec::with_capacity(5);
    for i in 0..h.len() {
        let cut = h[i];
        let mut rest: Vec<&Card> = h.into_iter().take(i).collect();
        let end = h.into_iter().skip(i + 1);
        rest.extend(end);
        res.push((cut, rest));
    }
    res
}
