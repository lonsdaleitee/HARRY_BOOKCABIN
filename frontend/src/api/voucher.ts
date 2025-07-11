import axios from 'axios'
import type { CheckVoucherRequest, CheckVoucherResponse, GenerateVoucherRequest, GenerateVoucherResponse } from '../types'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const checkVoucher = async (data: CheckVoucherRequest): Promise<CheckVoucherResponse> => {
  const response = await api.post<CheckVoucherResponse>('/api/check', data)
  return response.data
}

export const generateVoucher = async (data: GenerateVoucherRequest): Promise<GenerateVoucherResponse> => {
  const response = await api.post<GenerateVoucherResponse>('/api/generate', data)
  return response.data
}
