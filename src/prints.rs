use std::process;

pub fn no_args_finger() {
    println!("\u{261E}  no args");
}

pub fn err_finger(s: &str) {
    println!("\u{261E}  '{}' is not a recognized function, flag, or path", s);
    process::exit(1);
}

pub fn err_flag(s: &str) {
    println!("\u{261E}  '{}' is not a recognized flag", s);
    process::exit(1);
}

// '9': "\u277e >"
// 'finger': "\u261E",