use crate::argsutils::bus::LogicBus;
use crate::errs::errors;

use std::path::Path;

fn is_empty(args: &Vec<String>) -> bool {
    return args.len() == 0;
}

const VALID_FUNCS: [&str; 6] = ["puts", "list", "tree", "remove", "unpack", "zipto"];

fn is_func(r: &str) -> bool {
    return VALID_FUNCS.contains(&r);
}

fn is_flag(r: &str) -> bool {
    return r.contains("-");
}

pub fn is_path(s: &str) -> bool {
    let p = Path::new(s).canonicalize();
    match p {
        Ok(_r) => return true,
        Err(_e) => return false,
    }
}

fn is_diff(r: &str, diff_flag: bool) -> bool {
    if r.len() != 2 {
        return false;
    }

    let rg_chars: Vec<char> = r.chars().collect();
    let sec_char: char;

    match &rg_chars[0] {
        '0'..='9' => sec_char = rg_chars[1],
        _ => return false,
    }

    let valid_time_fmt = match sec_char {
        'd' | 'm' => true,
        _ => false,
    };

    if !valid_time_fmt {
        errors::diff_err_flag(&String::from(sec_char));
    }

    if !diff_flag {
        errors::fmt_before_diff_flag(&String::from(r));
    }

    return true;
}

pub fn parse_args(mut args: Vec<String>) -> LogicBus {
    let mut b = LogicBus {
        // defaults
        conf_path: String::from(""),
        function: String::from(""),
        verbosity: 0,
        paths: vec![],
        help: false,
        diff: false,
        diff_opts: String::from(""),
        recursive: false,
    };

    while !is_empty(&args) {
        let cur = &args[0];

        match cur {
            //ngl, but this "_" match syntax looks ugly, might be an antipattern
            _ if is_func(&cur) => b.assign_func(&cur),
            _ if is_flag(&cur) => b.assign_flag(&cur),
            _ if is_path(&cur) => b.add_path(&cur),
            _ if is_diff(&cur, b.diff) => b.add_diff_opt(&cur),
            _ => errors::err_finger(&cur),
        }

        &args.remove(0);
    }

    for p in &b.paths {
        println!("'{}' is in bus.paths", p);
    }

    return b;
}
