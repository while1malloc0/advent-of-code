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
    let in_file_path = PathBuf::from(format!("./data/{}", filename));
    let in_file = fs::canonicalize(&in_file_path).expect("could not canonicalize file");
    fs::read_to_string(in_file)
        .expect("could not read file")
        .to_string()
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

    let input = read_data_file(in_file_name);

    if cli.part == 1 {
        println!("{}", part_one(input))
    } else {
        println!("{}", part_two(input))
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn example_p1() {
        let input = read_data_file("example.txt");
        let got = part_one(input);
        let want = "<replace me>";
        assert_eq!(want, got);
    }

    #[test]
    #[ignore]
    fn example_p2() {
        let input = read_data_file("example.txt");
        let got = part_two(input);
        let want = "<replace me>";
        assert_eq!(want, got);
    }
}
