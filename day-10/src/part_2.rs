use std::collections::{HashMap, HashSet};
use std::fmt::{Debug, Formatter, Write};
use std::time::Instant;
use crate::part_1;
use crate::part_1::{Position, Joint};
use crate::part_2::Point::{Empty, Occupied, Path};

// Same as Position, just to keep track of what is scaled
#[derive(Hash, Eq, PartialEq, Copy, Clone)]
struct ScaledPosition {
    y: usize,
    x: usize
}
impl ScaledPosition {
    fn from_position(p: &Position) -> Self {
        ScaledPosition{y: p.y*3 + 1, x: p.x*3 + 1}
    }

    fn adjacent_points(&self) -> Vec<ScaledPosition> {
        let mut adjacent = vec![];
        for y in 0..3 {
            for x in 0..3 {
                if self.x + x == 0 || self.y + y == 0 {
                    continue
                }
                if y == 1 && x == 1 {
                    continue
                }
                adjacent.push(ScaledPosition{y: self.y + y - 1, x: self.x + x - 1})
            }
        }

        adjacent
    }
}

#[derive(Debug, Eq, PartialEq, Copy, Clone)]
enum Point {
    Empty,
    Path,
    Occupied,
}

struct Map {
    map: Vec<Vec<Point>>
}
impl Map {
    fn from_part_1(old_map: &part_1::Map, path_map: &HashSet<Position>) -> Self {
        let mut map = vec![vec![Empty; old_map.map[0].len() * 3]; old_map.map.len() * 3];
        for (y, row) in old_map.map.iter().enumerate() {
            for (x, square) in row.iter().enumerate() {
                let p = Position{y, x};
                let new_p = Position{y: y*3 + 1, x: x*3 + 1};

                match square {
                    Joint::G => {}
                    Joint::Start => {
                        // Idk what start connects to....?
                        // Maybe fill in from path?
                        for dy in 0..3 {
                            for dx in 0..3 {
                                map[new_p.y + dy - 1][new_p.x + dx - 1] = Path
                            }
                        }
                    }
                    Joint::Pipe(d1, d2) => {
                        let on_path = path_map.contains(&p);

                        let d1_p = d1.move_dir(&new_p);
                        let d2_p = d2.move_dir(&new_p); // Translation isn't working 0,0 still maps to corner

                        map[new_p.y][new_p.x] = if on_path { Path } else { Occupied };
                        map[d1_p.y][d1_p.x] = if on_path { Path } else { Occupied };
                        map[d2_p.y][d2_p.x] = if on_path { Path } else { Occupied };
                    }
                }
            }
        }

        Map{map}
    }

    fn get(&self, p: &ScaledPosition) -> &Point {
        self.map.get(p.y).expect("invalid y").get(p.x).expect("invalid x")
    }

    fn is_point_enclosed(&self, p: &ScaledPosition, already_seen: &mut HashMap<ScaledPosition, bool>) -> bool {
       let mut visited: HashSet<ScaledPosition> = HashSet::new();

        let mut to_visit = p.adjacent_points();
        let mut to_visit_index = 0;

        while to_visit_index < to_visit.len() {
            let visiting = *to_visit.get(to_visit_index).expect("somehow out of range");

            if already_seen.contains_key(&visiting) {
                let outcome = already_seen.get(&visiting).unwrap().clone();
                already_seen.insert(*p, outcome);
                to_visit.iter().for_each(|p| {
                    already_seen.insert(*p, outcome);
                });
                return outcome
            }

            if visited.contains(&visiting) {
                to_visit_index += 1;
                continue
            }

            if self.is_border_point(&visiting) {
                already_seen.insert(*p, false);
                to_visit.iter().for_each(|p| {
                    already_seen.insert(*p, false);
                });
                return false
            }

            if self.get(&visiting) == &Path {
                to_visit_index += 1;
                continue
            }
            visited.insert(visiting);
            // Store the adjacent points in a temporary vector
            let mut adjacent_points = visiting.adjacent_points();
            to_visit.append(&mut adjacent_points);

            to_visit_index += 1;
        }
        already_seen.insert(*p, true);
        to_visit.iter().for_each(|p| {
            already_seen.insert(*p, true);
        });

        true
    }

    fn is_border_point(&self, p: &ScaledPosition) -> bool {
        if p.y == self.map.len() -1 || p.y == 0 {
            return true;
        }
        if p.x == self.map[0].len() - 1 || p.x == 0 {
            return true;
        }
        return false;
    }
}

impl Debug for Map {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        for line in self.map.iter() {
            for point in line.iter() {
                match point {
                    Empty => f.write_char('.'),
                    Path => f.write_char('@'),
                    Occupied => f.write_char('x'),
                };
            }
            f.write_str("\n");
        }
        return Ok(())
    }
}

pub fn solve(original_map: part_1::Map, path: Vec<Position>) -> i32 {
    let path_map: HashSet<Position> = path.into_iter().collect();
    let sparse_map = Map::from_part_1(&original_map, &path_map);


    let mut outcome_map: HashMap<ScaledPosition, bool> = HashMap::new();

    let start = Instant::now();

    let mut count = 0;
    for (y, line) in original_map.map.iter().enumerate() {
        for (x, point) in line.iter().enumerate() {
            let position = Position{y, x};
            if path_map.contains(&position) {
                continue
            }
            if sparse_map.is_point_enclosed(&ScaledPosition::from_position(&position), &mut outcome_map) {
                count += 1;
            }
        }
    }

    println!("since {:?}", start.elapsed());
    // without map 10.737086542s

    count
}