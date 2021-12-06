fn main() {
    let mut game: BingoGame = include_str!("../../inputs/4.txt").into();
    let (mut winner, winning_number) = game.play().unwrap();
    let answer: u32 = winner.uncalled().map(|x| x.val).sum::<u32>() * winning_number;
    println!("Part 1: {}", answer);

    let mut p2: BingoGame = include_str!("../../inputs/4.txt").into();
    let (mut p2_winner, p2_winning_number) = p2.play_to_lose().unwrap();
    let p2_answer: u32 = p2_winner.uncalled().map(|x| x.val).sum::<u32>() * p2_winning_number;
    println!("Part 2: {}", p2_answer);
}

struct BingoGame {
    cards: Vec<BingoCard>,
    uncalled_nums: Vec<u32>,
}

impl BingoGame {
    fn play(&mut self) -> Result<(BingoCard, u32), &str> {
        for val in self.uncalled_nums.iter() {
            for i in 0..self.cards.len() {
                self.cards[i].mark(*val);
                if self.cards[i].won() {
                    return Ok((self.cards[i].clone(), *val));
                }
            }
        }
        Err("no winner found")
    }

    fn play_to_lose(&mut self) -> Result<(BingoCard, u32), &str> {
        let mut winners = vec![];
        let mut cards = self.cards.clone();
        println!("Starting len: {}", cards.len());
        for val in self.uncalled_nums.iter() {
            println!("Current len: {}", cards.len());
            let mut to_remove = vec![];
            for i in 0..cards.len() {
                cards[i].mark(*val);
                if cards[i].won() {
                    winners.push((cards[i].clone(), *val));
                    to_remove.push(i);
                }
            }
            let mut next: Vec<BingoCard> = vec![];
            for i in 0..cards.len() {
                if !to_remove.contains(&i) {
                    next.push(cards[i].clone());
                }
            }
            cards = next;
        }
        if winners.is_empty() {
            return Err("no winner found");
        }
        let (winner, winning_number) = &winners[winners.len() - 1];
        Ok((winner.clone(), winning_number.clone()))
    }
}

impl From<&str> for BingoGame {
    fn from(input: &str) -> Self {
        let splitted: Vec<&str> = input.split("\n\n").collect();
        let unparsed_nums: Vec<u32> = splitted[0]
            .split(',')
            .map(|x| x.parse().expect("could not parse num"))
            .collect();
        let cards: Vec<BingoCard> = splitted[1..]
            .into_iter()
            .map(|x| BingoCard::from(*x))
            .collect();
        BingoGame {
            cards: cards,
            uncalled_nums: unparsed_nums,
        }
    }
}

#[derive(Clone, Debug)]
struct BingoCard {
    spaces: Vec<Vec<BingoSpace>>,
}

impl BingoCard {
    fn rows(&mut self) -> Vec<Vec<BingoSpace>> {
        self.spaces.clone()
    }

    fn cols(&mut self) -> Vec<Vec<BingoSpace>> {
        let num_cols = self.spaces[0].len();
        let mut result: Vec<Vec<BingoSpace>> = vec![];
        for i in 0..num_cols {
            let mut row: Vec<BingoSpace> = vec![];
            for j in 0..self.spaces.len() {
                row.push(self.spaces[j][i]);
            }
            result.push(row);
        }
        result
    }

    fn uncalled(&mut self) -> impl Iterator<Item = BingoSpace> {
        let mut uncalled: Vec<BingoSpace> = vec![];
        for row in self.rows() {
            for space in row {
                if !space.called {
                    uncalled.push(space);
                }
            }
        }
        uncalled.into_iter()
    }

    fn mark(&mut self, val: u32) {
        for i in 0..self.spaces.len() {
            for j in 0..self.spaces[i].len() {
                if self.spaces[i][j].val == val {
                    self.spaces[i][j].called = true
                }
            }
        }
    }

    fn won(&mut self) -> bool {
        for row in self.rows() {
            if row.iter().all(|space| space.called) {
                return true;
            }
        }

        for col in self.cols() {
            if col.iter().all(|space| space.called) {
                return true;
            }
        }

        false
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
        let (mut winner, winning_number) = game.play().unwrap();
        assert_eq!(
            winner.uncalled().map(|x| x.val).sum::<u32>() * winning_number,
            want
        );
    }

    #[test]
    fn p2_e2e() {
        let mut game: BingoGame = include_str!("../../inputs/4.example.txt").into();
        let want = 1924;
        let (mut winner, winning_number) = game.play_to_lose().unwrap();
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
        let mut subject: BingoCard = input.into();
        let want = 300;
        let got: u32 = subject.uncalled().map(|x| x.val).sum();
        assert_eq!(got, want);
    }
}
