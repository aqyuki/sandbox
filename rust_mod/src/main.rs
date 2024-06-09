mod animal;

fn main() {
    let poti = animal::Dog {
        name: String::from("poti"),
        age: 10,
        fav: animal::Fav {
            food: String::from("bone"),
        },
    };

    let neko = animal::Mike {
        name: String::from("tama"),
        age: 5,
    };

    println!("{:?}", poti);
    println!("{:?}", neko);
}
