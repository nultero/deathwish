use crate::argsutils::bus::LogicBus;
use crate::cmds::nov_funcs::*;
use crate::cmds::nov_structs::NovemDir;
use crate::colors::prints::NOVEM_NINE;

use std::fs;

extern crate chrono;
use chrono::{DateTime, Local};

pub fn list(mut b: LogicBus, c: &str) {
    if b.paths.len() == 0 {
        println!("{} no filepaths passed to novem", NOVEM_NINE);
        return;
    }

    // run checks for flags / opts

    let _nov: NovemDir = rip_str_to_nov(read_conf(c));

    for i in 0..b.paths.len() {
        let md = fs::metadata(&b.paths[i]).unwrap();
        let modtime: DateTime<Local> = DateTime::from(md.modified().unwrap());
        let fmt_modtime = modtime.to_rfc2822();

        println!("{} >>> {:?}", &b.paths[i], fmt_modtime);

        b.paths[i] = b.paths[i].replace(&b.user_dir, "");

        println!("|( {}", &b.paths[i]);
    }

    // read_conf(c);

    // let n = Local::now().to_rfc2822();
    // println!("{}", n);
    // println!("{}", String::from(&n));
}