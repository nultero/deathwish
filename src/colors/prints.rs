pub const RED_FINGER: &str = "\x1b[31;1m\u{261E}\x1b[0m ";
pub const NOVEM_NINE: &str = "\x1b[32;1m\u{277e}\x1b[0m ";

pub fn no_args_finger() {
    println!("{}  no args", RED_FINGER);
}

pub fn blue(s: &str) -> String {
    let s = format!("\x1b[32;1;4m{}\x1b[0m", s);
    return s;
}

pub fn emph(s: &str) -> String {
    let s = format!("\x1b[;1;1m{}\x1b[0m", s);
    return s;
}
