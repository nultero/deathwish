use crate::errs::errors::throw_err;

pub fn get_user_dir() -> String {
    
    let user_dir = dirs::home_dir();

    if !user_dir.is_some() {
        throw_err("problem getting your home directory");
    }

    let user_dir = user_dir.unwrap();
    let user_dir = user_dir.to_str().unwrap();

    return user_dir.to_owned();
}