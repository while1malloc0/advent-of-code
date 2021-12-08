use std::collections::{HashMap, HashSet};

fn main() {
    let p1_input = include_str!("../../inputs/8.txt");
    let p1_answer = p1(p1_input);
    println!("Part 1: {}", p1_answer);
}

fn p1(input: &str) -> u32 {
    let lines = input.trim().split("\n");
    let four_digit_outputs: Vec<&str> =
        lines.map(|s| s.split("|").nth(1).unwrap().trim()).collect();
    let mut count = 0;
    for output in four_digit_outputs {
        count += output
            .split(" ")
            .map(|s| s.chars().collect::<Vec<char>>().len())
            .filter(|x| *x == 2 || *x == 3 || *x == 4 || *x == 7)
            .collect::<Vec<usize>>()
            .len() as u32;
    }
    count
}

fn p2(input: &str) -> u32 {
    panic!("not implemented yet");
}

struct Display {
    signal: Vec<String>,
    output: Vec<String>,
}

impl Display {
    fn uniques(&self) -> Vec<String> {
        let mut output = vec![];
        for val in &self.signal {
            let num_chars = val.chars().collect::<Vec<char>>().len();
            if [2, 3, 4, 7].contains(&num_chars) {
                output.push(val.clone());
            }
        }
        output
    }

    fn fives(&self) -> Vec<String> {
        let mut output = vec![];
        for val in &self.signal {
            let num_chars = val.chars().collect::<Vec<char>>().len();
            if num_chars == 5 {
                output.push(val.clone());
            }
        }
        output
    }

    fn sixes(&self) -> Vec<String> {
        let mut output = vec![];
        for val in &self.signal {
            let num_chars = val.chars().collect::<Vec<char>>().len();
            if num_chars == 6 {
                output.push(val.clone());
            }
        }
        output
    }

    fn solve(&self) -> HashMap<String, String> {
        let unis = self.solve_uniques();
        let mut reverse_unis = HashMap::new();
        for (k, v) in unis.clone() {
            reverse_unis.insert(v, k);
        }
        let fvs = self.solve_fives(reverse_unis);

        let mut output = unis.clone();
        output.extend(fvs);
        output
    }

    fn solve_uniques(&self) -> HashMap<String, String> {
        let mut stringified: HashMap<i32, &str> = HashMap::new();
        stringified.insert(2, "1");
        stringified.insert(3, "7");
        stringified.insert(4, "4");
        stringified.insert(7, "8");

        let mut output: HashMap<String, String> = HashMap::new();

        for val in self.uniques() {
            let num_chars = val.chars().collect::<Vec<char>>().len();
            let ins = stringified
                .get(&(num_chars as i32))
                .expect("non-unique found in uniques");
            output.insert(val.to_string(), ins.to_string());
        }

        output
    }

    fn solve_fives(&self, uniques: HashMap<String, String>) -> HashMap<String, String> {
        let mut output: HashMap<String, String> = HashMap::new();

        for val in self.fives() {
            let mut candidates: HashSet<String> = HashSet::new();
            candidates.insert("2".into());
            candidates.insert("3".into());
            candidates.insert("5".into());
            loop {
                if candidates.len() == 1 {
                    let remaining = candidates.clone().into_iter().collect::<Vec<String>>();
                    output.insert(val, remaining[0].to_string());
                    break;
                }

                let mut current_val_set: HashSet<String> = HashSet::new();
                let current_val_chars: Vec<String> = val.chars().map(|c| c.to_string()).collect();
                for c in current_val_chars {
                    current_val_set.insert(c);
                }

                let one_chars: Vec<String> = uniques
                    .get("1")
                    .unwrap()
                    .chars()
                    .map(|c| c.to_string())
                    .collect();
                let mut one_set: HashSet<String> = HashSet::new();
                for c in one_chars {
                    one_set.insert(c);
                }
                let num_one_overlaps = one_set
                    .intersection(&current_val_set)
                    .collect::<Vec<&String>>()
                    .len();
                if num_one_overlaps == 2 {
                    candidates.remove("2");
                    candidates.remove("5");
                } else if num_one_overlaps == 1 {
                    candidates.remove("3");
                } else {
                    unreachable!();
                }

                // do overlaps for 4 and 7

                break;
            }
        }

        output
    }
}

impl From<&str> for Display {
    fn from(input: &str) -> Self {
        let splitted: Vec<&str> = input.trim().split("|").collect();
        let signal: Vec<String> = splitted[0]
            .trim()
            .split(" ")
            .map(|s| s.trim().to_string())
            .collect();
        let output: Vec<String> = splitted[1]
            .trim()
            .split(" ")
            .map(|s| s.trim().to_string())
            .collect();
        Self { signal, output }
    }
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let input = include_str!("../../inputs/8.example.txt");
        let got = p1(input);
        let want = 26;
        assert_eq!(want, got);
    }

    #[ignore]
    #[test]
    fn p2_e2e() {
        let input = include_str!("../../inputs/8.example.txt");
        let got = p2(input);
        let want = 61229;
        assert_eq!(want, got);
    }

    #[test]
    fn parse_dislay() {
        let input = "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe";
        let got: Display = input.into();
        assert_eq!(
            got.signal,
            vec![
                "be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd",
                "edb"
            ]
            .into_iter()
            .map(str::to_string)
            .collect::<Vec<String>>()
        );
        assert_eq!(
            got.output,
            vec!["fdgacbe", "cefdb", "cefbgd", "gcbe"]
                .into_iter()
                .map(str::to_string)
                .collect::<Vec<String>>()
        );
    }

    #[test]
    fn uniques() {
        let subject: Display = "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe".into();
        let got = subject.uniques();
        let want = vec!["be", "cfbegad", "cgeb", "edb"];
        assert_eq!(got, want);
    }

    #[test]
    fn fives() {
        let subject: Display = "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe".into();
        let got = subject.fives();
        let want = vec!["fdcge", "fecdb", "fabcd"];
        assert_eq!(got, want);
    }

    #[test]
    fn sixes() {
        let subject: Display = "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe".into();
        let got = subject.sixes();
        let want = vec!["cbdgef", "fgaecd", "agebfd"];
        assert_eq!(got, want);
    }

    #[test]
    fn solve() {
        let subject: Display = "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe".into();
        let got = subject.solve();
        let mut want: HashMap<String, String> = HashMap::new();
        want.insert("be".into(), "1".into());
        want.insert("cfbegad".into(), "8".into());
        want.insert("cbdgef".into(), "9".into());
        want.insert("fgaecd".into(), "6".into());
        want.insert("cgeb".into(), "4".into());
        want.insert("fecdb".into(), "3".into());
        want.insert("fabcd".into(), "2".into());
        want.insert("edb".into(), "7".into());
        assert_eq!(got, want);
    }
}
