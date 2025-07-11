# Airline Voucher Seat Assignment - Frontend

A React + Vite frontend application for airline voucher seat assignment system.

## Features

- **Form Validation**: Uses Zod for robust form validation
- **State Management**: Jotai for efficient state management
- **Modern UI**: Clean, responsive design with CSS-in-JS
- **Testing**: Comprehensive tests with Vitest and React Testing Library
- **TypeScript**: Full type safety throughout the application

## Tech Stack

- **React 18** - UI framework
- **Vite** - Build tool and dev server
- **TypeScript** - Type safety
- **Zod** - Schema validation
- **Jotai** - State management
- **Axios** - HTTP client
- **date-fns** - Date utilities
- **Vitest** - Testing framework
- **React Testing Library** - Component testing

## Project Structure

```
src/
├── api/              # API service functions
├── components/       # React components
├── store/           # Jotai atoms for state management
├── test/            # Test files
├── types/           # TypeScript type definitions
├── utils/           # Utility functions
├── App.tsx          # Main app component
└── main.tsx         # App entry point
```

## Getting Started

### Prerequisites

- Node.js (version 18 or higher)
- npm or yarn

### Installation

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

4. Open your browser and visit `http://localhost:3000`

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run test` - Run tests
- `npm run test:watch` - Run tests in watch mode
- `npm run lint` - Run ESLint

## Usage

1. **Fill in crew details**: Enter crew name and ID
2. **Enter flight information**: Provide flight number and date (DD-MM-YY format)
3. **Select aircraft type**: Choose from ATR, Airbus 320, or Boeing 737 Max
4. **Generate vouchers**: Click the button to generate 3 random seat assignments

The application will:
- Validate the form data using Zod schemas
- Check if vouchers already exist for the flight/date
- Generate 3 unique random seats if no duplicates exist
- Display the generated seat numbers or error messages

## Form Validation

The application includes comprehensive form validation:

- **Required fields**: All fields must be filled
- **Crew name**: Must be non-empty string
- **Crew ID**: Must be non-empty string
- **Flight number**: Must be non-empty string (e.g., "GA102")
- **Flight date**: Must be in DD-MM-YY format (e.g., "09-07-25")
- **Aircraft type**: Must be one of the predefined aircraft types

## State Management

The application uses Jotai for state management with the following atoms:

- `formDataAtom` - Form input data
- `isLoadingAtom` - Loading state
- `generatedSeatsAtom` - Generated seat numbers
- `errorMessageAtom` - Error messages
- `successMessageAtom` - Success messages

## API Integration

The frontend communicates with the backend through two endpoints:

- `POST /api/check` - Check if vouchers exist for a flight/date
- `POST /api/generate` - Generate new voucher assignments

## Testing

Run the test suite:

```bash
npm run test
```

The tests cover:
- Component rendering and user interactions
- Form validation logic
- Date utility functions
- Type schema validation
- API integration (mocked)

## Aircraft Seat Layouts

The application supports three aircraft types with different seat configurations:

- **ATR**: Rows 1-18, Seats A, C, D, F (72 total seats)
- **Airbus 320**: Rows 1-32, Seats A, B, C, D, E, F (192 total seats)
- **Boeing 737 Max**: Rows 1-32, Seats A, B, C, D, E, F (192 total seats)

## Design Features

- **Responsive Design**: Works on desktop and mobile devices
- **Modern UI**: Clean, professional interface with gradient backgrounds
- **Loading States**: Visual feedback during API calls
- **Error Handling**: Clear error messages and validation feedback
- **Accessibility**: Proper form labels and semantic HTML

## Configuration

The API base URL can be configured in `src/api/voucher.ts`. By default, it points to `http://localhost:8080/api`.

## Browser Support

The application supports all modern browsers that support ES2020+ features.
