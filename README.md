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

### Conclusion
Rust is touted as being the paramount systems programming language. Granted I am not well versed with it, but I spent a _long_ time (much more than for the Go solution including all the optimization time) coming up with a solution that is 10x slower than the Go solution.

With that in mind, I personally think it's reasonable to say that Go's productivity increase is worth so much more than the correctness and performance that Rust guarantees. I'm sure there are cases where Go's garbage collector would get in the way, but this is not one of them. But, this is a sample size of 1 person and 1 problem. They are different hammers for different jobs, and in all reality should not be compared as they so often seem to be.

Things I ended up loving about Rust:
- `match` is incredibly powerful and packs a lot of expressiveness into very few lines of code
- Algebraic datatypes (Rust's `enum`) are incredible (especially paired with `match`)
- Writing Rust _forces_ you to know about your memory (is it on the stack or the heap? what owns this memory?). I think it's a good thing for all programmers to think about, rather than just allocating without understanding what's really going on.
- `cargo` seems to have a better story for dependency management than gomod does
    
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
