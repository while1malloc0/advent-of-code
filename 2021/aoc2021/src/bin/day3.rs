use std::collections::HashMap;

fn main() {
    // let report: Report = include_str!("../../inputs/3.example.txt").into();
    // let answer = report.gamma.decimal() * report.epsilon.decimal();
    // println!("{}", answer);

    let bsl: BitStringList = include_str!("../../inputs/3.example.txt").into();
    Report::calc_co2(&bsl);
}

#[derive(Debug, PartialEq, Clone)]
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

    pub fn len(&self) -> usize {
        self.0.len()
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

    pub fn most_common(&self, idx: usize) -> Option<u32> {
        let mut counts: HashMap<char, u32> = HashMap::new();
        for bs in &self.0 {
            let bit = bs.0.chars().nth(idx).unwrap();
            let current = counts.entry(bit).or_insert(0);
            *current += 1;
        }
        let max_val = counts.values().max().unwrap();
        for k in counts.keys() {
            if counts.get(k).unwrap() == max_val {
                return Some(k.to_string().parse().unwrap());
            }
        }
    }

    pub fn least_common(&self, idx: usize) -> u32 {
        let mc = self.most_common(idx);
        if mc == 0 {
            1
        } else {
            0
        }
    }

    pub fn str_len(&self) -> usize {
        return self.0[0].len();
    }
}

impl From<&str> for BitStringList {
    fn from(input: &str) -> Self {
        let bitstrings = input
            .trim()
            .split("\n")
            .map(|s| BitString::from(s))
            .collect();
        BitStringList(bitstrings)
    }
}

struct Report {
    pub gamma: BitString,
    pub epsilon: BitString,
    pub co2: BitString,
    pub oxygen: BitString,

    bitstrings: BitStringList,
}

impl Report {
    pub fn from_bitstrings(bitstrings: Vec<BitString>) -> Self {
        let bsl = BitStringList(bitstrings);
        let epsilon = Report::calc_episilon(&bsl);
        let gamma = Report::calc_gamma(&bsl);
        let co2 = Report::calc_co2(&bsl);
        let oxygen = Report::calc_oxygen(&bsl);
        Report {
            bitstrings: bsl,
            epsilon: epsilon,
            gamma: gamma,
            co2: co2,
            oxygen: oxygen,
        }
    }

    fn calc_gamma(bitstrings: &BitStringList) -> BitString {
        let mut out: BitString = "".into();
        for i in 0..bitstrings.str_len() {
            out.push_str(&bitstrings.most_common(i).to_string());
        }
        out
    }

    fn calc_episilon(bitstrings: &BitStringList) -> BitString {
        let gamma = Report::calc_gamma(bitstrings);
        gamma.flip_all()
    }

    fn calc_co2(bitstrings: &BitStringList) -> BitString {
        let mut candidates = bitstrings.0.clone();
        println!("calc_co2: {:?}", candidates);
        for i in 0..bitstrings.str_len() {
            let lc = bitstrings.least_common(i);
            candidates = candidates
                .into_iter()
                .filter(|x| {
                    x.0.chars()
                        .nth(i)
                        .unwrap()
                        .to_string()
                        .parse::<u32>()
                        .unwrap()
                        == lc
                })
                .collect();
        }
        candidates[0].clone()
    }

    fn calc_oxygen(bitstrings: &BitStringList) -> BitString {
        let mut candidates = bitstrings.0.clone();
        println!("{:?}", candidates);
        for i in 0..bitstrings.str_len() {
            let mc = bitstrings.most_common(i);
            candidates = candidates
                .into_iter()
                .filter(|x| {
                    x.0.chars()
                        .nth(i)
                        .unwrap()
                        .to_string()
                        .parse::<u32>()
                        .unwrap()
                        == mc
                })
                .collect();
        }
        candidates[0].clone()
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
    fn pt2_e2e() {
        let report: Report = include_str!("../../inputs/3.example.txt").into();
        let got = report.oxygen.decimal() * report.co2.decimal();
        let want = 230;
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

    #[test]
    fn bitstring_least_common() {
        let bsl: BitStringList = BitStringList(vec![
            BitString("1001".to_string()),
            BitString("1110".to_string()),
            BitString("0110".to_string()),
        ]);
        assert_eq!(bsl.least_common(0), 0);
        assert_eq!(bsl.least_common(1), 0);
        assert_eq!(bsl.least_common(2), 0);
        assert_eq!(bsl.least_common(3), 1);
    }

    #[test]
    fn report_calc_co2() {
        let bsl: BitStringList = include_str!("../../inputs/3.example.txt").into();
        assert_eq!(Report::calc_co2(&bsl).decimal(), 10);
    }

    #[test]
    fn report_calc_oxygen() {
        let bsl: BitStringList = include_str!("../../inputs/3.example.txt").into();
        assert_eq!(Report::calc_oxygen(&bsl).decimal(), 23);
    }
}
