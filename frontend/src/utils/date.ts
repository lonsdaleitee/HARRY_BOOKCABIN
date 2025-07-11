import { format } from 'date-fns'

export const formatDateForAPI = (dateString: string): string => {
  // Convert from DD-MM-YY to YYYY-MM-DD format
  const [day, month, year] = dateString.split('-')
  const fullYear = `20${year}` // Assuming 2000s
  return `${fullYear}-${month.padStart(2, '0')}-${day.padStart(2, '0')}`
}

export const getCurrentDate = (): string => {
  // Return current date in DD-MM-YY format
  return format(new Date(), 'dd-MM-yy')
}

export const formatDateInput = (value: string): string => {
  // Remove all non-numeric characters
  const numericOnly = value.replace(/\D/g, '')
  
  // Limit to 6 digits (DDMMYY)
  const limited = numericOnly.slice(0, 6)
  
  // Add dashes automatically
  if (limited.length >= 5) {
    return `${limited.slice(0, 2)}-${limited.slice(2, 4)}-${limited.slice(4, 6)}`
  } else if (limited.length >= 3) {
    return `${limited.slice(0, 2)}-${limited.slice(2, 4)}`
  } else if (limited.length >= 1) {
    return limited
  }
  
  return limited
}

export const formatFlightNumberInput = (value: string): string => {
  // Remove spaces and convert to uppercase
  const cleaned = value.replace(/\s/g, '').toUpperCase()
  
  // Only allow letters and numbers
  const alphaNumericOnly = cleaned.replace(/[^A-Z0-9]/g, '')
  
  // Separate letters and numbers
  const letters = alphaNumericOnly.replace(/[0-9]/g, '')
  const numbers = alphaNumericOnly.replace(/[A-Z]/g, '')
  
  // Limit letters to reasonable length (max 3 for airline codes)
  const limitedLetters = letters.slice(0, 3)
  // Limit numbers to 4 digits
  const limitedNumbers = numbers.slice(0, 4)
  
  // Combine them
  return limitedLetters + limitedNumbers
}
