use crate::argsutils::bus::LogicBus;
use crate::cmds::puts::puts;
use crate::cmds::list::list;
use crate::cmds::utils::get_user_dir;

use crate::cmds::config::{conf_exists, confirm_conf, create_conf, get_conf};

pub fn eval(mut b: LogicBus) {
    let c = get_conf();

    if conf_exists(&c) {
        b.user_dir = get_user_dir();
        b.add_dir_called_from();

        match b.function.as_str() {
            "puts" => puts(b, &c),
            "list" => list(b, &c),
            _ => println!("args should be validated, so this will never print"),
        }
    } else {
        if confirm_conf(&c) {
            create_conf(&c);
        } else {
            println!("exiting")
        }
    }
}
