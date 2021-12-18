use std::collections::HashMap;

fn main() {
    let p1_input = include_str!("../../inputs/10.txt");
    let p1_answer = p1(p1_input);
    println!("Part 1: {}", p1_answer);

    let p2_input = include_str!("../../inputs/10.txt");
    let p2_answer = p2(p2_input);
    println!("Part 1: {}", p2_answer);
}

fn p1(input: &str) -> u32 {
    // parse into lines
    let lines = input.lines();
    // find errant character
    let mut score = 0;
    // add to score
    for line in lines {
        let mut stack: Vec<char> = vec![];
        for c in line.chars() {
            if is_opener(c) {
                stack.push(c);
            } else {
                let opener = stack.pop().unwrap();
                if let Some(errant) = expect(&opener, &c) {
                    score += score_for(errant);
                }
            }
        }
    }
    // return score
    score
}

fn p2(input: &str) -> u64 {
    // extract incomplete lines
    let lines = input.lines().collect();
    let incompletes = extract_incompletes(lines);
    // get completion string for each
    let mut completions: Vec<String> = vec![];
    for line in incompletes {
        completions.push(completion_string_for(line));
    }
    // score completion strings
    let mut scores: Vec<u64> = vec![];
    for completion in completions {
        scores.push(ac_score_for(completion));
    }
    // figure out "middle" score
    let result = middle_of(scores);
    result
}

#[inline]
fn is_opener(c: char) -> bool {
    ['(', '{', '<', '['].contains(&c)
}

#[inline]
fn expect(lhs: &char, rhs: &char) -> Option<char> {
    let mut opener_and_closer: HashMap<char, char> = HashMap::new();
    opener_and_closer.insert('{', '}');
    opener_and_closer.insert('(', ')');
    opener_and_closer.insert('[', ']');
    opener_and_closer.insert('<', '>');

    let expected = opener_and_closer
        .get(lhs)
        .expect("unknown char given for left hand side of expectation");
    if *expected != *rhs {
        return Some(rhs.clone());
    }
    None
}

#[inline]
fn score_for(c: char) -> u32 {
    let mut scores: HashMap<char, u32> = HashMap::new();
    scores.insert(')', 3);
    scores.insert(']', 57);
    scores.insert('}', 1197);
    scores.insert('>', 25137);

    *scores.get(&c).unwrap()
}

#[inline]
fn extract_incompletes(input: Vec<&str>) -> Vec<&str> {
    let mut result = vec![];
    'outer: for line in input {
        let mut stack: Vec<char> = vec![];
        for c in line.chars() {
            if is_opener(c) {
                stack.push(c);
            } else {
                let opener = stack.pop().unwrap();
                if let Some(_) = expect(&opener, &c) {
                    continue 'outer;
                }
            }
        }
        result.push(line);
    }
    result
}

// were this production code, line extraction and autocompletion should probably
// be done in the same loop as syntax error checking, but it's not, so...
#[inline]
fn completion_string_for(input: &str) -> String {
    let mut stack: Vec<char> = vec![];
    for c in input.chars() {
        if is_opener(c) {
            stack.push(c);
        } else {
            // again, this should be a syntax error check, but I split it into
            // two functions not knowing what p2 would be. Oh well...
            stack.pop().unwrap();
        }
    }

    // at this point, the stack has all of our incomplete items
    let mut completions: HashMap<char, char> = HashMap::new();
    completions.insert('{', '}');
    completions.insert('(', ')');
    completions.insert('[', ']');
    completions.insert('<', '>');
    let mut result: String = "".into();
    loop {
        if let Some(c) = stack.pop() {
            result.push_str(&completions.get(&c).unwrap().to_string());
        } else {
            break;
        }
    }
    result
}

#[inline]
fn ac_score_for(input: String) -> u64 {
    let points = vec![')', ']', '}', '>'];
    let mut score: u64 = 0;
    for c in input.chars() {
        score *= 5;
        score += (points.iter().position(|r| *r == c).unwrap() + 1) as u64
    }
    score
}

#[inline]
fn middle_of(input: Vec<u64>) -> u64 {
    let mut tmp = input.clone();
    tmp.sort();
    let idx = ((tmp.len() / 2) as f32).floor() as usize;
    tmp[idx]
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let input = include_str!("../../inputs/10.example.txt");
        let got = p1(input);
        let want = 26397;
        assert_eq!(got, want);
    }

    #[test]
    fn p2_e2e() {
        let input = include_str!("../../inputs/10.example.txt");
        let got = p2(input);
        let want = 288957;
        assert_eq!(got, want);
    }

    #[test]
    fn test_is_opener() {
        assert_eq!(true, is_opener('('));
        assert_eq!(true, is_opener('{'));
        assert_eq!(true, is_opener('['));
        assert_eq!(true, is_opener('<'));

        assert_eq!(false, is_opener(')'));
        assert_eq!(false, is_opener('}'));
        assert_eq!(false, is_opener(']'));
        assert_eq!(false, is_opener('>'));
    }

    #[test]
    fn test_expect() {
        assert_eq!(expect(&'{', &'}'), None);
        assert_eq!(expect(&'{', &')'), Some(')'));
    }

    #[test]
    fn test_score_for() {
        assert_eq!(score_for(')'), 3);
        assert_eq!(score_for(']'), 57);
        assert_eq!(score_for('}'), 1197);
        assert_eq!(score_for('>'), 25137);
    }

    #[test]
    fn test_extract_incompletes() {
        let input: Vec<&str> = include_str!("../../inputs/10.example.txt")
            .lines()
            .collect();
        let want: Vec<&str> = vec![
            "[({(<(())[]>[[{[]{<()<>>",
            "[(()[<>])]({[<{<<[]>>(",
            "(((({<>}<{<{<>}{[]{[]{}",
            "{<[[]]>}<{[{[{[]{()[[[]",
            "<{([{{}}[<[[[<>{}]]]>[]]",
        ];
        let got = extract_incompletes(input);
        assert_eq!(got, want);
    }

    #[test]
    fn test_completion_string_for() {
        assert_eq!(
            completion_string_for("[({(<(())[]>[[{[]{<()<>>"),
            "}}]])})]"
        );
        assert_eq!(completion_string_for("[(()[<>])]({[<{<<[]>>("), ")}>]})");
        assert_eq!(
            completion_string_for("(((({<>}<{<{<>}{[]{[]{}"),
            "}}>}>))))"
        );
        assert_eq!(
            completion_string_for("{<[[]]>}<{[{[{[]{()[[[]"),
            "]]}}]}]}>"
        );
        assert_eq!(completion_string_for("<{([{{}}[<[[[<>{}]]]>[]]"), "])}>");
    }

    #[test]
    fn test_ac_score_for() {
        assert_eq!(ac_score_for("}}]])})]".into()), 288957);
        assert_eq!(ac_score_for(")}>]})".into()), 5566);
        assert_eq!(ac_score_for("}}>}>))))".into()), 1480781);
        assert_eq!(ac_score_for("]]}}]}]}>".into()), 995444);
        assert_eq!(ac_score_for("])}>".into()), 294);
    }

    #[test]
    fn test_middle_num() {
        let input = vec![288957, 5566, 1480781, 995444, 294];
        assert_eq!(middle_of(input), 288957);
    }
}
