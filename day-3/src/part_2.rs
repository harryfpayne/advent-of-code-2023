
use crate::part_1;
use crate::part_1::{Grid, Number};

pub(crate) fn solve(puzzle: &str) -> i32 {
    let grid = Grid::new(puzzle);
    let numbers = part_1::get_numbers(&grid);

    let mut sum = 0;
    for (y, row) in grid.chars.iter().enumerate() {
        for (x, c) in row.iter().enumerate() {
            let coord = part_1::Coord{y,x};
            if c == &'*' {
                let adjacent: Vec<&Number> = numbers.iter().filter(|n| n.is_adjacent(&coord)).collect();
                if adjacent.len() != 2 {
                    continue
                }
                sum += adjacent.iter().fold(1, |acc, n| acc * n.value)
            }
        }
    }

    sum
}