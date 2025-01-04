<h1 align="center">
  Dirmaker
</h1>

<p align="center">
  <a href="#-about">About</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#%EF%B8%8F-how-to-use">How to Use</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-contact">Contact</a>
</p>
<br>
<p>
<img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=49AA26&labelColor=000000">
</p>

# 📖 About

**Dirmaker** is an application developed in Go (Golang) that simplifies directory creation. It is designed to be a lightweight and efficient utility, perfect for organizing directories in projects or tasks.

# ⚙️ How to Use

1. Ensure you have [Go](https://golang.org/) installed on your machine.
2. Clone this repository:
   ```bash
   git clone https://github.com/seu-usuario/dirmaker.git
   ```
3. Compile the project:
   ```bash
   go build -o dirmaker
   ```
4. Add the generated binary to your PATH for easy access.

## How it works

```bash
  dirmaker create [root directory] [flags]
```

## Good To Know

#### Shorthands

```bash
--subdirs = -s | --files = -f
```

---

<br>

# Usage Examples

### Create a Directory with Subdirectories and Files:

```bash
  dirmaker create mydir --subdirs src,bin --files README.md,main.go
```

Dirmaker will create:

- A folder named mydir.
- Inside mydir, subfolders src and bin.
- Inside mydir, files README.md and main.go.

Generated structure:

```bash
mydir/
├── bin/
├── src/
├── README.md
└── main.go
```

### Create Files in the Current Directory:

```bash
dirmaker create ./ --files README.md,main.go
```

Dirmaker will create:

- Two files in the root directory.

Generated structure:

```bash
  .
  ├── README.md
  └── main.go
```

<br>

# 📝 License

This project is licensed under the MIT License. See the file LICENSE for more details.

# 🌐 Contact

<a href="https://www.linkedin.com/in/seu-perfil-linkedin" target="_blank"><img src="https://img.shields.io/badge/-LinkedIn-%230077B5?style=for-the-badge&logo=linkedin&logoColor=white" target="_blank"></a>
<a href="mailto:seu-email@gmail.com"><img src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white" target="_blank"></a>

<br>

---

Made with ♥ by Gabriel Gonçalves 🖖
