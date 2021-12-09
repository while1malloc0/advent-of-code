use std::collections::{HashMap, HashSet};

fn main() {
    let input = include_str!("../../inputs/9.txt");
    let p1_answer = p1(input);
    println!("Part 1: {}", p1_answer);
}

fn p1(input: &str) -> u32 {
    let heightmap: HeightMap = input.into();
    let low_points = heightmap.low_points();
    let risks = low_points.values().map(|x| x + 1);
    risks.sum()
}

struct HeightMap(HashMap<(u32, u32), u32>);

impl HeightMap {
    fn get(&self, x: i32, y: i32) -> Option<&u32> {
        if x < 0 || y < 0 {
            return None;
        }
        self.0.get(&(x as u32, y as u32))
    }

    fn neighbors(&self, x: i32, y: i32) -> Result<Vec<u32>, &str> {
        if x < 0 || y < 0 {
            return Err("X and Y cannot be less than 0");
        }

        let mut neighbors: Vec<u32> = vec![];

        let left = self.get(x - 1, y);
        if let Some(i) = left {
            neighbors.push(*i);
        }

        let top = self.get(x, y - 1);
        if let Some(i) = top {
            neighbors.push(*i);
        }

        let right = self.get(x + 1, y);
        if let Some(i) = right {
            neighbors.push(*i);
        }

        let bottom = self.get(x, y + 1);
        if let Some(i) = bottom {
            neighbors.push(*i);
        }

        Ok(neighbors)
    }

    fn low_points(&self) -> HashMap<(u32, u32), u32> {
        let mut low_points: HashMap<(u32, u32), u32> = HashMap::new();
        for (coord, val) in &self.0 {
            let (x, y) = coord;
            let neighbors = self.neighbors(*x as i32, *y as i32).unwrap();
            if is_smallest(*val, neighbors) {
                low_points.insert((*x, *y), *val);
            }
        }
        low_points
    }

    fn basins(&self) -> Vec<&HashSet<(u32, u32)>> {
        let result = vec![];

        let low_points = self.low_points();
        let stack: Vec<(u32, u32)> = low_points.keys().into_iter().map(|x| *x).collect();
        println!("{:?}", stack);

        result
    }
}

impl From<&str> for HeightMap {
    fn from(input: &str) -> Self {
        let mut content = HashMap::new();
        let lines = input.trim().lines();
        for (y, line) in lines.enumerate() {
            let vals = line.chars();
            for (x, c) in vals.enumerate() {
                let coord = (x as u32, y as u32);
                let num = c.to_string().parse::<u32>().unwrap();
                content.insert(coord, num);
            }
        }
        Self(content)
    }
}

fn is_smallest(lhs: u32, rhs: Vec<u32>) -> bool {
    let smallest_in_vec = rhs.iter().min().unwrap();
    return lhs < *smallest_in_vec;
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let input = include_str!("../../inputs/9.example.txt");
        let got = p1(input);
        let want = 15;
        assert_eq!(got, want);
    }

    #[test]
    fn test_parse() {
        let input = include_str!("../../inputs/9.example.txt");
        let subject: HeightMap = input.into();
        let got = subject.get(0, 0).unwrap();
        let want = 2;
        assert_eq!(*got, want);
    }

    #[test]
    fn test_neighbors() {
        let input = include_str!("../../inputs/9.example.txt");
        let subject: HeightMap = input.into();
        let got = subject.neighbors(0, 0).unwrap();
        let want = vec![1, 3];
        assert_eq!(got, want)
    }

    #[test]
    fn test_is_smallest() {
        let lhs: u32 = 0;
        let rhs: Vec<u32> = vec![1, 2, 3];
        let got = is_smallest(lhs, rhs);
        let want = true;
        assert_eq!(got, want);
    }

    #[test]
    fn test_low_points() {
        let input = include_str!("../../inputs/9.example.txt");
        let subject: HeightMap = input.into();
        let got = subject.low_points();
        let mut want: HashMap<(u32, u32), u32> = HashMap::new();
        want.insert((1, 0), 1);
        want.insert((9, 0), 0);
        want.insert((2, 2), 5);
        want.insert((6, 4), 5);
        assert_eq!(got, want);
    }

    #[test]
    fn test_basins() {
        let input = include_str!("../../inputs/9.example.txt");
        let subject: HeightMap = input.into();
        let got = subject.basins();
        let check = got[0];
        let mut want: HashSet<(u32, u32)> = HashSet::new();
        want.insert((0, 0));
        want.insert((0, 1));
        want.insert((1, 0));
        assert_eq!(check, &want);
    }
}
