use crate::argsutils::bus::LogicBus;
use crate::cmds::utils::get_user_dir;
use crate::cmds::puts::puts;

use crate::cmds::config::{
    get_conf,
    conf_exists,
    confirm_conf,
    create_conf
};

pub fn eval(mut b: LogicBus) {
    let c = get_conf();

    if conf_exists(&c) {
        b.user_dir = get_user_dir();

        match b.function.as_str() {
            "puts" => puts(b, &c),

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



