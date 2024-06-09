fn main() {
    let mut v = vec![0; 5];
    v.push(2);
    println!("{:?}", v);
    v.insert(1, 1);
    println!("{:?}", v);
    v.swap_remove(1);
    println!("{:?}", v);

    let v1: &[i32];
    v1 = &v;
    let v2 = &v1[2..4];

    println!("{:?}", v1);
    println!("{:?}", v2);
}
