use crate::lib::nov_structs::NovemDir;
use crate::errs::errors;
use std::fs;

pub fn read_conf(cf_path: &str) -> String {
    let s = fs::read_to_string(cf_path);
    if !s.is_ok() {
        errors::throw_err("problem opening novem data file");
    }

    return s.unwrap();
}

pub fn rip_str_to_nov(s: String) -> NovemDir {
    let mut nv = NovemDir::new();
    let mut lines = s.lines();
    nv.name = lines.next().unwrap().to_owned();

    println!("ripfn: {}", &nv.name);

    for ln in lines {
        if ln.contains("|") {
            println!("ripln: {}", &ln);
        } else {
            nv.dirs.push(rip_str_to_nov(ln.to_string()));
        }
    }

    return nv;
}
