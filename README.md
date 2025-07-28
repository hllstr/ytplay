# ytplay

Interactive YouTube audio player CLI written in Go.  
Search, download (best‐quality M4A) and play YouTube audio straight from your terminal.

---

## Features

- Interactive prompt — tidak perlu flag  
- Cari lagu berdasarkan keyword  
- Pilih dari 8 hasil teratas  
- Download audio .m4a kualitas tertinggi  
- Simpan otomatis ke `~/Downloads`  
- Putar langsung (macOS: afplay | Linux: mpv | Windows: PowerShell)  

---

## Prasyarat

- Go 1.20+ (untuk build dari source)  
- macOS: `afplay` (bawaan)  
- Linux: `mpv` (install manual)  
- Windows: PowerShell (bawaan)  

---

## Instalasi

### Build dari Source

```bash
git clone https://github.com/hllstr/ytplay.git
cd ytplay
go build -o ytplay
mv ytplay ~/bin/         # pastikan ~/bin ada di $PATH
```

### Menggunakan Pre-built Binary

1. Masuk ke halaman [Releases](https://github.com/hllstr/ytplay/releases).  
2. Download file `ytplay-<os>-<arch>`.  
3. Pindahkan `ytplay` ke direktori di `$PATH`.

---

## Cara Pakai

Jalankan perintah berikut dan ikuti prompt:

```bash
ytplay
```

1. **Masukkan kata kunci** lagu (misal: `ncs` atau `Coldplay Paradise`).  
2. **Pilih nomor** video dari 1–8.  
3. Program **mengunduh** audio `.m4a` terbaik ke `~/Downloads/[ID] <Title> - <Channel>.m4a`.  
4. Setelah selesai, **memutar** audio secara otomatis.  
5. Di akhir, program menanyakan “Cari lagi? (y/n)”.

---

## Contoh Sesi

```shell
$ ytplay
Masukkan kata kunci lagu: ncs
[1] NCS Release – Alan Walker – Fade
[2] NCS Release – Elektronomia – Limitless
[3] ...
Pilih nomor (1-8): 2
Mengunduh: Elektronomia – Limitless
Menyimpan ke /home/hllstr/Downloads/YTPLAY Title - Channel [abcdefg].m4a
Memutar: /home/hllstr/Downloads/YTPLAY Title - Channel.m4a
Cari lagi? (y/n): n
Makasih udah pake program ini <3 - hllstr.
```

---

## Struktur Proyek

```
ytplay/
├── main.go      # entrypoint interaktif
├── search.go    # wrapper ytsearch
├── download.go  # wrapper youtube/v2
├── player.go    # pemutar cross-platform
├── go.mod
└── .gitignore
```

---

## Kontribusi

1. Fork repository  
2. Buat branch fitur (`git checkout -b feat-foo`)  
3. Commit perubahan (`git commit -m "feat: add foo"`)  
4. Push ke branch-mu (`git push origin feat-foo`)  
5. Buka Pull Request  

---

## TODO

### Fitur & Fungsi
- [ ] Menambahkan kontrol audio (play/pause/stop/next/prev)
- [ ] Menambahkan fitur playlist untuk menyimpan dan mengelola daftar lagu  
- [ ] Menambahkan opsi penyimpanan sementara: file audio disimpan di `/temp/` agar tidak memenuhi folder Downloads  
- [x] Menambahkan sistem pengecekan file sebelum download: jika ID audio sudah ada di direktori, langsung play tanpa download ulang.  
- [x] Mendukung penyimpanan otomatis di `/sdcard/Download/YTPLAY` saat digunakan di Termux (Android)  

### Antarmuka & Pengalaman Pengguna
- [ ] Menambahkan menu utama agar lebih intuitif  
- [ ] Improve UI agar lebih mudah dinavigasi dan menarik  

### Kualitas Kode
- [ ] Refactor struktur program agar lebih modular dan mudah dipelihara

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
```
