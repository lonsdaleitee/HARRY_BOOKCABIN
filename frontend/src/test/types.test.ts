import { describe, it, expect } from 'vitest'
import { voucherFormSchema, AircraftType } from '../types'

describe('Voucher Form Schema', () => {
  it('validates valid form data', () => {
    const validData = {
      crewName: 'John Doe',
      crewId: '12345',
      flightNumber: 'GA102',
      flightDate: '09-07-25',
      aircraft: AircraftType.AIRBUS_320,
    }

    const result = voucherFormSchema.safeParse(validData)
    expect(result.success).toBe(true)
  })

  it('rejects empty crew name', () => {
    const invalidData = {
      crewName: '',
      crewId: '12345',
      flightNumber: 'GA102',
      flightDate: '09-07-25',
      aircraft: AircraftType.AIRBUS_320,
    }

    const result = voucherFormSchema.safeParse(invalidData)
    expect(result.success).toBe(false)
    if (!result.success) {
      expect(result.error.errors[0].message).toBe('Crew name is required')
    }
  })

  it('rejects invalid aircraft type', () => {
    const invalidData = {
      crewName: 'John Doe',
      crewId: '12345',
      flightNumber: 'GA102',
      flightDate: '09-07-25',
      aircraft: 'Invalid Aircraft' as AircraftType,
    }

    const result = voucherFormSchema.safeParse(invalidData)
    expect(result.success).toBe(false)
  })

  it('trims whitespace from string fields', () => {
    const dataWithWhitespace = {
      crewName: '  John Doe  ',
      crewId: '  12345  ',
      flightNumber: '  GA102  ',
      flightDate: '09-07-25',
      aircraft: AircraftType.AIRBUS_320,
    }

    const result = voucherFormSchema.safeParse(dataWithWhitespace)
    expect(result.success).toBe(true)
    if (result.success) {
      expect(result.data.crewName).toBe('John Doe')
      expect(result.data.crewId).toBe('12345')
      expect(result.data.flightNumber).toBe('GA102')
    }
  })
})
