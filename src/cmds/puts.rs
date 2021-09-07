use crate::argsutils::bus::LogicBus;
// use crate::cmds::nov_funcs::*;
use crate::cmds::nov_structs::NovemDir;
use crate::colors::prints::NOVEM_NINE;
use crate::colors::prints::emph;
use std::fs;

extern crate chrono;
use chrono::{DateTime, Local};

pub fn puts(b: LogicBus, _c: &str) {

    if b.paths.len() == 0 {
        println!("{} no filepaths passed to novem", NOVEM_NINE);
        return;
    }

    // the idea for puts is to create a searchable tree
    // and break each of the incoming paths
    // into pieces digestible into the Novem tree struct
    // and eventually commit these to the novem file
    // the user dir is a weird exception in that I don't
    // want to commit the home/ separately

    let nv = NovemDir::basic(b.user_dir.to_owned());

    for p in b.paths {
        let f = fs::metadata(p.to_owned()).unwrap();
        if f.is_dir() {
            if b.recursive {
                // _nv.name = f.to_owned();
                println!("skipping the dir walk for now");
                
            } else {
                println!("{}", format!("skipping dir '{}' -- recurse flag not set", emph(&p)));
            }

        } else if f.is_file() {
            let time = f.modified().unwrap();
            let time: DateTime<Local> = DateTime::from(time);
            let time = time.to_rfc2822();

            let s = format!("|||{}", time);
            let mut sub = p.replace(&nv.name, &"");
            sub.push_str(&s);
            nv.digest(sub);
        }
    }

    // run checks for flags / opts

    // let _nov: NovemDir = rip_str_to_nov(read_conf(c));

    // for i in 0..b.paths.len() {
    //     let md = fs::metadata(&b.paths[i]).unwrap();
    //     let modtime: DateTime<Local> = DateTime::from(md.modified().unwrap());
    //     let fmt_modtime = modtime.to_rfc2822();

    //     println!("{} >>> {:?}", &b.paths[i], fmt_modtime);

    //     b.paths[i] = b.paths[i].replace(&b.user_dir, "");

    //     println!("|( {}", &b.paths[i]);
    // }

    // read_conf(c);

    // let n = Local::now().to_rfc2822();
    // println!("{}", n);
    // println!("{}", String::from(&n));
}
