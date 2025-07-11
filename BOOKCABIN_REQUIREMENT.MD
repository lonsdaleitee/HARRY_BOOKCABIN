**Exam Question: Airline Voucher Seat Assignment App (React + .net/
GoLang)**

**Objective**

You are building a **voucher seat assignment web application** for an
airline campaign. For each flight, the airline wants to randomly assign
**3 unique seat numbers**, with specific seat maps depending on aircraft
type.

The crew must use the app to:

> 1.Input flight and crew details.
>
> 2.Generate **exactly 3 random seats** based on the aircraft layout.
>
> 3.Prevent duplicate assignments for the **same flight number on the
> same date**. 4.Persist the assignments to a database.

**Requirements**

**Frontend (React)**

Build a user interface that allows the crew to:

+-----------------------------------+-----------------------------------+
| •                                 | > Enter:\                         |
|                                   | > oCrew Name\                     |
| •                                 | > oCrew ID\                       |
|                                   | > oFlight number\                 |
|                                   | > oFlight date (DD -- MM - YY)\   |
|                                   | > oAircraft type (dropdown):      |
|                                   | >                                 |
|                                   | > ▪ ATR                           |
|                                   | >                                 |
|                                   | > ▪ Airbus 320                    |
|                                   | >                                 |
|                                   | > ▪ Boeing 737 Max\               |
|                                   | > Press a "Generate Vouchers"     |
|                                   | > button to:\                     |
|                                   | > oValidate if the flight already |
|                                   | > has assigned vouchers for that  |
|                                   | > date (API: POST /api/check)\    |
|                                   | > oIf not already generated:      |
|                                   | >                                 |
|                                   | > ▪ Call API (POST /api/generate) |
|                                   | >                                 |
|                                   | > ▪ Display the 3 randomly chosen |
|                                   | > seats                           |
+===================================+===================================+
+-----------------------------------+-----------------------------------+

Display an error if vouchers have already been generated for that
flight/date.

**Backend (.Net / GoLang )**

Implement a REST API with the following endpoints:

**POST /api/check**

> • Input:
>
> json\
> CopyEdit\
> { \"flightNumber\": \"GA102\", \"date\": \"2025-07-12\" }
>
> • Output:
>
> json\
> CopyEdit\
> { \"exists\": true }

**POST /api/generate**

> • Input:
>
> json\
> CopyEdit\
> {\
> \"name\": \"Sarah\",\
> \"id\": \"98123\",\
> \"flightNumber\": \"ID102\",\
> \"date\": \"2025-07-12\",\
> \"aircraft\": \"Airbus 320\"\
> }
>
> • Output:
>
> json\
> CopyEdit\
> {\
> \"success\": true,\
> \"seats\": \[\"3B\", \"7C\", \"14D\"\]\
> }

Save the assignment to SQLite (vouchers.db) with schema:

> **Field** **Type**\
> id INT (PK) AUTOINCREMENT crew_name TEXT\
> crew_id TEXT\
> flight_number TEXT\
> flight_date TEXT\
> aircraft_type TEXT\
> seat1, seat2, seat3 TEXT\
> created_at TEXT (timestamp)

**Seat Layouts Information**

> **Aircraft Type Rows Seats per Row**
>
> ATR 1-18 A, C, D, F
>
> Airbus 320 1-32 A, B, C, D, E, F
>
> Boeing 737 Max 1- 32 A, B, C, D, E, F

**Bonus**

+-----------------------------------+-----------------------------------+
| > •\                              | > Use parameterized SQL to        |
| > •\                              | > prevent injection Format error  |
| > •                               | > messages cleanly\               |
|                                   | > Modularize backend handlers for |
|                                   | > maintainability                 |
+===================================+===================================+
+-----------------------------------+-----------------------------------+

**Deliverables**

+-----------------------------------+-----------------------------------+
| > •\                              | > frontend/ --- React project\    |
| > •\                              | > backend/ --- GoLang server with |
| > •                               | > SQLite Instructions to run      |
|                                   | > (with or without Docker)        |
+===================================+===================================+
+-----------------------------------+-----------------------------------+
