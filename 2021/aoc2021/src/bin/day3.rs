fn main() {
    println!("working");
}

#[derive(Debug, PartialEq)]
struct BitString(String);

impl BitString {
    pub fn decimal(&self) -> i32 {
        self.0.chars().rev().enumerate().fold(0, |acc, (i, n)| {
            let val: i32 = n.to_string().parse().expect("could not parse into digit");
            acc + (i32::pow(2, i as u32) * val)
        })
    }

    pub fn flip_all(&self) -> Self {
        self.0
            .chars()
            .fold("".into(), |mut acc: BitString, s| {
                if s.to_string().eq("0") {
                    acc.push_str("1")
                } else {
                    acc.push_str("0")
                }
                acc
            })
            .into()
    }

    pub fn push_str(&mut self, s: &str) {
        self.0.push_str(s)
    }
}

impl From<&str> for BitString {
    fn from(input: &str) -> Self {
        return BitString(input.to_string());
    }
}

struct BitStringList(Vec<BitString>);

impl BitStringList {
    fn len(&self) -> usize {
        self.0.len()
    }

    fn enumerate(&self) -> std::iter::Enumerate<std::slice::Iter<BitString>> {
        self.0.iter().enumerate()
    }
}

impl BitStringList {
    pub fn most_common(&self, idx: usize) -> u32 {
        let sum: i32 = self
            .0
            .iter()
            .map(|bs| {
                let i: i32 =
                    bs.0.chars()
                        .nth(idx)
                        .unwrap()
                        .to_string()
                        .parse()
                        .expect("could not parse bitstring position");
                i
            })
            .sum();
        if sum > (self.0.len() / 2) as i32 {
            1
        } else {
            0
        }
    }
}

struct Report {
    pub gamma: BitString,
    pub epsilon: BitString,

    bitstrings: BitStringList,
}

impl Report {
    pub fn from_bitstrings(bitstrings: Vec<BitString>) -> Self {
        let bsl = BitStringList(bitstrings);
        let epsilon = Report::calc_episilon(&bsl);
        let gamma = Report::calc_gamma(&bsl);
        Report {
            bitstrings: bsl,
            epsilon: epsilon,
            gamma: gamma,
        }
    }

    fn calc_gamma(bitstrings: &BitStringList) -> BitString {
        BitString(
            bitstrings
                .0
                .iter()
                .enumerate()
                .fold(String::from(""), |mut acc, (i, _)| {
                    acc.push_str(&bitstrings.most_common(i).to_string());
                    acc
                }),
        )
    }

    fn calc_episilon(bitstrings: &BitStringList) -> BitString {
        let gamma = Report::calc_gamma(bitstrings);
        gamma.flip_all()
    }
}

impl From<&str> for Report {
    fn from(input: &str) -> Self {
        let bitstrings = input
            .trim()
            .split("\n")
            .map(|s| BitString::from(s))
            .collect();
        Report::from_bitstrings(bitstrings)
    }
}

mod test {
    use super::*;

    #[test]
    fn pt1_e2e() {
        let report: Report = include_str!("../../inputs/3.example.txt").into();
        let got = report.gamma.decimal() * report.epsilon.decimal();
        let want = 198;
        assert_eq!(got, want);
    }

    #[test]
    fn bitstring_decimal() {
        let bs: BitString = "10110".into();
        assert_eq!(bs.decimal(), 22);
    }

    #[test]
    fn bitstring_flipall() {
        let bs: BitString = "10010".into();
        assert_eq!(bs.flip_all(), "01101".into())
    }

    #[test]
    fn bitstring_list() {
        let bsl: BitStringList = BitStringList(vec![
            BitString("1001".to_string()),
            BitString("1110".to_string()),
            BitString("0110".to_string()),
        ]);
        assert_eq!(bsl.most_common(0), 1);
        assert_eq!(bsl.most_common(1), 1);
        assert_eq!(bsl.most_common(2), 1);
        assert_eq!(bsl.most_common(3), 0);
    }
}
