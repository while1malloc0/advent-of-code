use std::fs::read_to_string;
use std::collections::HashSet;

fn main() {
    let mut input: Vec<i32> = Vec::new();
    if let Ok(f) = read_to_string("../input") {
        for line in f.split("\n") {
            let result = line.parse::<i32>().unwrap();
            input.push(result);
        }
    }

    let mut seen: HashSet<i32> = HashSet::new();
    seen.insert(input[0]);
    for i in 1..input.len() {
        let target = 2020 - input[i];
        if seen.contains(&target) {
            println!("{}", input[i]*target);
            break;
        }
        seen.insert(input[i]);
    }
}
