use std::env;

mod cmds;
mod colors;
mod conf;
mod errs;
mod lib;

pub const CONFIG_PATH: &'static str = "~/.novem/"; // main dotfiles path

fn main() {
    let args: Vec<String> = env::args().skip(1).collect();

    if args.len() == 0 {
        colors::prints::no_args_finger(); // do nothing on no args -- not an error
        
    } else {
        let bus = cmds::args::parse_args(args);
        cmds::evals::eval(bus);
    }
}
