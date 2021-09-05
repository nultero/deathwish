use std::fs;

extern crate chrono;
use chrono::{DateTime, Local};

#[allow(dead_code)]
pub struct NovemFile {
    pub name: String,
    pub timestamp: String,
}

#[allow(dead_code)]
pub struct NovemDir {
    pub name: String,
    pub files: Vec<NovemFile>,
    pub dirs: Vec<NovemDir>,
}

impl NovemDir {
    // have to get the metadata passed in as param actually
    // won't be getting whole filepath in if recursive
    pub fn nv_add(&mut self, f: &str) {
        let _f = fs::metadata(f).unwrap();
        if _f.is_dir() {
            let mut _nv = NovemDir {
                name: f.to_owned(),
                files: Vec::new(),
                dirs: Vec::new(),
            };
            self.dirs.push(_nv);
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
