use std::cmp::Ordering;
use std::collections::HashMap;
use std::fmt::{Debug, Error, Formatter};

#[derive(PartialEq, Eq, Hash, PartialOrd, Clone)]
pub struct Card(i32);
impl Card {
    pub fn new(c: &char) -> Card {
        if c < &'2' {
            panic!("card too low")
        }
        if c.is_numeric() {
            return Card(c.to_string().parse().expect("invalid number"))
        }

        match c {
            &'T' => Card(10),
            &'J' => Card(11),
            &'Q' => Card(12),
            &'K' => Card(13),
            &'A' => Card(14),
            &_ => panic!("invalid card")
        }
    }

    fn is_wildcard(&self) -> bool {
        return self.0 == 11
    }
}


impl Debug for Card {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        if self.0 >= 10 {
            match self.0 {
                10 => f.write_str("T"),
                11 => f.write_str("J"),
                12 => f.write_str("Q"),
                13 => f.write_str("K"),
                14 => f.write_str("A"),
                _ => return Err(Error),
            }?;
            return Ok(())
        }

        f.write_str(&*self.0.to_string())?;
        Ok(())
    }
}

#[derive(PartialEq, PartialOrd, Debug, Eq)]
pub enum HandType {
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}

#[derive(Debug)]
pub struct Hand {
    pub cards: Vec<Card>,
}
impl Hand {
    pub fn get_kind(&self) -> HandType {
        let mut number_of_wildcards = 0;
        let mut freq: HashMap<&Card, i32> = HashMap::new();
        self.cards.iter().for_each(| card| {
            if card.is_wildcard() {
                number_of_wildcards += 1;
            } else {
                *freq.entry(card).or_default() += 1;
            }
        });

        let (largest_card, _) = freq.iter().max_by(|a, b| a.1.cmp(&b.1)).unwrap_or((&&Card(2), &0));
        *freq.entry(largest_card).or_default() += number_of_wildcards;

        match freq.len() {
            1 => HandType::FiveOfAKind,
            2 => {
                // four or full house
                if freq.values().any(|f| f == &4) {
                    return HandType::FourOfAKind
                }
                return HandType::FullHouse
            },
            3 => {
                // 3 or 2 pair
                if freq.values().any(|f| f == &3) {
                    return HandType::ThreeOfAKind
                }
                return HandType::TwoPair
            }
            4 => HandType::OnePair,
            _ => HandType::HighCard
        }
    }
}

impl PartialEq<Self> for Hand {
    fn eq(&self, other: &Self) -> bool {
        return false;
    }
}

impl PartialOrd<Self> for Hand {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        let s = self.get_kind();
        let o = other.get_kind();
        if s == o {
            for (i, card) in self.cards.iter().enumerate() {
                let other_card = other.cards.get(i).expect("different lengths");
                if card != other_card {
                    if card.is_wildcard() {
                        return Some(Ordering::Less)
                    }
                    if other_card.is_wildcard() {
                        return Some(Ordering::Greater)
                    }
                    return card.partial_cmp(&other_card);
                }
            }
            return Some(Ordering::Equal)
        }
        return s.partial_cmp(&o);
    }
}