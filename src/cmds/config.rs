use crate::cmds::utils::get_user_dir;
use crate::errs::errors::sys_err;

use crate::prints::{blue, emph, NOVEM_NINE};
use crate::CONFIG_PATH;

use std::fs;
use std::io::*;
use std::path::Path;

extern crate chrono;
use chrono::Local;

pub fn get_conf() -> String {
    let mut _home: String = String::default();

    // expand $HOME to full path -- canonicalize apparently does not do this
    let cf_chars: Vec<char> = CONFIG_PATH.chars().collect();
    if cf_chars[0] == '~' {
        _home = get_user_dir();
        let sl = &CONFIG_PATH[1..];
        _home = format!("{}{}", _home, sl);
    } else {
        _home = CONFIG_PATH.to_owned();
    }

    return format!("{}novem.txt", _home);
}

pub fn conf_exists(home: &str) -> bool {
    let cf = Path::new(&home).canonicalize();
    match cf {
        Ok(_p) => return true,
        Err(_e) => return false,
    }
}

pub fn confirm_conf(conf_path: &str) -> bool {
    let c = blue(conf_path);
    println!("conf does not seem to exist at `{}`", &c);

    let mut out = stdout();
    let emphasize_input = emph(&"[ Y / n ]");
    let prompt = format!("\ncreate conf at this location? {} ", &emphasize_input);
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

pub fn create_conf(conf_path: &str) {
    // conf_path format:
    // `/home/$user/.novem/novem.txt`

    let nv_dir_str = &conf_path.replace("novem.txt", "");
    let nv_dir = Path::new(&nv_dir_str);
    let nv_path = nv_dir.canonicalize();

    if !nv_path.is_ok() {
        fs::create_dir(&nv_dir).unwrap();
        let em = emph(&nv_dir_str);
        println!("{} created dir at `{}`", NOVEM_NINE, &em);
    } else {
        nv_path.unwrap();
    }

    let dir_str_spl: Vec<&str> = conf_path.split("/").filter(|s| s.len() > 0).collect();

    let user_dir = format!("{}{}", get_user_dir(), "\n");
    let mut nv_str = String::from(&user_dir);

    let now = Local::now().to_rfc2822();

    for i in 2..dir_str_spl.len() {
        let tabs = "\t".repeat(i - 1);
        let cur = &dir_str_spl[i];

        let substr = format!("{}{} | {}\n", tabs, cur, now);
        nv_str.push_str(&substr);
    }

    if nv_str.ends_with('\n') {
        nv_str = nv_str[..nv_str.len() - 1].to_string();
    }

    let res = fs::write(&conf_path, &nv_str);

    if res.is_ok() {
        println!("{} created config", NOVEM_NINE);
    } else {
        sys_err(res.err().unwrap());
    }
}
