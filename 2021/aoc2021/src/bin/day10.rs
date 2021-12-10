use std::collections::HashMap;

fn main() {
    let p1_input = include_str!("../../inputs/10.txt");
    let p1_answer = p1(p1_input);
    println!("Part 1: {}", p1_answer);
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
}
