#[derive(Debug)]
pub struct Dog {
    pub name: String,
    pub age: u64,
    pub fav: Fav,
}

#[derive(Debug)]
pub struct Fav {
    pub food: String,
}
