use std::env;

mod bus;
mod prints;
mod args;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() == 1 {
        prints::no_args_finger();

    } else { // args were passed

        let rgs = args[1..].to_owned();

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

        println!("{:?}", bus.conf_path);
    }


    // let b = bus::LogicBus {

    // }

}
