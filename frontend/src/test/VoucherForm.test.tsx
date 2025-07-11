import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { VoucherForm } from '../components/VoucherForm'
import { AircraftType } from '../types'

// Mock the API module
vi.mock('../api/voucher', () => ({
  checkVoucher: vi.fn(),
  generateVoucher: vi.fn(),
}))

import { checkVoucher, generateVoucher } from '../api/voucher'

const mockCheckVoucher = vi.mocked(checkVoucher)
const mockGenerateVoucher = vi.mocked(generateVoucher)

describe('VoucherForm', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('renders all form fields', () => {
    render(<VoucherForm />)
    
    expect(screen.getByLabelText(/crew name/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/crew id/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/flight number/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/flight date/i)).toBeInTheDocument()
    expect(screen.getByLabelText(/aircraft type/i)).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /generate vouchers/i })).toBeInTheDocument()
  })

  it('shows validation errors for empty fields', async () => {
    const user = userEvent.setup()
    render(<VoucherForm />)
    
    const submitButton = screen.getByRole('button', { name: /generate vouchers/i })
    await user.click(submitButton)
    
    await waitFor(() => {
      expect(screen.getByText(/crew name is required/i)).toBeInTheDocument()
      expect(screen.getByText(/crew id is required/i)).toBeInTheDocument()
      expect(screen.getByText(/flight number is required/i)).toBeInTheDocument()
      expect(screen.getByText(/flight date is required/i)).toBeInTheDocument()
    })
  })

  it('fills form and submits successfully', async () => {
    const user = userEvent.setup()
    
    mockCheckVoucher.mockResolvedValue({ exists: false })
    mockGenerateVoucher.mockResolvedValue({
      success: true,
      seats: ['3B', '7C', '14D']
    })
    
    render(<VoucherForm />)
    
    // Fill form
    await user.type(screen.getByLabelText(/crew name/i), 'John Doe')
    await user.type(screen.getByLabelText(/crew id/i), '12345')
    await user.type(screen.getByLabelText(/flight number/i), 'GA102')
    await user.type(screen.getByLabelText(/flight date/i), '09-07-25')
    await user.selectOptions(screen.getByLabelText(/aircraft type/i), AircraftType.AIRBUS_320)
    
    // Submit form
    await user.click(screen.getByRole('button', { name: /generate vouchers/i }))
    
    // Wait for success message and seats
    await waitFor(() => {
      expect(screen.getByText(/vouchers generated successfully/i)).toBeInTheDocument()
      expect(screen.getByText('3B')).toBeInTheDocument()
      expect(screen.getByText('7C')).toBeInTheDocument()
      expect(screen.getByText('14D')).toBeInTheDocument()
    })
    
    expect(mockCheckVoucher).toHaveBeenCalledWith({
      flightNumber: 'GA102',
      date: '2025-07-09'
    })
    
    expect(mockGenerateVoucher).toHaveBeenCalledWith({
      name: 'John Doe',
      id: '12345',
      flightNumber: 'GA102',
      date: '2025-07-09',
      aircraft: AircraftType.AIRBUS_320
    })
  })

  it('shows error when vouchers already exist', async () => {
    const user = userEvent.setup()
    
    mockCheckVoucher.mockResolvedValue({ exists: true })
    
    render(<VoucherForm />)
    
    // Fill form
    await user.type(screen.getByLabelText(/crew name/i), 'John Doe')
    await user.type(screen.getByLabelText(/crew id/i), '12345')
    await user.type(screen.getByLabelText(/flight number/i), 'GA102')
    await user.type(screen.getByLabelText(/flight date/i), '09-07-25')
    await user.selectOptions(screen.getByLabelText(/aircraft type/i), AircraftType.AIRBUS_320)
    
    // Submit form
    await user.click(screen.getByRole('button', { name: /generate vouchers/i }))
    
    // Wait for error message
    await waitFor(() => {
      expect(screen.getByText(/vouchers have already been generated/i)).toBeInTheDocument()
    })
    
    expect(mockGenerateVoucher).not.toHaveBeenCalled()
  })
})
