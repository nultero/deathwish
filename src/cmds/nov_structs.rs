
pub struct NovemFile {
    pub name: String,
    pub timestamp: String,
}

pub struct NovemDir {
    pub name: String,
    pub files: Vec<NovemFile>,
    pub dirs: Vec<NovemDir>,
}

impl NovemDir {

    #[allow(dead_code)]
    pub fn basic(user_dir: String) -> NovemDir {
        return NovemDir {
            name: user_dir,
            files: vec!(),
            dirs: vec!()
        }
    }

    #[allow(dead_code)]
    pub fn from(_path: String, _user_dir: String) -> NovemDir {
        return NovemDir::new();
    }

    #[allow(dead_code)]
    pub fn has_node(_e: String) -> bool {
        return true
    }

    pub fn new() -> NovemDir {
        return NovemDir {
            name: String::new(),
            files: Vec::new(),
            dirs: Vec::new(),
        };
    }

    #[allow(dead_code)]
    fn add() {

    }

    pub fn digest(&self, mut path: String) {

        if path.starts_with("/") {
            path = path[1..].to_string();
        }

        println!("{}", &path);
    }
}

#[allow(dead_code)]
fn break_root_path(mut path: String, user_dir: String) -> Vec<String> {

    let mut tmp = "".to_owned();
    if path.contains(&user_dir) {
        path = path.replace(&user_dir, "");
        tmp.push_str(&user_dir);
    } // else edge case -- moving to new machine, etc.

    let spl = path.split("/").filter(|s| s.len() > 1);
    let mut out: Vec<String> = vec!(tmp);

    for s in spl {
        out.push(s.to_owned());
    }

    return out;
}