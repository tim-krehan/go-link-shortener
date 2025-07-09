# Golang Link Shortener

A lightweight link shortener written in Go. As if there aren't enough already. 😉  
Full configurability, dynamic updates, and simple deployment.

## Features

- **URL Shortening**: Create custom short links to redirect anywhere.
- **Live Configuration Updates**: Changes in the short links config (`go-shorts.json`) take effect instantly—no need to restart.
- **Easy Deployment with Helm**: Get up and running quickly with Kubernetes using the provided Helm chart.

## File Structure

- **main.go** – Starts the HTTP server and initializes routes.
- **config.go** – Manages loading of service configuration.
- **handler.go** – Defines HTTP request handlers.
- **short.go** – Defines short link structure and logic.
- **shorts.go** – Handles loading, saving, and updating short links dynamically.

## Building the Package

1. Navigate to the project directory.
2. Run:
   ```bash
   go build
   ```

This will produce the binary `go-link-shortener`.

## Usage

1. Run the binary:
   ```bash
   ./go-link-shortener
   ```
2. Customize your configuration in `go-link-shortener.json`.
   - Change port
   - Define fallback link if no match is found
3. Update short links in `go-shorts.json` on the fly.
   - The app auto-detects file changes and reloads.

## Routes

### `/to/{slug}`

- Redirects to the original URL based on the short link slug.
- Example: `/to/git` ➜ `https://github.com/tim-krehan/go-link-shortener`

### `/list`

- Displays all configured short links with targets.

## Helm Installation 🚀

Deploy with Helm using the following commands:

```bash
helm repo add go-link-shortener https://tim-krehan.github.io/go-link-shortener
helm repo update
helm install go-link-shortener go-link-shortener/go-link-shortener --create-namespace --namespace link-shortener
```

📁 Example `values.yaml` configuration:  
[`charts/go-link-shortener/values.yaml`](https://github.com/tim-krehan/go-link-shortener/blob/master/charts/go-link-shortener/values.yaml)

### ⚠️ Required Configuration Values

You can set the following values to update your config:

```yaml
config:
  - slug: git
    target: https://github.com/tim-krehan/go-link-shortener
    description: Link shortener written in golang. As if there aren't enough allready.
```

You can add or remove short links as needed. Updates will be reflected immediately.

## Dependencies

- [fsnotify](https://github.com/fsnotify/fsnotify) – File system notifications for live config reload.

## License

This project is licensed under the [Apache License](https://github.com/tim-krehan/go-link-shortener/blob/master/LICENSE).
