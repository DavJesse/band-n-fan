# Groupie Trackers

Groupie Trackers is a web application that displays information about bands and artists using data from a given API. The project focuses on data manipulation and visualization, as well as implementing client-server interactions.

## Features

- Display information about bands and artists, including:
  - Names
  - Images
  - Year of activity
  - First album date
  - Members
- Show last and upcoming concert locations
- Display last and upcoming concert dates
- Visualize the relationships between artists, dates, and locations
- Implement client-server interactions for dynamic data retrieval

## Technologies Used

- Backend: Go
- Frontend: HTML, CSS, JavaScript
- Templates: Go's `html/template` package

## Project Structure
```
groupie-trackers/
├── internal/
│   ├── api/
│   │   ├── fetch_data.go
│   │   └── fetch_data_test.go
│   └── handlers/
│       ├── relation.go
│       └── ... (other handler files)
├── test/
│   └── ... (test files)
├── web/
│   ├── static/
│   │   ├── css/
│   │   ├── images/
│   │   └── js/
│   └── templates/
│       ├── relation.html
│       └── ... (other template files)
├── go.mod
├── main.go
└── README.md
```

## Getting Started

1. Clone the repository:
```
From Gitea:
git clone https://learn.zone01kisumu.ke/git/davodhiambo/groupie-tracker.git

From Github:
git clone https://github.com/DavJesse/band-n-fan.git
```

2. Navigate to the project directory:
```
From Gitea:
cd goupie-tracker

From Github:
cd band-n-fan
```

3. Run the application:
```
go run main.go
```

4. Open your web browser and visit `http://localhost:8080` (or the port specified in your main.go file).

## API Endpoints

The application uses the following API endpoints:

- `https://groupietrackers.herokuapp.com/api/artists`: Information about bands and artists
- `https://groupietrackers.herokuapp.com/api/locations`: Concert locations
- `https://groupietrackers.herokuapp.com/api/dates`: Concert dates
- `https://groupietrackers.herokuapp.com/api/relation`: Links between artists, dates, and locations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).

## Authors
[David Odhiambo](https://learn.zone01kisumu.ke/git/davodhiambo)

Apprentice Software Developer | Zone01 Kisumu

[Swabri Kanenje](https://learn.zone01kisumu.ke/git/skanenje)

Apprentice Software Developer | Zone01 Kisumu