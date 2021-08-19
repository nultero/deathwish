use crate::argsutils::bus::LogicBus;
use crate::cmds::funcs::nov_puts;
use crate::errs::errors::sys_err;
use crate::prints::{blue, emph, NOVEM_NINE};
use crate::CONFIG_PATH;
use dirs::home_dir;

use std::fs;
use std::io::*;
use std::path::Path;

pub fn eval(b: LogicBus) {
    let c = get_conf();

    if conf_exists(&c) {
        match b.function.as_str() {
            "puts" => nov_puts(b, &c),

            _ => println!("args should be validated, so this will never print"),
        }
    } else {
        if confirm_conf(&c) {
            create_conf(&c);
        } else {
            println!("exiting")
        }
    }
}

fn get_conf() -> String {
    let mut _home: String = String::default();

    // expand $HOME to full path -- canonicalize apparently does not do this
    let cf_chars: Vec<char> = CONFIG_PATH.chars().collect();
    if cf_chars[0] == '~' {
        let h = home_dir().unwrap();
        let h = h.to_str();
        _home = h.unwrap().to_owned();

        let sl = &CONFIG_PATH[1..];
        _home = format!("{}{}", _home, sl);
    } else {
        _home = CONFIG_PATH.to_owned();
    }

    return format!("{}novem.json", _home);
}

fn conf_exists(home: &str) -> bool {
    let cf = Path::new(&home).canonicalize();
    match cf {
        Ok(_p) => return true,
        Err(_e) => return false,
    }
}

fn confirm_conf(conf_path: &str) -> bool {
    let c = blue(conf_path);
    println!("conf does not seem to exist at `{}`", &c);

    let mut out = stdout();
    let emphasize = emph(&"[ Y / n ]");
    let prompt = format!("\ncreate conf at this location? {} ", &emphasize);
    out.write(prompt.as_bytes()).unwrap();
    out.flush().unwrap();

    let sin = stdin();
    let mut inp = String::from("");
    sin.read_line(&mut inp).unwrap();

    if !inp.to_lowercase().contains("n") {
        return true;
    }

    return false;
}

fn create_conf(conf_path: &str) {
    let dir = &conf_path.replace("novem.json", "");
    let d = Path::new(&dir);

    let p = d.canonicalize();

    if !p.is_ok() {
        fs::create_dir(&d).unwrap();
        let em = emph(&dir);
        println!("{} created dir at `{}`", NOVEM_NINE, &em);
    } else {
        p.unwrap();
    }

    let res = fs::write(&conf_path, &"");

    if res.is_ok() {
        println!("{} created config", NOVEM_NINE);
    } else {
        sys_err(res.err().unwrap());
    }
}
