package promo

type RewardLevel struct {
	ClientTargets     []int
	BuyerTargets      []int
	ClientRewards     map[int]Reward
	BuyerRewards      map[int]Reward
	ReferralBonusDays int
}

var LevelDiscounts = []int{13, 24, 32}

type Reward struct {
	Days     int
	Discount int
}

var RewardLevels = map[int]RewardLevel{
	1: {
		ClientTargets: []int{1, 3, 7, 11, 15, 20},
		BuyerTargets:  []int{1, 2, 4},
		ClientRewards: map[int]Reward{
			1:  {Days: 5},
			3:  {Days: 4},
			7:  {Days: 5},
			11: {Days: 8},
			15: {Days: 10},
			20: {Days: 12, Discount: 20},
		},
		BuyerRewards: map[int]Reward{
			1: {Days: 7},
			2: {Days: 13},
			4: {Days: 21},
		},
		ReferralBonusDays: 3,
	},
	2: {
		ClientTargets: []int{1, 3, 7, 12, 17, 24},
		BuyerTargets:  []int{1, 2, 3},
		ClientRewards: map[int]Reward{
			1:  {Days: 7},
			3:  {Days: 5},
			7:  {Days: 10},
			12: {Days: 13},
			17: {Days: 15},
			24: {Days: 8, Discount: 32},
		},
		BuyerRewards: map[int]Reward{
			1: {Days: 8},
			2: {Days: 17},
			3: {Days: 25},
		},
		ReferralBonusDays: 7,
	},
}
