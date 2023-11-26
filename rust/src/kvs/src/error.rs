use std::io;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum KvsError {
    #[error("{0}")]
    Io(#[from] io::Error),

    #[error("`{0}`")]
    Serde(#[from] serde_json::Error),

    /// Removing non-existent key error.
    #[error("key not found")]
    KeyNotFound,

    #[error("Unexpected command type")]
    UnexpectedCommandType,
}

/// Result type for kvs.
pub type Result<T> = std::result::Result<T, KvsError>;
