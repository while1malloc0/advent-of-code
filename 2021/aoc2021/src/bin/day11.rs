use std::collections::HashMap;
use std::fmt;

fn main() {
    unreachable!();
}

fn p1(input: &str) -> u64 {
    panic!("not yet implemented")
}

#[derive(PartialEq)]
struct Board(HashMap<(u8, u8), u8>);

impl Board {
    fn get(&self, x: u8, y: u8) -> Option<&u8> {
        self.0.get(&(x, y))
    }

    fn inc_all(&mut self) {
        for key in self.0.clone().keys() {
            let current = self.0.entry(*key).or_default();
            *current += 1;
        }
    }

    fn tick(&mut self) {
        self.inc_all();
        self.handle_flashes();
    }

    fn handle_flashes(&mut self) {
        let mut i = 0;
        loop {
            println!("{}: {:?}", i, self);
            let above_nine: Vec<(u8, u8)> = self
                .0
                .clone()
                .into_iter()
                .filter(|(_, v)| *v > 9)
                .map(|(k, _)| k)
                .collect();
            if above_nine.len() == 0 {
                break;
            }
            for coord in above_nine {
                // inc all neighbors
                let neighbors = Board::neighbors(coord.0, coord.1);
                for neighbor in neighbors {
                    let current = self.0.entry(neighbor).or_default();
                    *current += 1;
                }

                // set back to 0
                let current = self.0.entry(coord).or_default();
                *current = 0;
            }
            i += 1;
        }
    }

    fn neighbors(x: u8, y: u8) -> Vec<(u8, u8)> {
        let tmpx = x as i8;
        let tmpy = y as i8;
        let tmp: Vec<(i8, i8)> = vec![
            (tmpx - 1, tmpy - 1),
            (tmpx, tmpy - 1),
            (tmpx + 1, tmpy - 1),
            (tmpx - 1, tmpy),
            (tmpx + 1, tmpy),
            (tmpx - 1, tmpy + 1),
            (tmpx, tmpy + 1),
            (tmpx + 1, tmpy + 1),
        ];
        tmp.into_iter()
            .filter(|(x, _)| *x >= 0)
            .filter(|(x, _)| *x <= 9)
            .filter(|(_, y)| *y >= 0)
            .filter(|(_, y)| *y <= 9)
            .map(|(x, y)| (x as u8, y as u8))
            .collect()
    }
}

impl From<&str> for Board {
    fn from(input: &str) -> Self {
        let mut content: HashMap<(u8, u8), u8> = HashMap::new();
        for (y, line) in input.lines().enumerate() {
            for (x, c) in line.chars().enumerate() {
                let num = c.to_string().parse::<u8>().unwrap();
                content.insert((x as u8, y as u8), num);
            }
        }
        Self(content)
    }
}

impl fmt::Debug for Board {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut s: String = String::from("");
        s.push_str("\n");
        for y in 0..=9 {
            for x in 0..=9 {
                if let Some(val) = self.get(x, y) {
                    s.push_str(&val.to_string());
                }
            }
            s.push_str("\n");
        }
        f.write_str(&s.trim_end())
    }
}

mod test {
    use super::*;

    #[test]
    #[ignore]
    fn p1_e2e() {
        let input = include_str!("../../inputs/11.example.txt");
        let got = p1(input);
        let want = 1656;
        assert_eq!(got, want);
    }

    #[test]
    fn test_parse_board() {
        let input = include_str!("../../inputs/11.example.txt");
        let got: Board = input.into();
        assert_eq!(*got.get(0, 0).unwrap(), 5);
    }

    #[test]
    fn test_board_tick() {
        let input = r"
11111
19991
19191
19991
11111"
            .trim();
        let mut subject: Board = input.into();
        subject.inc_all();
        assert_eq!(*subject.get(0, 0).unwrap(), 2);
    }

    #[test]
    fn test_tick() {
        let input = r"
11111
19991
19191
19991
11111
        "
        .trim();
        let mut subject: Board = input.into();
        let want: Board = r"
34543
40004
50005
40004
34543
        "
        .trim()
        .into();

        subject.tick();
        assert_eq!(subject, want);
    }

    #[test]
    fn test_neighbors() {
        let got = Board::neighbors(1, 1);
        let want: Vec<(u8, u8)> = vec![
            (0, 0),
            (1, 0),
            (2, 0),
            (0, 1),
            (2, 1),
            (0, 2),
            (1, 2),
            (2, 2),
        ];
        assert_eq!(got, want);
    }

    #[test]
    fn test_neighbors_edge() {
        let got = Board::neighbors(0, 0);
        let want: Vec<(u8, u8)> = vec![(1, 0), (0, 1), (1, 1)];
        assert_eq!(got, want);
    }
}
