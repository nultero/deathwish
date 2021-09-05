use std::env;

mod argsutils {
    pub mod args;
    pub mod bus;
}
mod colors {
    pub mod prints;
}
mod cmds {
    pub mod config;
    pub mod evals;
    pub mod nov_structs;
    pub mod puts;
    pub mod utils;
}
mod errs {
    pub mod errors;
}

use argsutils::args;
use cmds::evals;
use colors::prints;

pub const CONFIG_PATH: &str = "~/.novem/"; // main dotfiles json path

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() == 1 {
        prints::no_args_finger(); // do nothing on no args -- not an error
    } else {
        let rgs = args[1..].to_vec();
        let bus = args::parse_args(rgs);
        evals::eval(bus);
    }
}
