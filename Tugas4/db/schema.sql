CREATE TABLE produk (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    nama       VARCHAR(100) NOT NULL,
    harga      BIGINT NOT NULL,
    stok       INT NOT NULL DEFAULT 0,
    keterangan TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO produk (nama, harga, stok, keterangan) VALUES
    ('Kopi Arabika Gayo 250g', 85000, 40, 'Biji kopi sangrai medium dari Aceh'),
    ('Teh Melati Tubruk', 12500, 120, NULL),
    ('Gula Aren Cair 500ml', 32000, 0, 'Stok habis, restock minggu depan'),
    ('Keripik Tempe Pedas', 15000, 75, NULL),
    ('Madu Hutan Sumbawa', 98000, 18, 'Madu murni, cocok untuk oleh-oleh');
