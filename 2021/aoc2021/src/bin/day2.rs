fn main() {
    let input = include_str!("../../inputs/2.txt");
    let directions = parse_directions(input);
    let mut coordinate = Coordinate::new();
    coordinate.apply(directions);
    let got = coordinate.position * coordinate.depth;
    println!("Part 1: {:?}", got);
}

#[derive(Debug)]
struct Coordinate {
    depth: i32,
    position: i32,
}

impl Coordinate {
    fn new() -> Self {
        return Coordinate {
            depth: 0,
            position: 0,
        };
    }

    fn apply(&mut self, vectors: Vec<Vector>) -> &Self {
        for v in vectors {
            match v.direction {
                Direction::Forward => self.position += v.velocity,
                Direction::Up => self.depth -= v.velocity,
                Direction::Down => self.depth += v.velocity,
            }
        }
        self
    }
}

#[derive(Debug)]
enum Direction {
    Forward,
    Up,
    Down,
}

impl From<&str> for Direction {
    fn from(s: &str) -> Self {
        match s {
            "forward" => Direction::Forward,
            "up" => Direction::Up,
            "down" => Direction::Down,
            _ => panic!("{:?} is not a valid direction", s),
        }
    }
}

#[derive(Debug)]
struct Vector {
    direction: Direction,
    velocity: i32,
}

impl From<&str> for Vector {
    fn from(s: &str) -> Self {
        let direction_and_velocity: Vec<&str> = s.split(" ").collect();
        let direction = Direction::from(direction_and_velocity[0]);
        let velocity = direction_and_velocity[1]
            .parse()
            .expect("could not parse velocity");
        Vector {
            direction: direction,
            velocity: velocity,
        }
    }
}

fn parse_directions(input: &str) -> Vec<Vector> {
    input
        .trim()
        .split("\n")
        .map(|line| Vector::from(line))
        .collect()
}

mod tests {
    use super::*;

    #[test]
    fn e2e() {
        const INPUT: &str = include_str!("../../inputs/2.example.txt");
        let mut coordinate = Coordinate::new();
        let directions = parse_directions(INPUT);
        coordinate.apply(directions);
        let got = coordinate.position * coordinate.depth;
        assert_eq!(got, 150);
    }
}
