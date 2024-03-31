# Golang Link Shortener

This is a simple link shortener implemented in Go.

## Features

- **URL Shortening**: Allows shortening long URLs to custom keys.
- **Configurability**: Enables customization of port and default target URL via a configuration file.
- **Monitoring Configuration Changes**: Monitors changes to the configuration file and updates the service's behavior accordingly.

## File Structure

- **main.go**: The main file that initializes the service, registers routes, and starts the HTTP server.
- **config.go**: Contains functions for loading and processing the configuration file.
- **handler.go**: Defines handler functions for HTTP requests.
- **short.go**: Represents a structure for short links and implements functions for handling short links.
- **shorts.go**: Contains the logic for loading, saving, and updating short links.

## Building the Package:

1. Navigate to the Project Directory: Open a terminal and change to the directory where your Go project is located.
2. Build the Package: Run the following command to build the package:
```bash
# This command compiles the Go code in your project and generates an executable binary file.
go build
```

## Usage

Once you've built the package, you can use the resulting binary as follows:
1. Run the Binary:<br>
Execute the generated binary file.
```bash
./go-link-shortener
```
2. **Customize Configuration**: Edit the `go-link-shortener.json` file to adjust the port and default target if no short is found. Updating this file requires a restart of the service.
3. **Updating Shorts**: Edit the `go-shorts.json` to update the shorts list on the fly. The application will detect it automatically and update it's running config.

## Routes

### /to/* Route:

- This route is responsible for redirecting users to the original URL associated with the provided short link slug.
- When a user navigates to this route with a specific short link slug appended to it, the application looks up the corresponding original URL and redirects the user to that URL.
- For example, accessing /to/my-short-link would redirect the user to the original URL associated with the short link slug my-short-link.

### /list Route:

- This route is used to display a list of all available short links along with their corresponding original URLs.
- When a user accesses this route, the application retrieves the list of short links from its storage and presents them in a human-readable format.
- Typically, this route is used for administrative purposes or for users to explore available short links.


## Dependencies

- **fsnotify**: This library is used to monitor changes to the configuration file.

## License

This project is licensed under the [Apache License](https://github.com/tim-krehan/go-link-shortener/tree/main?tab=Apache-2.0-1-ov-file).