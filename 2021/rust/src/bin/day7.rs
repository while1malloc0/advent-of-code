use std::collections::HashMap;

fn main() {
    let p1_input: Crabs = include_str!("../../inputs/7.txt").into();
    let p2_input = p1_input.clone();

    let p1_answer = p1(p1_input);
    println!("Part 1: {}", p1_answer);

    let p2_answer = p2(p2_input);
    println!("Part 2: {}", p2_answer);
}

fn p1(crabs: Crabs) -> u32 {
    crabs.min_fuel()
}

fn p2(crabs: Crabs) -> u32 {
    crabs.min_with_drag()
}

#[derive(Clone)]
struct Crabs {
    shuttles: Vec<u32>,
    counts: HashMap<u32, u32>,
    min: u32,
    max: u32,
}

impl Crabs {
    fn min_fuel(&self) -> u32 {
        let mut totals: HashMap<u32, u32> = HashMap::new();
        for num in self.min..=self.max {
            if totals.contains_key(&num) {
                continue;
            }
            let mut cost: HashMap<u32, i32> = HashMap::new();
            for i in 0..self.shuttles.len() {
                if cost.contains_key(&self.shuttles[i]) {
                    continue;
                }
                let num_of_shuttles = self.counts.get(&self.shuttles[i]).unwrap();
                let cost_for_shuttle: i32 =
                    ((self.shuttles[i] as i32) - (num as i32)).abs() * *num_of_shuttles as i32;
                cost.insert(self.shuttles[i], cost_for_shuttle);
            }
            let total_cost: i32 = cost.values().sum();
            totals.insert(num, total_cost as u32);
        }

        *totals.values().min().unwrap()
    }

    fn min_with_drag(&self) -> u32 {
        let mut totals: HashMap<u32, u32> = HashMap::new();
        for num in self.min..=self.max {
            if totals.contains_key(&num) {
                continue;
            }
            let mut cost: HashMap<u32, i32> = HashMap::new();
            for i in 0..self.shuttles.len() {
                if cost.contains_key(&self.shuttles[i]) {
                    continue;
                }
                let num_of_shuttles = self.counts.get(&self.shuttles[i]).unwrap();
                let distance: i32 = ((self.shuttles[i] as i32) - (num as i32)).abs();
                let drag = sigma(distance);
                let cost_for_shuttle: i32 = drag * (*num_of_shuttles as i32);
                cost.insert(self.shuttles[i], cost_for_shuttle);
            }
            let total_cost: i32 = cost.values().sum();
            totals.insert(num, total_cost as u32);
        }

        *totals.values().min().unwrap()
    }
}

impl From<&str> for Crabs {
    fn from(input: &str) -> Self {
        let shuttles: Vec<u32> = input
            .trim()
            .split(",")
            .map(|x| x.parse::<u32>().unwrap())
            .collect();
        let mut counts: HashMap<u32, u32> = HashMap::new();
        for num in shuttles.clone() {
            let count = counts.entry(num).or_insert(0);
            *count += 1;
        }

        let min = *shuttles.iter().min().unwrap();
        let max = *shuttles.iter().max().unwrap();

        Self {
            shuttles,
            counts,
            min,
            max,
        }
    }
}

fn sigma(n: i32) -> i32 {
    let mut sum = 0;
    for i in 1..=n {
        sum += i;
    }
    sum
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let input: Crabs = include_str!("../../inputs/7.example.txt").into();
        let got = p1(input);
        assert_eq!(got, 37);
    }

    #[test]
    fn p2_e2e() {
        let input: Crabs = include_str!("../../inputs/7.example.txt").into();
        let got = p2(input);
        assert_eq!(got, 168);
    }

    #[test]
    fn test_sigma() {
        let n = 11;
        let got = sigma(n);
        assert_eq!(got, 66);
    }
}
