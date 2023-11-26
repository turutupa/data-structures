use std::{
    collections::HashMap,
    fs,
    path::{Path, PathBuf},
};

use crate::{KvsError, Result};

pub struct KvStore {
    // path: PathBuf,
    store: HashMap<String, String>,
}

impl KvStore {
    pub fn open(path: impl Into<PathBuf>) -> Result<KvStore> {
        let path = path.into();
        fs::create_dir_all(&path)?;

        Ok(KvStore {
            store: HashMap::new(),
        })
    }

    pub fn set(&mut self, key: String, value: String) -> Result<()> {
        Ok(())
    }

    pub fn get(&mut self, key: String) -> Result<Option<String>> {
        Ok(Some("".to_string()))
    }

    pub fn remove(&mut self, key: String) -> Result<()> {
        Ok(())
    }
}

mod tests {
    use super::*;

    #[test]
    fn get_config_dir() {}
}
