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

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let input = include_str!("../../inputs/8.example.txt");
        let got = p1(input);
        let want = 26;
        assert_eq!(want, got);
    }
}
