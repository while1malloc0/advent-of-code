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

enum Shape {
    Rock,
    Paper,
    Scissors,
}

impl From<&str> for Shape {
    fn from(s: &str) -> Self {
        match s {
            "A" => Shape::Rock,
            "B" => Shape::Paper,
            "C" => Shape::Scissors,
            "X" => Shape::Rock,
            "Y" => Shape::Paper,
            "Z" => Shape::Scissors,
            _ => panic!("illegal char for game"),
        }
    }
}

impl Shape {
    fn points(self) -> u64 {
        match self {
            Shape::Rock => 1,
            Shape::Paper => 2,
            Shape::Scissors => 3,
        }
    }
}

enum MatchResult {
    Win,
    Lose,
    Draw,
}

impl From<&str> for MatchResult {
    fn from(s: &str) -> Self {
        match s {
            "X" => MatchResult::Lose,
            "Y" => MatchResult::Draw,
            "Z" => MatchResult::Win,
            _ => panic!("illegal char in match result parsing"),
        }
    }
}

impl MatchResult {
    fn points(self) -> u64 {
        match self {
            MatchResult::Win => 6,
            MatchResult::Lose => 0,
            MatchResult::Draw => 3,
        }
    }
}

struct Match {
    ours: Shape,
    theirs: Shape,
}

impl Match {
    fn score(self) -> u64 {
        let result;
        match self.ours {
            Shape::Rock => match self.theirs {
                Shape::Rock => result = MatchResult::Draw,
                Shape::Paper => result = MatchResult::Lose,
                Shape::Scissors => result = MatchResult::Win,
            },
            Shape::Paper => match self.theirs {
                Shape::Rock => result = MatchResult::Win,
                Shape::Paper => result = MatchResult::Draw,
                Shape::Scissors => result = MatchResult::Lose,
            },
            Shape::Scissors => match self.theirs {
                Shape::Rock => result = MatchResult::Lose,
                Shape::Paper => result = MatchResult::Win,
                Shape::Scissors => result = MatchResult::Draw,
            },
        }

        self.ours.points() + result.points()
    }
}

fn shape_from_result(shape: &Shape, result: MatchResult) -> Shape {
    match result {
        MatchResult::Win => match shape {
            Shape::Rock => Shape::Paper,
            Shape::Paper => Shape::Scissors,
            Shape::Scissors => Shape::Rock,
        },
        MatchResult::Lose => match shape {
            Shape::Rock => Shape::Scissors,
            Shape::Paper => Shape::Rock,
            Shape::Scissors => Shape::Paper,
        },
        MatchResult::Draw => match shape {
            Shape::Rock => Shape::Rock,
            Shape::Paper => Shape::Paper,
            Shape::Scissors => Shape::Scissors,
        },
    }
}

struct Game {
    matches: Vec<Match>,
}

struct StrategyBothMoves(String);
struct StrategyMovePlusResult(String);

impl From<StrategyBothMoves> for Game {
    fn from(input: StrategyBothMoves) -> Self {
        let mut matches = vec![];
        for line in input.0.lines() {
            let parts: Vec<&str> = line.split(" ").collect();
            let theirs: Shape = parts[0].into();
            let ours: Shape = parts[1].into();
            matches.push(Match { ours, theirs })
        }
        Game { matches }
    }
}

impl From<StrategyMovePlusResult> for Game {
    fn from(input: StrategyMovePlusResult) -> Self {
        let mut matches = vec![];
        for line in input.0.lines() {
            let parts: Vec<&str> = line.split(" ").collect();
            let theirs: Shape = parts[0].into();
            let desired_result: MatchResult = parts[1].into();
            let ours = shape_from_result(&theirs, desired_result);
            matches.push(Match { ours, theirs })
        }
        Game { matches }
    }
}

impl Game {
    fn play(self) -> u64 {
        let mut total = 0;
        for m in self.matches {
            total += m.score();
        }
        total
    }
}

fn part_one(input: String) -> String {
    let strategy = StrategyBothMoves(input);
    let game = Game::from(strategy);
    let result = game.play();
    format!("{}", result)
}

fn part_two(input: String) -> String {
    let strategy = StrategyMovePlusResult(input);
    let game = Game::from(strategy);
    let result = game.play();
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
        let want = "15";
        assert_eq!(want, got);
    }

    #[test]
    fn test_example_p2() {
        let input = read_data_file("example.txt");
        let got = part_two(input);
        let want = "12";
        assert_eq!(want, got);
    }
}
