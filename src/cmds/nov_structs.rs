use std::fs;

extern crate chrono;
use chrono::{DateTime, Local};

#[allow(dead_code)]
pub struct NovemFile {
    pub name: String,
    pub timestamp: String,
}

pub struct NovemDir {
    pub name: String,
    pub files: Vec<NovemFile>,
    pub dirs: Vec<NovemDir>,
}

impl NovemDir {
    pub fn new() -> NovemDir {
        return NovemDir {
            name: String::new(),
            files: Vec::new(),
            dirs: Vec::new(),
        };
    }

    // have to get the metadata passed in as param actually
    // won't be getting whole filepath in if recursive
    #[allow(dead_code)]
    pub fn nv_add(&mut self, f: &str, recurse: bool) {
        let _f = fs::metadata(f).unwrap();
        if _f.is_dir() {
            if recurse {
                let mut _nv = NovemDir::new();
                _nv.name = f.to_owned();
                self.dirs.push(_nv);
            } else {
                println!("");
            }

        } else if _f.is_file() {
            let time = _f.modified().unwrap();
            let time: DateTime<Local> = DateTime::from(time);
            let time = time.to_rfc2822();

            let _nf = NovemFile {
                name: f.to_owned(),
                timestamp: time,
            };
        }
    }
}
