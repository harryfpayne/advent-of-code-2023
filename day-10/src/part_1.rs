
#[derive(Debug, Copy, Clone, Hash, Eq, PartialEq)]
pub struct Position {
    pub(crate) y: usize,
    pub(crate) x: usize,
}

#[derive(Eq, PartialEq, Copy, Clone, Debug)]
pub enum Direction {
    N,E,S,W
}
impl Direction {
    pub fn move_dir(&self, p: &Position) -> Position {
        let mut p = p.clone();
        match self {
            Direction::N => p.y -= 1,
            Direction::E => p.x += 1,
            Direction::S => p.y += 1,
            Direction::W => p.x -= 1,
        }
        p
    }
    fn opposite(&self) -> Direction {
        match self {
            Direction::N => Direction::S,
            Direction::E => Direction::W,
            Direction::S => Direction::N,
            Direction::W => Direction::E,
        }
    }
}

#[derive(Debug, Eq, PartialEq)]
pub enum Joint {
    G,
    Start,
    Pipe(Direction, Direction),
}
impl Joint {
    fn get_movement_direction(&self, enter_side: &Direction) -> Option<Direction> {
        match self {
            Joint::G => None,
            Joint::Start => {
                Some(enter_side.opposite())
            },
            Joint::Pipe(d1, d2) => {
                if enter_side == d1 {
                    return Some(d2.clone())
                };
                if enter_side == d2 {
                    return Some(d1.clone())
                };
                None
            },
        }
    }
}

#[derive(Debug)]
pub struct Map {
    pub map: Vec<Vec<Joint>>,
}
impl Map {
    fn get(&self, p: &Position) -> &Joint {
        self.map.get(p.y).expect("invalid y").get(p.x).expect("invalid x")
    }

    fn follow(&self, enter_side: &Direction, p: &Position) -> Option<(Direction, Position)> {
        let j = self.get(p);
        if let Some(movement_direction) = j.get_movement_direction(enter_side) {
            return Some((movement_direction.opposite(), movement_direction.move_dir(p)))
        }
        return None
    }

    fn from_str(str: &str) -> (Self, Position) {
        let mut start_position = Position{y: 0, x: 0};
        let mut map = vec![];
        for (y, line) in str.lines().enumerate() {
            let mut map_line = vec![];
            for (x, c) in line.chars().enumerate() {
                map_line.push(match c {
                    '|' => Joint::Pipe(Direction::N, Direction::S),
                    '-' => Joint::Pipe(Direction::E, Direction::W),
                    'L' => Joint::Pipe(Direction::N, Direction::E),
                    'J' => Joint::Pipe(Direction::N, Direction::W),
                    '7' => Joint::Pipe(Direction::S, Direction::W),
                    'F' => Joint::Pipe(Direction::S, Direction::E),
                    '.' => Joint::G,
                    'S' => {
                        start_position = Position{y, x};
                        Joint::Start
                    },
                    _ => panic!("invalid char")
                })
            }

            map.push(map_line)
        }


        (Map{map}, start_position)
    }
}

pub fn solve(puzzle: &str) -> (Map, Vec<Position>, i32) {
    let (map, start_position) = Map::from_str(puzzle);

    for start_direction in [
        Direction::N,
        Direction::E,
        Direction::S,
        Direction::W,
    ] {
        let mut path = vec![];
        let mut position = start_position;
        let mut coming_from_direction = start_direction;
        let mut found_error = false;
        while map.get(&position) != &Joint::Start || path.len() == 0 {
            path.push(position.clone());
            if let Some((next_direction, next_position)) = map.follow(&mut coming_from_direction, &mut position) {
                position = next_position;
                coming_from_direction = next_direction;
            } else {
                found_error = true;
                break;
            }
        }

        if !found_error {
            let len = path.len() as i32;
            return (map, path, len / 2);
        }
    }
    panic!("no solution found")
}