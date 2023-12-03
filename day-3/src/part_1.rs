pub(crate) struct Grid {
    pub(crate) chars: Vec<Vec<char>>
}
impl Grid {
    pub(crate) fn new(s: &str) -> Self {
        let chars: Vec<Vec<char>> = s.lines().map(|line| line.chars().collect()).collect();
        Grid{chars}
    }
    fn has_neighbouring_symbol(&self, c: &Coord) -> bool {
        for Coord{y, x} in c.adjacent_coords() {
            if self.chars.get(y).is_some_and(|row| {
                row.get(x).is_some_and(|char| {
                    !char.is_alphanumeric() && char != &'.'
                })
            }) {
                return true
            }
        }
        false
    }

}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Coord {
    pub(crate) y: usize,
    pub(crate) x: usize,
}
impl Coord {
    fn adjacent_coords(&self) -> Vec<Coord> {
        let mut coords = vec![];

        // I'm having to do 0->3 when I want to do -1->2
        // because usize can't be negative
        // -1 when doing get, and check at start I don't index negative
        for y in 0..3 {
            for x in 0..3 {
                if (self.x == 0 && x == 0) || (self.y == 0 && y == 0) { // Can't index negative
                    continue
                }
                if x == 1 && y == 1 {
                    continue
                }

                // subtract 1 from each to account for 0->3 instead of -1->2
                coords.push(Coord{
                    y: self.y + y - 1,
                    x: self.x + x - 1,
                })
            }
        }

        coords
    }
}

#[derive(Debug, Clone)]
pub(crate) struct Number {
    pub(crate) value: i32,
    position: Vec<Coord>,
}
impl Number {
    const fn new() -> Number {
        Number{
            value: 0,
            position: vec![],
        }
    }
    fn is_empty(&self) -> bool {
        self.value == 0 && self.position.len() == 0
    }

    pub fn is_adjacent(&self, c: &Coord) -> bool {
        c.adjacent_coords().iter().any(|c| self.position.contains(c))
    }
}

pub(crate) fn get_numbers(grid: &Grid) -> Vec<Number> {
    let mut number_positions = vec![];

    for (y, line) in grid.chars.iter().enumerate() {
        let mut current_number = Number::new();

        for (x, char) in line.iter().enumerate() {
            // Go through each char
            if char.is_numeric() { // If it's a number - add to current_number
                let val = char.to_string().parse::<i32>().expect("invalid number");
                current_number.value *= 10;
                current_number.value += val;
                current_number.position.push(Coord{y, x})
            } else if !current_number.is_empty() {
                // If not a number and there is a current_number (i.e. not 0)
                // push current_number to list and reset
                number_positions.push(current_number);
                current_number = Number::new();
            }
        }

        if !current_number.is_empty() {
            number_positions.push(current_number);
        }
    }

    number_positions
}

pub(crate) fn solve(puzzle: &str) -> i32 {
    let grid = Grid::new(puzzle);
    let numbers = get_numbers(&grid);

    let mut count = 0;
    for number in numbers {
        if number.position.iter().any(|c| grid.has_neighbouring_symbol(c)) {
            count += number.value
        }
    }

    count
}