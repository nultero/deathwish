use std::fs;

pub fn read_conf(cf: &str) {
    let s = fs::read_to_string(cf).unwrap();
    let j: serde_json::Value = serde_json::from_str(&s).unwrap();
    let j = j.as_object().unwrap();

    // still need to iterate over the json

    let _dots = NovemDir {
        name: String::from(""),
        files: vec![],
    };

    // for i in j {
    //     dots.files.push()
    // }
    println!("{:?}", j);
}

pub struct NovemFile {
    pub name: String,
    pub timestamp: String,
}

pub struct NovemDir {
    pub name: String,
    pub files: Vec<NovemFile>,
}
