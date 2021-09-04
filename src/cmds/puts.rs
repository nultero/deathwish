use crate::argsutils::bus::LogicBus;
use crate::cmds::nov_structs::read_conf;
use crate::colors::prints::NOVEM_NINE;

use std::fs;

extern crate chrono;

use chrono::{Local, DateTime};

pub fn puts(mut b: LogicBus, c: &str) {
    if b.paths.len() == 0 {
        println!("{} no filepaths passed to novem", NOVEM_NINE);
        return;
    }

    println!("usr root: {}", &b.user_dir.to_string());

    for i in 0..b.paths.len() {

        let md = fs::metadata(&b.paths[i]).unwrap();
        let modtime: DateTime<Local> = DateTime::from(md.modified().unwrap());
        let fmt_modtime = modtime.to_rfc2822();

        println!("{} >>> {:?}", &b.paths[i], fmt_modtime);

        b.paths[i] = b.paths[i].replace(&b.user_dir.to_string(), "");
    }

    read_conf(c);


    let n = Local::now().to_rfc2822();
    println!("{}", n);
    println!("{}", String::from(&n));
}
