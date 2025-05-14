function getMonthlyPrice(tier) {
    switch (tier) {
        case "basic":
            return 10000; // $100.00 in pennies
        case "premium":
            return 15000; // $150.00 in pennies
        case "enterprise":
            return 50000; // $500.00 in pennies
        default:
            return 0; // Return 0 pennies for invalid tiers
    }
}

// don't touch below this line
export { getMonthlyPrice };
