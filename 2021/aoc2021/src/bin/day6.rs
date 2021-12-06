fn main() {
    let mut p1_subject: FishSchool = include_str!("../../inputs/6.txt").into();
    p1_subject.tick_n(80);
    println!("Part 1: {}", p1_subject.len());

    let mut p2_subject: FishSchool = include_str!("../../inputs/6.txt").into();
    p2_subject.tick_n(256);
    println!("Part 2: {}", p2_subject.len());
}

struct FishSchool(Vec<u8>);

impl FishSchool {
    fn tick_n(&mut self, n: u32) {
        for _ in 0..n {
            let mut next = vec![];
            let mut to_spawn = 0;
            for i in self.0.clone() {
                // case 1: counter is 0, spawn, next is 6
                if i == 0 {
                    to_spawn += 1;
                    next.push(6);
                } else {
                    // case 2: decrement
                    next.push(i - 1);
                }
            }

            for _ in 0..to_spawn {
                next.push(8);
            }

            self.0 = next;
        }
    }

    fn len(&self) -> usize {
        self.0.len()
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
        let mut subject: FishSchool = include_str!("../../inputs/6.example.txt").into();
        subject.tick_n(80);
        assert_eq!(subject.len(), 5934);
    }

    #[test]
    fn p2_e2e() {
        let mut subject: FishSchool = include_str!("../../inputs/6.example.txt").into();
        subject.tick_n(256);
        assert_eq!(subject.len(), 26984457539);
    }
}
