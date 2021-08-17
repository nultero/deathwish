use std::env;

mod bus;
mod prints;
mod args;
mod evals;
mod nov;


pub const CONFIG_PATH: &str = "~/.novem/"; // main dotfiles json path


fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() == 1 {
        prints::no_args_finger(); // do nothing on no args -- not an error

    } else {

        // finish out the flags' breakdown

        let rgs = args[1..].to_vec();

        let mut bus = bus::LogicBus {// defaults
            conf_path:      String::from(""),
            function:       String::from(""),
            verbosity:      0,
            paths:          vec![],
            help:           false,
            diff:           false,
            diff_opts:      String::from(""),
            recursive:      false
        };

        bus = args::parse_args(rgs, bus);

        evals::eval(bus);
    }
}
