## Key Value Store

Project proposed in https://github.com/pingcap/talent-plan/tree/master/courses/rust/projects/project-2

Taken from the previous repo, the expected behavior:

- The cargo project, kvs, builds a command-line key-value store client called kvs, which in turn calls into a library called kvs.
- The kvs executable supports the following command line arguments:
    - `kvs set <KEY> <VALUE>` Set the value of a string key to a string. Print an error and return a non-zero exit code on failure.
    - `kvs get <KEY>` Get the string value of a given string key. Print an error and return a non-zero exit code on failure.
    - `kvs rm <KEY>` Remove a given key. Print an error and return a non-zero exit code on failure.
    - `kvs -v` Print the version

The kvs library contains a type, KvStore, that supports the following methods:
- `KvStore::set(&mut self, key: String, value: String) -> Result<()>` Set the value of a string key to a string. Return an error if the value is not written successfully.
- `KvStore::get(&mut self, key: String) -> Result<Option<String>>` Get the string value of a string key. If the key does not exist, return None. Return an error if the value is not read successfully.
- `KvStore::remove(&mut self, key: String) -> Result<()>` Remove a given key. Return an error if the key does not exist or is not removed successfully.
- `KvStore::open(path: impl Into<PathBuf>) -> Result<KvStore>` Open the KvStore at a given path. Return the KvStore.
