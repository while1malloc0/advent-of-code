use std::collections::HashMap;

fn main() {
    println!("Working");
}

#[derive(Clone)]
struct Line {
    start: (u32, u32),
    end: (u32, u32),
}

impl Line {
    fn coords(&self) -> Result<Vec<(u32, u32)>, &str> {
        let mut result: Vec<(u32, u32)> = vec![];

        // TODO need to handle case where start < end for some coordinates (e.g. they go down or left)
        panic!("start here");
        if self.start.1 == self.end.1 {
            // case 1: line is horizontal
            for i in self.start.0..=self.end.0 {
                result.push((i as u32, self.start.1))
            }
        } else if self.start.0 == self.end.0 {
            // case 2: line is vertical
            for i in self.start.1..=self.end.1 {
                result.push((self.start.0, i as u32));
            }
        } else {
            // case 3: line is diagnal, and we don't do that here
            return Err("we only do horizontal or vertical lines here");
        }

        Ok(result)
    }
}

impl From<&str> for Line {
    fn from(input: &str) -> Self {
        let start_and_end: Vec<&str> = input.split("->").map(str::trim).collect();
        let start = start_and_end[0];
        let end = start_and_end[1];
        let (start_x, start_y) = match &start
            .split(",")
            .map(str::parse::<u32>)
            .map(Result::unwrap)
            .collect::<Vec<u32>>()[..]
        {
            &[x, y, ..] => (x, y),
            _ => unreachable!(),
        };

        let (end_x, end_y) = match &end
            .split(",")
            .map(str::parse::<u32>)
            .map(Result::unwrap)
            .collect::<Vec<u32>>()[..]
        {
            &[x, y, ..] => (x, y),
            _ => unreachable!(),
        };

        Self {
            start: (start_x, start_y),
            end: (end_x, end_y),
        }
    }
}

struct Lines(Vec<Line>);

impl Lines {
    fn horizontal(&self) -> Vec<Line> {
        self.0
            .clone()
            .into_iter()
            .filter(|line| line.start.1 == line.end.1)
            .collect()
    }

    fn vertical(&self) -> Vec<Line> {
        self.0
            .clone()
            .into_iter()
            .filter(|line| line.start.0 == line.end.0)
            .collect()
    }
}

impl From<&str> for Lines {
    fn from(input: &str) -> Self {
        let lines: Vec<Line> = input.split("\n").map(Line::from).collect();
        Lines(lines)
    }
}

fn p1(lines: Lines) -> u32 {
    let mut coords: HashMap<(u32, u32), u32> = HashMap::new();
    for line in lines.horizontal() {
        for coord in line.coords().unwrap() {
            let count = coords.entry(coord).or_insert(0);
            *count += 1;
        }
    }

    for line in lines.vertical() {
        for coord in line.coords().unwrap() {
            let count = coords.entry(coord).or_insert(0);
            *count += 1;
        }
    }

    println!("{:?}", coords.values());
    coords
        .values()
        .filter(|x| **x >= 2)
        .collect::<Vec<&u32>>()
        .len() as u32
}

mod test {
    use super::*;

    #[test]
    #[ignore]
    fn p1_e2e() {
        let subject: Lines = include_str!("../../inputs/5.example.txt").into();
        let got = p1(subject);
        assert_eq!(got, 5);
    }

    #[test]
    fn parse_line() {
        let input = "0,9 -> 5,9";
        let got: Line = input.into();
        assert_eq!(got.start, (0, 9));
        assert_eq!(got.end, (5, 9));
    }

    #[test]
    fn lines_horizontal() {
        let subject: Lines = include_str!("../../inputs/5.example.txt").into();
        let got = subject.horizontal();
        assert_eq!(got.len(), 4);
        assert_eq!(got[0].start.1, 9);
        assert_eq!(got[1].start.1, 4);
        assert_eq!(got[2].start.1, 9);
        assert_eq!(got[3].start.1, 4);
    }

    #[test]
    fn lines_vertical() {
        let subject: Lines = include_str!("../../inputs/5.example.txt").into();
        let got = subject.vertical();
        assert_eq!(got.len(), 2);
        assert_eq!(got[0].start.0, 2);
        assert_eq!(got[1].start.0, 7);
    }

    #[test]
    fn line_coords() {
        let horizontal: Line = "0,9 -> 5,9".into();
        let got = horizontal.coords().unwrap();
        assert_eq!(got.len(), 6);
        assert_eq!(got[0], (0, 9));
        assert_eq!(got[1], (1, 9));
        assert_eq!(got[2], (2, 9));
        assert_eq!(got[3], (3, 9));
        assert_eq!(got[4], (4, 9));
        assert_eq!(got[5], (5, 9));

        let vertical: Line = "2,2 -> 2,1".into();
        let got = vertical.coords().unwrap();
        assert_eq!(got.len(), 2);
        assert_eq!(got[0], (2, 2));
        assert_eq!(got[1], (2, 1));
    }
}
