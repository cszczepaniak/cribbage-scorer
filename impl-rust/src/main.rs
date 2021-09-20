use cribbage_scorer::score;
use crossbeam_utils::thread;
use itertools::Itertools;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::Arc;

fn main() {
    let mut all_scores: [AtomicUsize; 30] = Default::default();
    for i in 0..30 {
        all_scores[i] = AtomicUsize::new(0);
    }
    let all_scores = Arc::new(all_scores);
    let deck = cribbage_scorer::cards::new_deck();
    let combs = deck.iter().combinations(5).collect::<Vec<_>>();
    let num_workers = num_cpus::get();
    let chunk_size = (combs.len() + num_workers - 1) / num_workers;

    thread::scope(|scope| {
        for ch in combs.chunks(chunk_size) {
            let all_scores_clone = all_scores.clone();
            let ch_clone = ch.clone();
            scope.spawn(move |_| {
                for hand_vec in ch_clone {
                    let all = [
                        *hand_vec[0],
                        *hand_vec[1],
                        *hand_vec[2],
                        *hand_vec[3],
                        *hand_vec[4],
                    ];
                    // manually unwrapping here is probably fastest
                    all_scores_clone
                        [score::score_hand(&[all[0], all[1], all[2], all[3]], all[4], false)]
                    .fetch_add(1, Ordering::Relaxed);
                    all_scores_clone
                        [score::score_hand(&[all[4], all[0], all[1], all[2]], all[3], false)]
                    .fetch_add(1, Ordering::Relaxed);
                    all_scores_clone
                        [score::score_hand(&[all[3], all[4], all[0], all[1]], all[2], false)]
                    .fetch_add(1, Ordering::Relaxed);
                    all_scores_clone
                        [score::score_hand(&[all[2], all[3], all[4], all[0]], all[1], false)]
                    .fetch_add(1, Ordering::Relaxed);
                    all_scores_clone
                        [score::score_hand(&[all[1], all[2], all[3], all[4]], all[0], false)]
                    .fetch_add(1, Ordering::Relaxed);
                }
            });
        }
    })
    .unwrap();
    for (i, occ) in all_scores.iter().enumerate() {
        println!("{:02}: {}", i, occ.load(Ordering::Relaxed));
    }
}
