use std::collections::HashMap;

fn main() {
    let p1_in: Lines = include_str!("../../inputs/5.txt").into();
    let p1_answer = p1(p1_in);
    println!("Part 1: {}", p1_answer);

    let p2_in: Lines = include_str!("../../inputs/5.txt").into();
    let p2_answer = p2(p2_in);
    println!("Part 2: {}", p2_answer);
}

#[derive(Clone)]
struct Line {
    start: (u32, u32),
    end: (u32, u32),
}

impl Line {
    fn coords(&self) -> Result<Vec<(u32, u32)>, &str> {
        let mut result: Vec<(u32, u32)> = vec![];

        if self.start.1 == self.end.1 {
            // case 1: line is horizontal
            // handle cases going right to left
            let (start, end) = if self.start.0 > self.end.0 {
                (self.end.0, self.start.0)
            } else {
                (self.start.0, self.end.0)
            };
            for i in start..=end {
                result.push((i as u32, self.start.1))
            }
        } else if self.start.0 == self.end.0 {
            // case 2: line is vertical
            // handle cases going right to left
            let (start, end) = if self.start.1 > self.end.1 {
                (self.end.1, self.start.1)
            } else {
                (self.start.1, self.end.1)
            };
            for i in start..=end {
                result.push((self.start.0, i as u32));
            }
        } else {
            // get start and end x and y
            // if x1 < x2, x++, otherwise x--
            // if y1 < y2, y++, otherwise y--
            let (mut x, mut y) = (self.start.0, self.start.1);
            let (endx, endy) = (self.end.0, self.end.1);

            result.push((x, y));
            loop {
                if x == endx && y == endy {
                    break;
                }
                if x < endx {
                    x += 1;
                } else {
                    x -= 1;
                }

                if y < endy {
                    y += 1;
                } else {
                    y -= 1;
                }

                result.push((x, y))
            }
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

    fn diagonal(&self) -> Vec<Line> {
        self.0
            .clone()
            .into_iter()
            .filter(|line| line.start.0 != line.end.0 && line.start.1 != line.end.1)
            .collect()
    }
}

impl From<&str> for Lines {
    fn from(input: &str) -> Self {
        let lines: Vec<Line> = input.trim().split("\n").map(Line::from).collect();
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

    coords
        .values()
        .filter(|x| **x >= 2)
        .collect::<Vec<&u32>>()
        .len() as u32
}

fn p2(lines: Lines) -> u32 {
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
    for line in lines.diagonal() {
        for coord in line.coords().unwrap() {
            let count = coords.entry(coord).or_insert(0);
            *count += 1;
        }
    }

    coords
        .values()
        .filter(|x| **x >= 2)
        .collect::<Vec<&u32>>()
        .len() as u32
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let subject: Lines = include_str!("../../inputs/5.example.txt").into();
        let got = p1(subject);
        assert_eq!(got, 5);
    }

    #[test]
    fn p2_e2e() {
        let subject: Lines = include_str!("../../inputs/5.example.txt").into();
        let got = p2(subject);
        assert_eq!(got, 12);
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

        let diagonal: Line = "1,1 -> 3,3".into();
        let got = diagonal.coords().unwrap();
        assert_eq!(got.len(), 3);

        let other_diagonal: Line = "9,7 -> 7,9".into();
        let got = other_diagonal.coords().unwrap();
        assert_eq!(got.len(), 3);
    }
}
