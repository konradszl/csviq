# csviq

A terminal UI for reading, editing and querying CSV files.

`csviq` opens a CSV in an interactive table you can navigate, edit and filter, import data from other formats such as JSON, and query it with familiar SQL syntax without leaving your terminal.

**Status: early development.**  

## Install
> [!NOTE]
> Requires go **1.26** or later.

```sh
go install github.com/konradszl/csviq@latest
```

Or build from source:
```sh
git clone https://github.com/konradszl/csviq.git
cd csviq
go build .
```

## Usage
```sh
csviq data.csv
```

## Roadmap
- [ ] View a CSV in a scrollable table
- [ ] Cell editing and saving
- [ ] Import from JSON
- [ ] Query with SQL syntax
- [ ] Export to CSV / JSON
- [ ] Large-file streaming

## Development
```sh
go test ./...
go vet ./...
gofmt -l .
```

## Contributing
Issues and pull requests are welcome. This is a learning project as much as a tool, so questions about the design are welcome too.

## License
[MIT](LICENSE)
