use std::collections::HashMap;

fn main() {
    let p1_subject: FishSchool = include_str!("../../inputs/6.txt").into();
    let p1_answer = p1_subject.len_after(80);
    println!("Part 1: {}", p1_answer);

    let p2_subject: FishSchool = include_str!("../../inputs/6.txt").into();
    let p2_answer = p2_subject.len_after(256);
    println!("Part 2: {}", p2_answer);
}

struct FishSchool(Vec<u8>);

impl FishSchool {
    fn len_after(&self, n: u32) -> u64 {
        let mut counts: HashMap<u8, u64> = HashMap::new();
        for num in self.0.clone() {
            let count = counts.entry(num).or_insert(0);
            *count += 1;
        }

        for _ in 0..n {
            let mut dup = counts.clone();
            let num_to_spawn = dup.entry(0).or_insert(0);
            for i in 0..8 {
                let mut dup = counts.clone();
                let next = dup.entry(i + 1).or_insert(0);
                let fish = counts.entry(i).or_insert(0);
                *fish = *next;
            }
            counts.insert(8, *num_to_spawn);
            let ready = counts.entry(6).or_insert(0);
            *ready += *num_to_spawn;
        }

        counts.values().sum()
    }
}

impl From<&str> for FishSchool {
    fn from(input: &str) -> Self {
        let fish: Vec<u8> = input
            .trim()
            .split(",")
            .map(|s| s.parse().unwrap())
            .collect();
        FishSchool(fish)
    }
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let subject: FishSchool = include_str!("../../inputs/6.example.txt").into();
        assert_eq!(subject.len_after(80), 5934);
    }

    #[test]
    fn p2_e2e() {
        let subject: FishSchool = include_str!("../../inputs/6.example.txt").into();
        assert_eq!(subject.len_after(256), 26984457539);
    }
}
