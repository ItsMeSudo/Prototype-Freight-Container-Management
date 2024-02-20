# Prototype Freight Container (PFC)

Prototype Freight Container (PFC) is a full-stack application designed to manage freight container logistics. This documentation provides instructions for building, running, and modifying the application. 

## Backend

The backend of PFC is developed in GoLang, providing a REST API for managing container data. 

### Building from Source

1. Navigate to the `backend` folder.
2. Run `go build` to build the executable.

### Development

To run the backend in development mode:

```bash
go run main.go
```

### Optional Modifiers

After building the backend, you can use the following optional modifiers:

```bash
./backend -h
```

This will display the help message including available options.

By default, the backend also serves the built frontend.

### Endpoint documentation
[Available](https://github.com/ItsMeSudo/es2025-s09-r1-151/blob/main/backend/README.md) in the backend folder

## Frontend

The frontend of PFC is built using Vite with React TypeScript, Tailwind CSS, and Radix UI components.

### Building from Source

1. Navigate to the `frontend` folder.
2. Run `npm install` to install dependencies.
3. Run `npm run build` to build the frontend.

### Running

To run the frontend in development mode:

```bash
npm run dev
```

### Documentation for Technologies

- [Vite](https://vitejs.dev/)
- [React](https://reactjs.org/docs/getting-started.html)
- [TypeScript](https://www.typescriptlang.org/docs/)
- [Tailwind CSS](https://tailwindcss.com/docs)
- [Radix UI](https://www.radix-ui.com/)

## Prebuilt Versions

Prebuilt versions of the backend for Linux and Windows, as well as the frontend, are available in the `prebuilt` folder.

### Running Prebuilt Versions

1. Ensure the prebuilt folders for both backend and frontend are in the same directory.
   for example the tree needs to look like this:
    ```
    -folder
      --dist(folder)
      --backend-executable
    ```
3. Navigate to the prebuilt backend folder.
4. Run the executable:

   - For Linux: `./backend`
   - For Windows: `./backend.exe`

## Releases

Alternatively, if you prefer not to build from source, you can download fully built versions of PFC:

- [Windows Release](https://github.com/ItsMeSudo/es2025-s09-r1-151/releases/download/v1.0.0/windows.rar)
- [Linux Release](https://github.com/ItsMeSudo/es2025-s09-r1-151/releases/download/v1.0.0/linux.rar)

Simply extract the downloaded file and run the appropriate executable as mentioned above.

## Contributors

- [ItsMeSudo](https://github.com/ItsMeSudo)

