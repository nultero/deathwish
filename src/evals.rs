
use dirs::home_dir;
use crate::CONFIG_PATH;
use crate::bus::LogicBus;
use crate::prints::NOVEM_NINE;
use crate::nov::{NovemDir};

use std::path::Path;
use std::fs;

pub fn eval(b: LogicBus) {

    let c = get_conf_from_const();

    if conf_exists(&c) {

        match b.function.as_str() {

            "puts" => nov_puts(b, &c),

            _ => println!("{} novelty", NOVEM_NINE)
        }

    } else {
        // create conf
    }
}

fn get_conf_from_const() -> String {
    let mut _home: String = String::default();

    // expand $HOME to full path -- canonicalize apparently does not do this
    let cf_chars: Vec<char> = CONFIG_PATH.chars().collect();
    if cf_chars[0] == '~' {

        let h = home_dir().unwrap();
        let h = h.to_str();
        _home = h.unwrap().to_owned();

        let sl = &CONFIG_PATH[1..];
        _home = format!("{}{}", _home, sl);

    } else {
        _home = CONFIG_PATH.to_owned();
    }

    return format!("{}novem.json", _home);
}

fn conf_exists(home: &str) -> bool {
    let cf = Path::new(&home).canonicalize();
    match cf {
        Ok(_p) => return true,
        Err(_e) => return false
    }
}

fn read_conf(cf: &str)  {

    let s = fs::read_to_string(cf).unwrap();
    let j: serde_json::Value = serde_json::from_str(&s).unwrap();
    let j = j.as_object().unwrap();

    // still need to iterate over the json

    let _dots = NovemDir {
        name: String::from(""),
        files: vec!(),
    };

    // for i in j {
    //     dots.files.push()
    // }
    println!("{:?}", j);
}

fn nov_puts(_b: LogicBus, c: &str) {
    println!("{} puts functions successfully", NOVEM_NINE);

    read_conf(c);
}