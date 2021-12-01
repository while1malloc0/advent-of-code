fn main() {
    let input: Vec<i32> = include_str!("../../inputs/1.1.txt")
        .to_string()
        .trim()
        .split("\n")
        .map(|i| i.parse::<i32>().expect("could not parse input to int"))
        .collect();

    {
        println!("Part 1:");
        // non-clever solution, just use a for-loop
        let mut count = 0;
        for i in 0..input.len() - 1 {
            if input[i] < input[i + 1] {
                count += 1;
            }
        }
        println!("Simple: {:?}", count);

        // clever (or maybe just obtuse?) solution, slice input into a sliding
        // 2-window, filter for tuples that match our criteria, and see how many
        // there are
        let opt = input
            .clone()
            .as_slice()
            .windows(2)
            .filter(|x| x[0] < x[1])
            .count();
        println!("Clever: {:?}", opt);
    }

    {
        println!("Part 2:");

        // Simple: use a for loop
        let v: Vec<&[i32]> = input.as_slice().windows(3).collect();
        let mut count = 0;
        for i in 0..v.len() - 1 {
            if v[i].iter().sum::<i32>() < v[i + 1].iter().sum::<i32>() {
                count += 1
            }
        }
        println!("Simple: {:?}", count);

        // Clever (again, or maybe just obtuse): repeat our slice trick above
        let opt = input
            .as_slice()
            .windows(3)
            .map(|x| x.iter().sum()) // transform into a slice of sums...
            .collect::<Vec<i32>>()
            .as_slice()
            .windows(2) // ...then repeat the same trick as above
            .filter(|x| x[0] < x[1])
            .count();
        println!("Clever: {:?}", opt);
    }
}
