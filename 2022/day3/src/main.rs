use clap::Parser;
use std::collections::HashSet;
use std::fs;
use std::path::PathBuf;

#[derive(Parser)]
#[command()]
struct CLI {
    #[arg(value_parser = clap::value_parser!(u16).range(1..3))]
    part: u16,
    #[arg(short, long)]
    example: bool,
}

fn read_data_file(filename: &str) -> String {
    let in_file_path = PathBuf::from(format!("./data/{}", filename));
    let in_file = fs::canonicalize(&in_file_path).expect("could not canonicalize file");
    fs::read_to_string(in_file)
        .expect("could not read file")
        .to_string()
}

fn calc_intersection(lhs: &String, rhs: &String) -> Option<char> {
    let mut lh_hash: HashSet<char> = HashSet::new();
    for c in lhs.chars() {
        lh_hash.insert(c);
    }

    let mut rh_hash: HashSet<char> = HashSet::new();
    for c in rhs.chars() {
        rh_hash.insert(c);
    }

    let inter = lh_hash.intersection(&rh_hash).collect::<Vec<&char>>();

    if inter.is_empty() {
        None
    } else {
        Some(inter[0].clone())
    }
}

fn calc_intersection_three(fst: &String, snd: &String, trd: &String) -> Option<char> {
    let mut fst_hash: HashSet<char> = HashSet::new();
    for c in fst.chars() {
        fst_hash.insert(c);
    }

    let mut snd_hash: HashSet<char> = HashSet::new();
    for c in snd.chars() {
        snd_hash.insert(c);
    }

    let mut trd_hash: HashSet<char> = HashSet::new();
    for c in trd.chars() {
        trd_hash.insert(c);
    }

    let mut inter = &fst_hash & &snd_hash;
    inter = &inter & &trd_hash;

    if inter.is_empty() {
        None
    } else {
        Some(inter.iter().collect::<Vec<&char>>()[0].clone())
    }
}

fn convert(input: Option<char>) -> usize {
    match input {
        None => 0,
        Some(subject) => {
            let lowers: Vec<char> = ('a'..='z').collect();
            let uppers: Vec<char> = ('A'..='Z').collect();
            let both = [lowers, uppers].concat();
            let result = both.iter().position(|c| c == &subject).unwrap() + 1;
            result
        }
    }
}

fn part_one(input: String) -> String {
    let mut result = 0;
    for line in input.lines() {
        let size = line.len() / 2;
        let left_half: String = line[0..size].into();
        let right_half: String = line[size..].into();
        let intersection = calc_intersection(&left_half, &right_half);
        result += convert(intersection);
    }
    format!("{}", result)
}

fn part_two(input: String) -> String {
    let mut result = 0;
    let groups = input.lines().map(|s| s.to_owned()).collect::<Vec<String>>();
    let gps = groups.chunks(3);
    for g in gps {
        let inter = calc_intersection_three(&g[0], &g[1], &g[2]);
        result += convert(inter);
    }
    format!("{}", result)
}

fn main() {
    let cli = CLI::parse();

    let mut in_file_name = "input.txt";
    if cli.example {
        in_file_name = "example.txt";
    }

    let input = read_data_file(in_file_name);

    if cli.part == 1 {
        println!("{}", part_one(input))
    } else {
        println!("{}", part_two(input))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example_p1() {
        let input = read_data_file("example.txt");
        let got = part_one(input);
        let want = "157";
        assert_eq!(got, want);
    }

    #[test]
    fn test_example_p2() {
        let input = read_data_file("example.txt");
        let got = part_two(input);
        let want = "70";
        assert_eq!(got, want);
    }
}
