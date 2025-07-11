#!/bin/bash

# Clean Integration Test for Airline Voucher Backend
# This script creates a fresh database for testing

set -e

echo "🧪 Airline Voucher Backend Integration Tests (Clean)"
echo "===================================================="

# Create a test database
TEST_DB="./test_vouchers.db"
MAIN_DB="./vouchers.db"

# Backup main database if it exists
if [ -f "$MAIN_DB" ]; then
    cp "$MAIN_DB" "${MAIN_DB}.backup"
    echo "📦 Backed up existing database"
fi

# Remove existing test database
rm -f "$TEST_DB"

# Temporarily replace main database with test database for testing
mv "$MAIN_DB" "${MAIN_DB}.temp" 2>/dev/null || true
cp "$TEST_DB" "$MAIN_DB" 2>/dev/null || touch "$MAIN_DB"

echo "🧹 Using clean test database"

# Run the integration tests
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
        elif [ -z "$expected_content" ]; then
            echo -e "${GREEN}✓ PASS${NC}"
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

# Cleanup function
cleanup() {
    echo ""
    echo "🧹 Cleaning up..."
    
    # Restore original database
    rm -f "$MAIN_DB"
    if [ -f "${MAIN_DB}.temp" ]; then
        mv "${MAIN_DB}.temp" "$MAIN_DB"
        echo "📦 Restored original database"
    elif [ -f "${MAIN_DB}.backup" ]; then
        mv "${MAIN_DB}.backup" "$MAIN_DB"
        echo "📦 Restored backup database"
    fi
    
    # Remove test database
    rm -f "$TEST_DB"
}

# Set trap to cleanup on exit
trap cleanup EXIT

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
    '{"flightNumber": "CLEAN001", "date": "2025-12-31"}' \
    200 '"exists":false'

# Test 3: Generate voucher for ATR
test_api "Generate Voucher - ATR" "POST" "$API_URL/generate" \
    '{"name": "John Doe", "id": "12345", "flightNumber": "CLEAN001", "date": "2025-12-31", "aircraft": "ATR"}' \
    200 '"success":true'

# Test 4: Check voucher now exists
test_api "Check Existing Voucher" "POST" "$API_URL/check" \
    '{"flightNumber": "CLEAN001", "date": "2025-12-31"}' \
    200 '"exists":true'

# Test 5: Try to generate duplicate voucher
test_api "Generate Duplicate Voucher" "POST" "$API_URL/generate" \
    '{"name": "Jane Doe", "id": "67890", "flightNumber": "CLEAN001", "date": "2025-12-31", "aircraft": "ATR"}' \
    409 "Voucher already exists"

echo ""
echo "📊 Test Summary"
echo "=============="
echo -e "${GREEN}✓ Core functionality tests passed!${NC}"
echo ""
echo "🎯 Backend Features Verified:"
echo "  ✓ Health check endpoint"
echo "  ✓ Check voucher existence (fresh database)"
echo "  ✓ Generate vouchers successfully"
echo "  ✓ Prevent duplicate voucher generation"
echo "  ✓ Database persistence"
echo ""
echo -e "${GREEN}🎉 Backend is working correctly!${NC}"
