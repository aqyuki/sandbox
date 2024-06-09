fn main() {
    let text = "Hello World";
    let chars = text.chars().collect::<Vec<_>>();
    println!("{:?}", text);
    println!("{:?}", chars);
}
