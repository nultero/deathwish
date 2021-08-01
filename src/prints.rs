use std::process;

const RED_FINGER: &str = "\x1b[31;1m\u{261E}\x1b[0m ";
pub const NOVEM_NINE: &str = "\x1b[32;1m\u{277e}\x1b[0m ";

pub fn no_args_finger() {
    println!("{}  no args", RED_FINGER);
}

pub fn err_finger(s: &str) {
    println!("{} '{}' is not a recognized function, flag, or path", RED_FINGER, s);
    process::exit(1);
}

pub fn err_flag(s: &str) {
    println!("{} '{}' is not a recognized flag", RED_FINGER, s);
    process::exit(1);
}

pub fn diff_err_flag(s: &str) {
    println!("{} '{}' is not a recognized time format for the -d / diff flag", RED_FINGER, s);
    process::exit(1);
}

pub fn diff_exists_err() {
    println!("{} too many diff options passed in", RED_FINGER);
    process::exit(1);
}

pub fn fmt_before_diff_flag(s: &str) {
    println!("{} -d / diff option '{}' passed in before diff flag", RED_FINGER, s);
    process::exit(1);
}