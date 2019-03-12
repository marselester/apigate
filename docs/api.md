FORMAT: 1A

# Travel

Travel is a demo API allowing clients to book a trip (rent a car, reserve a hotel).
Make sure to provide `X-Request-ID` header for requests deduplication.

## Group Cars

## Cars For Rent [/cars]
### List All Cars [GET]
<a name="cars"></a>

+ Response 200 (application/json)

        [
            {
                "id": "cfb6f7a5-4591-4f5c-8b17-9a1b10f98ada",
                "name": "Toyota Yaris",
                "price": "30"
            },
            {
                "id": "afad6e6c-ef7f-4dc9-bc0f-ce74d5392175",
                "name": "Honda Civic",
                "price": "40"
            }
        ]

## Cars Bookings [/cars/bookings]
### Book a Car [POST]

You can book a car listed at [cars](#cars) endpoint.

+ car_id (string) - The car you would like to rent.

+ Request (application/json)

        {
            "car_id": "cfb6f7a5-4591-4f5c-8b17-9a1b10f98ada"
        }

+ Response 201 (application/json)

        {
            "id": "9e0d65f5-9de2-4428-9bee-1f3967f05129",
            "car_id": "cfb6f7a5-4591-4f5c-8b17-9a1b10f98ada",
            "status": "confirmed"
        }

## Car Booking [/cars/bookings/{booking_id}]

A booking has the following attributes:

+ id - ID of the booking.
+ car_id - The car ID you would like to rent.
+ created_at - An ISO8601 date when the car was book.
+ status - Current status of the booking.

### Car Booking Details [GET]

+ Parameters
    + booking_id: `"9e0d65f5-9de2-4428-9bee-1f3967f05129"` (string) - ID of the booking.

+ Response 200 (application/json)

        {
            "id": "9e0d65f5-9de2-4428-9bee-1f3967f05129",
            "car_id": "cfb6f7a5-4591-4f5c-8b17-9a1b10f98ada",
            "status": "confirmed"
        }

## Group Hotels

## Hotels to Stay [/hotels]
### List All Hotels [GET]
<a name="hotels"></a>

+ Response 200 (application/json)

        [
            {
                "id": "046d471d-70c7-4595-80cc-266d3e6e07fa",
                "name": "Holiday Inn",
                "price": "35"
            },
            {
                "id": "f183e115-7efb-49c1-b338-ed5265bc8431",
                "name": "Four Seasons",
                "price": "45"
            }
        ]

## Hotel Bookings [/hotels/bookings]
### Book a Hotel [POST]

You can book a hotel room listed at [hotels](#hotels) endpoint.

+ hotel_id (string) - The hotel you would like to book a room at.

+ Request (application/json)

        {
            "hotel_id": "046d471d-70c7-4595-80cc-266d3e6e07fa"
        }

+ Response 201 (application/json)

        {
            "id": "7b4fc183-ee67-494d-9715-3510c6d8f2ef",
            "hotel_id": "046d471d-70c7-4595-80cc-266d3e6e07fa",
            "status": "confirmed"
        }

## Hotel Booking [/hotels/bookings/{booking_id}]

A booking has the following attributes:

+ id - ID of the booking.
+ hotel_id - The hotel ID you would like to stay.
+ created_at - An ISO8601 date when the room was book.
+ status - Current status of the booking.

### Hotel Booking Details [GET]

+ Parameters
    + booking_id: `"7b4fc183-ee67-494d-9715-3510c6d8f2ef"` (string) - ID of the booking.

+ Response 200 (application/json)

        {
            "id": "7b4fc183-ee67-494d-9715-3510c6d8f2ef",
            "hotel_id": "046d471d-70c7-4595-80cc-266d3e6e07fa",
            "status": "confirmed"
        }
