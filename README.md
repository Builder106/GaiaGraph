# GaiaGraph

**GaiaGraph** is a climate data API that aggregates and provides access to global environmental metrics from various sources, including OpenWeather and NASA Power. This API allows users to retrieve combined climate data for a specified country.

---

## Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Endpoints](#endpoints)
4. [License](#license)

---

## Installation

1. **Clone the repository:**
   ```
   git clone https://github.com/Builder106/GaiaGraph
   cd GaiaGraph
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up the database:**
   ```
   # Ensure PostgreSQL is installed and running
   createdb climate_data
   ```

4. **Create a .env file in the root directory with the following content:**
   ```
   OPENWEATHER_API_KEY=your_openweather_api_key
   ```

5. **Run the server:**
   ```
   go run main.go
   ```

## Usage

Once the server is running, you can access the API at `http://localhost:8080`.

### Example Request

To get combined climate data for a specific country, make a GET request to the /climate-data endpoint with the country query parameter:

```
curl "http://localhost:8080/climate-data?country=Germany"
```

## Endpoints

`GET /climate-data`

Retrieved combined climate data for a specific country

### Query Parameters
* `country` (required): The name of the country for which to retrieve climate data.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
