use std::fs;
use std::path::PathBuf;
use clap::Parser;

#[derive(Parser)]
#[command()]
struct CLI {
    #[arg(value_parser = clap::value_parser!(u16).range(1..2))]
    part: u16,
    #[arg(short, long)]
    example: bool,
}

fn part_one(input: String) -> String {
    input
}

fn part_two(input: String) -> String {
    input
}

fn main() {
    let cli = CLI::parse();

    let mut in_file_name = "input.txt";
    if cli.example {
        in_file_name = "example.txt";
    }

    let in_file_path = PathBuf::from(format!("../data/{}", in_file_name));
    let in_file = fs::canonicalize(&in_file_path).expect("could not canonicalize file");
    let input = fs::read_to_string(in_file).expect("could not read file").to_string();

    if cli.part == 1 {
        println!("{}", part_one(input))
    } else {
        println!("{}", part_two(input))
    }
}
