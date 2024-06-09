fn main() {
    let v = 0..10;
    println!("{:?}", v);
    let v2 = v.filter(|x| x % 2 == 0).collect::<Vec<_>>();
    println!("{:?}", v2);

    for x in &v2 {
        println!("{}", x);
    }
}
