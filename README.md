# Restaurant Service

This is a service that provides information about restaurants and food options based on location. The service is built using the Golang Gin framework and uses CSV data for restaurant information.

## Setup

### Prerequisites

To run this service, make sure you have the following installed:

- Go (1.14+)
- Git
- Hoppscotch (for testing)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/restaurant-service.git
    cd restaurant-service
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

The service should now be running at http://localhost:8080.

## Testing with Hoppscotch

1. Install and open Hoppscotch (https://hoppscotch.io/).

2. Import the provided `hoppscotch_collection.json` file into Hoppscotch. This file contains pre-defined API requests for testing.

3. Configure the base URL to `http://localhost:8080`.

4. Use the imported collection to send requests to test various features of the service.

## Features

### 1. List Restaurants in a City

Endpoint: `GET /restaurants-in-city?city_code=<city_code>`

- Parameters: `city_code` (string)
- Returns a list of restaurants in the specified city.


![](/images/city.png)

### 2. List Food Options Near a Location

Endpoint: `GET /food-options-near?latitude=<latitude>&longitude=<longitude>`

- Parameters: `latitude` (float), `longitude` (float)
- Returns a list of food options near the specified location (within 5 kilometers).

![](/images/food-option.png)

### 3. List Restaurants Sorted by Rating

Endpoint: `GET /restaurants-sorted-by-rating?city_code=<city_code>`

- Parameters: `city_code` (string)
- Returns a list of restaurants in the specified city, sorted by rating.

![](/images/sorted.png)

### 4. Filter by Table Booking

Endpoint: `GET /filter-by-table-booking?city_code=<city_code>&has_table_booking=true`

- Parameters: `city_code` (string), `has_table_booking` (boolean)
- Returns a list of restaurants in the specified city that offer table booking.

![](/images/table-booking.png)
### 5. Filter by Online Delivery

Endpoint: `GET /filter-by-online-delivery?city_code=<city_code>&has_online_delivery=true`

- Parameters: `city_code` (string), `has_online_delivery` (boolean)
- Returns a list of restaurants in the specified city that offer online delivery.

![](/images/online.png)

### 6. Filter by Cuisines

Endpoint: `GET /filter-by-cuisines?city_code=<city_code>&cuisines=<cuisine_name>`

- Parameters: `city_code` (string), `cuisines` (string)
- Returns a list of restaurants in the specified city that serve the specified cuisine.

![](/images/cuisine.png)
### 7. CORS Support

The service includes CORS middleware to allow requests from different origins.

### 8. CSV Data Parsing

The service loads restaurant data from the provided CSV file and maps it to the defined model structure.

## Contributions

Feel free to contribute to this project by submitting pull requests or reporting issues.

Happy coding!