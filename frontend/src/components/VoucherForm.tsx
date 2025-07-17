import React, { useState, useEffect } from 'react'
import { useAtom } from 'jotai'
import { voucherFormSchema, fieldSchemas, type VoucherFormData, AircraftType } from '../types'
import { 
  formDataAtom, 
  isLoadingAtom, 
  generatedSeatsAtom, 
  errorMessageAtom, 
  successMessageAtom,
  currentVoucherAtom,
  isRegeneratingAtom
} from '../store/atoms'
import { checkVoucher, generateVoucher, getVoucher, regenerateSeat } from '../api/voucher'
import { formatDateForAPI, formatDateInput, formatFlightNumberInput } from '../utils/date'
import { ZodError } from 'zod'
import './VoucherForm.css'

export const VoucherForm: React.FC = () => {
  const [formData, setFormData] = useAtom(formDataAtom)
  const [isLoading, setIsLoading] = useAtom(isLoadingAtom)
  const [generatedSeats, setGeneratedSeats] = useAtom(generatedSeatsAtom)
  const [errorMessage, setErrorMessage] = useAtom(errorMessageAtom)
  const [successMessage, setSuccessMessage] = useAtom(successMessageAtom)
  const [currentVoucher, setCurrentVoucher] = useAtom(currentVoucherAtom)
  const [isRegenerating, setIsRegenerating] = useAtom(isRegeneratingAtom)
  const [validationErrors, setValidationErrors] = useState<Record<string, string>>({})

  // Check for existing voucher when flight details change
  useEffect(() => {
    const checkExistingVoucher = async () => {
      if (formData.flightNumber && formData.flightDate) {
        try {
          const formattedDate = formatDateForAPI(formData.flightDate)
          const result = await getVoucher({
            flightNumber: formData.flightNumber,
            date: formattedDate,
          })
          
          if (result.exists && result.voucher) {
            setCurrentVoucher({
              exists: true,
              seats: [result.voucher.seat1, result.voucher.seat2, result.voucher.seat3],
              flightNumber: formData.flightNumber,
              date: formattedDate,
            })
            setGeneratedSeats([result.voucher.seat1, result.voucher.seat2, result.voucher.seat3])
          } else {
            setCurrentVoucher(null)
            setGeneratedSeats([])
          }
        } catch (error) {
          // Silently handle errors - voucher might not exist yet
          setCurrentVoucher(null)
          setGeneratedSeats([])
        }
      } else {
        setCurrentVoucher(null)
        setGeneratedSeats([])
      }
    }

    // Debounce the check to avoid too many API calls
    const timeoutId = setTimeout(checkExistingVoucher, 500)
    return () => clearTimeout(timeoutId)
  }, [formData.flightNumber, formData.flightDate, setCurrentVoucher, setGeneratedSeats])

  const handleInputChange = (field: keyof VoucherFormData, value: string) => {
    let processedValue = value
    
    // Special handling for flight date field
    if (field === 'flightDate') {
      processedValue = formatDateInput(value)
    }
    // Special handling for flight number field
    else if (field === 'flightNumber') {
      processedValue = formatFlightNumberInput(value)
    }
    
    const updatedData = { ...formData, [field]: processedValue }
    setFormData(updatedData)
    
    // Validate only the specific field that changed
    try {
      fieldSchemas[field].parse(processedValue)
      // If validation passes, clear error for this field
      setValidationErrors(prev => {
        const newErrors = { ...prev }
        delete newErrors[field]
        return newErrors
      })
    } catch (error) {
      if (error instanceof ZodError) {
        // Set error for this specific field
        setValidationErrors(prev => ({
          ...prev,
          [field]: error.errors[0]?.message || 'Invalid value'
        }))
      }
    }
  }

  const clearMessages = () => {
    setErrorMessage(null)
    setSuccessMessage(null)
  }

  const handleRegenerateSeat = async (seatPosition: number) => {
    if (!currentVoucher) return

    clearMessages()
    setIsRegenerating(true)

    try {
      const result = await regenerateSeat({
        flightNumber: currentVoucher.flightNumber,
        date: currentVoucher.date,
        seatPosition: seatPosition,
      })

      if (result.success) {
        // Update the seats display
        setGeneratedSeats(result.allSeats)
        setCurrentVoucher({
          ...currentVoucher,
          seats: result.allSeats,
        })
        setSuccessMessage(`Seat ${seatPosition} regenerated successfully! New seat: ${result.newSeat}`)
      } else {
        setErrorMessage('Failed to regenerate seat. Please try again.')
      }
    } catch (error) {
      console.error('Error regenerating seat:', error)
      setErrorMessage('An error occurred while regenerating the seat. Please try again.')
    } finally {
      setIsRegenerating(false)
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    clearMessages()
    
    // If vouchers already exist, show message instead of generating new ones
    if (currentVoucher?.exists) {
      setErrorMessage('Vouchers have already been generated for this flight on the selected date. Use the regenerate buttons to change individual seats.')
      return
    }
    
    // Validate the entire form before submission
    try {
      voucherFormSchema.parse(formData)
      setValidationErrors({}) // Clear any remaining errors
    } catch (error) {
      if (error instanceof ZodError) {
        const errors: Record<string, string> = {}
        error.errors.forEach(err => {
          if (err.path[0]) {
            errors[err.path[0] as string] = err.message
          }
        })
        setValidationErrors(errors)
        return
      }
    }

    setIsLoading(true)

    try {
      // First check if vouchers already exist
      const formattedDate = formatDateForAPI(formData.flightDate)
      const checkResult = await checkVoucher({
        flightNumber: formData.flightNumber,
        date: formattedDate,
      })

      if (checkResult.exists) {
        setErrorMessage('Vouchers have already been generated for this flight on the selected date.')
        return
      }

      // Generate new vouchers
      const generateResult = await generateVoucher({
        name: formData.crewName,
        id: formData.crewId,
        flightNumber: formData.flightNumber,
        date: formattedDate,
        aircraft: formData.aircraft,
      })

      if (generateResult.success) {
        setGeneratedSeats(generateResult.seats)
        setCurrentVoucher({
          exists: true,
          seats: generateResult.seats,
          flightNumber: formData.flightNumber,
          date: formattedDate,
        })
        setSuccessMessage('Vouchers generated successfully!')
      } else {
        setErrorMessage('Failed to generate vouchers. Please try again.')
      }
    } catch (error) {
      console.error('Error:', error)
      setErrorMessage('An error occurred while processing your request. Please try again.')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="voucher-form-container">
      <h1>Airline Voucher Seat Assignment</h1>
      
      <form onSubmit={handleSubmit} className="voucher-form">
        <div className="form-group">
          <label htmlFor="crewName">
            Crew Name<span className="required-asterisk">*</span>
          </label>
          <input
            type="text"
            id="crewName"
            value={formData.crewName}
            onChange={(e) => handleInputChange('crewName', e.target.value)}
            className={validationErrors.crewName ? 'error' : ''}
            disabled={isLoading}
            placeholder='e.g John Doe'
          />
          {validationErrors.crewName && (
            <span className="error-text">{validationErrors.crewName}</span>
          )}
        </div>

        <div className="form-group">
          <label htmlFor="crewId">
            Crew ID<span className="required-asterisk">*</span>
          </label>
          <input
            type="text"
            id="crewId"
            value={formData.crewId}
            onChange={(e) => handleInputChange('crewId', e.target.value)}
            className={validationErrors.crewId ? 'error' : ''}
            disabled={isLoading}
            placeholder='e.g 2025071101'
          />
          {validationErrors.crewId && (
            <span className="error-text">{validationErrors.crewId}</span>
          )}
        </div>

        {/* No specific regulation for flight number, but usually it follows IATA format https://www.iata.org/en/publications/directories/code-search
        *
        * For example, GA102
        * Two alphabetic characters followed by 1-4 numeric characters
        * */}
        <div className="form-group">
          <label htmlFor="flightNumber">
            Flight Number<span className="required-asterisk">*</span>
          </label>
          <input
            type="text"
            id="flightNumber"
            value={formData.flightNumber}
            onChange={(e) => handleInputChange('flightNumber', e.target.value)}
            className={validationErrors.flightNumber ? 'error' : ''}
            disabled={isLoading}
            placeholder="e.g., GA102"
            maxLength={7}
            style={{ textTransform: 'uppercase' }}
          />
          {validationErrors.flightNumber && (
            <span className="error-text">{validationErrors.flightNumber}</span>
          )}
        </div>

        <div className="form-group">
          <label htmlFor="flightDate">
            Flight Date (DD-MM-YY)<span className="required-asterisk">*</span>
          </label>
          <input
            type="text"
            id="flightDate"
            value={formData.flightDate}
            onChange={(e) => handleInputChange('flightDate', e.target.value)}
            className={validationErrors.flightDate ? 'error' : ''}
            disabled={isLoading}
            placeholder="09-07-25"
            maxLength={8}
            inputMode="numeric"
            pattern="[0-9\-]*"
          />
          {validationErrors.flightDate && (
            <span className="error-text">{validationErrors.flightDate}</span>
          )}
        </div>

        <div className="form-group">
          <label htmlFor="aircraft">
            Aircraft Type<span className="required-asterisk">*</span>
          </label>
          <select
            id="aircraft"
            value={formData.aircraft}
            onChange={(e) => handleInputChange('aircraft', e.target.value)}
            className={validationErrors.aircraft ? 'error' : ''}
            disabled={isLoading}
          >
            <option value="">Select aircraft type</option>
            <option value={AircraftType.ATR}>ATR</option>
            <option value={AircraftType.AIRBUS_320}>Airbus 320</option>
            <option value={AircraftType.BOEING_737_MAX}>Boeing 737 Max</option>
          </select>
          {validationErrors.aircraft && (
            <span className="error-text">{validationErrors.aircraft}</span>
          )}
        </div>

        <button
          type="submit"
          className="submit-button"
          disabled={isLoading || isRegenerating}
        >
          {isLoading 
            ? 'Generating...' 
            : currentVoucher?.exists 
              ? 'Vouchers Already Generated' 
              : 'Generate Vouchers'
          }
        </button>
      </form>

      {errorMessage && (
        <div className="message error-message">
          {errorMessage}
        </div>
      )}

      {successMessage && (
        <div className="message success-message">
          {successMessage}
        </div>
      )}

      {generatedSeats.length > 0 && (
        <div className="seats-display">
          <h2>
            {currentVoucher?.exists ? 'Current Seat Assignments' : 'Generated Seat Numbers'}
          </h2>
          <div className="seats-list">
            {generatedSeats.map((seat, index) => (
              <div key={index} className="seat-item">
                <span className="seat-number">{seat}</span>
                {currentVoucher?.exists && (
                  <button
                    type="button"
                    className="regenerate-btn"
                    onClick={() => handleRegenerateSeat(index + 1)}
                    disabled={isRegenerating || isLoading}
                    title={`Regenerate seat ${index + 1}`}
                  >
                    {isRegenerating ? '...' : '🔄'}
                  </button>
                )}
              </div>
            ))}
          </div>
          {currentVoucher?.exists && (
            <p className="regeneration-info">
              Click the 🔄 button next to any seat to generate a new random seat assignment.
            </p>
          )}
        </div>
      )}
    </div>
  )
}
