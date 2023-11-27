use clap::Parser;
use std::collections::VecDeque;
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

#[derive(PartialEq, Debug, Clone, Copy)]
struct Instruction {
    num: usize,
    from: usize,
    to: usize,
}

impl TryFrom<&str> for Instruction {
    type Error = &'static str;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        let digits: Vec<&str> = value
            .split(" ")
            .filter(|s| s != &"move" && s != &"from" && s != &"to")
            .collect();
        if digits.len() != 3 {
            return Err("could not parse string");
        }
        let converted: Vec<usize> = digits.iter().map(|s| s.parse().unwrap()).collect();

        Ok(Instruction {
            num: converted[0],
            from: converted[1],
            to: converted[2],
        })
    }
}

type Stacks = VecDeque<Vec<char>>;

fn parse_stacks(input: &str) -> Stacks {
    let mut lns = input.lines().rev().collect::<VecDeque<&str>>();
    let mut result: Stacks = VecDeque::new();
    let stack_len = lns
        .pop_front()
        .unwrap()
        .chars()
        .filter(|c| c.is_numeric())
        .collect::<Vec<char>>()
        .len();
    for _ in 0..stack_len {
        let empty: Vec<char> = vec![];
        result.push_back(empty);
    }
    for l in lns {
        if l == "" {
            continue;
        }
        let mut chunks: VecDeque<char> = l.chars().collect();
        for i in 0..stack_len {
            // consume [
            chunks.pop_front();
            // take char
            let c = chunks.pop_front().unwrap();
            if !c.eq(&' ') {
                result[i].push(c);
            }
            // consume ]
            chunks.pop_front();
            // consume space
            chunks.pop_front();
        }
    }
    result
}

fn top_of_stacks(stacks: Stacks) -> String {
    let mut raw: Vec<char> = vec![];
    for mut s in stacks {
        if let Some(c) = s.pop() {
            raw.push(c);
        }
    }
    raw.iter().collect()
}

fn consume_instruction(state: Stacks, instruction: Instruction) -> Stacks {
    let mut result: Stacks = state.clone();
    for _ in 0..instruction.num {
        if let Some(item) = result[instruction.from - 1].pop() {
            result[instruction.to - 1].push(item);
        }
    }
    result
}

type InstructionStack = Vec<Instruction>;

fn parse_instruction_stack(input: &str) -> InstructionStack {
    let mut result: InstructionStack = vec![];
    let lns: Vec<&str> = input.trim().lines().rev().collect();
    for l in lns {
        let parsed: Instruction = l.try_into().unwrap();
        result.push(parsed);
    }
    result
}

fn part_one(input: String) -> String {
    let parts: Vec<&str> = input.split("\n\n").collect();
    if parts.len() != 2 {
        panic!("problem parsing file");
    }
    let instructions = parse_instruction_stack(parts[1]);
    let mut stacks = parse_stacks(parts[0]);
    for ins in instructions {
        stacks = consume_instruction(stacks, ins);
    }
    format!("{}", top_of_stacks(stacks))
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
        let want = "CMZ";
        assert_eq!(want, got);
    }

    #[test]
    fn consume_instruction() {
        let input = Instruction {
            num: 1,
            from: 2,
            to: 1,
        };
        let stacks = VecDeque::from([vec!['Z', 'N'], vec!['M', 'C', 'D'], vec!['P']]);
        let want = VecDeque::from([vec!['Z', 'N', 'D'], vec!['M', 'C'], vec!['P']]);
        let got = super::consume_instruction(stacks, input);
        assert_eq!(want, got);
    }

    #[test]
    fn parse_instruction_stack() {
        let input = "
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";
        let parsed: InstructionStack = super::parse_instruction_stack(input);
        let got = parsed.last().unwrap().to_owned();
        let want = Instruction {
            num: 1,
            from: 2,
            to: 1,
        };
        assert_eq!(want, got);
    }

    #[test]
    fn parse_stacks() {
        let input = "
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3";
        let parsed = super::parse_stacks(input);
        let got = &parsed[0];
        let want = &vec!['Z', 'N'];
        assert_eq!(want, got);
    }
}
