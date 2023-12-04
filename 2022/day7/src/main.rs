use clap::Parser;
use std::collections::{HashMap, HashSet};
use std::fs;
use std::path::{Path, PathBuf};

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

#[derive(PartialEq, Clone, Debug)]
enum TerminalParseState {
    LS,
    Command,
}

#[derive(PartialEq, Clone, Debug)]
struct TerminalState {
    pwd: String,
    dirs: HashSet<String>,
    fs: HashMap<&'static str, Vec<&'static str>>,
    parse_state: TerminalParseState,
}

impl TerminalState {
    fn new() -> Self {
        let pwd = "/".to_string();
        let dirs = HashSet::from(["/".to_string()]);
        let fs = HashMap::from([("/", vec![])]);
        let parse_state = TerminalParseState::Command;
        TerminalState {
            pwd,
            dirs,
            fs,
            parse_state,
        }
    }
}

fn execute(instruction: &str, state: &TerminalState) -> TerminalState {
    let mut state = state.clone();
    let parts: Vec<&str> = instruction.split(" ").map(|s| s.trim()).collect();
    let mode = parts[0];
    if mode == "$" {
        state.parse_state = TerminalParseState::Command;
    }

    match state.parse_state {
        TerminalParseState::Command => {
            let cmd = parts[1];
            if cmd == "cd" {
                let target = parts[2].to_owned();
                state.pwd = state.pwd + &target.to_owned();
            } else if cmd == "ls" {
                state.parse_state = TerminalParseState::LS;
            }
        }
        TerminalParseState::LS => {
            if mode == "dir" {
                let dir = state.pwd + &parts[1].to_owned();
                state.dirs.insert(dir);
            }
        }
    }

    state
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
    use super::*;

    #[test]
    fn example_p1() {
        let input = read_data_file("example.txt");
        let got = part_one(input);
        let want = "95437";
        assert_eq!(want, got);
    }

    #[test]
    fn parse_cd() {
        let input = "$ cd a";
        let starting = TerminalState::new();

        let want = TerminalState {
            pwd: "/a".to_string(),
            dirs: HashSet::new(),
            fs: HashMap::new(),
            parse_state: TerminalParseState::Command,
        };

        let got = execute(input, &starting);
        assert_eq!(want.pwd, got.pwd);
        assert_eq!(want.parse_state, got.parse_state);
    }

    #[test]
    fn parse_ls() {
        let input = "$ ls";
        let starting = TerminalState::new();

        let want = TerminalState {
            pwd: "/".to_string(),
            dirs: HashSet::new(),
            fs: HashMap::new(),
            parse_state: TerminalParseState::LS,
        };

        let got = execute(input, &starting);
        assert_eq!(want.parse_state, got.parse_state);
    }

    #[test]
    fn ls_dir() {
        let input = "dir a";
        let mut starting = TerminalState::new();
        starting.parse_state = TerminalParseState::LS;

        let want = TerminalState {
            pwd: "/".to_string(),
            dirs: HashSet::from(["/a".to_owned()]),
            fs: HashMap::new(),
            parse_state: TerminalParseState::LS,
        };

        let got = execute(input, &starting);
        assert_eq!(want.parse_state, got.parse_state);
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
