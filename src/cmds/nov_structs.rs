// use std::fs;

pub fn read_conf(cf: &str) {
    println!("cf : {}", &cf);

    // let s = fs::read_to_string(cf).unwrap();

    // still need to iterate over the json

    // let _dots = NovemDir {
    //     name: String::from(""),
    //     files: vec![],
    // };

    // for i in j {
    //     dots.files.push()
    // }
    // println!("{:?}", _dots);
}

#[allow(dead_code)]
pub struct NovemFile {
    pub name: String,
    pub timestamp: String,
}

#[allow(dead_code)]
pub struct NovemDir {
    pub name: String,
    pub files: Vec<NovemFile>,
}
