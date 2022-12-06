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

#[derive(Debug, PartialEq, Copy, Clone)]
struct SectionAssignment {
    start: u16,
    end: u16,
}

impl TryFrom<&str> for SectionAssignment {
    type Error = &'static str;

    fn try_from(s: &str) -> Result<Self, Self::Error> {
        let parts: Vec<&str> = s.split("-").collect();
        if parts.len() != 2 {
            return Err("could not parse string");
        }

        let start: u16 = parts[0].parse().unwrap();
        let end: u16 = parts[1].parse().unwrap();

        Ok(SectionAssignment { start, end })
    }
}

impl SectionAssignment {
    fn contains(self, other: &SectionAssignment) -> bool {
        self.start <= other.start && self.end >= other.end
    }

    fn overlaps_with(self, other: &SectionAssignment) -> bool {
        (self.start <= other.start && self.end >= other.start)
            || (self.start <= other.end && self.end >= other.end)
    }
}

fn line_contains_overlap(input: &str) -> bool {
    let sections: Vec<SectionAssignment> = input
        .split(",")
        .map(|s| SectionAssignment::try_from(s).unwrap())
        .collect();

    sections[0].contains(&sections[1]) || sections[1].contains(&sections[0])
}

fn line_contains_any_overlap(input: &str) -> bool {
    let sections: Vec<SectionAssignment> = input
        .split(",")
        .map(|s| SectionAssignment::try_from(s).unwrap())
        .collect();

    sections[0].overlaps_with(&sections[1]) || sections[1].overlaps_with(&sections[0])
}

fn part_one(input: String) -> String {
    let result = input
        .lines()
        .filter(|s| line_contains_overlap(s))
        .collect::<Vec<&str>>()
        .len();
    format!("{}", result)
}

fn part_two(input: String) -> String {
    let result = input
        .lines()
        .filter(|s| line_contains_any_overlap(s))
        .collect::<Vec<&str>>()
        .len();
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
    fn example_p1() {
        let input = read_data_file("example.txt");
        let got = part_one(input);
        let want = "2";
        assert_eq!(got, want);
    }

    #[test]
    fn parse_section_assignment() {
        let got: SectionAssignment = "2-4".try_into().unwrap();
        let want = SectionAssignment { start: 2, end: 4 };
        assert_eq!(got, want);
    }

    #[test]
    fn section_assignment_contains_other() {
        let lhs = SectionAssignment { start: 2, end: 4 };
        let rhs = SectionAssignment { start: 1, end: 5 };
        let got = lhs.contains(&rhs);
        let want = false;
        assert_eq!(want, got);
    }

    #[test]
    fn line_contains_overlap__no_overlap() {
        let input = "2-4,6-8";
        let got = line_contains_overlap(input);
        let want = false;
        assert_eq!(want, got);
    }

    #[test]
    fn line_contains_overlap__has_overlap() {
        let input = "2-8,3-7";
        let got = line_contains_overlap(input);
        let want = true;
        assert_eq!(want, got);
    }

    #[test]
    fn line_contains_overlap__has_overlap_reverse() {
        let input = "6-6,4-6";
        let got = line_contains_overlap(input);
        let want = true;
        assert_eq!(want, got);
    }

    #[test]
    fn example_p2() {
        let input = read_data_file("example.txt");
        let got = part_two(input);
        let want = "4";
        assert_eq!(got, want);
    }

    #[test]
    fn assignment_overlaps() {
        let lhs = SectionAssignment { start: 5, end: 7 };
        let rhs = SectionAssignment { start: 7, end: 9 };
        let got = lhs.overlaps_with(&rhs);
        let want = true;
        assert_eq!(want, got);
    }

    #[test]
    fn line_contains_any_overlap() {
        let input = "5-7,7-9";
        let got = super::line_contains_any_overlap(input);
        let want = true;
        assert_eq!(want, got);
    }
}
