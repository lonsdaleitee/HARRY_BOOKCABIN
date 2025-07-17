import { z } from 'zod'

// Aircraft types
export const AircraftType = {
  ATR: 'ATR',
  AIRBUS_320: 'Airbus 320',
  BOEING_737_MAX: 'Boeing 737 Max',
} as const

export type AircraftType = typeof AircraftType[keyof typeof AircraftType]

// Date validation helper
const dateRegex = /^\d{2}-\d{2}-\d{2}$/
// Flight number validation helper (IATA format: 2+ letters followed by 1-4 numbers)
const flightNumberRegex = /^[A-Z]{2,}[0-9]{1,4}$/

// Form validation schema
export const voucherFormSchema = z.object({
  crewName: z.string().min(1, 'Crew name is required').trim(),
  crewId: z.string().min(1, 'Crew ID is required').trim(),
  flightNumber: z.string()
    .min(1, 'Flight number is required')
    .regex(flightNumberRegex, 'Flight number must have at least 2 letters followed by 1-4 numbers (e.g., GA102)')
    .trim()
    .transform((val) => val.toUpperCase()),
  flightDate: z.string()
    .min(1, 'Flight date is required')
    .refine((date) => {
      // Check if input contains any alphabetical characters
      if (/[a-zA-Z]/.test(date)) {
        return false
      }
      return true
    }, 'Only numerical characters are allowed for the date')
    .refine((date) => {
      // Check if it matches the DD-MM-YY format
      if (!dateRegex.test(date)) {
        return false
      }
      return true
    }, 'Date must be in DD-MM-YY format')
    .refine((date) => {
      if (!date.match(dateRegex)) return false
      const [day, month] = date.split('-').map(Number)
      if (day < 1 || day > 31) return false
      if (month < 1 || month > 12) return false
      return true
    }, 'Please enter a valid date'),
  aircraft: z.enum([AircraftType.ATR, AircraftType.AIRBUS_320, AircraftType.BOEING_737_MAX], {
    errorMap: () => ({ message: 'Please select an aircraft type' })
  }),
})

// Individual field schemas for per-field validation
export const fieldSchemas = {
  crewName: voucherFormSchema.shape.crewName,
  crewId: voucherFormSchema.shape.crewId,
  flightNumber: voucherFormSchema.shape.flightNumber,
  flightDate: voucherFormSchema.shape.flightDate,
  aircraft: voucherFormSchema.shape.aircraft,
}

export type VoucherFormData = z.infer<typeof voucherFormSchema>

// API response types
export interface CheckVoucherResponse {
  exists: boolean
}

export interface GenerateVoucherResponse {
  success: boolean
  seats: string[]
}

// Get voucher types
export interface GetVoucherRequest {
  flightNumber: string
  date: string
}

export interface Voucher {
  id: number
  crew_name: string
  crew_id: string
  flight_number: string
  flight_date: string
  aircraft_type: string
  seat1: string
  seat2: string
  seat3: string
  created_at: string
}

export interface GetVoucherResponse {
  voucher: Voucher | null
  exists: boolean
}

// Regenerate seat types
export interface RegenerateSeatRequest {
  flightNumber: string
  date: string
  seatPosition: number // 1, 2, or 3
}

export interface RegenerateSeatResponse {
  success: boolean
  newSeat: string
  allSeats: string[]
}

// API request types
export interface CheckVoucherRequest {
  flightNumber: string
  date: string
}

export interface GenerateVoucherRequest {
  name: string
  id: string
  flightNumber: string
  date: string
  aircraft: AircraftType
}
