CREATE DATABASE spread;

USE spread;

CREATE TABLE files (
    file_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    file_name VARCHAR(100) NOT NULL,
    file_type VARCHAR(30) NOT NULL,
    file_size BIGINT UNSIGNED NOT NULL,
    version INT UNSIGNED NOT NULL DEFAULT 1,
    uploaded_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted INT UNSIGNED NOT NULL DEFAULT 0
);

CREATE TABLE chunks (
    chunk_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    file_id INT UNSIGNED NOT NULL,
    node_id VARCHAR(255) NOT NULL,
    chunk_index INT NOT NULL,
    chunk_size BIGINT NOT NULL,
    chunk_hash VARCHAR(64) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (file_id) REFERENCES files (file_id)
);

CREATE TABLE storage_nodes (
    node_id VARCHAR(255) NOT NULL PRIMARY KEY,
    port INT NOT NULL,
    ip_address VARCHAR(45),
    location VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_heartbeat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);