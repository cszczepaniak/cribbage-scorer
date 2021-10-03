# Cribbage Scorer
A collection of automatic hand scorers for the card game of cribbage.

My initial goal was to be able to score every cribbage hand. Now this is a good place for me to learn new languages and understand how well my solutions are performing against others.

## Results
<details>
  <summary>See the latest results</summary>

```
===========================================================================
impl-go

real	0m0.331s
user	0m0.000s
sys	0m0.015s

===========================================================================
impl-rust

real	0m3.057s
user	0m0.000s
sys	0m0.015s
```
  
Caveat: I'm sure Rust could be as fast or faster than Go -- I'm just not well-versed enough in it to squeeze out all of that performance!
</details>


## Running All Solutions
To run all of the solutions, run
```bash
./run.sh
```
The solutions are benchmarked in a VERY rudimentary way (literally just using `time`), and the results will be written to `results.txt`.

## Adding a Solution
A solution must fulfill the following requirements:
1. It must be in a folder with the prefix `impl-`
1. It must contain an executable bash script called `build.sh` which build the executable and puts it in the folder `./bin/impl-*` (relative to your implementation's directory, and with the executable having the same name as this directory)

If these two things are satisfied, `run.sh` will know what to do with your solution.
