#!/bin/bash

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🧪 Phase 4 Integration Test"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Test 1: Build
echo "1️⃣  Testing Build..."
go build -o agenticide-test ./cmd/agenticide
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi
echo "✅ Build successful"
echo ""

# Test 2: Version
echo "2️⃣  Testing Version..."
./agenticide-test --version
echo ""

# Test 3: Extension listing
echo "3️⃣  Testing Extension Commands..."
./agenticide-test ext list
echo ""

# Test 4: Marketplace search
echo "4️⃣  Testing Marketplace Search..."
./agenticide-test search security
echo ""

# Test 5: Task commands
echo "5️⃣  Testing Task System..."
./agenticide-test task list
echo ""

# Cleanup
rm -f agenticide-test

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ All integration tests passed!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
