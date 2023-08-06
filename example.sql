

CREATE TABLE IF NOT EXISTS fakultas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS jurusan (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    fakultas_id INT NOT NULL,
    FOREIGN KEY (fakultas_id) REFERENCES fakultas(id)
);

CREATE TABLE IF NOT EXISTS matkulfakultas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    sks INT NOT NULL,
    semestermin INT NOT NULL,
    fakultas_id INT NOT NULL,
    prediksinilai VARCHAR(255) NOT NULL,
    FOREIGN KEY (fakultas_id) REFERENCES fakultas(id)
);

CREATE TABLE IF NOT EXISTS matkuljurusan (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    sks INT NOT NULL,
    semestermin INT NOT NULL,
    jurusan_id INT NOT NULL,
    prediksinilai VARCHAR(255) NOT NULL,
    FOREIGN KEY (jurusan_id) REFERENCES jurusan(id)
);