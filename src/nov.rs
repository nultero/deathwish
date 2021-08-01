
pub struct NovemFile {
    pub name:           String,
    pub timestamp:      String,
}

pub struct NovemDir {
    pub name:   String,
    pub files:  Vec<NovemFile>
}