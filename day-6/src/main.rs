

fn parse(s: &str) -> Vec<(i64, i64)> {
    let (times, distances) = s.split_once("\n").expect("invalid str");

    let (_, times) = times.split_once(":").expect("invalid time");
    let (_, distances) = distances.split_once(":").expect("invalid time");

    let times: Vec<i64> = times.trim().split_ascii_whitespace().filter_map(|v| v.parse().ok()).collect();
    let distances: Vec<i64> = distances.trim().split_ascii_whitespace().filter_map(|v| v.parse().ok()).collect();

    times.into_iter().zip(distances.into_iter()).collect()
}

fn parse_2(s: &str) -> Vec<(i64, i64)> {
    let (times, distances) = s.split_once("\n").expect("invalid str");

    let (_, times) = times.split_once(":").expect("invalid time");
    let (_, distances) = distances.split_once(":").expect("invalid time");

    let times: Vec<i64> = times.trim().replace(" ", "").split_ascii_whitespace().filter_map(|v| v.parse().ok()).collect();
    let distances: Vec<i64> = distances.trim().replace(" ", "").split_ascii_whitespace().filter_map(|v| v.parse().ok()).collect();

    times.into_iter().zip(distances.into_iter()).collect()
}

fn calculate_possible_wins((time, distance): (i64, i64)) -> i64 {
    let mut count = 0;
    for time_held in 0..time {
        let remaining_time = time - time_held;
        let distance_travelled = remaining_time * time_held;
        if distance_travelled > distance {
            count += 1;
        }
    }
    count
}

fn main() {
    let records = parse(PUZZLE);
    let mut margin = 1;
    for record in records {
        margin *= calculate_possible_wins(record)
    }
    println!("Part 1: {}", margin);

    let records = parse_2(PUZZLE);
    let mut margin = 1;
    for record in records {
        margin *= calculate_possible_wins(record)
    }
    println!("Part 2: {}", margin)
}

const TEST: &str = "\
Time:      7  15   30
Distance:  9  40  200";

const PUZZLE: &str = "\
Time:        55     82     64     90
Distance:   246   1441   1012   1111
";