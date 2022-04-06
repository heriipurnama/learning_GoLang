package main

type DiscountCalculator struct {
	minimumPuchaseAmount int
	discountAmount       int
}

func NewDiscountCalculator(minimumPuchaseAmount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		minimumPuchaseAmount: minimumPuchaseAmount,
		discountAmount:       discountAmount,
	}
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPuchaseAmount {
		return purchaseAmount - c.discountAmount
	}

	return purchaseAmount
}
