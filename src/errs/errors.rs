use crate::colors;

use colors::prints::RED_FINGER;
use std::process;

use std::io::Error;

pub fn sys_err(e: Error) {
    println!("{} {}", RED_FINGER, &e);
    process::exit(1);
}

pub fn throw_err(err_description: &str) {
    println!("{} {}", RED_FINGER, err_description);
    process::exit(1);
}

pub fn err_finger(s: &str) {
    println!(
        "{} '{}' is not a recognized function, flag, or path",
        RED_FINGER, s
    );
    process::exit(1);
}

pub fn err_flag(s: &str) {
    println!("{} '{}' is not a recognized flag", RED_FINGER, s);
    process::exit(1);
}

pub fn diff_err_flag(s: &str) {
    println!(
        "{} '{}' is not a recognized time format for the -d / diff flag",
        RED_FINGER, s
    );
    process::exit(1);
}

pub fn diff_exists_err() {
    println!("{} too many diff options passed in", RED_FINGER);
    process::exit(1);
}

pub fn fmt_before_diff_flag(s: &str) {
    println!(
        "{} -d / diff option '{}' passed in before diff flag",
        RED_FINGER, s
    );
    process::exit(1);
}
