# SRT to LRC Converter

This is a simple command-line tool written in Go that converts SubRip subtitle (.srt) files into Lyric (.lrc) files.

## Usage

To use this tool, you need to have Go installed on your machine. Once you have Go installed, you can run the program with the following command:

```bash
go run srt2lrc.go input.srt output.lrc
```

Replace `input.srt` with the path to your .srt file, and `output.lrc` with the path where you want the .lrc file to be saved.

## Functionality

The program works by parsing the .srt file and converting each subtitle into the .lrc format. The .lrc format is a simple text format that is often used for lyrics of songs, but it can also be used for subtitles.

The program assumes that the .srt file is structured like this:

```
1
00:00:20,000 --> 00:00:24,400
Example subtitle
```

Each subtitle is converted into two lines in the .lrc file: one line for the start time and text, and one line for the end time with a blank line.

## Limitations

This tool is a simple example and might not work with all .srt files, as the .srt format can have many variations. It also does not handle errors, so it might fail if the .srt file is not formatted correctly.

## License

This project is licensed under the terms of the MIT license.