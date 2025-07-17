import { atom } from 'jotai'
import type { VoucherFormData } from '../types'
import { AircraftType } from '../types'

// Form state
export const formDataAtom = atom<VoucherFormData>({
  crewName: '',
  crewId: '',
  flightNumber: '',
  flightDate: '',
  aircraft: AircraftType.ATR,
})

// Loading state
export const isLoadingAtom = atom(false)

// Generated seats
export const generatedSeatsAtom = atom<string[]>([])

// Error message
export const errorMessageAtom = atom<string | null>(null)

// Success message
export const successMessageAtom = atom<string | null>(null)

// Current voucher state (for regeneration)
export const currentVoucherAtom = atom<{
  exists: boolean
  seats: string[]
  flightNumber: string
  date: string
} | null>(null)

// Regeneration loading state
export const isRegeneratingAtom = atom(false)
