import { describe, it, expect } from 'vitest'
import { formatDateForAPI, getCurrentDate } from '../utils/date'

describe('Date Utils', () => {
  describe('formatDateForAPI', () => {
    it('converts DD-MM-YY to YYYY-MM-DD format', () => {
      expect(formatDateForAPI('09-07-25')).toBe('2025-07-09')
      expect(formatDateForAPI('01-12-25')).toBe('2025-12-01')
      expect(formatDateForAPI('31-01-24')).toBe('2024-01-31')
    })

    it('handles single digit dates and months', () => {
      expect(formatDateForAPI('5-7-25')).toBe('2025-07-05')
      expect(formatDateForAPI('1-1-25')).toBe('2025-01-01')
    })
  })

  describe('getCurrentDate', () => {
    it('returns date in DD-MM-YY format', () => {
      const result = getCurrentDate()
      // Check format: should be DD-MM-YY
      expect(result).toMatch(/^\d{2}-\d{2}-\d{2}$/)
    })
  })
})
