use clap::Parser;
use std::collections::{HashSet, VecDeque};
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

fn unique(s: &[char]) -> bool {
    let uniq: HashSet<&char> = HashSet::from_iter(s.iter());
    uniq.len() == s.len()
}

// TODO: a lot of this could be cleaned up, but I'm still stuck on that damn Day
// 5 crate parsing
fn part_one(input: String) -> String {
    let mut chs: VecDeque<char> = input.chars().collect();
    let mut seen: Vec<char> = vec![];

    // add first four
    seen.push(chs.pop_front().unwrap());
    seen.push(chs.pop_front().unwrap());
    seen.push(chs.pop_front().unwrap());
    seen.push(chs.pop_front().unwrap());

    loop {
        let last_four = seen.windows(4).last().unwrap();
        if unique(last_four) {
            break;
        }
        seen.push(chs.pop_front().unwrap());
    }

    seen.len().to_string()
}

fn part_two(input: String) -> String {
    let mut chs: VecDeque<char> = input.chars().collect();
    let mut seen: Vec<char> = vec![];

    // add first 14
    for _ in 0..14 {
        seen.push(chs.pop_front().unwrap());
    }

    loop {
        let last_fourteen = seen.windows(14).last().unwrap();
        if unique(last_fourteen) {
            break;
        }
        seen.push(chs.pop_front().unwrap());
    }

    seen.len().to_string()
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
    fn example_p1() {
        let input = read_data_file("example.txt");
        let got = part_one(input);
        let want = "7";
        assert_eq!(want, got);
    }

    #[test]
    fn example_p2() {
        let input = read_data_file("example.txt");
        let got = part_two(input);
        let want = "19";
        assert_eq!(want, got);
    }
}
