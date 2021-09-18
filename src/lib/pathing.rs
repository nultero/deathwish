use crate::errs::errors::throw_err;
use crate::CONFIG_PATH;

pub fn get_user_dir() -> String {
    let user_dir = dirs::home_dir();

    if !user_dir.is_some() {
        throw_err("problem getting your home directory");
    }

    let user_dir = user_dir.unwrap();
    let user_dir = user_dir.to_str().unwrap();

    return user_dir.to_owned();
}

pub fn expand_tilde_home() -> String {

    let mut _home = String::new(); 
    // compiler complains 'home' is overwritten
    // before being read -- that's why it's here

    // expand $HOME to full path -- canonicalize does not do this
    let first_char = CONFIG_PATH.chars().nth(0).unwrap();
    if first_char == '~' {
        _home = get_user_dir();
        let sl = &CONFIG_PATH[1..];
        _home = format!("{}{}", _home, sl);
    } else {
        _home = CONFIG_PATH.to_owned();
    }

    return _home;
}