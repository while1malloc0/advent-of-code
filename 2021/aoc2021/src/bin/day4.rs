use std::collections::VecDeque;

fn main() {
    println!("working!");
}

struct BingoGame {
    cards: Vec<BingoCard>,
    uncalled_nums: VecDeque<u32>,
}

impl BingoGame {
    fn play(&mut self) -> Result<(BingoCard, u32), &str> {
        for i in 0..self.uncalled_nums {}
    }
}

impl From<&str> for BingoGame {
    fn from(input: &str) -> Self {
        let splitted: Vec<&str> = input.split("\n\n").collect();
        let unparsed_nums: VecDeque<u32> = splitted[0]
            .split(',')
            .map(|x| x.parse().expect("could not parse num"))
            .collect();
        let cards: Vec<BingoCard> = splitted[1..]
            .into_iter()
            .map(|x| BingoCard::from(*x))
            .collect();
        BingoGame {
            cards,
            uncalled_nums: unparsed_nums,
        }
    }
}

#[derive(Clone, Debug)]
struct BingoCard {
    spaces: Vec<Vec<BingoSpace>>,
}

impl BingoCard {
    fn rows(&self) -> impl Iterator<Item = &Vec<BingoSpace>> {
        self.spaces.iter()
    }

    fn cols(&self) -> impl Iterator<Item = Vec<BingoSpace>> {
        let num_cols = self.spaces[0].len();
        let mut result: Vec<Vec<BingoSpace>> = vec![];
        for i in 0..num_cols {
            let mut row: Vec<BingoSpace> = vec![];
            for j in 0..self.spaces.len() {
                row.push(self.spaces[j][i]);
            }
            result.push(row);
        }
        result.into_iter()
    }

    fn uncalled(&self) -> impl Iterator<Item = BingoSpace> {
        let mut uncalled: Vec<BingoSpace> = vec![];
        for row in self.rows() {
            for space in row {
                if !space.called {
                    uncalled.push(*space);
                }
            }
        }
        uncalled.into_iter()
    }
}

impl From<&str> for BingoCard {
    fn from(input: &str) -> Self {
        let mut spaces: Vec<Vec<BingoSpace>> = vec![];
        for row in input.trim().split('\n') {
            let mut space_row: Vec<BingoSpace> = vec![];
            for num in row.replace("  ", " ").trim().split(' ') {
                let val: u32 = num.trim().parse().expect("could not parse value");
                space_row.push(BingoSpace { val, called: false })
            }
            spaces.push(space_row);
        }
        BingoCard { spaces }
    }
}

#[derive(Clone, Copy, Debug)]
struct BingoSpace {
    called: bool,
    val: u32,
}

mod test {
    use super::*;

    #[test]
    fn p1_e2e() {
        let mut game: BingoGame = include_str!("../../inputs/4.example.txt").into();
        let want = 4512;
        let (winner, winning_number) = game.play().unwrap();
        assert_eq!(
            winner.uncalled().map(|x| x.val).sum::<u32>() * winning_number,
            want
        );
    }

    #[test]
    fn bingo_card_uncalled_sum() {
        let input = r"
22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
"
        .trim();
        let subject: BingoCard = input.into();
        let want = 300;
        let got: u32 = subject.uncalled().map(|x| x.val).sum();
        assert_eq!(got, want);
    }
}
