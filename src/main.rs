use std::env;

mod bus;
mod prints;
mod args;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() == 1 {
        prints::no_args_finger();

    } else { // args were passed

        let mut rgs = args[1..].to_owned();

        let mut bus = bus::LogicBus {
            conf_path: String::from(""),

        }

        while !args::is_empty(&rgs) {
            let cur = &rgs[0];
            println!("{} ", cur);
            &rgs.remove(0);
        }

        println!("{:?}", rgs)
    }


    // let b = bus::LogicBus {

    // }

}
