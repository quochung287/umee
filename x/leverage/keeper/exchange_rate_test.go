package keeper_test

//sdk "github.com/cosmos/cosmos-sdk/types"

func (s *IntegrationTestSuite) TestExchangeToken() {
	//uDenom := s.tk.FromTokenToUTokenDenom(s.ctx, umeeDenom)

	//addr1 := s.setupAccount(umeeDenom, 1000, 1000, 0, false)

	// test manipulate leverage module token balance by mint/burn
	// test manipulate IBC (or "other") module balance by mint/burn
	// SET test manipulate reserve amounts
	// SET test manipulate borrow amounts - by interest scalar and total borrows
	// SET manipulate uToken supply, including zero case
	// test other denoms do not affect umee denom

	// make exchange rate testing function that, after each scenario, manually checks multiplier,
	// then exchanges tokens and back, checking amounts.

	// make test shortcut for "function must not change exchange rate"
}

/*
func (s *IntegrationTestSuite) TestXXX() {
	uDenom := s.tk.FromTokenToUTokenDenom(s.ctx, umeeDenom)

	// set u/umee collateral amount of empty account address (error)
	err := s.tk.SetCollateralAmount(s.ctx, sdk.AccAddress{}, sdk.NewInt64Coin(uDenom, 0))
	s.Require().EqualError(err, "empty address")

	addr := sdk.AccAddress{0x01}

	// force invalid denom
	err = s.tk.SetCollateralAmount(s.ctx, addr, sdk.Coin{Denom: "", Amount: sdk.ZeroInt()})
	s.Require().EqualError(err, "0: invalid asset")

	// set u/umee collateral amount
	s.Require().NoError(s.tk.SetCollateralAmount(s.ctx, addr, sdk.NewInt64Coin(uDenom, 10)))

	// confirm effect
	collateral := s.tk.GetCollateralAmount(s.ctx, addr, uDenom)
	s.Require().Equal(sdk.NewInt64Coin(uDenom, 10), collateral)

	// set u/umee collateral amount to zero
	s.Require().NoError(s.tk.SetCollateralAmount(s.ctx, addr, sdk.NewInt64Coin(uDenom, 0)))

	// confirm effect
	collateral = s.tk.GetCollateralAmount(s.ctx, addr, uDenom)
	s.Require().Equal(sdk.NewInt64Coin(uDenom, 0), collateral)

	// set unregistered token collateral amount
	s.Require().NoError(s.tk.SetCollateralAmount(s.ctx, addr, sdk.NewInt64Coin("abcd", 10)))

	// confirm effect
	collateral = s.tk.GetCollateralAmount(s.ctx, addr, "abcd")
	s.Require().Equal(sdk.NewInt64Coin("abcd", 10), collateral)

	// set unregistered token collateral amount to zero
	s.Require().NoError(s.tk.SetCollateralAmount(s.ctx, addr, sdk.NewInt64Coin("abcd", 0)))

	// confirm effect
	collateral = s.tk.GetCollateralAmount(s.ctx, addr, "abcd")
	s.Require().Equal(sdk.NewInt64Coin("abcd", 0), collateral)

	// we do not test empty denom, as that will cause a panic
}
*/
