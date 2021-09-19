use cribbage_scorer::cards::Card;
use itertools::Itertools;

fn main() {
    let mut num = 0;
    for hand in cribbage_scorer::cards::new_deck()
        .into_iter()
        .combinations(5)
    {
        for _ in hands_from_cards(&hand) {
            num += 1
        }
    }
    println!("{:?}", num);
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
