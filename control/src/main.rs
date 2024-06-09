fn main() {
    let f = 1;
    let _v: Box<dyn Iterator<Item = i32>> = if f == 1 {
        Box::new(0..10)
    } else {
        Box::new((0..10).filter(|x| x % 2 == 0))
    };
}
