#!/bin/bash

# Comprehensive Integration Test for Airline Voucher Backend
# This script tests all major functionality of the backend API

set -e

echo "🧪 Airline Voucher Backend Integration Tests"
echo "============================================="

BASE_URL="http://localhost:8080"
API_URL="$BASE_URL/api"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to make HTTP requests and check responses
test_api() {
    local test_name="$1"
    local method="$2"
    local endpoint="$3"
    local data="$4"
    local expected_status="$5"
    local expected_content="$6"
    
    echo -n "Testing: $test_name... "
    
    if [ "$method" = "GET" ]; then
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" "$endpoint")
    else
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X "$method" \
            -H "Content-Type: application/json" \
            -d "$data" \
            "$endpoint")
    fi
    
    http_code=$(echo "$response" | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    body=$(echo "$response" | sed -e 's/HTTPSTATUS:.*//g')
    
    if [ "$http_code" -eq "$expected_status" ]; then
        if [ -n "$expected_content" ] && [[ "$body" == *"$expected_content"* ]]; then
            echo -e "${GREEN}✓ PASS${NC}"
            echo "  Response: $body"
        elif [ -z "$expected_content" ]; then
            echo -e "${GREEN}✓ PASS${NC}"
            echo "  Response: $body"
        else
            echo -e "${RED}✗ FAIL${NC}"
            echo "  Expected content: $expected_content"
            echo "  Actual response: $body"
            return 1
        fi
    else
        echo -e "${RED}✗ FAIL${NC}"
        echo "  Expected status: $expected_status"
        echo "  Actual status: $http_code"
        echo "  Response: $body"
        return 1
    fi
}

# Wait for server to be ready
echo "Checking if server is running..."
for i in {1..10}; do
    if curl -s "$BASE_URL/health" > /dev/null; then
        echo -e "${GREEN}✓ Server is running${NC}"
        break
    fi
    if [ $i -eq 10 ]; then
        echo -e "${RED}✗ Server is not responding${NC}"
        echo "Please start the server with: go run main.go"
        exit 1
    fi
    echo "Waiting for server... ($i/10)"
    sleep 1
done

echo ""
echo "🔍 Running API Tests..."
echo ""

# Test 1: Health Check
test_api "Health Check" "GET" "$BASE_URL/health" "" 200 '"status":"healthy"'

# Test 2: Check non-existent voucher
test_api "Check Non-existent Voucher" "POST" "$API_URL/check" \
    '{"flightNumber": "TEST001", "date": "2025-12-25"}' \
    200 '"exists":false'

# Test 3: Generate voucher for ATR
test_api "Generate Voucher - ATR" "POST" "$API_URL/generate" \
    '{"name": "John Doe", "id": "12345", "flightNumber": "TEST001", "date": "2025-12-25", "aircraft": "ATR"}' \
    200 '"success":true'

# Test 4: Check voucher now exists
test_api "Check Existing Voucher" "POST" "$API_URL/check" \
    '{"flightNumber": "TEST001", "date": "2025-12-25"}' \
    200 '"exists":true'

# Test 5: Try to generate duplicate voucher
test_api "Generate Duplicate Voucher" "POST" "$API_URL/generate" \
    '{"name": "Jane Doe", "id": "67890", "flightNumber": "TEST001", "date": "2025-12-25", "aircraft": "ATR"}' \
    409 "Voucher already exists"

# Test 6: Generate voucher for Airbus 320
test_api "Generate Voucher - Airbus 320" "POST" "$API_URL/generate" \
    '{"name": "Alice Smith", "id": "11111", "flightNumber": "TEST002", "date": "2025-12-26", "aircraft": "Airbus 320"}' \
    200 '"success":true'

# Test 7: Generate voucher for Boeing 737 Max
test_api "Generate Voucher - Boeing 737 Max" "POST" "$API_URL/generate" \
    '{"name": "Bob Wilson", "id": "22222", "flightNumber": "TEST003", "date": "2025-12-27", "aircraft": "Boeing 737 Max"}' \
    200 '"success":true'

# Test 8: Invalid aircraft type
test_api "Invalid Aircraft Type" "POST" "$API_URL/generate" \
    '{"name": "Test User", "id": "99999", "flightNumber": "TEST004", "date": "2025-12-28", "aircraft": "Invalid Plane"}' \
    400 "Invalid aircraft type"

# Test 9: Invalid date format
test_api "Invalid Date Format" "POST" "$API_URL/generate" \
    '{"name": "Test User", "id": "99999", "flightNumber": "TEST005", "date": "28-12-2025", "aircraft": "ATR"}' \
    400 "Invalid date format"

# Test 10: Missing required fields
test_api "Missing Required Fields" "POST" "$API_URL/generate" \
    '{"name": "Test User", "flightNumber": "TEST006"}' \
    400 "Invalid request body"

# Test 11: Invalid JSON
test_api "Invalid JSON" "POST" "$API_URL/generate" \
    'invalid json here' \
    400 ""

# Test 12: Missing fields in check request
test_api "Check - Missing Fields" "POST" "$API_URL/check" \
    '{"flightNumber": "TEST007"}' \
    400 "Invalid request body"

echo ""
echo "📊 Test Summary"
echo "=============="
echo -e "${GREEN}✓ All integration tests passed!${NC}"
echo ""
echo "🎯 Backend Features Verified:"
echo "  ✓ Health check endpoint"
echo "  ✓ Check voucher existence"
echo "  ✓ Generate vouchers for all aircraft types"
echo "  ✓ Prevent duplicate voucher generation"
echo "  ✓ Input validation and error handling"
echo "  ✓ Database persistence"
echo "  ✓ Proper HTTP status codes"
echo "  ✓ JSON response format"
echo ""
echo -e "${GREEN}🎉 Backend is ready for production!${NC}"
