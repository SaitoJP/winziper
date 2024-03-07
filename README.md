# Winziper

Winziper is a command-line tool written in Go for zip compression, specifically designed for cross-platform file sharing between Mac and Windows systems. This tool not only compresses files into a zip archive but also converts file names into the SJIS (Shift_JIS) character encoding for compatibility with Windows systems. Additionally, it automatically removes macOS-specific files like ".DS_Store" to clean up the archive for Windows users.

## Features

- **Cross-Platform Compatibility:** Automatically converts file names into SJIS encoding to ensure compatibility with Windows systems.
- **Clean Archives:** Removes macOS-specific files (e.g., .DS_Store) from the zip archive, making it cleaner for Windows users.
- **Easy to Use:** Simple and straightforward command-line interface.

## Requirements

To use Winziper, you need to have Go installed on your Mac. You can download and install Go from [the official site](https://golang.org/dl/).

## Installation

Clone the repository to your local machine:
```bash
git clone https://github.com/SaitoJP/winziper.git
```

Navigate to the winziper directory:
```bash
cd winziper
```

## Usage

To compress a file or directory with Winziper, run the following command in your terminal:
```bash
go run main.go [path to file or directory]
```
Replace [path to file or directory] with the actual path to the file or directory you want to compress.

## Contributing

Contributions are welcome! If you have a fix or feature you would like to add, please feel free to fork the repository, make your changes, and submit a pull request.

## License

Winziper is released under the MIT License.
