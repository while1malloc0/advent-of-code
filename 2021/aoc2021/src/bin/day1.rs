fn main() {
    let input_raw = include_str!("../../inputs/1.1.txt").to_string();
    let input: Vec<i32> = input_raw
        .trim()
        .split("\n")
        .map(|i| i.parse::<i32>().expect("could not parse input to int"))
        .collect();

    {
        println!("Part 1:");
        let mut count = 0;
        for i in 0..input.len() - 1 {
            if input[i] < input[i + 1] {
                count += 1;
            }
        }
        println!("{:?}", count);
    }

    {
        println!("Part 2:");
        let input_windows = input[..].windows(3);
        let mut v = vec![];
        for window in input_windows {
            v.push(window);
        }
        let mut count = 0;
        for i in 0..v.len() - 1 {
            if v[i].iter().sum::<i32>() < v[i + 1].iter().sum::<i32>() {
                count += 1
            }
        }
        println!("{:?}", count);
    }
}
