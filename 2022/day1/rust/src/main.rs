use clap::Parser;
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
    let in_file_path = PathBuf::from(format!("../data/{}", filename));
    let in_file = fs::canonicalize(&in_file_path).expect("could not canonicalize file");
    fs::read_to_string(in_file)
        .expect("could not read file")
        .to_string()
}

fn cell_sum(cell: &str) -> i64 {
    let nums = cell.split("\n");
    let parsed: Vec<i64> = nums.map(|num| num.parse::<i64>().unwrap()).collect();
    parsed.iter().sum()
}

fn part_one(input: String) -> String {
    let cells = input.split("\n\n");
    let sums = cells.map(|cell| cell_sum(cell));
    let answer = sums.max().unwrap();
    answer.to_string()
}

fn part_two(input: String) -> String {
    let cells = input.split("\n\n");
    let sums = cells.map(|cell| cell_sum(cell));
    let mut sorted: Vec<i64> = sums.collect();
    sorted.sort();
    sorted.reverse();
    let top_three: Vec<i64> = sorted[0..3].into();
    let answer: i64 = top_three.iter().sum();
    answer.to_string()
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
        let want = "24000";
        let got = part_one(input);
        assert_eq!(want, got)
    }

    #[test]
    fn example_p2() {
        let input = read_data_file("example.txt");
        let want = "45000";
        let got = part_two(input);
        assert_eq!(want, got);
    }
}
